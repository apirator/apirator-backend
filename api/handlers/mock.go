package handlers

import (
	"encoding/json"
	"github.com/apirator/apirator-backend/internal/errors"
	ms "github.com/apirator/apirator-backend/pkg/mock"
	"github.com/apirator/apirator-backend/pkg/service"
	"github.com/labstack/echo/v4"
)

type MockHandler struct {
	service *service.MockService
}

func (mh *MockHandler) AddMock(c echo.Context) error {
	defer c.Request().Body.Close()
	mock := &ms.Mock{}
	if err := json.NewDecoder(c.Request().Body).Decode(mock); err != nil {
		return errors.Wrap(err)
	}
	nm, err := mh.service.Add(c.Param("namespace"), mock)
	if err != nil {
		return errors.Wrap(err)
	}
	return c.JSON(201, nm)

}

func (mh *MockHandler) ListMock(c echo.Context) error {
	data, err := mh.service.List(c.Param("namespace"))
	if err != nil {
		return c.JSON(500, nil)
	}
	return c.JSON(200, data)
}

func NewMockHandler(service *service.MockService) *MockHandler {
	return &MockHandler{
		service: service,
	}
}
