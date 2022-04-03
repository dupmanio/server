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

package route

import (
	"github.com/dupman/server/api/controller"
	"github.com/dupman/server/api/middleware"
	"github.com/dupman/server/lib"
	"github.com/qor/roles"
)

// WebsiteRoutes data type.
type WebsiteRoutes struct {
	handler           lib.RequestHandler
	logger            lib.Logger
	websiteController controller.WebsiteController
	rbacMiddleware    middleware.RBACMiddleware
}

// Setup sets up WebsiteRoutes.
func (r WebsiteRoutes) Setup() {
	r.logger.Debug("Setting up Website route")

	group := r.handler.Gin.Group("/website")

	group.GET(
		"/",
		r.rbacMiddleware.Handler(roles.Allow(roles.Read, "user")),
		r.websiteController.All,
	)
	group.POST(
		"/",
		r.rbacMiddleware.Handler(roles.Allow(roles.CRUD, "user")),
		r.websiteController.Create,
	)
}

// NewWebsiteRoutes creates WebsiteRoutes.
func NewWebsiteRoutes(
	handler lib.RequestHandler,
	logger lib.Logger,
	websiteController controller.WebsiteController,
	rbacMiddleware middleware.RBACMiddleware,
) WebsiteRoutes {
	return WebsiteRoutes{
		handler:           handler,
		logger:            logger,
		websiteController: websiteController,
		rbacMiddleware:    rbacMiddleware,
	}
}
