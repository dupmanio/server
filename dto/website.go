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
// swagger:model WebsiteOnCreate
type WebsiteOnCreate struct {
	// Website URL
	//
	// required: true
	// example: https://example.com
	URL string `json:"url" binding:"required,url"`

	// Website security Token
	//
	// required: true
	// example: h^djncU878*jKCN&87I#DK
	Token string `json:"token" binding:"required"`
}

// WebsiteOnResponse represents website response payload.
// swagger:model WebsiteOnResponse
type WebsiteOnResponse struct {
	// Website ID
	//
	// required: true
	// swagger:strfmt uuid
	ID uuid.UUID `json:"id" binding:"required"`

	// Website creation date and time
	//
	// required: true
	CreatedAt time.Time `json:"createdAt" binding:"required"`

	// Website update date and time
	//
	// required: true
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`

	// Website URL
	//
	// required: true
	// example: https://example.com
	URL string `json:"url" binding:"required"`
}

// WebsitesOnResponse represents multiple WebsiteOnResponse-s.
// swagger:model WebsitesOnResponse
type WebsitesOnResponse []WebsiteOnResponse

// WebsiteOnSystemResponse represents website response payload for the system route.
// swagger:model WebsiteOnSystemResponse
type WebsiteOnSystemResponse struct {
	// Website ID
	//
	// required: true
	// swagger:strfmt uuid
	ID uuid.UUID `json:"id" binding:"required"`

	// Website URL
	//
	// required: true
	// example: https://example.com
	URL string `json:"url" binding:"required"`

	// Website Token
	//
	// required: true
	// example: gwadadad1p...1Pjshdjw==
	Token string `json:"token" binding:"required"`
}

// WebsitesOnSystemResponse represents multiple WebsiteOnSystemResponse-s.
// swagger:model WebsitesOnSystemResponse
type WebsitesOnSystemResponse []WebsiteOnSystemResponse
