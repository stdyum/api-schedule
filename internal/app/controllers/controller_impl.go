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

	groupIds := make([]uuid.UUID, 0, len(groupsMap))
	for id := range groupsMap {
		groupIds = append(groupIds, id)
	}

	roomsIds := make([]uuid.UUID, 0, len(roomsMap))
	for id := range roomsMap {
		roomsIds = append(roomsIds, id)
	}

	subjectsIds := make([]uuid.UUID, 0, len(subjectsMap))
	for id := range subjectsMap {
		subjectsIds = append(subjectsIds, id)
	}

	teachersIds := make([]uuid.UUID, 0, len(teachersMap))
	for id := range teachersMap {
		teachersIds = append(teachersIds, id)
	}

	groups, err := c.registry.GetGroupsByIds(ctx, enrollment, groupIds)
	if err != nil {
		return dto.ScheduleResponseDTO{}, err
	}

	rooms, err := c.registry.GetRoomsByIds(ctx, enrollment, roomsIds)
	if err != nil {
		return dto.ScheduleResponseDTO{}, err
	}

	subjects, err := c.registry.GetSubjectsByIds(ctx, enrollment, subjectsIds)
	if err != nil {
		return dto.ScheduleResponseDTO{}, err
	}

	teachers, err := c.registry.GetTeachersByIds(ctx, enrollment, teachersIds)
	if err != nil {
		return dto.ScheduleResponseDTO{}, err
	}

	lessons := uslices.MapFunc(lessonsRaw, func(item entities.ScheduleLesson) schedule.Lesson {
		return schedule.Lesson{
			ID:             item.ID,
			StudyPlaceId:   item.StudyPlaceId,
			Group:          groups[item.GroupId],
			Room:           rooms[item.RoomId],
			Subject:        subjects[item.SubjectId],
			Teacher:        teachers[item.TeacherId],
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
