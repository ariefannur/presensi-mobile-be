package controller

import (
	model "presensi-mobile/models"

	"github.com/gofiber/fiber/v2"
)

// GetVersion godoc
// @Summary      get app version
// @Description  get version app of api
// @Success      200  {object}  model.Version
// @Router       /version [get]
func GetVersion(c *fiber.Ctx) error {

	version := model.Version{
		Name: "1.0.0.alpha-1",
		Code: 1,
	}
	return c.JSON(version)

}
