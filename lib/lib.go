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

package lib

import "go.uber.org/fx"

// Module exports lib module.
var Module = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(NewDatabase),
	fx.Provide(GetLogger),
	fx.Provide(NewRequestHandler),
)
