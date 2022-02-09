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
	"time"

	"github.com/google/uuid"
)

// UserAccount represents user accounts data.
type UserAccount struct {
	ID        uuid.UUID `json:"id" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	Email     string    `json:"email" binding:"required"`
}

// UserLogin represents user login payload.
type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserRegister represents registration login payload.
type UserRegister struct {
	Username  string `json:"username" binding:"required,unique_username"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email,unique_email"`
	Password  string `json:"password" binding:"required,min=8"`
}
