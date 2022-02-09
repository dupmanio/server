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
	"github.com/dupman/server/model"
	"github.com/dupman/server/resources"
	"github.com/dupman/server/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

// AuthController data type.
type AuthController struct {
	AbstractController
	authService service.JWTAuthService
	userService service.UserService
}

// NewAuthController creates a new AuthController.
func NewAuthController(
	authService service.JWTAuthService,
	userService service.UserService,
) AuthController {
	return AuthController{
		authService: authService,
		userService: userService,
	}
}

// Login authenticates user.
func (a AuthController) Login(c *gin.Context) {
	var credentials *dto.UserLogin

	if err := c.ShouldBind(&credentials); err != nil {
		a.httpService.HTTPValidationError(c, err)

		return
	}

	user, err := a.userService.GetUserByUsernameOrEmail(credentials.Username)
	if err != nil {
		a.httpService.HTTPError(c, http.StatusUnauthorized, resources.InvalidCredentials)

		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		a.httpService.HTTPError(c, http.StatusUnauthorized, resources.InvalidCredentials)

		return
	}

	token, err := a.authService.GenerateToken(&user)
	if err != nil {
		a.httpService.HTTPError(c, http.StatusUnauthorized, resources.UnableToCreateToken)

		return
	}

	a.httpService.HTTPResponse(c, http.StatusOK, token)
}

// Register creates a new user.
func (a AuthController) Register(c *gin.Context) {
	var (
		userModel = model.User{}

		userRaw        *dto.UserRegister
		userAccount    dto.UserAccount
		hashedPassword []byte
		err            error
	)

	if err = c.ShouldBind(&userRaw); err != nil {
		a.httpService.HTTPValidationError(c, err)

		return
	}

	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(userRaw.Password), bcrypt.DefaultCost); err != nil {
		a.httpService.HTTPError(c, http.StatusInternalServerError, resources.FailedHashingPassword)

		return
	}

	_ = copier.Copy(&userModel, &userRaw)
	userModel.Password = string(hashedPassword)

	if err = a.userService.CreateUser(&userModel); err != nil {
		a.httpService.HTTPError(c, http.StatusInternalServerError, err.Error())

		return
	}

	_ = copier.Copy(&userAccount, &userModel)
	a.httpService.HTTPResponse(c, http.StatusCreated, userAccount)
}
