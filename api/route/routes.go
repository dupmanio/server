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

package route

import "go.uber.org/fx"

// Module exports route module.
var Module = fx.Options(
	fx.Provide(NewAccountRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes.
type Routes []Route

// Route interface.
type Route interface {
	Setup()
}

// NewRoutes creates a new Routes.
func NewRoutes(
	accountRoutes AccountRoutes,
	authRoutes AuthRoutes,
) Routes {
	return Routes{
		accountRoutes,
		authRoutes,
	}
}

// Setup sets up Routes.
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
