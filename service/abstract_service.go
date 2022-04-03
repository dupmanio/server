/*
 * This file is part of the dupman/server project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>
 */

package service

import (
	"math"

	"github.com/dupman/server/helper"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AbstractService data structure.
type AbstractService struct{}

func (s AbstractService) paginate(
	value interface{},
	pagination *helper.Pagination,
	db *gorm.DB,
) func(db *gorm.DB) *gorm.DB {
	var totalItems int64

	db.Model(value).Count(&totalItems)

	pagination.TotalItems = totalItems
	pagination.TotalPages = int(math.Ceil(float64(totalItems) / float64(pagination.GetLimit())))

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}

func (s AbstractService) withUser(userID uuid.UUID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id", userID)
	}
}
