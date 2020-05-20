//+build wireinject

package main

import (
	"github.com/apirator/apirator-backend/api/servers"
	"github.com/apirator/apirator-backend/pkg"
	"github.com/google/wire"
)

func Application() (sever *servers.Server, err error) {
	wire.Build(pkg.Container)
	return nil, nil
}
