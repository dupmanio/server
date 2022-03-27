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
	"github.com/qor/roles"
)

// SystemRoutes data type.
type SystemRoutes struct {
	handler          lib.RequestHandler
	logger           lib.Logger
	systemController controller.SystemController
	rbacMiddleware   middleware.RBACMiddleware
}

// Setup sets up SystemRoutes.
func (r SystemRoutes) Setup() {
	r.logger.Debug("Setting up System route")

	group := r.handler.Gin.Group("/system")

	group.Use(r.rbacMiddleware.Handler(roles.Allow(roles.CRUD, "admin", "service")))
	{
		group.GET("/websites", r.systemController.Websites)
	}
}

// NewSystemRoutes creates SystemRoutes.
func NewSystemRoutes(
	handler lib.RequestHandler,
	logger lib.Logger,
	systemController controller.SystemController,
	rbacMiddleware middleware.RBACMiddleware,
) SystemRoutes {
	return SystemRoutes{
		handler:          handler,
		logger:           logger,
		systemController: systemController,
		rbacMiddleware:   rbacMiddleware,
	}
}
