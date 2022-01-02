package controller

import (
	"presensi-mobile/database"
	"presensi-mobile/models"

	"github.com/gofiber/fiber/v2"
)

// USER
// register / setup
// api/register
// upload user csv
// login
// api/login
// profile
// api/profile/{nis}

func InsertUser(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})

	} else {
		user := new(models.User)

		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    404,
				"message": err.Error(),
			})
		}

		if err := user.IsValid(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    404,
				"message": err,
			})
		}
		err := db.InsertUsers(user)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    404,
				"message": err.Error(),
			})
		} else {
			return c.JSON(fiber.Map{
				"code":    200,
				"message": "success",
			})
		}

	}
}
