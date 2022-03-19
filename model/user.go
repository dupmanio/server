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

package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents user entity.
type User struct {
	Base
	Username  string `gorm:"unique"`
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
	Websites  []Website
	KeyPairID uuid.UUID `gorm:"->;<-:create"`
	KeyPair   KeyPair   `gorm:"->;<-:create"`
	Roles     []Role    `gorm:"many2many:user_roles"`
}

func (e *User) BeforeCreate(tx *gorm.DB) (err error) {
	if err = e.Base.BeforeCreate(tx); err != nil {
		return err
	}

	e.KeyPair = KeyPair{
		PrivateKey: "tmp",
	}

	e.Roles = []Role{
		{Name: "user"},
	}

	return err
}

func (e *User) GetRoles() (roles []string) {
	for _, role := range e.Roles {
		roles = append(roles, role.Name)
	}

	return roles
}
