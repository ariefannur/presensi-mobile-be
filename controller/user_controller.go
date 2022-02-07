package controller

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"presensi-mobile/database"
	"presensi-mobile/models"
	"presensi-mobile/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func InsertUser(c *fiber.Ctx) error {
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
		user := new(models.Users)

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
		result := inserUserByModel(db, user)
		if result != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    404,
				"message": err,
			})
		}

		return c.JSON(fiber.Map{
			"code":    200,
			"message": "success",
		})
	}

}

func inserUserByModel(db *database.Queries, user *models.Users) error {
	err := db.InsertUsers(user)

	if err != nil {
		return err
	}

	return nil
}

func InserCSVFileUsers(c *fiber.Ctx) error {
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
		file, err := c.FormFile("csv_file")

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		} else {
			file, err := moveTmpFile(file)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   true,
					"message": err.Error(),
				})
			}
			listUsers := openCSVFile(file)
			for _, user := range listUsers {
				inserUserByModel(db, user)
			}
			os.Remove(file.Name())
			return c.JSON(fiber.Map{
				"code":    200,
				"message": "success",
			})
		}
	}
}

func moveTmpFile(file *multipart.FileHeader) (*os.File, error) {
	tempFile, err := ioutil.TempFile("tmp", "file_*.csv")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer tempFile.Close()
	// write this byte array to our temporary file

	dataFile, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	byteContainer, err := ioutil.ReadAll(dataFile)
	tempFile.Write(byteContainer)
	return tempFile, nil
}

func openCSVFile(file *os.File) []*models.Users {
	csvFile, err := os.Open(file.Name())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		message := fmt.Sprintf("gagal %s", err)
		fmt.Println(message)
	}
	list := make([]*models.Users, 0)
	for index, line := range csvLines {
		if index > 0 {
			user := models.Users{
				Id:        line[0],
				Name:      line[1],
				Password:  "123",
				Email:     line[2],
				User_Type: line[3],
			}
			list = append(list, &user)
		}
	}
	file.Close()
	return list
}

func GetUsers(c *fiber.Ctx) error {
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
		var users interface{}
		var err error
		id := c.Params("id")

		if id != "" {
			fmt.Println(id)
			user, error := db.GetUsersById(id)
			users = &user
			err = error
		} else {
			data, _error := db.GetUsers()
			err = _error
			users = &data
		}

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"code":    200,
			"message": "success",
			"data":    users,
		})
	}
}

func ChangePassword(c *fiber.Ctx) error {
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
		user := new(models.Users)

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
		_, err := db.ChangePassword(user.Password, user.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    404,
				"message": err,
			})
		}

		return c.JSON(fiber.Map{
			"code":    200,
			"message": "success",
		})
	}

}
