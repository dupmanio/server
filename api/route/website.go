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

// WebsiteRoutes data type.
type WebsiteRoutes struct {
	handler           lib.RequestHandler
	logger            lib.Logger
	websiteController controller.WebsiteController
	authMiddleware    middleware.JWTAuthMiddleware
}

// Setup sets up WebsiteRoutes.
func (r WebsiteRoutes) Setup() {
	r.logger.Debug("Setting up Website route")

	group := r.handler.Gin.Group("/website").Use(r.authMiddleware.Handler())

	group.GET("/", r.websiteController.All)
	group.POST("/", r.websiteController.Create)
}

// NewWebsiteRoutes creates WebsiteRoutes.
func NewWebsiteRoutes(
	handler lib.RequestHandler,
	logger lib.Logger,
	websiteController controller.WebsiteController,
	authMiddleware middleware.JWTAuthMiddleware,
) WebsiteRoutes {
	return WebsiteRoutes{
		handler:           handler,
		logger:            logger,
		websiteController: websiteController,
		authMiddleware:    authMiddleware,
	}
}
