package controllers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-schedule/internal/app/dto"
	"github.com/stdyum/api-schedule/internal/app/repositories"
	"github.com/stdyum/api-schedule/internal/modules/types_registry"
)

type Controller interface {
	Schedule(ctx context.Context, enrollment models.Enrollment, column string, columnId uuid.UUID, from, to time.Time) (dto.ScheduleResponseDTO, error)
	CreateScheduleMeta(ctx context.Context, enrollment models.Enrollment, request dto.CreateScheduleMetaRequestDTO) (dto.CreateScheduleMetaResponseDTO, error)

	CreateLessons(ctx context.Context, enrollment models.Enrollment, request dto.CreateLessonsRequestDTO) (dto.CreateLessonsResponseDTO, error)
	UpdateLesson(ctx context.Context, enrollment models.Enrollment, request dto.UpdateLessonRequestDTO) error
	DeleteLessonById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) error

	CreateLessonsGeneral(ctx context.Context, enrollment models.Enrollment, request dto.CreateLessonsGeneralRequestDTO) (dto.CreateLessonsGeneralResponseDTO, error)
	UpdateLessonGeneral(ctx context.Context, enrollment models.Enrollment, request dto.UpdateLessonGeneralRequestDTO) error
	DeleteLessonGeneralById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) error
}

type controller struct {
	repository repositories.Repository
	registry   types_registry.IController
}

func New(repository repositories.Repository, registry types_registry.IController) Controller {
	return &controller{repository: repository, registry: registry}
}
