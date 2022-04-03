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
	"github.com/dupman/server/lib"
)

// AuthRoutes data type.
type AuthRoutes struct {
	handler        lib.RequestHandler
	logger         lib.Logger
	authController controller.AuthController
}

// Setup sets up AuthRoutes.
func (r AuthRoutes) Setup() {
	r.logger.Debug("Setting up Auth routes")

	group := r.handler.Gin.Group("/auth")

	group.POST("/token", r.authController.Token)
	group.POST("/register", r.authController.Register)
}

// NewAuthRoutes creates AuthRoutes.
func NewAuthRoutes(
	handler lib.RequestHandler,
	logger lib.Logger,
	authController controller.AuthController,
) AuthRoutes {
	return AuthRoutes{
		handler:        handler,
		logger:         logger,
		authController: authController,
	}
}
