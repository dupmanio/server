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
	"net/url"

	"github.com/go-playground/validator/v10"
)

// URLValidator data structure.
type URLValidator struct{}

// NewURLValidator creates a new URLValidator.
func NewURLValidator() URLValidator {
	return URLValidator{}
}

// Validate validates field.
func (v URLValidator) Validate(fl validator.FieldLevel) bool {
	u, err := url.Parse(fl.Field().String())

	return err == nil && u.Scheme != "" && u.Host != ""
}
