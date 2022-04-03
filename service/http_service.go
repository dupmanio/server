/*
 * This file is part of the dupman/server project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>
 */

package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dupman/server/constant"
	"github.com/dupman/server/dto"
	"github.com/dupman/server/helper"
	"github.com/dupman/server/model"
	"github.com/dupman/server/resources"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// HTTPService data type.
type HTTPService struct {
	userService UserService
}

// NewHTTPService creates a new HTTPService.
func NewHTTPService(userService UserService) HTTPService {
	return HTTPService{
		userService: userService,
	}
}

// HTTPError sends HTTP error response.
func (s HTTPService) HTTPError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, dto.HTTPError{Code: code, Error: message})
}

// HTTPResponse sends HTTP response.
func (s HTTPService) HTTPResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, dto.HTTPResponse{Code: code, Data: data})
}

// HTTPPaginatedResponse sends HTTP response with pagination.
func (s HTTPService) HTTPPaginatedResponse(c *gin.Context, code int, data interface{}, pagination helper.Pagination) {
	c.JSON(code, dto.HTTPResponse{Code: code, Data: data, Pagination: pagination})
}

// HTTPValidationError sends HTTP validation response.
func (s HTTPService) HTTPValidationError(c *gin.Context, err error) {
	s.HTTPError(c, http.StatusBadRequest, s.NormalizeHTTPValidationError(err))
}

// NormalizeHTTPValidationError Normalizes the HTTP validation error.
func (s HTTPService) NormalizeHTTPValidationError(err error) []string {
	var validationErr validator.ValidationErrors
	if errors.As(err, &validationErr) {
		return s.formatValidationErrors(validationErr)
	}

	return []string{err.Error()}
}

// CurrentUserID returns the current user's ID.
func (s HTTPService) CurrentUserID(ctx *gin.Context) uuid.UUID {
	if id, err := uuid.Parse(ctx.GetString(constant.UserIDKey)); err == nil {
		return id
	}

	return uuid.Nil
}

// CurrentUser returns the current user.
func (s HTTPService) CurrentUser(ctx *gin.Context) (user model.User, err error) {
	if user, err = s.userService.Get(s.CurrentUserID(ctx)); err != nil {
		user.Roles = []model.Role{{Name: "anonymous"}}

		return user, err
	}

	return user, nil
}

// Paginate returns pagination object from the request.
func (s HTTPService) Paginate(ctx *gin.Context) *helper.Pagination {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	return &helper.Pagination{
		Limit: limit,
		Page:  page,
	}
}

func (s HTTPService) formatValidationErrors(validationErrors validator.ValidationErrors) (errors []string) {
	for _, fieldError := range validationErrors {
		var errorMessage string

		switch fieldError.Tag() {
		case "required":
			errorMessage = fmt.Sprintf(resources.KeyIsRequired, fieldError.Field())
		case "min":
			errorMessage = fmt.Sprintf(resources.ValueIsLessThenMin, fieldError.Field(), fieldError.Param())
		case "email":
			errorMessage = fmt.Sprintf(resources.ValueIsNotEmail, fieldError.Field())
		case "unique_username":
			errorMessage = resources.UsernameIsTaken
		case "unique_email":
			errorMessage = resources.EmailIsTaken
		case "url":
			errorMessage = fmt.Sprintf(resources.ValueIsNotURL, fieldError.Field())
		default:
			errorMessage = fieldError.Error()
		}

		errors = append(errors, errorMessage)
	}

	return errors
}
