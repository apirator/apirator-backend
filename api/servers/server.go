package servers

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	server   *echo.Echo
	register *RegisterAPI
}

func (s *Server) Run() {
	e := s.server
	s.register.registerAPIs()
	e.Use(middleware.Logger())
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Logger.Fatal(e.Start(":9999"))
}

func NewServer(echo *echo.Echo, register *RegisterAPI) *Server {
	return &Server{
		server:   echo,
		register: register,
	}
}
