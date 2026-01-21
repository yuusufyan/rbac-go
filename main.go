package main

import (
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"

	"rbac-go/internal/infrastructure/configs"
	"rbac-go/internal/infrastructure/utils"
)

func main() {
	log := utils.NewLogger()

	configs.LoadEnv()
	configs.LoadConfig()
	app := fiber.New()

	if configs.AppConfig == nil {
		panic("CONFIG IS NIL — LoadConfig() FAILED")
	}

	app.Use(fiberLogger.New(fiberLogger.Config{
		Format:     "[${time}] ${ip} - ${status} ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	log.Info("Server starting...")

	app.Listen(":" + configs.AppConfig.AppPort)
}
