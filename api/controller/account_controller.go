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

package controller

import (
	"net/http"

	"github.com/dupman/server/dto"
	"github.com/dupman/server/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// AccountController data type.
type AccountController struct {
	httpService service.HTTPService
}

// NewAccountController creates a new AccountController.
func NewAccountController(httpService service.HTTPService) AccountController {
	return AccountController{
		httpService: httpService,
	}
}

// GetCurrentAccount gets authenticated account.
// swagger:operation GET /account Account currentUser
//
// Get current authenticated user.
//
// ---
// Security:
// - OAuth2PasswordBearer: []
//
// responses:
//   200:
//     description: Ok
//     schema:
//         $ref: "#/definitions/UserAccount"
//   401:
//     description: Access Denied
//     schema:
//         $ref: "#/definitions/HTTPError"
func (c AccountController) GetCurrentAccount(ctx *gin.Context) {
	var userAccount dto.UserAccount

	user, _ := c.httpService.CurrentUser(ctx)
	_ = copier.Copy(&userAccount, &user)

	c.httpService.HTTPResponse(ctx, http.StatusOK, userAccount)
}
