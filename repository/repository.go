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

import "go.uber.org/fx"

// Module exports repository module.
var Module = fx.Options(
	fx.Provide(NewUserRepository),
	fx.Provide(NewRepositories),
)

// Repositories contains multiple repositories.
type Repositories []IRepository

// IRepository repository interface.
type IRepository interface {
	Setup()
}

// NewRepositories creates a new Repositories.
func NewRepositories(
	userRepository UserRepository,
) Repositories {
	return Repositories{
		userRepository,
	}
}

// Setup sets up Repositories.
func (r Repositories) Setup() {
	for _, repository := range r {
		repository.Setup()
	}
}
