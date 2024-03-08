package controllers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-schedule/internal/app/dto"
	"github.com/stdyum/api-schedule/internal/app/entities"
	schedule "github.com/stdyum/api-schedule/internal/app/models"
)

func (c *controller) Schedule(ctx context.Context, enrollment models.Enrollment, column string, columnId uuid.UUID, from, to time.Time) (dto.ScheduleResponseDTO, error) {
	col, _ := entities.ColumnFromString(column)

	lessonsRaw, err := c.repository.GetSchedule(ctx, enrollment.StudyPlaceId, col, columnId, from, to)
	if err != nil {
		return dto.ScheduleResponseDTO{}, err
	}

	groupsMap := make(map[uuid.UUID]bool)
	roomsMap := make(map[uuid.UUID]bool)
	subjectsMap := make(map[uuid.UUID]bool)
	teachersMap := make(map[uuid.UUID]bool)
	for _, lesson := range lessonsRaw {
		groupsMap[lesson.GroupId] = true
		roomsMap[lesson.RoomId] = true
		subjectsMap[lesson.SubjectId] = true
		teachersMap[lesson.TeacherId] = true
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

	typesModels, err := c.registry.GetTypesByIds(ctx, enrollment, typesIds)
	if err != nil {
		return dto.ScheduleResponseDTO{}, err
	}

	lessons := uslices.MapFunc(lessonsRaw, func(item entities.ScheduleLesson) schedule.Lesson {
		return schedule.Lesson{
			ID:             item.ID,
			StudyPlaceId:   item.StudyPlaceId,
			Group:          typesModels.GroupsIds[item.GroupId],
			Room:           typesModels.RoomsIds[item.RoomId],
			Subject:        typesModels.SubjectsIds[item.SubjectId],
			Teacher:        typesModels.TeachersIds[item.TeacherId],
			StartTime:      item.StartTime,
			EndTime:        item.EndTime,
			LessonIndex:    item.LessonIndex,
			PrimaryColor:   item.PrimaryColor,
			SecondaryColor: item.SecondaryColor,
		}
	})

	return dto.ScheduleResponseDTO{
		Lessons: uslices.MapFunc(lessons, func(item schedule.Lesson) dto.ScheduleLessonResponseDTO {
			return dto.ScheduleLessonResponseDTO{
				ID:           item.ID,
				StudyPlaceId: item.StudyPlaceId,
				Group: dto.ScheduleLessonGroupResponseDTO{
					ID:   item.Group.ID,
					Name: item.Group.Name,
				},
				Room: dto.ScheduleLessonRoomResponseDTO{
					ID:   item.Room.ID,
					Name: item.Room.Name,
				},
				Subject: dto.ScheduleLessonSubjectResponseDTO{
					ID:   item.Subject.ID,
					Name: item.Subject.Name,
				},
				Teacher: dto.ScheduleLessonTeacherResponseDTO{
					ID:   item.Teacher.ID,
					Name: item.Teacher.Name,
				},
				StartTime:      item.StartTime,
				EndTime:        item.EndTime,
				LessonIndex:    item.LessonIndex,
				PrimaryColor:   item.PrimaryColor,
				SecondaryColor: item.SecondaryColor,
			}
		}),
		Info: dto.ScheduleInfoResponseDTO{
			StudyPlaceId: enrollment.StudyPlaceId,
			Column:       string(col),
			ColumnName:   "--",
			StartDate:    from,
			EndDate:      to,
		},
	}, nil
}

func (c *controller) ScheduleGeneral(ctx context.Context, enrollment models.Enrollment, column string, columnId uuid.UUID) (dto.ScheduleGeneralResponseDTO, error) {
	return dto.ScheduleGeneralResponseDTO{}, nil
}

func (c *controller) CreateScheduleMeta(ctx context.Context, enrollment models.Enrollment, request dto.CreateScheduleMetaRequestDTO) (dto.CreateScheduleMetaResponseDTO, error) {
	if err := enrollment.Permissions.Assert(models.PermissionSchedule); err != nil {
		return dto.CreateScheduleMetaResponseDTO{}, err
	}

	meta := uslices.MapFunc(request.List, func(item dto.CreateScheduleMetaEntryRequestDTO) entities.Schedule {
		return entities.Schedule{
			ID:           uuid.New(),
			StudyPlaceId: enrollment.StudyPlaceId,
			Date:         item.Date,
			Status:       item.Status,
		}
	})

	if err := c.repository.CreateScheduleMeta(ctx, meta); err != nil {
		return dto.CreateScheduleMetaResponseDTO{}, err
	}

	return dto.CreateScheduleMetaResponseDTO{
		List: uslices.MapFunc(meta, func(item entities.Schedule) dto.CreateScheduleMetaEntryResponseDTO {
			return dto.CreateScheduleMetaEntryResponseDTO{
				ID:           item.ID,
				StudyPlaceId: item.StudyPlaceId,
				Date:         item.Date,
				Status:       item.Status,
			}
		}),
	}, nil
}
