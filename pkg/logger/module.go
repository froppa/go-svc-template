package logger

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Module provides a *zap.Logger to an Fx application and
// ensures any buffered logs are flushed on shutdown.
var Module = fx.Module(
	"logger",
	fx.Provide(NewLogger),
	fx.Invoke(RegisterLoggerCleanup),
)

// NewLogger builds a development‐mode logger.
func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return cfg.Build()
}

// RegisterLoggerCleanup registers a Lifecycle hook that flushes
// any buffered log entries on application stop. Sync errors
// (e.g. “inappropriate ioctl for device” on stderr) are dropped.
func RegisterLoggerCleanup(lc fx.Lifecycle, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			_ = logger.Sync()
			return nil
		},
	})
}
