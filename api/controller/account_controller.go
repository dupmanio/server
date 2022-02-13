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

	"github.com/dupman/server/constant"
	"github.com/dupman/server/dto"
	"github.com/dupman/server/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

// AccountController data type.
type AccountController struct {
	AbstractController
	userService service.UserService
}

// NewAccountController creates a new AccountController.
func NewAccountController(userService service.UserService) AccountController {
	return AccountController{
		userService: userService,
	}
}

// GetCurrentAccount gets authenticated account.
// swagger:operation GET /account Account currentUser
//
// Get current authenticated user.
//
// ---
// Security:
// - OAuth2PasswordBearer:
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
func (a AccountController) GetCurrentAccount(c *gin.Context) {
	var userAccount dto.UserAccount

	id, _ := uuid.Parse(c.GetString(constant.UserIDKey))
	user, _ := a.userService.GetUser(id)
	_ = copier.Copy(&userAccount, &user)

	a.httpService.HTTPResponse(c, http.StatusOK, userAccount)
}
