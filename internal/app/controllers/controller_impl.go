package controllers

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-schedule/internal/app/dto"
	"github.com/stdyum/api-schedule/internal/app/entities"
	schedule "github.com/stdyum/api-schedule/internal/app/models"
)

func (c *controller) Schedule(ctx context.Context, enrollment models.Enrollment, column string, columnId uuid.UUID, from, to time.Time) (dto.ScheduleResponseDTO, error) {
	col, ok := entities.ColumnFromString(column)
	if !ok {
		return dto.ScheduleResponseDTO{}, errors.New("no such column")
	}

	lessonsRaw, err := c.repository.GetSchedule(ctx, enrollment.StudyPlaceId, col, columnId, from, to)
	if err != nil {
		return dto.ScheduleResponseDTO{}, err
	}

	lessons := uslices.MapFunc(lessonsRaw, func(item entities.ScheduleLesson) schedule.Lesson {
		return schedule.Lesson{
			ID:           item.ID,
			StudyPlaceId: item.StudyPlaceId,
			Group: models.Group{
				ID: item.GroupId,
			},
			Room: models.Room{
				ID: item.RoomId,
			},
			Subject: models.Subject{
				ID: item.SubjectId,
			},
			Teacher: models.Teacher{
				ID: item.TeacherId,
			},
			StartDateTime:  item.StartTime,
			EndDateTime:    item.EndTime,
			DayIndex:       item.DayIndex,
			LessonIndex:    item.LessonIndex,
			PrimaryColor:   item.PrimaryColor,
			SecondaryColor: item.SecondaryColor,
		}
	})

	types, err := c.fillScheduleTypes(ctx, enrollment, lessons)
	if err != nil {
		return dto.ScheduleResponseDTO{}, err
	}

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
				StartTime:      item.StartDateTime,
				EndTime:        item.EndDateTime,
				LessonIndex:    item.LessonIndex,
				PrimaryColor:   item.PrimaryColor,
				SecondaryColor: item.SecondaryColor,
			}
		}),
		Info: dto.ScheduleInfoResponseDTO{
			StudyPlaceId: enrollment.StudyPlaceId,
			Column:       col.String(),
			ColumnName:   col.Name(types, columnId),
			StartDate:    from,
			EndDate:      to,
		},
	}, nil
}

func (c *controller) ScheduleGeneral(ctx context.Context, enrollment models.Enrollment, column string, columnId uuid.UUID) (dto.ScheduleGeneralResponseDTO, error) {
	col, ok := entities.ColumnFromString(column)
	if !ok {
		return dto.ScheduleGeneralResponseDTO{}, errors.New("no such column")
	}

	lessonsRaw, err := c.repository.GetScheduleGeneral(ctx, enrollment.StudyPlaceId, col, columnId)
	if err != nil {
		return dto.ScheduleGeneralResponseDTO{}, err
	}

	lessons := uslices.MapFunc(lessonsRaw, func(item entities.LessonGeneral) schedule.Lesson {
		return schedule.Lesson{
			ID:           item.ID,
			StudyPlaceId: item.StudyPlaceId,
			Group: models.Group{
				ID: item.GroupId,
			},
			Room: models.Room{
				ID: item.RoomId,
			},
			Subject: models.Subject{
				ID: item.SubjectId,
			},
			Teacher: models.Teacher{
				ID: item.TeacherId,
			},
			StartTime:      item.StartTime,
			EndTime:        item.EndTime,
			DayIndex:       item.DayIndex,
			LessonIndex:    item.LessonIndex,
			PrimaryColor:   item.PrimaryColor,
			SecondaryColor: item.SecondaryColor,
		}
	})

	types, err := c.fillScheduleTypes(ctx, enrollment, lessons)
	if err != nil {
		return dto.ScheduleGeneralResponseDTO{}, err
	}

	return dto.ScheduleGeneralResponseDTO{
		Lessons: uslices.MapFunc(lessons, func(item schedule.Lesson) dto.ScheduleLessonGeneralResponseDTO {
			return dto.ScheduleLessonGeneralResponseDTO{
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
				DayIndex:       item.DayIndex,
				LessonIndex:    item.LessonIndex,
				PrimaryColor:   item.PrimaryColor,
				SecondaryColor: item.SecondaryColor,
			}
		}),
		Info: dto.ScheduleInfoResponseDTO{
			StudyPlaceId: enrollment.StudyPlaceId,
			Column:       col.String(),
			ColumnName:   col.Name(types, columnId),
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
