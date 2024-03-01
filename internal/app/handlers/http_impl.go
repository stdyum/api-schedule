package handlers

import (
	netHttp "net/http"

	"github.com/stdyum/api-common/hc"
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

	schedule, err := h.controller.Schedule(ctx, enrollment, column, columnId, from, to)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, schedule)
}
