package internal

import (
	"github.com/stdyum/api-common/grpc/clients"
	"github.com/stdyum/api-common/server"
	"github.com/stdyum/api-schedule/internal/app"
	"github.com/stdyum/api-schedule/internal/app/controllers"
	"github.com/stdyum/api-schedule/internal/config"
	"github.com/stdyum/api-schedule/internal/modules/types_registry"
)

func Configure() (server.Routes, controllers.Controller, error) {
	db, err := config.ConnectToDatabase(config.Config.Database)
	if err != nil {
		return server.Routes{}, nil, err
	}

	studyPlacesServer, err := config.ConnectToStudyPlacesServer(config.Config.StudyPlacesGRpc)
	if err != nil {
		return server.Routes{}, nil, err
	}
	clients.StudyPlacesGRpcClient = studyPlacesServer

	typesRegistryClient, err := config.ConnectToTypesRegistryServer(config.Config.TypesRegistryGRpc)
	if err != nil {
		return server.Routes{}, nil, err
	}

	typesRegistry := types_registry.New(typesRegistryClient)

	routes, ctrl, err := app.New(db, typesRegistry)
	if err != nil {
		return server.Routes{}, nil, err
	}

	routes.Ports = config.Config.Ports
	return routes, ctrl, nil
}
