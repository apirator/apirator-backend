package handlers

import "github.com/google/wire"

var ApplicationHandlers = wire.NewSet(
	NewMockHandler,
)
