package validators

import (
	"context"

	"github.com/stdyum/api-schedule/internal/app/dto"
)

type Validator interface {
	ValidateCreateLessonEntryRequest(ctx context.Context, request dto.CreateLessonEntryRequestDTO) error
	ValidateCreateLessonsRequest(ctx context.Context, request dto.CreateLessonsRequestDTO) error
	ValidateUpdateLessonRequest(ctx context.Context, request dto.UpdateLessonRequestDTO) error
	ValidateDeleteLessonRequest(ctx context.Context, request dto.DeleteLessonRequestDTO) error

	ValidateCreateLessonGeneralEntryRequest(ctx context.Context, request dto.CreateLessonGeneralEntryRequestDTO) error
	ValidateCreateLessonsGeneralRequest(ctx context.Context, request dto.CreateLessonsGeneralRequestDTO) error
	ValidateUpdateLessonGeneralRequest(ctx context.Context, request dto.UpdateLessonGeneralRequestDTO) error
	ValidateDeleteLessonGeneralRequest(ctx context.Context, request dto.DeleteLessonGeneralRequestDTO) error

	ValidateCreateScheduleMetaEntryRequest(ctx context.Context, request dto.CreateScheduleMetaEntryRequestDTO) error
	ValidateCreateScheduleMetaRequest(ctx context.Context, request dto.CreateScheduleMetaRequestDTO) error
}

type validator struct {
}

func New() Validator {
	return &validator{}
}
