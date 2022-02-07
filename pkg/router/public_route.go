package router

import (
	"os"
	controller "presensi-mobile/controller"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group(os.Getenv("GROUP_PATH"))
	route.Get("/version", controller.GetVersion)

}
