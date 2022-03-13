package handler

import (
	"net/http"

	"github.com/bangumi/server/internal/errgo"
	"github.com/bangumi/server/internal/strparse"

	"github.com/bangumi/server/domain"
	"github.com/bangumi/server/model"
	"github.com/bangumi/server/web/res"
	"github.com/bangumi/server/web/util"
	"github.com/gofiber/fiber/v2"
)

func (h Handler) ListPersonRevision(c *fiber.Ctx) error {
	page, err := getPageQuery(c, episodeDefaultLimit, defaultMaxPageLimit)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "bad query args: "+err.Error())
	}
	personID, err := strparse.Uint32(c.Query("person_id"))
	if err != nil || personID <= 0 {
		return fiber.NewError(http.StatusBadRequest, "bad person_id: "+c.Query("person_id"))
	}

	return h.listPersonRevision(c, uint32(personID), page)
}

func (h Handler) listPersonRevision(c *fiber.Ctx, personID domain.PersonIDType, page pageQuery) error {
	var response = res.Paged{
		Limit:  page.Limit,
		Offset: page.Offset,
	}
	count, err := h.r.CountPersonRelated(c.Context(), personID)
	if err != nil {
		return errgo.Wrap(err, "episode.Count")
	}

	if count == 0 {
		response.Data = []int{}
		return c.JSON(response)
	}

	if int64(page.Offset) >= count {
		return fiber.NewError(http.StatusBadRequest, "offset if greater than count")
	}

	response.Total = count

	revisions, err := h.r.ListPersonRelated(c.Context(), personID, page.Limit, page.Offset)

	if err != nil {
		return errgo.Wrap(err, "revision.ListPersonRelated")
	}

	data := make([]res.Revision, len(revisions))

	for i, r := range revisions {
		data[i] = convertModelRevision(&r)
	}
	response.Data = data
	return c.JSON(response)
}

func (h Handler) GetPersionRevision(c *fiber.Ctx) error {
	id, err := strparse.Uint32(c.Params("id"))
	if err != nil || id <= 0 {
		return fiber.NewError(http.StatusBadRequest, "bad id: "+c.Params("id"))
	}
	r, err := h.r.GetPersonRelated(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(res.Error{
			Title:   "Not Found",
			Details: util.DetailFromRequest(c),
		})
	}
	return c.JSON(r)

}

func convertModelRevision(r *model.Revision) res.Revision {
	return res.Revision{
		ID:      r.ID,
		Type:    r.Type,
		Summary: r.Summary,
		Creator: res.Creator{
			Username: r.Creator.Username,
			Nickname: r.Creator.Nickname,
		},
		CreatedAt: r.CreatedAt,
		Data:      r.Data,
	}
}
