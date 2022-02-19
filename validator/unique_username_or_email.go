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

package validator

import (
	"github.com/dupman/server/service"
	"github.com/go-playground/validator/v10"
)

// UniqueUsernameOrEmailValidator data structure.
type UniqueUsernameOrEmailValidator struct {
	userService service.UserService
}

// NewUniqueUsernameOrEmailValidator creates a new NewUniqueUsernameOrEmailValidator.
func NewUniqueUsernameOrEmailValidator(userService service.UserService) UniqueUsernameOrEmailValidator {
	return UniqueUsernameOrEmailValidator{
		userService: userService,
	}
}

// Validate validates field.
func (v UniqueUsernameOrEmailValidator) Validate(fl validator.FieldLevel) bool {
	_, err := v.userService.GetByUsernameOrEmail(fl.Field().String())

	return err != nil
}
