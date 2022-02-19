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
	sqlType "github.com/dupman/server/sql/type"
	"github.com/google/uuid"
)

// Website represents website entity.
type Website struct {
	Base
	URL    string
	Token  sqlType.WebsiteToken
	UserID uuid.UUID
}
