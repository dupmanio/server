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

package seeder

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewUserSeeder),
	fx.Provide(NewSeeders),
)

type Seeder interface {
	Up()
	Down()
}

type Seeders []Seeder

func NewSeeders(userSeeder UserSeeder) Seeders {
	return Seeders{
		userSeeder,
	}
}

func (s Seeders) Up() {
	for _, seeder := range s {
		seeder.Up()
	}
}

func (s Seeders) Down() {
	for _, seeder := range s {
		seeder.Down()
	}
}
