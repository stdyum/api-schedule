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

type DeleteLessonRequestDTO struct {
	ID   uuid.UUID `json:"id"`
	Date time.Time `json:"date"`
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

type DeleteLessonGeneralRequestDTO struct {
	ID       uuid.UUID `json:"id"`
	DayIndex int       `json:"dayIndex"`
}

type CreateScheduleMetaEntryRequestDTO struct {
	Date   time.Time `json:"date"`
	Status string    `json:"status"`
}

type CreateScheduleMetaRequestDTO struct {
	List []CreateScheduleMetaEntryRequestDTO `json:"list"`
}

type GetScheduleRequestDTO struct {
	Column   string    `json:"column"`
	ColumnId uuid.UUID `json:"columnId"`
	From     time.Time `json:"from"`
	To       time.Time `json:"to"`
}

type GetScheduleGeneralRequestDTO struct {
	Column   string    `json:"column"`
	ColumnId uuid.UUID `json:"columnId"`
}

type EntriesFilterRequestDTO struct {
	TeacherId uuid.UUID
	GroupIds  []uuid.UUID
	SubjectId uuid.UUID
}
