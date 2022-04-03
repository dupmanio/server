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

package helper

import (
	"context"
	"net/http"
	"testing"

	"github.com/dupman/server/api/controller"
	"github.com/dupman/server/api/middleware"
	"github.com/dupman/server/api/route"
	"github.com/dupman/server/lib"
	"github.com/dupman/server/repository"
	"github.com/dupman/server/service"
	"github.com/dupman/server/test/seeder"
	"github.com/dupman/server/validator"
	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

type TesterSetup struct {
	handler *gin.Engine
	Seeders seeder.Seeders
	Expect  *httpexpect.Expect
}

var testerSetup TesterSetup

func BootstrapSuite(t *testing.T, testSuite suite.TestingSuite) {
	t.Helper()

	logger := fx.WithLogger(func(logger lib.Logger) fxevent.Logger {
		return logger.GetFxLogger()
	})

	module := fx.Options(
		controller.Module,
		middleware.Module,
		route.Module,
		lib.Module,
		repository.Module,
		service.Module,
		validator.Module,
		seeder.Module,
		fx.Invoke(bootstrapWrapper(t, testSuite)),
	)

	app := fx.New(module, logger)

	startCtx, cancel := context.WithTimeout(context.Background(), app.StartTimeout())
	defer cancel()

	if err := app.Start(startCtx); err != nil {
		panic(err)
	}
}

func SetupTester(t *testing.T) *TesterSetup {
	t.Helper()

	testerSetup.Expect = httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(testerSetup.handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	return &testerSetup
}

func bootstrapWrapper(t *testing.T, testSuite suite.TestingSuite) func(
	lib.RequestHandler,
	route.Routes,
	middleware.Middlewares,
	repository.Repositories,
	validator.Validators,
	seeder.Seeders,
) {
	t.Helper()

	return func(
		handler lib.RequestHandler,
		routes route.Routes,
		middlewares middleware.Middlewares,
		repositories repository.Repositories,
		validators validator.Validators,
		seeders seeder.Seeders,
	) {
		middlewares.Setup()
		routes.Setup()
		repositories.Setup()
		validators.Setup()

		testerSetup = TesterSetup{
			handler: handler.Gin,
			Seeders: seeders,
		}

		suite.Run(t, testSuite)
	}
}
