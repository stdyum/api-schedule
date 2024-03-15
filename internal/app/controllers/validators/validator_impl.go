package validators

import (
	"context"
	"time"

	"github.com/stdyum/api-common/uvalidator"
	"github.com/stdyum/api-schedule/internal/app/dto"
)

func (v *validator) ValidateCreateLessonEntryRequest(_ context.Context, request dto.CreateLessonEntryRequestDTO) error {
	return uvalidator.V(
		uvalidator.UUIDNotNil("groupId", request.GroupId),
		uvalidator.UUIDNotNil("roomId", request.RoomId),
		uvalidator.UUIDNotNil("subjectId", request.SubjectId),
		uvalidator.UUIDNotNil("teacherId", request.TeacherId),
		uvalidator.RangeInt("lessonIndex", request.LessonIndex, 0, 100),
		uvalidator.TimeAfter("startTime", request.StartTime, time.Now()),
		uvalidator.TimeAfter("endTime", request.EndTime, request.StartTime),
	)
}

func (v *validator) ValidateCreateLessonsRequest(ctx context.Context, request dto.CreateLessonsRequestDTO) error {
	for _, r := range request.List {
		if err := v.ValidateCreateLessonEntryRequest(ctx, r); err != nil {
			return err
		}
	}

	return uvalidator.V(
		uvalidator.SliceNotEmpty("list", request.List),
	)
}

func (v *validator) ValidateUpdateLessonRequest(ctx context.Context, request dto.UpdateLessonRequestDTO) error {
	if err := v.ValidateCreateLessonEntryRequest(ctx, request.CreateLessonEntryRequestDTO); err != nil {
		return err
	}

	return uvalidator.V(
		uvalidator.UUIDNotNil("id", request.ID),
	)
}

func (v *validator) ValidateDeleteLessonRequest(_ context.Context, request dto.DeleteLessonRequestDTO) error {
	return uvalidator.V(
		uvalidator.UUIDNotNil("id", request.ID),
	)
}

func (v *validator) ValidateCreateLessonGeneralEntryRequest(ctx context.Context, request dto.CreateLessonGeneralEntryRequestDTO) error {
	return uvalidator.V(
		uvalidator.UUIDNotNil("groupId", request.GroupId),
		uvalidator.UUIDNotNil("roomId", request.RoomId),
		uvalidator.UUIDNotNil("subjectId", request.SubjectId),
		uvalidator.UUIDNotNil("teacherId", request.TeacherId),
		uvalidator.RangeInt("lessonIndex", request.LessonIndex, 0, 100),
		uvalidator.DurationBeforeOrEqual("endTime", request.EndTime, time.Hour*24),
		uvalidator.DurationAfter("endTime", request.EndTime, request.StartTime),
	)
}

func (v *validator) ValidateCreateLessonsGeneralRequest(ctx context.Context, request dto.CreateLessonsGeneralRequestDTO) error {
	for _, r := range request.List {
		if err := v.ValidateCreateLessonGeneralEntryRequest(ctx, r); err != nil {
			return err
		}
	}

	return uvalidator.V(
		uvalidator.IsNotNil("list", request),
	)
}

func (v *validator) ValidateUpdateLessonGeneralRequest(ctx context.Context, request dto.UpdateLessonGeneralRequestDTO) error {
	if err := v.ValidateCreateLessonGeneralEntryRequest(ctx, request.CreateLessonGeneralEntryRequestDTO); err != nil {
		return err
	}

	return uvalidator.V(
		uvalidator.UUIDNotNil("id", request.ID),
	)
}

func (v *validator) ValidateDeleteLessonGeneralRequest(_ context.Context, request dto.DeleteLessonGeneralRequestDTO) error {
	return uvalidator.V(
		uvalidator.UUIDNotNil("id", request.ID),
		uvalidator.RangeInt("dayIndex", request.DayIndex, 0, 7),
	)
}

func (v *validator) ValidateCreateScheduleMetaEntryRequest(_ context.Context, request dto.CreateScheduleMetaEntryRequestDTO) error {
	return uvalidator.V(
		uvalidator.DateAfterOrEqual("date", request.Date, time.Now()),
		uvalidator.StringIn("status", request.Status, "updated", "not_updated", "in_progress"),
	)
}

func (v *validator) ValidateCreateScheduleMetaRequest(ctx context.Context, request dto.CreateScheduleMetaRequestDTO) error {
	for _, r := range request.List {
		if err := v.ValidateCreateScheduleMetaEntryRequest(ctx, r); err != nil {
			return err
		}
	}

	return uvalidator.V(
		uvalidator.IsNotNil("list", request),
	)
}
