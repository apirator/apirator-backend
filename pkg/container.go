package pkg

import (
	"github.com/apirator/apirator-backend/api/handlers"
	"github.com/apirator/apirator-backend/api/servers"
	"github.com/apirator/apirator-backend/internal/apimock"
	"github.com/apirator/apirator-backend/pkg/service"
	"github.com/google/wire"
)

var Container = wire.NewSet(
	handlers.ApplicationHandlers,
	service.ApplicationServices,
	apimock.ApplicationClients,
	servers.ApplicationServers,
)
