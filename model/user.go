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

// User represents user entity.
type User struct {
	Base
	Username  string `gorm:"unique"`
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
}
