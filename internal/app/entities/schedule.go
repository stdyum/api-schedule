package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/entities"
)

type Schedule struct {
	entities.Timed

	ID           uuid.UUID
	StudyPlaceId uuid.UUID
	Date         time.Time
	Status       string
}

type ScheduleLesson struct {
	ID             uuid.UUID
	StudyPlaceId   uuid.UUID
	GroupId        uuid.UUID
	RoomId         uuid.UUID
	SubjectId      uuid.UUID
	TeacherId      uuid.UUID
	StartTime      time.Time
	EndTime        time.Time
	DayIndex       int
	LessonIndex    int
	PrimaryColor   string
	SecondaryColor string
}
