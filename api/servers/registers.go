package servers

import (
	"github.com/apirator/apirator-backend/api/handlers"
	infra "github.com/apirator/apirator-backend/internal/logger"
	"github.com/labstack/echo/v4"
)

type RegisterAPI struct {
	mockHandler *handlers.MockHandler
	echo        *echo.Echo
	healthCheck *handlers.HealthCheckHandler
}

func (r *RegisterAPI) registerAPIs() {
	infra.Logger.Info("Starting registering APIs...")
	r.echo.POST("/api/:namespace/mocks", r.mockHandler.AddMock)
	r.echo.GET("/api/:namespace/mocks", r.mockHandler.ListMock)
	r.echo.GET("/health", r.healthCheck.Health)
	infra.Logger.Info("APIs registered successfully.")
}

func NewRegister(handler *handlers.MockHandler, echo *echo.Echo, health *handlers.HealthCheckHandler) *RegisterAPI {
	return &RegisterAPI{
		mockHandler: handler,
		echo:        echo,
		healthCheck: health,
	}
}

func NewEcho() *echo.Echo {
	return echo.New()
}
