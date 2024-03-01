package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateLessonEntryResponseDTO struct {
	ID             uuid.UUID `json:"id"`
	StudyPlaceId   uuid.UUID `json:"studyPlaceId"`
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

type CreateLessonsResponseDTO struct {
	List []CreateLessonEntryResponseDTO `json:"list"`
}

type CreateLessonGeneralEntryResponseDTO struct {
	ID             uuid.UUID     `json:"id"`
	StudyPlaceId   uuid.UUID     `json:"studyPlaceId"`
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

type CreateLessonsGeneralResponseDTO struct {
	List []CreateLessonGeneralEntryResponseDTO `json:"list"`
}

type CreateScheduleMetaEntryResponseDTO struct {
	ID           uuid.UUID `json:"id"`
	StudyPlaceId uuid.UUID `json:"studyPlaceId"`
	Date         time.Time `json:"date"`
	Status       string    `json:"status"`
}

type CreateScheduleMetaResponseDTO struct {
	List []CreateScheduleMetaEntryResponseDTO `json:"list"`
}
