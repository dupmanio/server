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

// KeyPairRepository data type.
type KeyPairRepository struct {
	lib.Database
	logger lib.Logger
}

// NewKeyPairRepository creates a new KeyPairRepository.
func NewKeyPairRepository(db lib.Database, logger lib.Logger) KeyPairRepository {
	return KeyPairRepository{
		Database: db,
		logger:   logger,
	}
}

// Setup sets up KeyPairRepository.
func (r KeyPairRepository) Setup() {
	r.logger.Debug("Setting up Key Pair repository")

	err := r.Database.AutoMigrate(&model.KeyPair{})
	if err != nil {
		r.logger.Error(err)
	}
}
