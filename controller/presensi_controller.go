package controller

import (
	"presensi-mobile/database"
	"presensi-mobile/models"
	"presensi-mobile/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// MAIN
// presensi
// api/presence
// filter date
// api/my-presence

func CreatePresensi(c *fiber.Ctx) error {
	authMap := utils.CheckValidToken(c)

	if authMap != nil {
		if authMap["code"] == 500 {
			return c.Status(fiber.StatusInternalServerError).JSON(authMap)
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(authMap)
		}

	}

	var presensi models.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(500).SendString("Error: " + err.Error())
	}
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err := db.CreatePresensi(&presensi); err != nil {
		return c.Status(500).SendString("Error: " + err.Error())
	}

	return c.Status(200).SendString("Success")

}
