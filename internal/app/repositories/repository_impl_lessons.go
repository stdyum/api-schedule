package repositories

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/stdyum/api-schedule/internal/app/entities"
)

func (r *repository) CreateLessons(ctx context.Context, lessons []entities.Lesson) error {
	query := "BEGIN BATCH"

	var args []any
	for _, lesson := range lessons {
		query += `
INSERT INTO schedule.lessons 
    (id, study_place_id, group_id, room_id, subject_id, teacher_id, date, start_time, end_time, lesson_index, primary_color, secondary_color, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, dateOf(now()), dateOf(now()));
`

		args = append(args, []any{
			gocql.UUID(lesson.ID),
			gocql.UUID(lesson.StudyPlaceId),
			gocql.UUID(lesson.GroupId),
			gocql.UUID(lesson.RoomId),
			gocql.UUID(lesson.SubjectId),
			gocql.UUID(lesson.TeacherId),
			lesson.StartTime,
			lesson.StartTime,
			lesson.EndTime,
			lesson.LessonIndex,
			lesson.PrimaryColor,
			lesson.SecondaryColor,
		}...)
	}
	query += "APPLY BATCH;"

	return r.database.Query(query, args...).WithContext(ctx).Exec()
}

func (r *repository) UpdateLesson(ctx context.Context, lesson entities.Lesson) error {
	return r.database.Query(`
UPDATE schedule.lessons SET  
	group_id = ?,
	room_id = ?,
	subject_id = ?,
	teacher_id = ?,
	date = toDate(?),
	start_time = ?,
	end_time = ?,
	lesson_index = ?,
	primary_color = ?,
	secondary_color = ?,
	updated_at = dateOf(now())
WHERE 
    id = ? AND study_place_id = ?
`,
		gocql.UUID(lesson.GroupId),
		gocql.UUID(lesson.RoomId),
		gocql.UUID(lesson.SubjectId),
		gocql.UUID(lesson.TeacherId),
		lesson.StartTime,
		lesson.StartTime,
		lesson.EndTime,
		lesson.LessonIndex,
		lesson.PrimaryColor,
		lesson.SecondaryColor,
		gocql.UUID(lesson.ID),
		gocql.UUID(lesson.StudyPlaceId),
	).WithContext(ctx).Exec()
}

func (r *repository) DeleteLessonById(ctx context.Context, studyPlaceId uuid.UUID, id uuid.UUID) error {
	return r.database.Query(`
DELETE FROM schedule.lessons WHERE id = ? AND study_place_id = ?
`,
		gocql.UUID(id),
		gocql.UUID(studyPlaceId),
	).WithContext(ctx).Exec()
}
