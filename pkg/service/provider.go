package service

import "github.com/google/wire"

var ApplicationServices = wire.NewSet(
	NewMockService,
)
