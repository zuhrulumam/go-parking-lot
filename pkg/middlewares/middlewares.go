package middlewares

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/zuhrulumam/go-parking-lot/pkg/ctxkeys"
	"go.uber.org/zap"
)

func RequestContextMiddleware(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Generate correlation ID if not present
		correlationID := c.Get("X-Correlation-ID", uuid.New().String())

		// Create context with values
		ctx := context.WithValue(c.Context(), ctxkeys.CtxKeyCorrelationID, correlationID)
		ctx = context.WithValue(ctx, ctxkeys.CtxKeyApp, "parking-service")
		ctx = context.WithValue(ctx, ctxkeys.CtxKeyRuntime, "go")
		ctx = context.WithValue(ctx, ctxkeys.CtxKeyEnv, "production") // or from env
		ctx = context.WithValue(ctx, ctxkeys.CtxKeyAppVersion, "v1.0.0")
		ctx = context.WithValue(ctx, ctxkeys.CtxKeyPath, c.Path())
		ctx = context.WithValue(ctx, ctxkeys.CtxKeyMethod, c.Method())
		ctx = context.WithValue(ctx, ctxkeys.CtxKeyIP, c.IP())
		ctx = context.WithValue(ctx, ctxkeys.CtxKeyPort, c.Port())
		ctx = context.WithValue(ctx, ctxkeys.CtxKeySrcIP, c.Context().RemoteAddr().String())
		ctx = context.WithValue(ctx, ctxkeys.CtxKeyHeader, c.GetReqHeaders())

		// Store context
		c.Locals("ctx", ctx)

		// Continue
		err := c.Next()

		// End time
		duration := time.Since(start)

		// Response status and logging
		logger.Info("HTTP Request",
			zap.String("path", c.Path()),
			zap.String("method", c.Method()),
			zap.String("correlation_id", correlationID),
			zap.Int("status", c.Response().StatusCode()),
			zap.String("duration", duration.String()),
			zap.Any("header", c.GetReqHeaders()),
		)

		return err
	}
}
