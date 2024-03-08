package app

import (
	"github.com/gocql/gocql"
	"github.com/stdyum/api-common/server"
	"github.com/stdyum/api-schedule/internal/app/controllers"
	"github.com/stdyum/api-schedule/internal/app/errors"
	"github.com/stdyum/api-schedule/internal/app/handlers"
	"github.com/stdyum/api-schedule/internal/app/repositories"
	"github.com/stdyum/api-schedule/internal/modules/types_registry"
)

func New(database *gocql.Session, registry types_registry.Controller) (server.Routes, controllers.Controller, error) {
	repo := repositories.New(database)

	ctrl := controllers.New(repo, registry)

	errors.Register()

	httpHndl := handlers.NewHTTP(ctrl)
	grpcHndl := handlers.NewGRPC(ctrl)

	routes := server.Routes{
		GRPC: grpcHndl,
		HTTP: httpHndl,
	}

	return routes, ctrl, nil
}
