package servers

import (
	"github.com/apirator/apirator-backend/api/handlers"
	"github.com/labstack/echo/v4"
)

type Register struct {
	MockHandler *handlers.MockHandler
	echo        *echo.Echo
}

func (r *Register) registerAPIs() {
	r.echo.POST("/api/{namespace}/mocks", r.MockHandler.AddMock)
	r.echo.GET("/api/{namespace}/mocks", r.MockHandler.ListMock)
}

func NewRegister(handler *handlers.MockHandler, echo *echo.Echo) *Register {
	return &Register{
		MockHandler: handler,
		echo:        echo,
	}
}

func NewEcho() *echo.Echo {
	return echo.New()
}
