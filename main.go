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

package main

import (
	"github.com/dupman/server/bootstrap"
	"github.com/dupman/server/lib"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	logger := fx.WithLogger(func(logger lib.Logger) fxevent.Logger {
		return logger.GetFxLogger()
	})

	fx.New(bootstrap.Module, logger).Run()
}
