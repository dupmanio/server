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

package middleware

import "go.uber.org/fx"

// Module exports middleware module.
var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewJWTAuthMiddleware),
	fx.Provide(NewRBACMiddleware),
	fx.Provide(NewMiddlewares),
)

// IMiddleware middleware interface.
type IMiddleware interface {
	Setup()
}

// Middlewares contains multiple middleware.
type Middlewares []IMiddleware

// NewMiddlewares creates a new Middlewares
// Register the middleware that should be applied directly (globally).
func NewMiddlewares(corsMiddleware CorsMiddleware, authMiddleware JWTAuthMiddleware) Middlewares {
	return Middlewares{
		corsMiddleware,
		authMiddleware,
	}
}

// Setup sets up Middlewares.
func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
