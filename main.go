package main

import (
	"presensi-mobile/pkg/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	app := fiber.New()
	router.SwaggerRouter(app)
	router.PublicRoutes(app)
	router.NotFoundRoute(app)

	app.Listen(":9090")
}
