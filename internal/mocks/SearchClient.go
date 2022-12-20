// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	accessor "github.com/bangumi/server/internal/web/accessor"

	fiber "github.com/gofiber/fiber/v2"

	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"
)

// SearchClient is an autogenerated mock type for the Client type
type SearchClient struct {
	mock.Mock
}

type SearchClient_Expecter struct {
	mock *mock.Mock
}

func (_m *SearchClient) EXPECT() *SearchClient_Expecter {
	return &SearchClient_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: ctx, auth
func (_m *SearchClient) Handle(ctx *fiber.Ctx, auth *accessor.Accessor) error {
	ret := _m.Called(ctx, auth)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx, *accessor.Accessor) error); ok {
		r0 = rf(ctx, auth)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchClient_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type SearchClient_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - ctx *fiber.Ctx
//   - auth *accessor.Accessor
func (_e *SearchClient_Expecter) Handle(ctx interface{}, auth interface{}) *SearchClient_Handle_Call {
	return &SearchClient_Handle_Call{Call: _e.mock.On("Handle", ctx, auth)}
}

func (_c *SearchClient_Handle_Call) Run(run func(ctx *fiber.Ctx, auth *accessor.Accessor)) *SearchClient_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*fiber.Ctx), args[1].(*accessor.Accessor))
	})
	return _c
}

func (_c *SearchClient_Handle_Call) Return(_a0 error) *SearchClient_Handle_Call {
	_c.Call.Return(_a0)
	return _c
}

// OnSubjectDelete provides a mock function with given fields: ctx, id
func (_m *SearchClient) OnSubjectDelete(ctx context.Context, id model.SubjectID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.SubjectID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchClient_OnSubjectDelete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnSubjectDelete'
type SearchClient_OnSubjectDelete_Call struct {
	*mock.Call
}

// OnSubjectDelete is a helper method to define mock.On call
//   - ctx context.Context
//   - id model.SubjectID
func (_e *SearchClient_Expecter) OnSubjectDelete(ctx interface{}, id interface{}) *SearchClient_OnSubjectDelete_Call {
	return &SearchClient_OnSubjectDelete_Call{Call: _e.mock.On("OnSubjectDelete", ctx, id)}
}

func (_c *SearchClient_OnSubjectDelete_Call) Run(run func(ctx context.Context, id model.SubjectID)) *SearchClient_OnSubjectDelete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SubjectID))
	})
	return _c
}

func (_c *SearchClient_OnSubjectDelete_Call) Return(_a0 error) *SearchClient_OnSubjectDelete_Call {
	_c.Call.Return(_a0)
	return _c
}

// OnSubjectUpdate provides a mock function with given fields: ctx, id
func (_m *SearchClient) OnSubjectUpdate(ctx context.Context, id model.SubjectID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.SubjectID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchClient_OnSubjectUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnSubjectUpdate'
type SearchClient_OnSubjectUpdate_Call struct {
	*mock.Call
}

// OnSubjectUpdate is a helper method to define mock.On call
//   - ctx context.Context
//   - id model.SubjectID
func (_e *SearchClient_Expecter) OnSubjectUpdate(ctx interface{}, id interface{}) *SearchClient_OnSubjectUpdate_Call {
	return &SearchClient_OnSubjectUpdate_Call{Call: _e.mock.On("OnSubjectUpdate", ctx, id)}
}

func (_c *SearchClient_OnSubjectUpdate_Call) Run(run func(ctx context.Context, id model.SubjectID)) *SearchClient_OnSubjectUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SubjectID))
	})
	return _c
}

func (_c *SearchClient_OnSubjectUpdate_Call) Return(_a0 error) *SearchClient_OnSubjectUpdate_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewSearchClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewSearchClient creates a new instance of SearchClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSearchClient(t mockConstructorTestingTNewSearchClient) *SearchClient {
	mock := &SearchClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
