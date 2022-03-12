package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetPersionRevision(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (h Handler) ListPersonRevision(c *fiber.Ctx) error {
	_, err := getPageQuery(c, episodeDefaultLimit, defaultMaxPageLimit)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "bad query args: "+err.Error())
	}
	return c.JSON(nil)
}
