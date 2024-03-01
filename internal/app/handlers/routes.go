package handlers

import (
	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-common/http/middlewares"
	"github.com/stdyum/api-common/proto/impl/studyplaces"
	"google.golang.org/grpc"
)

func (h *http) ConfigureRoutes() *hc.Engine {
	engine := hc.New()
	engine.Use(hc.Recovery())

	v1 := engine.Group("api/v1", hc.Logger(), middlewares.ErrorMiddleware())

	v1.Use(middlewares.EnrollmentAuthMiddleware()).GET("schedule", h.Schedule)

	{
		lessonsGroup := v1.Group("lessons").Use(middlewares.EnrollmentAuthMiddleware())

		lessonsGroup.POST("", h.CreateLesson)
		lessonsGroup.PUT("", h.UpdateLesson)
		lessonsGroup.DELETE(":id", h.DeleteLesson)
	}

	return engine
}

func (h *gRPC) ConfigureRoutes() *grpc.Server {
	grpcServer := grpc.NewServer()
	studyplaces.RegisterStudyplacesServer(grpcServer, h)
	return grpcServer
}
