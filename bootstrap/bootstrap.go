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

package bootstrap

import (
	"context"
	"fmt"

	"github.com/dupman/server/api/controller"
	"github.com/dupman/server/api/middleware"
	"github.com/dupman/server/api/route"
	"github.com/dupman/server/lib"
	"github.com/dupman/server/repository"
	"github.com/dupman/server/service"
	"github.com/dupman/server/validator"
	"go.uber.org/fx"
)

// Module exports bootstrap module.
var Module = fx.Options(
	controller.Module,
	middleware.Module,
	route.Module,
	lib.Module,
	repository.Module,
	service.Module,
	validator.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes route.Routes,
	config lib.Config,
	middlewares middleware.Middlewares,
	logger lib.Logger,
	database lib.Database,
	repositories repository.Repositories,
	validators validator.Validators,
) {
	connection, _ := database.DB.DB()

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Debug("Starting Application")
			logger.Info("\n" +
				"     _                                           \n" +
				"  __| | _   _  _ __   _ __ ___    __ _  _ __     \n" +
				" / _` || | | || '_ \\ | '_ ` _ \\  / _` || '_ \\ \n" +
				"| (_| || |_| || |_+ || | | | | || (_| || | | |   \n" +
				" \\__,_| \\__,_|| .__/ |_| |_| |_| \\__,_||_| |_|\n" +
				"              |_|                                ")

			go func() {
				middlewares.Setup()
				routes.Setup()
				repositories.Setup()
				validators.Setup()

				err := handler.Gin.Run(fmt.Sprintf("%s:%d", config.Server.ListenAddr, config.Server.Port))
				if err != nil {
					logger.Fatal(err)
				}
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			logger.Debug("Stopping Application")
			connection.Close()

			return nil
		},
	})
}
