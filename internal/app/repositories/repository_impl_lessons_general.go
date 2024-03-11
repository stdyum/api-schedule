package repositories

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/stdyum/api-schedule/internal/app/entities"
)

func (r *repository) CreateGeneralLessons(ctx context.Context, lessons []entities.LessonGeneral) error {
	query := "BEGIN BATCH"

	var args []any
	for _, lesson := range lessons {
		query += `
INSERT INTO schedule.lessons_general 
    (id, study_place_id, group_id, room_id, subject_id, teacher_id, start_time, end_time, day_index, lesson_index, primary_color, secondary_color, created_at, updated_at)
VALUES 
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, dateOf(now()), dateOf(now()));
`

		args = append(args, []any{
			gocql.UUID(lesson.ID),
			gocql.UUID(lesson.StudyPlaceId),
			gocql.UUID(lesson.GroupId),
			gocql.UUID(lesson.RoomId),
			gocql.UUID(lesson.SubjectId),
			gocql.UUID(lesson.TeacherId),
			lesson.StartTime,
			lesson.EndTime,
			lesson.DayIndex,
			lesson.LessonIndex,
			lesson.PrimaryColor,
			lesson.SecondaryColor,
		}...)
	}
	query += "APPLY BATCH;"

	return r.database.Query(query, args...).WithContext(ctx).Exec()
}

func (r *repository) UpdateGeneralLesson(ctx context.Context, lesson entities.LessonGeneral) error {
	return r.database.Query(`
UPDATE schedule.lessons_general SET  
	group_id = ?,
	room_id = ?,
	subject_id = ?,
	teacher_id = ?,
	start_time = ?,
	end_time = ?,
	lesson_index = ?,
	primary_color = ?,
	secondary_color = ?,
	updated_at = dateOf(now())
WHERE 
    study_place_id = ? AND day_index = ? AND id = ? 
IF EXISTS
`,
		gocql.UUID(lesson.GroupId),
		gocql.UUID(lesson.RoomId),
		gocql.UUID(lesson.SubjectId),
		gocql.UUID(lesson.TeacherId),
		lesson.StartTime,
		lesson.EndTime,
		lesson.LessonIndex,
		lesson.PrimaryColor,
		lesson.SecondaryColor,
		gocql.UUID(lesson.StudyPlaceId),
		lesson.DayIndex,
		gocql.UUID(lesson.ID),
	).WithContext(ctx).Exec()
}

func (r *repository) DeleteGeneralLessonById(ctx context.Context, studyPlaceId uuid.UUID, dayIndex int, id uuid.UUID) error {
	return r.database.Query(`
DELETE FROM schedule.lessons_general WHERE study_place_id = ? AND day_index = ? AND id = ?
`,
		gocql.UUID(studyPlaceId),
		dayIndex,
		gocql.UUID(id),
	).WithContext(ctx).Exec()
}
