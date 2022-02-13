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

package dto

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	errTokenExpired          = errors.New("token has expired")
	errTokenUsedBeforeIssued = errors.New("token used before issued")
)

// JWTResponse represents JWT token data.
// swagger:model JWTResponse
type JWTResponse struct {
	// JWT Access Token
	//
	// required: true
	// example: eyJhbGciOiJSUz...dAlCslnO3YqiCA
	AccessToken string `json:"access_token"`

	// JWT Token Type
	//
	// required: true
	// example: Bearer
	TokenType string `json:"token_type"`

	// JWT Token Expires In
	//
	// required: true
	// example: 3600
	ExpiresIn int64 `json:"expires_in"`
}

// JWTClaims represents JWT token claim.
type JWTClaims struct {
	ID        uuid.UUID `json:"sub,omitempty"`
	ExpiresAt int64     `json:"exp,omitempty"`
	IssuedAt  int64     `json:"iat,omitempty"`
}

// Valid validates JWTClaims.
func (c JWTClaims) Valid() (err error) {
	now := time.Now()
	if now.After(time.Unix(c.ExpiresAt, 0)) {
		err = errTokenExpired
	}

	if now.Before(time.Unix(c.IssuedAt, 0)) {
		err = errTokenUsedBeforeIssued
	}

	return err
}
