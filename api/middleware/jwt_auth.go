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
	"net/http"
	"strings"

	"github.com/dupman/server/constant"
	"github.com/dupman/server/lib"
	"github.com/dupman/server/model"
	"github.com/dupman/server/service"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware data type.
type JWTAuthMiddleware struct {
	handler     lib.RequestHandler
	logger      lib.Logger
	httpService service.HTTPService
	authService service.JWTAuthService
}

// NewJWTAuthMiddleware creates a new JWTAuthMiddleware.
func NewJWTAuthMiddleware(
	handler lib.RequestHandler,
	logger lib.Logger,
	httpService service.HTTPService,
	authService service.JWTAuthService,
) JWTAuthMiddleware {
	return JWTAuthMiddleware{
		handler:     handler,
		logger:      logger,
		httpService: httpService,
		authService: authService,
	}
}

// Setup sets up jwt auth middleware.
func (m JWTAuthMiddleware) Setup() {
	m.logger.Debug("Setting up JWT Auth middleware")

	m.handler.Gin.Use(m.Handler())
}

// Handler handles middleware functionality.
func (m JWTAuthMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error

		var user model.User

		token := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
		if token != "" {
			if user, err = m.authService.Authorize(token); err != nil {
				m.httpService.HTTPError(ctx, http.StatusUnauthorized, err.Error())
				ctx.Abort()

				return
			}
		}

		ctx.Set(constant.UserIDKey, user.ID.String())
	}
}
