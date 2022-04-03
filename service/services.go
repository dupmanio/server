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

import "go.uber.org/fx"

// Module exports service module.
var Module = fx.Options(
	fx.Provide(NewHTTPService),
	fx.Provide(NewJWTAuthService),
	fx.Provide(NewUserService),
	fx.Provide(NewWebsiteService),
)
