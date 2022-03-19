/*
 * This file is part of the dupman/server project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>, March 2022
 */

package repository

import (
	"github.com/dupman/server/lib"
	"github.com/dupman/server/model"
)

// RoleRepository data type.
type RoleRepository struct {
	lib.Database
	logger lib.Logger
}

// NewRoleRepository creates a new RoleRepository.
func NewRoleRepository(db lib.Database, logger lib.Logger) RoleRepository {
	return RoleRepository{
		Database: db,
		logger:   logger,
	}
}

// Setup sets up RoleRepository.
func (r RoleRepository) Setup() {
	r.logger.Debug("Setting up Role repository")

	err := r.Database.AutoMigrate(&model.Role{})
	if err != nil {
		r.logger.Error(err)
	}

	// Setup Roles.
	roles := []model.Role{
		{Name: "anonymous", Description: "Unauthenticated user"},
		{Name: "user", Description: "Authenticated User"},
		{Name: "admin", Description: "Administrator"},
		{Name: "service", Description: "Service User"},
	}

	for i := range roles {
		err = r.Database.Create(&roles[i]).Error
		if err != nil {
			r.logger.Warn(err)
		}
	}
}
