package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"log/slog"
	"net/http"
	"runtime/debug"
)

func ZapLoggerMiddleware(zapLogger *zap.Logger) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {

			// Panic recovery
			defer func() {
				if r := recover(); r != nil {
					recoveredErr := fmt.Errorf("panic recovered: %v", r)

					zapLogger.Error("Recovered from panic",
						zap.String("method", c.Request().Method),
						zap.String("uri", c.Request().RequestURI),
						zap.Any("panic", r),
						zap.ByteString("debug", debug.Stack()))

					// Return generic 500 error
					_ = c.JSON(http.StatusInternalServerError, echo.Map{
						"message": "Internal Server Error",
					})
					err = recoveredErr
				}
			}()

			// Execute next middleware/handler
			err = next(c)

			// Response logging
			status := c.Response().Status
			uri := c.Request().RequestURI
			method := c.Request().Method
			responseBody, _ := c.Get("responseBody").(string)

			var rawJson json.RawMessage
			_ = json.Unmarshal([]byte(responseBody), &rawJson)

			if err != nil {
				zapLogger.Error("Request error",
					zap.String("method", c.Request().Method),
					zap.String("uri", c.Request().RequestURI),
					zap.Int("status", c.Response().Status),
					//	zap.Any("log error with Any", err),
					//	zap.Error(err),
					zap.String("log error with fmt", fmt.Sprintf("%+v", err)),
				)

			} else {
				zapLogger.Info("Request success",
					zap.String("method", method),
					zap.String("uri", uri),
					zap.Int("status", status),
					zap.Any("response", rawJson),
				)
			}

			return err
		}
	}
}

func SlogLoggerMiddleware(slogLogger *slog.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogHost:     true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {

			responseBody := c.Get("responseBody")
			var responseJSON interface{}

			if v.Error == nil {
				if bodyStr, ok := responseBody.(string); ok {
					if err := json.Unmarshal([]byte(bodyStr), &responseJSON); err == nil {
						slogLogger.LogAttrs(context.Background(), slog.LevelInfo, "request",
							slog.String("uri", v.URI),
							slog.Int("status", v.Status),
							slog.Any("response", responseJSON),
						)
					} else {
						fmt.Println("v.Error content: ", err)
					}
				}
			} else {
				slogLogger.LogAttrs(context.Background(), slog.LevelError, v.Error.Error(),
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			}
			return nil
		},
	})
}
