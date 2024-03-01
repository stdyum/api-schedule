package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/entities"
)

type Lesson struct {
	entities.Timed

	ID             uuid.UUID
	StudyPlaceId   uuid.UUID
	GroupId        uuid.UUID
	RoomId         uuid.UUID
	SubjectId      uuid.UUID
	TeacherId      uuid.UUID
	Date           time.Time
	StartTime      time.Time
	EndTime        time.Time
	LessonIndex    int
	PrimaryColor   string
	SecondaryColor string
}
