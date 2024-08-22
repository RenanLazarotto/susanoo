package main

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"tsukuyomi/config"
	"tsukuyomi/routers"
)

func main() {
	config.SetupLog()
	log.Info("Starting...")
	config := config.Load()
	log.Debug("Config loaded")

	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(recover.New())
	app.Use(cors.New())

	log.Debug("Created fiber app", "middlewares", []string{"recover", "cors"})

	routers.SetupRouter(app, config)

	log.Debug("Configured routes")

	log.Infof("%s listening on port %d", config.App.Name, config.App.Port)
	log.Fatalf("can't start application: %v", app.Listen(fmt.Sprintf(":%d", config.App.Port)))
}
