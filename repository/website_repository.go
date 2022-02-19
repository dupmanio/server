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

// WebsiteRepository data type.
type WebsiteRepository struct {
	lib.Database
	logger lib.Logger
}

// NewWebsiteRepository creates a new WebsiteRepository.
func NewWebsiteRepository(db lib.Database, logger lib.Logger) WebsiteRepository {
	return WebsiteRepository{
		Database: db,
		logger:   logger,
	}
}

// Setup sets up WebsiteRepository.
func (r WebsiteRepository) Setup() {
	r.logger.Debug("Setting up Website repository")

	err := r.Database.AutoMigrate(&model.Website{})
	if err != nil {
		r.logger.Error(err)
	}
}
