package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"log/slog"
)

func ZapLoggerMiddleware(zapLogger *zap.Logger) echo.MiddlewareFunc {

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {

			responseBody := c.Get("responseBody")
			var responseJSON interface{}

			if v.Error == nil {
				if bodyStr, ok := responseBody.(string); ok {
					if err := json.Unmarshal([]byte(bodyStr), &responseJSON); err == nil {
						zapLogger.Info("request",
							zap.String("URI", v.URI),
							zap.Int("status", v.Status),
							zap.Any("response", responseJSON),
						)
					}
				}
			} else {
				fmt.Println("v.Error content: ", v.Error)

				zapLogger.Error(v.Error.Error(),
					zap.String("URI", v.URI),
					zap.Int("status", v.Status))
			}
			return nil
		},
	})
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
