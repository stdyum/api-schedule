package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
)

type Lesson struct {
	ID             uuid.UUID
	StudyPlaceId   uuid.UUID
	Type           string
	Group          models.Group
	Room           models.Room
	Subject        models.Subject
	Teacher        models.Teacher
	StartDateTime  time.Time
	EndDateTime    time.Time
	StartTime      time.Duration
	EndTime        time.Duration
	DayIndex       int
	LessonIndex    int
	PrimaryColor   string
	SecondaryColor string
}
