package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-schedule/internal/app/entities"
)

func (r *repository) GetSchedule(ctx context.Context, studyPlaceId uuid.UUID, column entities.Column, columnId uuid.UUID, from, to time.Time) ([]entities.ScheduleLesson, error) {
	type day struct {
		isCurrent bool
		lessons   []entities.ScheduleLesson
	}

	daysAmount := int(to.Sub(from).Hours() / 24)
	days := make(map[time.Time]*day, daysAmount)
	for i := 0; i < daysAmount; i++ {
		days[from.AddDate(0, 0, i)] = &day{}
	}

	//language=SQL
	query := `
SELECT id, study_place_id, date, status FROM schedule.schedule
	WHERE date > ? AND date < ? AND study_place_id = ?
`

	scanner := r.database.Query(query,
		from,
		to,
		gocql.UUID(studyPlaceId),
	).WithContext(ctx).Iter().Scanner()

	schedule, err := databases.ScanArray(scanner, r.scanSchedule)
	if err != nil {
		return nil, err
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	currentLessonsTime := make([]time.Time, 0, len(schedule))
	for _, s := range schedule {
		if s.Status == "updated" {
			d := days[s.Date]
			d.isCurrent = true
			currentLessonsTime = append(currentLessonsTime, s.Date)
		}
	}

	generalLessonsIndexes := make([]int, 0, 7)
	{
		generalLessonsIndexesMap := make(map[int]bool)
		for t, d := range days {
			if d.isCurrent {
				continue
			}

			weekday := int(t.Weekday())
			if _, ok := generalLessonsIndexesMap[weekday]; !ok {
				generalLessonsIndexes = append(generalLessonsIndexes, weekday)
				generalLessonsIndexesMap[weekday] = true
			}

			if len(generalLessonsIndexes) == 7 {
				break
			}
		}
	}

	//language=SQL
	query = fmt.Sprintf(`
SELECT id, study_place_id, group_id, room_id, subject_id, teacher_id, date, start_time, end_time, lesson_index, primary_color, secondary_color  FROM schedule.lessons 
    WHERE study_place_id = ? AND date IN ? AND %s = ?
`, column)

	scanner = r.database.Query(query,
		gocql.UUID(studyPlaceId),
		currentLessonsTime,
		gocql.UUID(columnId),
	).WithContext(ctx).Iter().Scanner()

	currentLessonsRaw, err := databases.ScanArray(scanner, r.scanLesson)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	//language=SQL
	query = fmt.Sprintf(`
SELECT id, study_place_id, group_id, room_id, subject_id, teacher_id, start_time, end_time, day_index, lesson_index, primary_color, secondary_color FROM lessons_general 
    WHERE study_place_id = ? AND day_index IN ? AND %s = ?
`, column)

	scanner = r.database.Query(query,
		gocql.UUID(studyPlaceId),
		generalLessonsIndexes,
		gocql.UUID(columnId),
	).WithContext(ctx).Iter().Scanner()

	generalLessonsRaw, err := databases.ScanArray(scanner, r.scanLessonGeneral)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	currentLessons := uslices.MapFunc(currentLessonsRaw, func(item entities.Lesson) entities.ScheduleLesson {
		return entities.ScheduleLesson{
			ID:             item.ID,
			StudyPlaceId:   item.StudyPlaceId,
			GroupId:        item.GroupId,
			RoomId:         item.RoomId,
			SubjectId:      item.SubjectId,
			TeacherId:      item.TeacherId,
			StartTime:      item.StartTime,
			EndTime:        item.EndTime,
			LessonIndex:    item.LessonIndex,
			PrimaryColor:   item.PrimaryColor,
			SecondaryColor: item.SecondaryColor,
		}
	})

	generalLessons := make(map[int][]entities.LessonGeneral, 7)
	for _, lesson := range generalLessonsRaw {
		generalLessons[lesson.DayIndex] = append(generalLessons[lesson.DayIndex], lesson)
	}

	total := len(currentLessonsRaw)
	for t, d := range days {
		if d.isCurrent {
			continue
		}

		lessons := uslices.MapFunc(generalLessons[int(t.Weekday())], func(item entities.LessonGeneral) entities.ScheduleLesson {
			return entities.ScheduleLesson{
				ID:             item.ID,
				StudyPlaceId:   item.StudyPlaceId,
				GroupId:        item.GroupId,
				RoomId:         item.RoomId,
				SubjectId:      item.SubjectId,
				TeacherId:      item.TeacherId,
				StartTime:      t.Add(item.StartTime),
				EndTime:        t.Add(item.EndTime),
				LessonIndex:    item.LessonIndex,
				PrimaryColor:   item.PrimaryColor,
				SecondaryColor: item.SecondaryColor,
			}
		})
		d.lessons = lessons
		total += len(d.lessons)
	}

	lessons := make([]entities.ScheduleLesson, 0, total)
	lessons = append(lessons, currentLessons...)
	for _, scheduleLessons := range days {
		lessons = append(lessons, scheduleLessons.lessons...)
	}

	return lessons, nil
}

func (r *repository) GetScheduleGeneral(ctx context.Context, studyPlaceId uuid.UUID, column entities.Column, columnId uuid.UUID) ([]entities.LessonGeneral, error) {
	//language=SQL
	query := fmt.Sprintf(`
SELECT id, study_place_id, group_id, room_id, subject_id, teacher_id, start_time, end_time, day_index, lesson_index, primary_color, secondary_color FROM lessons_general 
    WHERE study_place_id = ? AND %s = ?
`, column)

	scanner := r.database.Query(query,
		gocql.UUID(studyPlaceId),
		gocql.UUID(columnId),
	).WithContext(ctx).Iter().Scanner()

	lessons, err := databases.ScanArray(scanner, r.scanLessonGeneral)
	if err != nil {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		return nil, err
	}

	return lessons, nil
}

func (r *repository) CreateScheduleMeta(ctx context.Context, meta []entities.Schedule) error {
	query := "BEGIN BATCH"

	var args []any
	for _, metaEntry := range meta {
		query += `
INSERT INTO schedule.schedule 
    (id, study_place_id, date, status, created_at, updated_at)
VALUES (?, ?, ?, ?, dateOf(now()), dateOf(now()));
`

		args = append(args, []any{
			gocql.UUID(metaEntry.ID),
			gocql.UUID(metaEntry.StudyPlaceId),
			metaEntry.Date,
			metaEntry.Status,
		}...)
	}
	query += "APPLY BATCH;"

	return r.database.Query(query, args...).WithContext(ctx).Exec()
}
