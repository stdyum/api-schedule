package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/grpc"
	"github.com/stdyum/api-common/proto/impl/schedule"
	"github.com/stdyum/api-schedule/internal/app/dto"
)

func (h *gRPC) GetLessonById(ctx context.Context, uuidWrapper *schedule.UUID) (*schedule.Lesson, error) {
	enrollmentUser, err := grpc.EnrollmentAuth(ctx, uuidWrapper.Token, uuidWrapper.StudyPlaceId)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(uuidWrapper.Uuid)
	if err != nil {
		return nil, err
	}

	lesson, err := h.controller.GetLessonById(ctx, enrollmentUser.Enrollment, id)
	if err != nil {
		return nil, err
	}

	return &schedule.Lesson{
		Id:             lesson.ID.String(),
		StudyPlaceId:   lesson.StudyPlaceId.String(),
		GroupId:        lesson.GroupId.String(),
		RoomId:         lesson.RoomId.String(),
		SubjectId:      lesson.SubjectId.String(),
		TeacherId:      lesson.TeacherId.String(),
		StartTime:      lesson.StartTime.Unix(),
		EndTime:        lesson.EndTime.Unix(),
		LessonIndex:    int32(lesson.LessonIndex),
		PrimaryColor:   lesson.PrimaryColor,
		SecondaryColor: lesson.SecondaryColor,
	}, nil
}

func (h *gRPC) GetLessons(ctx context.Context, filter *schedule.EntriesFilter) (*schedule.Lessons, error) {
	enrollmentUser, err := grpc.EnrollmentAuth(ctx, filter.Token, filter.StudyPlaceId)
	if err != nil {
		return nil, err
	}

	teacherId, err := uuid.Parse(filter.TeacherId)
	if err != nil {
		return nil, err
	}

	groupId, err := uuid.Parse(filter.GroupId)
	if err != nil {
		return nil, err
	}

	subjectId, err := uuid.Parse(filter.SubjectId)
	if err != nil {
		return nil, err
	}

	in := dto.EntriesFilterRequestDTO{
		TeacherId: teacherId,
		GroupId:   groupId,
		SubjectId: subjectId,
	}

	lessons, err := h.controller.GetLessons(ctx, enrollmentUser.Enrollment, in)
	if err != nil {
		return nil, err
	}

	out := schedule.Lessons{
		List: make([]*schedule.Lesson, len(lessons)),
	}

	for i, lesson := range lessons {
		out.List[i] = &schedule.Lesson{
			Id:             lesson.ID.String(),
			StudyPlaceId:   lesson.StudyPlaceId.String(),
			GroupId:        lesson.GroupId.String(),
			RoomId:         lesson.RoomId.String(),
			SubjectId:      lesson.SubjectId.String(),
			TeacherId:      lesson.TeacherId.String(),
			StartTime:      lesson.StartTime.Unix(),
			EndTime:        lesson.EndTime.Unix(),
			LessonIndex:    int32(lesson.LessonIndex),
			PrimaryColor:   lesson.PrimaryColor,
			SecondaryColor: lesson.SecondaryColor,
		}
	}

	return &out, nil
}

func (h *gRPC) GetUniqueEntries(ctx context.Context, filter *schedule.EntriesFilter) (*schedule.Entries, error) {
	enrollmentUser, err := grpc.EnrollmentAuth(ctx, filter.Token, filter.StudyPlaceId)
	if err != nil {
		return nil, err
	}

	teacherId, err := uuid.Parse(filter.TeacherId)
	if err != nil {
		return nil, err
	}

	groupId, err := uuid.Parse(filter.GroupId)
	if err != nil {
		return nil, err
	}

	subjectId, err := uuid.Parse(filter.SubjectId)
	if err != nil {
		return nil, err
	}

	in := dto.EntriesFilterRequestDTO{
		TeacherId: teacherId,
		GroupId:   groupId,
		SubjectId: subjectId,
	}

	entries, err := h.controller.GetUniqueEntries(ctx, enrollmentUser.Enrollment, in)
	if err != nil {
		return nil, err
	}

	out := schedule.Entries{
		List: make([]*schedule.Entry, len(entries)),
	}

	for i, entry := range entries {
		out.List[i] = &schedule.Entry{
			TeacherId: entry.TeacherId.String(),
			GroupId:   entry.GroupId.String(),
			SubjectId: entry.SubjectId.String(),
		}
	}

	return &out, nil
}
