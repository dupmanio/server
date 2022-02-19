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
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

// Module exports validator module.
var Module = fx.Options(
	fx.Provide(NewUniqueUsernameOrEmailValidator),
	fx.Provide(NewURLValidator),
	fx.Provide(NewValidators),
)

// Validators contains multiple validators.
type Validators map[string]IValidator

// IValidator interface.
type IValidator interface {
	Validate(fl validator.FieldLevel) bool
}

// NewValidators creates a new Validators.
func NewValidators(
	uniqueUsernameOrEmail UniqueUsernameOrEmailValidator,
	urlValidator URLValidator,
) Validators {
	return Validators{
		"unique_username": uniqueUsernameOrEmail,
		"unique_email":    uniqueUsernameOrEmail,
		"url":             urlValidator,
	}
}

// Setup sets up Validators.
func (v Validators) Setup() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for tag, validationHandler := range v {
			_ = val.RegisterValidation(tag, validationHandler.Validate)
		}
	}
}
