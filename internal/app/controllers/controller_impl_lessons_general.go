package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-schedule/internal/app/dto"
	"github.com/stdyum/api-schedule/internal/app/entities"
)

func (c *controller) CreateLessonsGeneral(ctx context.Context, enrollment models.Enrollment, request dto.CreateLessonsGeneralRequestDTO) (dto.CreateLessonsGeneralResponseDTO, error) {
	if err := c.validator.ValidateCreateLessonsGeneralRequest(ctx, request); err != nil {
		return dto.CreateLessonsGeneralResponseDTO{}, err
	}

	if err := enrollment.Permissions.Assert(models.PermissionSchedule); err != nil {
		return dto.CreateLessonsGeneralResponseDTO{}, err
	}

	lessons := make([]entities.LessonGeneral, len(request.List))
	for i, lessonDTO := range request.List {
		lessons[i] = entities.LessonGeneral{
			ID:             uuid.New(),
			StudyPlaceId:   enrollment.StudyPlaceId,
			GroupId:        lessonDTO.GroupId,
			RoomId:         lessonDTO.RoomId,
			SubjectId:      lessonDTO.SubjectId,
			TeacherId:      lessonDTO.TeacherId,
			StartTime:      lessonDTO.StartTime,
			EndTime:        lessonDTO.EndTime,
			DayIndex:       lessonDTO.DayIndex,
			LessonIndex:    lessonDTO.LessonIndex,
			PrimaryColor:   lessonDTO.PrimaryColor,
			SecondaryColor: lessonDTO.SecondaryColor,
		}
	}

	if err := c.repository.CreateGeneralLessons(ctx, lessons); err != nil {
		return dto.CreateLessonsGeneralResponseDTO{}, err
	}

	return dto.CreateLessonsGeneralResponseDTO{
		List: uslices.MapFunc(lessons, func(item entities.LessonGeneral) dto.CreateLessonGeneralEntryResponseDTO {
			return dto.CreateLessonGeneralEntryResponseDTO{
				ID:             item.ID,
				StudyPlaceId:   item.StudyPlaceId,
				GroupId:        item.GroupId,
				RoomId:         item.RoomId,
				SubjectId:      item.SubjectId,
				TeacherId:      item.TeacherId,
				StartTime:      item.StartTime,
				EndTime:        item.EndTime,
				DayIndex:       item.DayIndex,
				LessonIndex:    item.LessonIndex,
				PrimaryColor:   item.PrimaryColor,
				SecondaryColor: item.SecondaryColor,
			}
		}),
	}, nil
}

func (c *controller) UpdateLessonGeneral(ctx context.Context, enrollment models.Enrollment, request dto.UpdateLessonGeneralRequestDTO) error {
	if err := c.validator.ValidateUpdateLessonGeneralRequest(ctx, request); err != nil {
		return err
	}

	if err := enrollment.Permissions.Assert(models.PermissionSchedule); err != nil {
		return err
	}

	lesson := entities.LessonGeneral{
		ID:             request.ID,
		StudyPlaceId:   enrollment.StudyPlaceId,
		GroupId:        request.GroupId,
		RoomId:         request.RoomId,
		SubjectId:      request.SubjectId,
		TeacherId:      request.TeacherId,
		StartTime:      request.StartTime,
		EndTime:        request.EndTime,
		DayIndex:       request.DayIndex,
		LessonIndex:    request.LessonIndex,
		PrimaryColor:   request.PrimaryColor,
		SecondaryColor: request.SecondaryColor,
	}

	return c.repository.UpdateGeneralLesson(ctx, lesson)
}

func (c *controller) DeleteLessonGeneralById(ctx context.Context, enrollment models.Enrollment, request dto.DeleteLessonGeneralRequestDTO) error {
	if err := c.validator.ValidateDeleteLessonGeneralRequest(ctx, request); err != nil {
		return err
	}

	if err := enrollment.Permissions.Assert(models.PermissionSchedule); err != nil {
		return err
	}

	return c.repository.DeleteGeneralLessonById(ctx, enrollment.StudyPlaceId, request.DayIndex, request.ID)
}
