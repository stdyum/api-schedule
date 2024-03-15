package handlers

import (
	netHttp "net/http"

	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-schedule/internal/app/dto"
)

func (h *http) Schedule(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	column := ctx.Query("column")
	columnId, err := ctx.QueryUUID("columnId")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	from, err := ctx.QueryTime("from")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	to, err := ctx.QueryTime("to")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	request := dto.GetScheduleRequestDTO{
		Column:   column,
		ColumnId: columnId,
		From:     from,
		To:       to,
	}

	schedule, err := h.controller.Schedule(ctx, enrollment, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, schedule)
}

func (h *http) ScheduleGeneral(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	column := ctx.Query("column")
	columnId, err := ctx.QueryUUID("columnId")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	request := dto.GetScheduleGeneralRequestDTO{
		Column:   column,
		ColumnId: columnId,
	}

	schedule, err := h.controller.ScheduleGeneral(ctx, enrollment, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, schedule)
}
