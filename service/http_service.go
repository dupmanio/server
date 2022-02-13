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

package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dupman/server/dto"
	"github.com/dupman/server/resources"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// HTTPService data type.
type HTTPService struct{}

// NewHTTPService creates a new HTTPService.
func NewHTTPService() HTTPService {
	return HTTPService{}
}

// HTTPError sends HTTP error response.
func (s HTTPService) HTTPError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, dto.HTTPError{Code: code, Error: message})
}

// HTTPResponse sends HTTP response.
func (s HTTPService) HTTPResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, dto.HTTPResponse{Code: code, Data: data})
}

// HTTPValidationError sends HTTP validation response.
func (s HTTPService) HTTPValidationError(c *gin.Context, err error) {
	var validationErr validator.ValidationErrors
	if errors.As(err, &validationErr) {
		s.HTTPError(c, http.StatusBadRequest, s.formatValidationErrors(validationErr))

		return
	}

	s.HTTPError(c, http.StatusBadRequest, err.Error())
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
		default:
			errorMessage = fieldError.Error()
		}

		errors = append(errors, errorMessage)
	}

	return errors
}
