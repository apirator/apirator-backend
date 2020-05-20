package servers

import "github.com/google/wire"

var ApplicationServers = wire.NewSet(
	NewRegister,
	NewEcho,
	NewServer,
)
