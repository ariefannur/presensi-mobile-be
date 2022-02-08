package controller

import (
	"fmt"
	"presensi-mobile/database"
	"presensi-mobile/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// GetVersion godoc
// @Summary      Login to api
// @Description  to login and create token
// @Success      200  {object}  models.Session
// @Param        email     body    string true        "email"
// @Param        password  body    string true        "password"
// @Router       /login [post]
func Login(c *fiber.Ctx) error {

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
		token, _ := utils.CreateJWTToken(user.Email, device)

		db.SessionQueries.CreateSession(user.Id, token, device)
		fmt.Println(user.Id)
		session, _ := db.GetSession(user.Id)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    200,
			"message": "Login success",
			"data":    session,
		})
	}

}

// GetVersion godoc
// @Summary      Logout clear session
// @Description  to logout clear session
// @Success      200
// @Param        user_id     body    string true        "user_id"
// @Router       /logout [post]
func Logout(c *fiber.Ctx) error {
	authMap := utils.CheckValidToken(c)

	if authMap != nil {
		if authMap["code"] == 500 {
			return c.Status(fiber.StatusInternalServerError).JSON(authMap)
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(authMap)
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
