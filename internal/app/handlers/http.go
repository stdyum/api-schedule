package handlers

import (
	"github.com/stdyum/api-common/hc"
	confHttp "github.com/stdyum/api-common/http"
	"github.com/stdyum/api-schedule/internal/app/controllers"
)

type HTTP interface {
	confHttp.Routes

	Schedule(ctx *hc.Context)
	ScheduleGeneral(ctx *hc.Context)

	CreateLesson(ctx *hc.Context)
	UpdateLesson(ctx *hc.Context)
	DeleteLesson(ctx *hc.Context)

	CreateLessonGeneral(ctx *hc.Context)
	UpdateLessonGeneral(ctx *hc.Context)
	DeleteLessonGeneral(ctx *hc.Context)
}

type http struct {
	controller controllers.Controller
}

func NewHTTP(controller controllers.Controller) HTTP {
	return &http{
		controller: controller,
	}
}
