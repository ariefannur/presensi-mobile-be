package router

import "github.com/gofiber/fiber/v2"

func StaticFile(a *fiber.App) {
	a.Static("/", "./dir/photo")
}
