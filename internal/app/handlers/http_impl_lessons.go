package handlers

import (
	netHttp "net/http"

	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-schedule/internal/app/dto"
)

func (h *http) CreateLesson(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var request dto.CreateLessonsRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	lesson, err := h.controller.CreateLessons(ctx, enrollment, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusCreated, lesson)
}

func (h *http) UpdateLesson(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var request dto.UpdateLessonRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := h.controller.UpdateLesson(ctx, enrollment, request); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}

func (h *http) DeleteLesson(ctx *hc.Context) {
	enrollment := ctx.Enrollment()
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	date, err := ctx.QueryTime("date")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	request := dto.DeleteLessonRequestDTO{
		ID:   id,
		Date: date,
	}

	if err = h.controller.DeleteLessonById(ctx, enrollment, request); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}
