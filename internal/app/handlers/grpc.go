package handlers

import (
	"github.com/stdyum/api-common/grpc"
	"github.com/stdyum/api-schedule/internal/app/controllers"
)

type GRPC interface {
	grpc.Routes
}

type gRPC struct {
	controller controllers.Controller
}

func NewGRPC(controller controllers.Controller) GRPC {
	return &gRPC{
		controller: controller,
	}
}
