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
// swagger:model UserAccount
type UserAccount struct {
	// User ID
	//
	// required: true
	// swagger:strfmt uuid
	ID uuid.UUID `json:"id" binding:"required"`

	// User creation date and time
	//
	// required: true
	CreatedAt time.Time `json:"createdAt" binding:"required"`

	// User update date and time
	//
	// required: true
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`

	// Username
	//
	// required: true
	// example: j_doe
	Username string `json:"username" binding:"required"`

	// User First Name
	//
	// required: true
	// example: John
	FirstName string `json:"firstName" binding:"required"`

	// User Last Name
	//
	// required: true
	// example: Doe
	LastName string `json:"lastName" binding:"required"`

	// User Email
	//
	// required: true
	// example: j_doe@dup.man
	Email string `json:"email" binding:"required"`
}

// UserLogin represents user login payload.
// swagger:model UserLogin
type UserLogin struct {
	// Username or email
	//
	// required: true
	// example: j_doe@dup.man
	Username string `json:"username" binding:"required"`

	// User password
	//
	// required: true
	// example: pa$$w0rd
	Password string `json:"password" binding:"required"`
}

// UserRegister represents registration login payload.
// swagger:model UserRegister
type UserRegister struct {
	// Username
	//
	// required: true
	// example: j_doe@dup.man
	Username string `json:"username" binding:"required,unique_username"`

	// User First Name
	//
	// required: true
	// example: John
	FirstName string `json:"firstName" binding:"required"`

	// User Last Name
	//
	// required: true
	// example: Doe
	LastName string `json:"lastName" binding:"required"`

	// User Email
	//
	// required: true
	// example: j_doe@dup.man
	Email string `json:"email" binding:"required,email,unique_email"`

	// User password
	//
	// required: true
	// example: pa$$w0rd
	// minimum: 8
	Password string `json:"password" binding:"required,min=8"`
}
