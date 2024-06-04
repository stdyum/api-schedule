package handlers

import (
	"github.com/stdyum/api-common/grpc"
	"github.com/stdyum/api-common/proto/impl/schedule"
	"github.com/stdyum/api-schedule/internal/app/controllers"
)

type GRPC interface {
	grpc.Routes
	schedule.ScheduleServer
}

type gRPC struct {
	schedule.UnimplementedScheduleServer

	controller controllers.Controller
}

func NewGRPC(controller controllers.Controller) GRPC {
	return &gRPC{
		controller: controller,
	}
}
