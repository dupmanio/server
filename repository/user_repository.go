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

package repository

import (
	"github.com/dupman/server/lib"
	"github.com/dupman/server/model"
)

// UserRepository data type.
type UserRepository struct {
	lib.Database
	logger lib.Logger
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(db lib.Database, logger lib.Logger) UserRepository {
	return UserRepository{
		Database: db,
		logger:   logger,
	}
}

// Setup sets up UserRepository.
func (r UserRepository) Setup() {
	r.logger.Debug("Setting up User repository")

	err := r.Database.AutoMigrate(&model.User{})
	if err != nil {
		r.logger.Error(err)
	}
}
