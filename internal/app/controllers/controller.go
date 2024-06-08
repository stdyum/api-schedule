package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-schedule/internal/app/controllers/validators"
	"github.com/stdyum/api-schedule/internal/app/dto"
	"github.com/stdyum/api-schedule/internal/app/repositories"
	"github.com/stdyum/api-schedule/internal/modules/types_registry"
)

type Controller interface {
	Schedule(ctx context.Context, enrollment models.Enrollment, request dto.GetScheduleRequestDTO) (dto.ScheduleResponseDTO, error)
	ScheduleGeneral(ctx context.Context, enrollment models.Enrollment, request dto.GetScheduleGeneralRequestDTO) (dto.ScheduleGeneralResponseDTO, error)
	CreateScheduleMeta(ctx context.Context, enrollment models.Enrollment, request dto.CreateScheduleMetaRequestDTO) (dto.CreateScheduleMetaResponseDTO, error)

	GetLessonById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.LessonResponseDTO, error)
	GetLessons(ctx context.Context, enrollment models.Enrollment, filter dto.EntriesFilterRequestDTO) ([]dto.LessonResponseDTO, error)
	CreateLessons(ctx context.Context, enrollment models.Enrollment, request dto.CreateLessonsRequestDTO) (dto.CreateLessonsResponseDTO, error)
	UpdateLesson(ctx context.Context, enrollment models.Enrollment, request dto.UpdateLessonRequestDTO) error
	DeleteLessonById(ctx context.Context, enrollment models.Enrollment, request dto.DeleteLessonRequestDTO) error

	CreateLessonsGeneral(ctx context.Context, enrollment models.Enrollment, request dto.CreateLessonsGeneralRequestDTO) (dto.CreateLessonsGeneralResponseDTO, error)
	UpdateLessonGeneral(ctx context.Context, enrollment models.Enrollment, request dto.UpdateLessonGeneralRequestDTO) error
	DeleteLessonGeneralById(ctx context.Context, enrollment models.Enrollment, request dto.DeleteLessonGeneralRequestDTO) error

	GetUniqueEntries(ctx context.Context, enrollment models.Enrollment, filter dto.EntriesFilterRequestDTO, cursor string, limit int) (dto.EntriesFilterResponseDTO, error)
}

type controller struct {
	validator validators.Validator

	repository repositories.Repository
	registry   types_registry.Controller
}

func New(validator validators.Validator, repository repositories.Repository, registry types_registry.Controller) Controller {
	return &controller{
		validator:  validator,
		repository: repository,
		registry:   registry,
	}
}
