/*
 * This file is part of the dupman/server project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>, March 2022
 */

package middleware

import (
	"net/http"

	"github.com/dupman/server/resources"
	"github.com/dupman/server/service"
	"github.com/gin-gonic/gin"
	"github.com/qor/roles"
)

// RBACMiddleware data type.
type RBACMiddleware struct {
	httpService service.HTTPService
}

// NewRBACMiddleware creates a new RBACMiddleware.
func NewRBACMiddleware(httpService service.HTTPService) RBACMiddleware {
	return RBACMiddleware{
		httpService: httpService,
	}
}

// Setup sets up the RBAC middleware.
func (m RBACMiddleware) Setup() {}

// Handler handles middleware functionality.
func (m RBACMiddleware) Handler(permission *roles.Permission) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var permissionMode roles.PermissionMode

		user, _ := m.httpService.CurrentUser(ctx)
		userRolesRaw := user.GetRoles()

		userRoles := make([]interface{}, len(userRolesRaw))
		for i, role := range userRolesRaw {
			userRoles[i] = role
		}

		switch ctx.Request.Method {
		case "GET":
			permissionMode = roles.Read
		case "POST":
			permissionMode = roles.Create
		case "PUT":
			permissionMode = roles.Update
		case "DELETE":
			permissionMode = roles.Delete
		default:
			permissionMode = roles.Read
		}

		if !permission.HasPermission(permissionMode, userRoles...) {
			m.httpService.HTTPError(ctx, http.StatusUnauthorized, resources.AccessDenied)
			ctx.Abort()

			return
		}
	}
}
