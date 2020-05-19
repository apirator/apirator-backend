package apimock

import "github.com/google/wire"

var ApplicationClients = wire.NewSet(
	NewRestConfig,
	NewApiratorClient,
)
