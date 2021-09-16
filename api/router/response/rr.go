package response

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// MakeSuccessJSON return success response
func MakeSuccessJSON(c *fiber.Ctx, data interface{}) error {
	return c.JSON(data)
}

func MakeSuccessString(c *fiber.Ctx, data interface{}) error {
	return c.SendString(data.(string))
}

// MakeFail return fail response
func MakeFail(c *fiber.Ctx, message interface{}) error {
	return c.Status(http.StatusBadRequest).SendString(message.(string))
}
