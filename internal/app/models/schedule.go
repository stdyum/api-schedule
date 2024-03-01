package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
)

type Lesson struct {
	ID             uuid.UUID
	StudyPlaceId   uuid.UUID
	Group          models.Group
	Room           models.Room
	Subject        models.Subject
	Teacher        models.Teacher
	StartTime      time.Time
	EndTime        time.Time
	LessonIndex    int
	PrimaryColor   string
	SecondaryColor string
}
