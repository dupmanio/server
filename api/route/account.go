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

import (
	"github.com/dupman/server/api/controller"
	"github.com/dupman/server/api/middleware"
	"github.com/dupman/server/lib"
)

// AccountRoutes data type.
type AccountRoutes struct {
	handler           lib.RequestHandler
	logger            lib.Logger
	accountController controller.AccountController
	authMiddleware    middleware.JWTAuthMiddleware
}

// Setup sets up AccountRoutes.
func (r AccountRoutes) Setup() {
	r.logger.Debug("Setting up Account route")

	api := r.handler.Gin.Group("/account").Use(r.authMiddleware.Handler())

	api.GET("/", r.accountController.GetCurrentAccount)
}

// NewAccountRoutes creates AccountRoutes.
func NewAccountRoutes(
	handler lib.RequestHandler,
	logger lib.Logger,
	accountController controller.AccountController,
	authMiddleware middleware.JWTAuthMiddleware,
) AccountRoutes {
	return AccountRoutes{
		handler:           handler,
		logger:            logger,
		accountController: accountController,
		authMiddleware:    authMiddleware,
	}
}
