package controller

import (
	"fmt"
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
	userId := strconv.FormatInt(presensi.User_Id, 10)

	// check presensi today
	dataPresensi, err := db.CheckPrensensiToday(userId)
	fmt.Println(err)
	if dataPresensi.ID != 0 && err == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    404,
			"message": "User sudah melakukan presensi hari ini",
		})
	} else {

		fileTmp, errFile := utils.MoveTmpFile(file, utils.Photo_Path, utils.GetFormatPhotoName(userId))

		if errFile != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    404,
				"message": err.Error(),
			})
		}
		presensi.Foto = fileTmp.Name()[9:len(fileTmp.Name())]

		if err := db.CreatePresensi(&presensi); err != nil {
			return c.Status(500).SendString("Error: " + err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    200,
			"message": "Success",
		})
	}

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

	userId := c.Params("user_id")

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	presensi, err := db.GetPresensi(userId)
	if err != nil {
		return c.Status(500).SendString("Error: " + err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": 200,
		"data": presensi,
	})

}
