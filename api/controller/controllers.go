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

package controller

import "go.uber.org/fx"

// Module exports controller module.
var Module = fx.Options(
	fx.Provide(NewAccountController),
	fx.Provide(NewAuthController),
	fx.Provide(NewSystemController),
	fx.Provide(NewWebsiteController),
)
