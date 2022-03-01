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
	"github.com/dupman/server/lib"
)

// SystemRoutes data type.
type SystemRoutes struct {
	handler          lib.RequestHandler
	logger           lib.Logger
	systemController controller.SystemController
}

// Setup sets up SystemRoutes.
func (r SystemRoutes) Setup() {
	r.logger.Debug("Setting up System route")

	// @todo: implement security middleware.
	group := r.handler.Gin.Group("/system")

	group.GET("/websites", r.systemController.Websites)
}

// NewSystemRoutes creates SystemRoutes.
func NewSystemRoutes(
	handler lib.RequestHandler,
	logger lib.Logger,
	systemController controller.SystemController,
) SystemRoutes {
	return SystemRoutes{
		handler:          handler,
		logger:           logger,
		systemController: systemController,
	}
}
