package handlers

import (
	netHttp "net/http"

	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-schedule/internal/app/dto"
)

func (h *http) CreateLessonGeneral(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var request dto.CreateLessonsGeneralRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	lesson, err := h.controller.CreateLessonsGeneral(ctx, enrollment, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusCreated, lesson)
}

func (h *http) UpdateLessonGeneral(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var request dto.UpdateLessonGeneralRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := h.controller.UpdateLessonGeneral(ctx, enrollment, request); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}

func (h *http) DeleteLessonGeneral(ctx *hc.Context) {
	enrollment := ctx.Enrollment()
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	dayIndex, err := ctx.QueryInt("dayIndex")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	if err = h.controller.DeleteLessonGeneralById(ctx, enrollment, dayIndex, id); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}
