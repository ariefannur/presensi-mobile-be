package main

import (
	"fmt"
	"presensi-mobile/pkg/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	app := fiber.New()
	router.SwaggerRouter(app)
	router.PublicRoutes(app)
	router.PrivateRoutes(app)
	router.NotFoundRoute(app)
	fmt.Println("START")
	app.Listen(":9090")
}
