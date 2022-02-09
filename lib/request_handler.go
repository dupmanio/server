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

package lib

import (
	"log"

	"github.com/gin-gonic/gin"
)

// RequestHandler function.
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates a new request handler.
func NewRequestHandler(config Config, logger Logger) RequestHandler {
	gin.DefaultWriter = logger.GetGinLogger()

	if config.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	if err := engine.SetTrustedProxies(config.Server.TrustedProxies); err != nil {
		log.Fatalln(err)
	}

	return RequestHandler{Gin: engine}
}
