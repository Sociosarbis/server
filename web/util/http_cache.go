package util

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CacheControl(c *fiber.Ctx, seconds int) {
	c.Set("Cache-Control", fmt.Sprintf("public, max-age=%d", seconds))
}
