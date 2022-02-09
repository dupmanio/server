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
	"strings"

	"go.uber.org/zap"
)

// GinLogger structure.
type GinLogger struct {
	*Logger
}

// GetGinLogger get the gin logger.
func (l Logger) GetGinLogger() GinLogger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)

	return GinLogger{
		Logger: newSugaredLogger(logger),
	}
}

// Write interface implementation for gin-framework.
func (l GinLogger) Write(data []byte) (n int, err error) {
	message := strings.TrimSuffix(string(data), "\n")

	switch {
	case strings.HasPrefix(message, "[GIN-debug]"):
		l.Debug(message)
	case strings.HasPrefix(message, "[GIN-error]"):
		l.Error(message)
	default:
		l.Info(message)
	}

	return len(data), nil
}
