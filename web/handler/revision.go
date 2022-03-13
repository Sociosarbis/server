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
		return errgo.Wrap(err, "revision.CountPersonRelated")
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

	creatorMap, err := h.u.GetByIDs(c.Context(), listUniqueCreatorID(revisions)...)

	if err != nil {
		return errgo.Wrap(err, "user.GetByIDs")
	}
	for i, r := range revisions {
		data[i] = convertModelRevision(r, creatorMap)
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

	creatorMap, err := h.u.GetByIDs(c.Context(), r.CreatorID)

	if err != nil {
		return errgo.Wrap(err, "user.GetByIDS")
	}

	return c.JSON(convertModelRevision(r, creatorMap))

}

func listUniqueCreatorID(revisions []*model.Revision) []domain.IDType {
	m := map[domain.IDType]bool{}
	ret := []domain.IDType{}
	for _, r := range revisions {
		if _, ok := m[r.CreatorID]; !ok {
			m[r.CreatorID] = true
			ret = append(ret, r.CreatorID)
		}
	}
	return ret
}

func convertModelRevision(r *model.Revision, creatorMap map[domain.IDType]model.User) res.Revision {
	creator := creatorMap[r.CreatorID]
	return res.Revision{
		ID:      r.ID,
		Type:    r.Type,
		Summary: r.Summary,
		Creator: res.Creator{
			Username: creator.UserName,
			Nickname: creator.UserName,
		},
		CreatedAt: r.CreatedAt,
		Data:      r.Data,
	}
}
