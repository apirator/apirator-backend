package handlers

import "github.com/labstack/echo/v4"

type HealthCheckHandler struct {
}

func (h *HealthCheckHandler) Health(c echo.Context) error {
	return c.JSON(201, &status{
		Status: "UP",
	})
}

type status struct {
	Status string `json:"status"`
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}
