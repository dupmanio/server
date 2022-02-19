/*
 * This file is part of the dupman/server project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>, February 2022
 */

package middleware

import (
	"net/http"
	"strings"

	"github.com/dupman/server/constant"
	"github.com/dupman/server/model"
	"github.com/dupman/server/resources"
	"github.com/dupman/server/service"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware data type.
type JWTAuthMiddleware struct {
	httpService service.HTTPService
	authService service.JWTAuthService
}

// NewJWTAuthMiddleware creates a new JWTAuthMiddleware.
func NewJWTAuthMiddleware(httpService service.HTTPService, authService service.JWTAuthService) JWTAuthMiddleware {
	return JWTAuthMiddleware{
		httpService: httpService,
		authService: authService,
	}
}

// Setup sets up jwt auth middleware.
func (m JWTAuthMiddleware) Setup() {}

// Handler handles middleware functionality.
func (m JWTAuthMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
		if token != "" {
			var (
				user model.User
				err  error
			)

			if user, err = m.authService.Authorize(token); err != nil {
				m.httpService.HTTPError(ctx, http.StatusUnauthorized, err.Error())
				ctx.Abort()

				return
			}

			ctx.Set(constant.UserIDKey, user.ID.String())
			ctx.Next()

			return
		}

		m.httpService.HTTPError(ctx, http.StatusUnauthorized, resources.AccessDenied)
		ctx.Abort()
	}
}
