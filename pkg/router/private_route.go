package router

import (
	"os"
	controller "presensi-mobile/controller"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group(os.Getenv("GROUP_PATH"))
	route.Post("/add-users", controller.InsertUser)
	route.Post("/change-password", controller.ChangePassword)
	route.Get("/users", controller.GetUsers)
	route.Get("/users/:id?", controller.GetUsers)
	route.Post("/add-users/group", controller.InserCSVFileUsers)

	route.Post("/logout", controller.Logout)
	route.Post("/refresh-token", controller.RefreshToken)

	route.Post("/add-presensi", controller.CreatePresensi)
	route.Get("/presensi/:user_id?", controller.GetPresensiByUserId)
}
