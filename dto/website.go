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

// WebsiteOnCreate represents website creation payload.
type WebsiteOnCreate struct {
	URL   string `json:"url" binding:"required,url"`
	Token string `json:"token" binding:"required"`
}

// WebsiteOnResponse represents website response payload.
type WebsiteOnResponse struct {
	ID        uuid.UUID `json:"id" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`
	URL       string    `json:"url" binding:"required"`
}

// WebsitesOnResponse represents multiple WebsiteOnResponse-s.
type WebsitesOnResponse []WebsiteOnResponse
