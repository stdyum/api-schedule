package repositories

import (
	"context"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-schedule/internal/app/entities"
)

func (r *repository) GetLessonByID(ctx context.Context, studyPlaceId uuid.UUID, id uuid.UUID) (entities.Lesson, error) {
	scanner := r.database.Query(`
SELECT id, study_place_id, group_id, room_id, subject_id, teacher_id, date, start_time, end_time, lesson_index, primary_color, secondary_color 
FROM schedule.lessons
WHERE study_place_id = ? AND id = ? ALLOW FILTERING`,
		gocql.UUID(studyPlaceId),
		gocql.UUID(id),
	).
		WithContext(ctx)

	return r.scanLesson(scanner)
}

func (r *repository) GetLessons(ctx context.Context, studyPlaceId uuid.UUID, teacherId uuid.UUID, subjectId uuid.UUID, groupIds []uuid.UUID) ([]entities.Lesson, error) {
	query := `
SELECT id, study_place_id, group_id, room_id, subject_id, teacher_id, date, start_time, end_time, lesson_index, primary_color, secondary_color
FROM schedule.lessons
WHERE study_place_id = ?`
	params := []any{gocql.UUID(studyPlaceId)}

	if teacherId != uuid.Nil {
		query += "AND teacher_id = ? "
		params = append(params, gocql.UUID(teacherId))
	}

	if subjectId != uuid.Nil {
		query += "AND subject_id = ? "
		params = append(params, gocql.UUID(subjectId))
	}

	if len(groupIds) > 0 {
		query += "AND group_id IN ? "
		params = append(params, uslices.MapFunc(groupIds, func(item uuid.UUID) gocql.UUID {
			return gocql.UUID(item)
		}))
	}

	query += "ALLOW FILTERING"

	scanner := r.database.Query(query, params...).
		WithContext(ctx).
		Iter().
		Scanner()

	return databases.ScanArray(scanner, r.scanLesson)
}

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
	//todo set date

	//language=SQL
	err := r.database.Query(`
UPDATE schedule.lessons SET  
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
    study_place_id = ? AND date = ? AND id = ? 
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
		lesson.StartTime,
		gocql.UUID(lesson.ID),
	).WithContext(ctx).Exec()

	return err
}

func (r *repository) DeleteLessonById(ctx context.Context, studyPlaceId uuid.UUID, date time.Time, id uuid.UUID) error {
	return r.database.Query(`
DELETE FROM schedule.lessons WHERE study_place_id = ? AND date = ? AND id = ?
`,
		gocql.UUID(studyPlaceId),
		date,
		gocql.UUID(id),
	).WithContext(ctx).Exec()
}
