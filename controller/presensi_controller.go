package controller

import (
	"presensi-mobile/database"
	"presensi-mobile/models"
	"presensi-mobile/pkg/utils"
	"strconv"

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
	file, err := c.FormFile("img_presensi")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    404,
			"message": err.Error(),
		})
	}

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

	userId := strconv.FormatInt(presensi.UserID, 10)
	fileTmp, errFile := utils.MoveTmpFile(file, utils.Photo_Path, utils.GetFormatPhotoName(userId))

	if errFile != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    404,
			"message": err.Error(),
		})
	}
	presensi.Foto = fileTmp.Name()

	if err := db.CreatePresensi(&presensi); err != nil {
		return c.Status(500).SendString("Error: " + err.Error())
	}

	return c.Status(200).SendString("Success")

}

func GetPresensiByUserId(c *fiber.Ctx) error {
	authMap := utils.CheckValidToken(c)

	if authMap != nil {
		if authMap["code"] == 500 {
			return c.Status(fiber.StatusInternalServerError).JSON(authMap)
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(authMap)
		}
	}

	var presensi models.Presensi
	userId := c.Params("user_id")

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if _, err := db.GetPresensi(userId); err != nil {
		return c.Status(500).SendString("Error: " + err.Error())
	}

	return c.JSON(presensi)

}
