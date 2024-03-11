package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateLessonEntryRequestDTO struct {
	GroupId        uuid.UUID `json:"groupId"`
	RoomId         uuid.UUID `json:"roomId"`
	SubjectId      uuid.UUID `json:"subjectId"`
	TeacherId      uuid.UUID `json:"teacherId"`
	StartTime      time.Time `json:"startTime"`
	EndTime        time.Time `json:"endTime"`
	LessonIndex    int       `json:"lessonIndex"`
	PrimaryColor   string    `json:"primaryColor"`
	SecondaryColor string    `json:"secondaryColor"`
}

type CreateLessonsRequestDTO struct {
	List []CreateLessonEntryRequestDTO `json:"list"`
}

type UpdateLessonRequestDTO struct {
	ID uuid.UUID `json:"id"`
	CreateLessonEntryRequestDTO
}

type CreateLessonGeneralEntryRequestDTO struct {
	GroupId        uuid.UUID     `json:"groupId"`
	RoomId         uuid.UUID     `json:"roomId"`
	SubjectId      uuid.UUID     `json:"subjectId"`
	TeacherId      uuid.UUID     `json:"teacherId"`
	StartTime      time.Duration `json:"startTime"`
	EndTime        time.Duration `json:"endTime"`
	DayIndex       int           `json:"dayIndex"`
	LessonIndex    int           `json:"lessonIndex"`
	PrimaryColor   string        `json:"primaryColor"`
	SecondaryColor string        `json:"secondaryColor"`
}

type CreateLessonsGeneralRequestDTO struct {
	List []CreateLessonGeneralEntryRequestDTO `json:"list"`
}

type UpdateLessonGeneralRequestDTO struct {
	ID uuid.UUID `json:"id"`
	CreateLessonGeneralEntryRequestDTO
}

type CreateScheduleMetaEntryRequestDTO struct {
	Date   time.Time `json:"date"`
	Status string    `json:"status"`
}

type CreateScheduleMetaRequestDTO struct {
	List []CreateScheduleMetaEntryRequestDTO `json:"list"`
}
