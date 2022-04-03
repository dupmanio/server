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

package lib

import (
	"strings"

	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// FxLogger structure.
type FxLogger struct {
	*Logger
}

// GetFxLogger get the fx logger.
func (l Logger) GetFxLogger() FxLogger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)

	return FxLogger{
		Logger: newSugaredLogger(logger),
	}
}

// LogEvent logs the given event to the provided Zap logger.
func (l FxLogger) LogEvent(event fxevent.Event) {
	switch event := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Logger.Debug("OnStart hook executing",
			zap.String("callee", event.FunctionName),
			zap.String("caller", event.CallerName),
		)
	case *fxevent.OnStartExecuted:
		if event.Err != nil {
			l.Logger.Debug("OnStart hook failed",
				zap.String("callee", event.FunctionName),
				zap.String("caller", event.CallerName),
				zap.Error(event.Err),
			)
		} else {
			l.Logger.Debug("OnStart hook executed",
				zap.String("callee", event.FunctionName),
				zap.String("caller", event.CallerName),
				zap.String("runtime", event.Runtime.String()),
			)
		}
	case *fxevent.OnStopExecuting:
		l.Logger.Debug("OnStop hook executing",
			zap.String("callee", event.FunctionName),
			zap.String("caller", event.CallerName),
		)
	case *fxevent.OnStopExecuted:
		if event.Err != nil {
			l.Logger.Debug("OnStop hook failed",
				zap.String("callee", event.FunctionName),
				zap.String("caller", event.CallerName),
				zap.Error(event.Err),
			)
		} else {
			l.Logger.Debug("OnStop hook executed",
				zap.String("callee", event.FunctionName),
				zap.String("caller", event.CallerName),
				zap.String("runtime", event.Runtime.String()),
			)
		}
	case *fxevent.Supplied:
		l.Logger.Debug("supplied", zap.String("type", event.TypeName), zap.Error(event.Err))
	case *fxevent.Provided:
		for _, rtype := range event.OutputTypeNames {
			l.Logger.Debug("provided",
				zap.String("constructor", event.ConstructorName),
				zap.String("type", rtype),
			)
		}

		if event.Err != nil {
			l.Logger.Error("error encountered while applying options",
				zap.Error(event.Err))
		}
	case *fxevent.Invoking:
		// Do not log stack as it will make logs hard to read.
		l.Logger.Debug("invoking",
			zap.String("function", event.FunctionName))
	case *fxevent.Invoked:
		if event.Err != nil {
			l.Logger.Debug("invoke failed",
				zap.Error(event.Err),
				zap.String("stack", event.Trace),
				zap.String("function", event.FunctionName))
		}
	case *fxevent.Stopping:
		l.Logger.Debug("received signal",
			zap.String("signal", strings.ToUpper(event.Signal.String())))
	case *fxevent.Stopped:
		if event.Err != nil {
			l.Logger.Error("stop failed", zap.Error(event.Err))
		}
	case *fxevent.RollingBack:
		l.Logger.Debug("start failed, rolling back", zap.Error(event.StartErr))
	case *fxevent.RolledBack:
		if event.Err != nil {
			l.Logger.Error("rollback failed", zap.Error(event.Err))
		}
	case *fxevent.Started:
		if event.Err != nil {
			l.Logger.Error("start failed", zap.Error(event.Err))
		} else {
			l.Logger.Debug("started")
		}
	case *fxevent.LoggerInitialized:
		if event.Err != nil {
			l.Logger.Error("custom logger initialization failed", zap.Error(event.Err))
		} else {
			l.Logger.Debug("initialized custom fxevent.Logger", zap.String("function", event.ConstructorName))
		}
	}
}
