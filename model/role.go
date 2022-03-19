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

package model

// Role represents user role entity.
type Role struct {
	Name        string `gorm:"type:varchar(255);primary_key"`
	Description string `gorm:"type:varchar(255);"`
}
