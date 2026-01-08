package main

import (
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"

	"rbac-go/internal/infrastructure/utils"
)

func main() {
	log := utils.NewLogger()

	app := fiber.New()

	app.Use(fiberLogger.New(fiberLogger.Config{
		Format:     "[${time}] ${ip} - ${status} ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	log.Info("Server starting...")

	app.Listen(":3000")
}
