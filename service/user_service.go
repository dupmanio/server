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
	"database/sql"

	"github.com/dupman/server/model"
	"github.com/dupman/server/repository"
	"github.com/google/uuid"
)

// UserService data structure.
type UserService struct {
	repository repository.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(repository repository.UserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

// Get gets one user.
func (s UserService) Get(id uuid.UUID) (user model.User, err error) {
	return user, s.repository.Joins("KeyPair").First(&user, id).Error
}

// CreateUser creates the user.
func (s UserService) CreateUser(user *model.User) error {
	return s.repository.Create(&user).Error
}

// GetByUsernameOrEmail loads the user by username or email.
func (s UserService) GetByUsernameOrEmail(username string) (user model.User, err error) {
	return user, s.repository.
		Where("username = @username OR email = @username", sql.Named("username", username)).
		Joins("KeyPair").
		First(&user).Error
}
