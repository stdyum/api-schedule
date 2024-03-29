package repositories

import (
	"context"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/stdyum/api-schedule/internal/app/entities"
)

type Repository interface {
	GetSchedule(ctx context.Context, studyPlaceId uuid.UUID, column entities.Column, columnId uuid.UUID, from, to time.Time) ([]entities.ScheduleLesson, error)
	GetScheduleGeneral(ctx context.Context, studyPlaceId uuid.UUID, column entities.Column, columnId uuid.UUID) ([]entities.LessonGeneral, error)
	CreateScheduleMeta(ctx context.Context, meta []entities.Schedule) error

	CreateLessons(ctx context.Context, lesson []entities.Lesson) error
	UpdateLesson(ctx context.Context, lesson entities.Lesson) error
	DeleteLessonById(ctx context.Context, studyPlaceId uuid.UUID, date time.Time, id uuid.UUID) error

	CreateGeneralLessons(ctx context.Context, lesson []entities.LessonGeneral) error
	UpdateGeneralLesson(ctx context.Context, lesson entities.LessonGeneral) error
	DeleteGeneralLessonById(ctx context.Context, studyPlaceId uuid.UUID, dayIndex int, id uuid.UUID) error
}

type repository struct {
	database *gocql.Session
}

func New(database *gocql.Session) Repository {
	return &repository{
		database: database,
	}
}
