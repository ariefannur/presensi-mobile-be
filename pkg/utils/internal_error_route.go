package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetInternalError(c *fiber.Ctx, err error) error {
	fmt.Printf("ERROR %+v\n", err)
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	} else {
		return nil
	}

}
