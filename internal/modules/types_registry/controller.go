package types_registry

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
)

type IController interface {
	GetGroupsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) (map[uuid.UUID]models.Group, error)
	GetRoomsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) (map[uuid.UUID]models.Room, error)
	GetSubjectsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) (map[uuid.UUID]models.Subject, error)
	GetTeachersByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) (map[uuid.UUID]models.Teacher, error)
}

type controller struct {
	repository iRepository
}

func newController(repository iRepository) IController {
	return &controller{repository: repository}
}

func (c *controller) GetGroupsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) (map[uuid.UUID]models.Group, error) {
	return c.repository.GetGroupsByIds(ctx, enrollment.Token, enrollment.StudyPlaceId, ids)
}

func (c *controller) GetRoomsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) (map[uuid.UUID]models.Room, error) {
	return c.repository.GetRoomsByIds(ctx, enrollment.Token, enrollment.StudyPlaceId, ids)
}

func (c *controller) GetSubjectsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) (map[uuid.UUID]models.Subject, error) {
	return c.repository.GetSubjectsByIds(ctx, enrollment.Token, enrollment.StudyPlaceId, ids)
}

func (c *controller) GetTeachersByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) (map[uuid.UUID]models.Teacher, error) {
	return c.repository.GetTeachersByIds(ctx, enrollment.Token, enrollment.StudyPlaceId, ids)
}
