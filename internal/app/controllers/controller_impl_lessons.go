package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-schedule/internal/app/dto"
	"github.com/stdyum/api-schedule/internal/app/entities"
	"github.com/stdyum/api-schedule/pkg/uuuid"
)

func (c *controller) GetLessonById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.LessonResponseDTO, error) {
	lesson, err := c.repository.GetLessonByID(ctx, enrollment.StudyPlaceId, id)
	if err != nil {
		return dto.LessonResponseDTO{}, err
	}

	return dto.LessonResponseDTO{
		ID:             lesson.ID,
		StudyPlaceId:   lesson.StudyPlaceId,
		GroupId:        lesson.GroupId,
		RoomId:         lesson.RoomId,
		SubjectId:      lesson.SubjectId,
		TeacherId:      lesson.TeacherId,
		StartTime:      lesson.StartTime,
		EndTime:        lesson.EndTime,
		LessonIndex:    lesson.LessonIndex,
		PrimaryColor:   lesson.PrimaryColor,
		SecondaryColor: lesson.SecondaryColor,
	}, nil
}

func (c *controller) GetLessons(ctx context.Context, enrollment models.Enrollment, filter dto.EntriesFilterRequestDTO) ([]dto.LessonResponseDTO, error) {
	lessons, err := c.repository.GetLessons(ctx, enrollment.StudyPlaceId, filter.TeacherId, filter.SubjectId, filter.GroupIds)
	if err != nil {
		return nil, err
	}

	return uslices.MapFunc(lessons, func(item entities.Lesson) dto.LessonResponseDTO {
		return dto.LessonResponseDTO{
			ID:             item.ID,
			StudyPlaceId:   item.StudyPlaceId,
			GroupId:        item.GroupId,
			RoomId:         item.RoomId,
			SubjectId:      item.SubjectId,
			TeacherId:      item.TeacherId,
			StartTime:      item.StartTime,
			EndTime:        item.EndTime,
			LessonIndex:    item.LessonIndex,
			PrimaryColor:   item.PrimaryColor,
			SecondaryColor: item.SecondaryColor,
		}
	}), nil
}

func (c *controller) CreateLessons(ctx context.Context, enrollment models.Enrollment, request dto.CreateLessonsRequestDTO) (dto.CreateLessonsResponseDTO, error) {
	//if err := c.validator.ValidateCreateLessonsRequest(ctx, request); err != nil {
	//	return dto.CreateLessonsResponseDTO{}, err
	//}

	if err := enrollment.Permissions.Assert(models.PermissionSchedule); err != nil {
		return dto.CreateLessonsResponseDTO{}, err
	}

	lessons := make([]entities.Lesson, len(request.List))
	entries := make([]entities.UniqueEntry, len(request.List))
	for i, lessonDTO := range request.List {
		lessons[i] = entities.Lesson{
			ID:             uuid.New(),
			StudyPlaceId:   enrollment.StudyPlaceId,
			GroupId:        lessonDTO.GroupId,
			RoomId:         lessonDTO.RoomId,
			SubjectId:      lessonDTO.SubjectId,
			TeacherId:      lessonDTO.TeacherId,
			StartTime:      lessonDTO.StartTime,
			EndTime:        lessonDTO.EndTime,
			LessonIndex:    lessonDTO.LessonIndex,
			PrimaryColor:   lessonDTO.PrimaryColor,
			SecondaryColor: lessonDTO.SecondaryColor,
		}

		entries[i] = entities.UniqueEntry{
			Id:           uuuid.UUIDsToBase56("", lessonDTO.GroupId, lessonDTO.SubjectId, lessonDTO.TeacherId),
			StudyPlaceId: enrollment.StudyPlaceId,
			GroupId:      lessonDTO.GroupId,
			SubjectId:    lessonDTO.SubjectId,
			TeacherId:    lessonDTO.TeacherId,
		}
	}

	if err := c.repository.CreateLessons(ctx, lessons); err != nil {
		return dto.CreateLessonsResponseDTO{}, err
	}

	if err := c.repository.CreateUniqueEntries(ctx, entries); err != nil {
		return dto.CreateLessonsResponseDTO{}, err
	}

	return dto.CreateLessonsResponseDTO{
		List: uslices.MapFunc(lessons, func(item entities.Lesson) dto.CreateLessonEntryResponseDTO {
			return dto.CreateLessonEntryResponseDTO{
				ID:             item.ID,
				StudyPlaceId:   item.StudyPlaceId,
				GroupId:        item.GroupId,
				RoomId:         item.RoomId,
				SubjectId:      item.SubjectId,
				TeacherId:      item.TeacherId,
				StartTime:      item.StartTime,
				EndTime:        item.EndTime,
				LessonIndex:    item.LessonIndex,
				PrimaryColor:   item.PrimaryColor,
				SecondaryColor: item.SecondaryColor,
			}
		}),
	}, nil
}

func (c *controller) UpdateLesson(ctx context.Context, enrollment models.Enrollment, request dto.UpdateLessonRequestDTO) error {
	if err := c.validator.ValidateUpdateLessonRequest(ctx, request); err != nil {
		return err
	}

	if err := enrollment.Permissions.Assert(models.PermissionSchedule); err != nil {
		return err
	}

	lesson := entities.Lesson{
		ID:             request.ID,
		StudyPlaceId:   enrollment.StudyPlaceId,
		GroupId:        request.GroupId,
		RoomId:         request.RoomId,
		SubjectId:      request.SubjectId,
		TeacherId:      request.TeacherId,
		StartTime:      request.StartTime,
		EndTime:        request.EndTime,
		LessonIndex:    request.LessonIndex,
		PrimaryColor:   request.PrimaryColor,
		SecondaryColor: request.SecondaryColor,
	}

	return c.repository.UpdateLesson(ctx, lesson)
}

func (c *controller) DeleteLessonById(ctx context.Context, enrollment models.Enrollment, request dto.DeleteLessonRequestDTO) error {
	if err := c.validator.ValidateDeleteLessonRequest(ctx, request); err != nil {
		return err
	}

	if err := enrollment.Permissions.Assert(models.PermissionSchedule); err != nil {
		return err
	}

	return c.repository.DeleteLessonById(ctx, enrollment.StudyPlaceId, request.Date, request.ID)
}
