package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/entities"
)

type LessonGeneral struct {
	entities.Timed

	ID             uuid.UUID
	StudyPlaceId   uuid.UUID
	GroupId        uuid.UUID
	RoomId         uuid.UUID
	SubjectId      uuid.UUID
	TeacherId      uuid.UUID
	StartTime      time.Duration
	EndTime        time.Duration
	DayIndex       int
	LessonIndex    int
	PrimaryColor   string
	SecondaryColor string
}
