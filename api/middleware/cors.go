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

package middleware

import (
	"time"

	"github.com/dupman/server/lib"
	"github.com/gin-contrib/cors"
)

const corsMaxAge = 12 * time.Hour

// CorsMiddleware data type.
type CorsMiddleware struct {
	handler lib.RequestHandler
	logger  lib.Logger
	config  lib.CORSConfig
}

// NewCorsMiddleware creates a new CorsMiddleware.
func NewCorsMiddleware(handler lib.RequestHandler, logger lib.Logger, config lib.Config) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
		logger:  logger,
		config:  config.CORS,
	}
}

// Setup sets up cors middleware.
func (m CorsMiddleware) Setup() {
	m.logger.Debug("Setting up CORS middleware")

	m.handler.Gin.Use(cors.New(cors.Config{
		AllowOrigins:     m.config.Origins,
		AllowMethods:     m.config.Methods,
		AllowHeaders:     m.config.Headers,
		AllowCredentials: true,
		MaxAge:           corsMaxAge,
	}))
}
