package types_registry

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/proto/impl/types_registry"
	"github.com/stdyum/api-common/uslices"
)

type iRepository interface {
	GetGroupsByIds(ctx context.Context, token string, studyPlaceId uuid.UUID, ids []uuid.UUID) (map[uuid.UUID]models.Group, error)
	GetRoomsByIds(ctx context.Context, token string, studyPlaceId uuid.UUID, ids []uuid.UUID) (map[uuid.UUID]models.Room, error)
	GetSubjectsByIds(ctx context.Context, token string, studyPlaceId uuid.UUID, ids []uuid.UUID) (map[uuid.UUID]models.Subject, error)
	GetTeachersByIds(ctx context.Context, token string, studyPlaceId uuid.UUID, ids []uuid.UUID) (map[uuid.UUID]models.Teacher, error)
}

type repository struct {
	client types_registry.TypesRegistryClient
}

func newRepository(client types_registry.TypesRegistryClient) iRepository {
	return &repository{client: client}
}

func (r *repository) GetGroupsByIds(ctx context.Context, token string, studyPlaceId uuid.UUID, ids []uuid.UUID) (map[uuid.UUID]models.Group, error) {
	response, err := r.client.GetGroupsByIds(ctx, r.tokenAndUUIDsToIdsList(token, studyPlaceId, ids))
	if err != nil {
		return nil, err
	}

	groups := make(map[uuid.UUID]models.Group, len(response.List))
	for id, group := range response.List {
		uid, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}

		groups[uid] = models.Group{
			ID:   uid,
			Name: group.Name,
		}
	}

	return groups, nil
}

func (r *repository) GetRoomsByIds(ctx context.Context, token string, studyPlaceId uuid.UUID, ids []uuid.UUID) (map[uuid.UUID]models.Room, error) {
	response, err := r.client.GetRoomsByIds(ctx, r.tokenAndUUIDsToIdsList(token, studyPlaceId, ids))
	if err != nil {
		return nil, err
	}

	rooms := make(map[uuid.UUID]models.Room, len(response.List))
	for id, room := range response.List {
		uid, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}

		rooms[uid] = models.Room{
			ID:   uid,
			Name: room.Name,
		}
	}

	return rooms, nil
}

func (r *repository) GetSubjectsByIds(ctx context.Context, token string, studyPlaceId uuid.UUID, ids []uuid.UUID) (map[uuid.UUID]models.Subject, error) {
	response, err := r.client.GetSubjectsByIds(ctx, r.tokenAndUUIDsToIdsList(token, studyPlaceId, ids))
	if err != nil {
		return nil, err
	}

	subjects := make(map[uuid.UUID]models.Subject, len(response.List))
	for id, subject := range response.List {
		uid, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}

		subjects[uid] = models.Subject{
			ID:   uid,
			Name: subject.Name,
		}
	}

	return subjects, nil
}

func (r *repository) GetTeachersByIds(ctx context.Context, token string, studyPlaceId uuid.UUID, ids []uuid.UUID) (map[uuid.UUID]models.Teacher, error) {
	response, err := r.client.GetTeachersByIds(ctx, r.tokenAndUUIDsToIdsList(token, studyPlaceId, ids))
	if err != nil {
		return nil, err
	}

	teachers := make(map[uuid.UUID]models.Teacher, len(response.List))
	for id, teacher := range response.List {
		uid, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}

		teachers[uid] = models.Teacher{
			ID:   uid,
			Name: teacher.Name,
		}
	}

	return teachers, nil
}

func (r *repository) uuidsToString(uuids []uuid.UUID) []string {
	return uslices.MapFunc(uuids, func(item uuid.UUID) string {
		return item.String()
	})
}

func (r *repository) tokenAndUUIDsToIdsList(token string, studyPlaceId uuid.UUID, uuids []uuid.UUID) *types_registry.IdList {
	return &types_registry.IdList{
		Token:        token,
		StudyPlaceId: studyPlaceId.String(),
		Ids:          r.uuidsToString(uuids),
	}
}
