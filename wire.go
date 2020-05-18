//+build wireinject

package main

import (
	"github.com/apirator/apirator-backend/pkg"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func Application() (e *echo.Echo, err error) {
	wire.Build(pkg.Container)
	return nil, nil
}
