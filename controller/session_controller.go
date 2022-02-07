package controller

import (
	"presensi-mobile/database"
	"presensi-mobile/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	authMap := utils.CheckValidToken(c)

	if authMap != nil {
		if authMap["code"] == 500 {
			return c.Status(fiber.StatusInternalServerError).JSON(authMap)
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(authMap)
		}

	}

	db, err := database.OpenDBConnection()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})

	} else {
		email := c.FormValue("email")
		password := c.FormValue("password")
		device := string(c.Request().Header.Peek("device"))

		user, err := db.UserQueries.Login(email, password)

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"code":    404,
				"message": "Login Gagal",
			})
		}
		token, err := utils.CreateJWTToken(user.Email, device)

		db.SessionQueries.CreateSession(user.Id, token, device)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    200,
			"message": "Login success",
		})
	}

}

func Logout(c *fiber.Ctx) error {
	authMap := utils.CheckValidToken(c)

	if authMap != nil {
		if authMap["code"] == 500 {
			return c.Status(fiber.StatusInternalServerError).JSON(authMap)
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(authMap)
		}

	}
	db, err := database.OpenDBConnection()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})

	} else {
		userId := c.FormValue("user_id")

		if err := db.SessionQueries.Logout(userId); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    404,
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    200,
			"message": "Logout success",
		})
	}
}

func RefreshToken(c *fiber.Ctx) error {

	authMap := utils.CheckValidToken(c)

	if authMap != nil {
		if authMap["code"] == 401 {
			return doRefresh(c)
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(authMap)
		}
	}

	return doRefresh(c)

}

func doRefresh(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})

	} else {
		userId := c.FormValue("user_id")
		device := string(c.Request().Header.Peek("device"))

		token, err := utils.CreateJWTToken(userId, device)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    404,
				"message": err.Error(),
			})
		}

		db.SessionQueries.RefreshToken(userId, token)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    200,
			"message": "Refresh token success",
		})
	}
}
