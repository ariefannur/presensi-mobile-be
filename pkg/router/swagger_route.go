package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	_ "presensi-mobile/docs"
)

func SwaggerRouter(a *fiber.App) {
	router := a.Group("swagger")

	router.Get("*", swagger.Handler)
}
