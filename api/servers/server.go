package servers

import (
	infra "github.com/apirator/apirator-backend/internal/logger"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	server   *echo.Echo
	register *RegisterAPI
}

func (s *Server) Run() {
	infra.Logger.Info("Starting apirator backend...")
	s.register.registerAPIs()
	infra.Logger.Info("Configuring middlewares...")
	s.server.Use(middleware.Logger())
	p := prometheus.NewPrometheus("apirator_backend", nil)
	p.Use(s.server)
	s.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	infra.Logger.Info("Middlewares configured successfully.")
	infra.Logger.Info("apirator backend started")
	s.server.Logger.Fatal(s.server.Start(":9999"))
}

func NewServer(echo *echo.Echo, register *RegisterAPI) *Server {
	return &Server{
		server:   echo,
		register: register,
	}
}
