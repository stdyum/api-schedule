package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	schedule "github.com/stdyum/api-schedule/internal/app/models"
)

func (c *controller) fillScheduleTypes(ctx context.Context, enrollment models.Enrollment, lessons []schedule.Lesson) (models.TypesModels, error) {
	groupsMap := make(map[uuid.UUID]bool)
	roomsMap := make(map[uuid.UUID]bool)
	subjectsMap := make(map[uuid.UUID]bool)
	teachersMap := make(map[uuid.UUID]bool)

	for _, lesson := range lessons {
		groupsMap[lesson.Group.ID] = true
		roomsMap[lesson.Room.ID] = true
		subjectsMap[lesson.Subject.ID] = true
		teachersMap[lesson.Teacher.ID] = true
	}

	typesIds := models.TypesIds{
		GroupsIds:   make([]uuid.UUID, 0, len(groupsMap)),
		RoomsIds:    make([]uuid.UUID, 0, len(roomsMap)),
		SubjectsIds: make([]uuid.UUID, 0, len(subjectsMap)),
		TeachersIds: make([]uuid.UUID, 0, len(teachersMap)),
	}

	for id := range groupsMap {
		typesIds.GroupsIds = append(typesIds.GroupsIds, id)
	}
	for id := range roomsMap {
		typesIds.RoomsIds = append(typesIds.RoomsIds, id)
	}
	for id := range subjectsMap {
		typesIds.SubjectsIds = append(typesIds.SubjectsIds, id)
	}
	for id := range teachersMap {
		typesIds.TeachersIds = append(typesIds.TeachersIds, id)
	}

	types, err := c.registry.GetTypesByIds(ctx, enrollment, typesIds)
	if err != nil {
		return models.TypesModels{}, err
	}

	for i := range lessons {
		lessons[i].Group = types.GroupsIds[lessons[i].Group.ID]
		lessons[i].Room = types.RoomsIds[lessons[i].Room.ID]
		lessons[i].Subject = types.SubjectsIds[lessons[i].Subject.ID]
		lessons[i].Teacher = types.TeachersIds[lessons[i].Teacher.ID]
	}

	return types, nil
}
