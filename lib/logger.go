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
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger structure.
type Logger struct {
	*zap.SugaredLogger
}

var (
	globalLogger *Logger
	zapLogger    *zap.Logger
)

// GetLogger get the logger.
func GetLogger(config Config) Logger {
	if globalLogger == nil {
		logger := newLogger(config.Env)
		globalLogger = &logger
	}

	return *globalLogger
}

func newSugaredLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

func newLogger(environment string) Logger {
	var config zap.Config

	if environment == "prod" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.Encoding = "console"
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapLogger, _ = config.Build()
	logger := newSugaredLogger(zapLogger)

	return *logger
}
