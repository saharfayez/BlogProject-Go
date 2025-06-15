package server

import (
	"go.uber.org/zap"
	appMiddleware "goproject/internal/app/middleware"
)

func Serve() {
	e := registerRoutes()
	appMiddleware.InitLogger(zap.DebugLevel)
	e.Logger.Fatal(e.Start(":8080"))
}
