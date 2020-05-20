package servers

import (
	"github.com/apirator/apirator-backend/api/handlers"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	server   *echo.Echo
	handlers *handlers.MockHandler
}

func (s *Server) Run() {
	e := s.server
	e.Use(middleware.Logger())
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Logger.Fatal(e.Start(":9999"))
}

func NewServer(echo *echo.Echo, handler *handlers.MockHandler) *Server {
	return &Server{
		server:   echo,
		handlers: handler,
	}
}
