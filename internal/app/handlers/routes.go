package handlers

import (
	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-common/http/middlewares"
	"google.golang.org/grpc"
)

func (h *http) ConfigureRoutes() *hc.Engine {
	engine := hc.New()
	engine.Use(hc.Recovery())

	v1 := engine.Group("api/v1", hc.Logger(), middlewares.ErrorMiddleware())

	v1.Use(middlewares.EnrollmentAuthMiddleware()).GET("schedule", h.Schedule)
	v1.Use(middlewares.EnrollmentAuthMiddleware()).GET("schedule/general", h.ScheduleGeneral)

	{
		lessonsGroup := v1.Group("lessons").Use(middlewares.EnrollmentAuthMiddleware())

		lessonsGroup.POST("", h.CreateLesson)
		lessonsGroup.PUT("", h.UpdateLesson)
		lessonsGroup.DELETE(":id", h.DeleteLesson)
	}

	{
		lessonsGeneralGroup := v1.Group("lessons/general").Use(middlewares.EnrollmentAuthMiddleware())

		lessonsGeneralGroup.POST("", h.CreateLessonGeneral)
		lessonsGeneralGroup.PUT("", h.UpdateLessonGeneral)
		lessonsGeneralGroup.DELETE(":id", h.DeleteLessonGeneral)
	}

	return engine
}

func (h *gRPC) ConfigureRoutes() *grpc.Server {
	return nil
}
