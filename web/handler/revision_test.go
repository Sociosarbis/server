package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/bangumi/server/domain"
	"github.com/bangumi/server/internal/test"
	"github.com/bangumi/server/model"
)

func TestHandler_ListPersionRevision_HappyPath(t *testing.T) {
	t.Parallel()
	m := &domain.MockRevisionRepo{}
	m.EXPECT().ListPersonRelated(mock.Anything, uint32(9), 30, 0).Return([]*model.Revision{{ID: 348475}}, nil)
	m.EXPECT().CountPersonRelated(mock.Anything, uint32(9)).Return(30, nil)

	app := test.GetWebApp(t, test.Mock{RevisionRepo: m})

	req := httptest.NewRequest(http.MethodGet, "/v0/revisions/persons?person_id=9", http.NoBody)

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHandler_GetPersionRevision_HappyPath(t *testing.T) {
	t.Parallel()
	m := &domain.MockRevisionRepo{}
	m.EXPECT().GetPersonRelated(mock.Anything, uint32(348475)).Return(&model.Revision{ID: 348475}, nil)

	app := test.GetWebApp(t, test.Mock{RevisionRepo: m})

	req := httptest.NewRequest(http.MethodGet, "/v0/revisions/persons/348475", http.NoBody)

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)
}
