package fiberx

import (
	"github.com/gopkg-dev/pkg/fiberx/middleware"
	"github.com/gopkg-dev/pkg/unique"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func NewServer(appName string) *fiber.App {

	app := fiber.New(fiber.Config{
		AppName:               appName,
		ServerHeader:          "fiberx-server",
		DisableStartupMessage: false,
		ErrorHandler:          DefaultErrorHandler,
	})

	app.Use(csrf.New())
	app.Use(cors.New())
	app.Use(middleware.Timer())

	app.Use(requestid.New(requestid.Config{
		ContextKey: RequestIDKey,
		Generator: func() string {
			return unique.MustUUID().String()
		},
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	app.Use(recover.New(recover.Config{
		Next:              nil,
		EnableStackTrace:  true,
		StackTraceHandler: nil,
	}))

	return app
}
