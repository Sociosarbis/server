package util

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func BadParam(c *fiber.Ctx, name string) error {
	return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("bad param %s: %s", name, c.Params(name)))
}

func BadQuery(c *fiber.Ctx, name string) error {
	return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("bad query %s: %s", name, c.Query(name)))
}
