package repositories

import (
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases"
	"github.com/stdyum/api-schedule/internal/app/entities"
)

func (r *repository) scanSchedule(sc databases.Scan) (schedule entities.Schedule, err error) {
	var entryId gocql.UUID
	var entryStudyPlaceId gocql.UUID

	if err = sc.Scan(&entryId, &entryStudyPlaceId, &schedule.Date, &schedule.Status); err != nil {
		return
	}

	schedule.ID = uuid.UUID(entryId)
	schedule.StudyPlaceId = uuid.UUID(entryStudyPlaceId)

	return
}

func (r *repository) scanLesson(sc databases.Scan) (lesson entities.Lesson, err error) {
	var entryId gocql.UUID
	var entryStudyPlaceId gocql.UUID
	var groupId gocql.UUID
	var roomId gocql.UUID
	var subjectId gocql.UUID
	var teacherId gocql.UUID

	if err = sc.Scan(
		&entryId,
		&entryStudyPlaceId,
		&groupId,
		&roomId,
		&subjectId,
		&teacherId,
		&lesson.Date,
		&lesson.StartTime,
		&lesson.EndTime,
		&lesson.LessonIndex,
		&lesson.PrimaryColor,
		&lesson.SecondaryColor,
	); err != nil {
		return
	}

	lesson.ID = uuid.UUID(entryId)
	lesson.StudyPlaceId = uuid.UUID(entryStudyPlaceId)
	lesson.GroupId = uuid.UUID(groupId)
	lesson.RoomId = uuid.UUID(roomId)
	lesson.SubjectId = uuid.UUID(subjectId)
	lesson.TeacherId = uuid.UUID(teacherId)

	return
}

func (r *repository) scanLessonGeneral(sc databases.Scan) (lesson entities.LessonGeneral, err error) {
	var entryId gocql.UUID
	var entryStudyPlaceId gocql.UUID
	var groupId gocql.UUID
	var roomId gocql.UUID
	var subjectId gocql.UUID
	var teacherId gocql.UUID

	if err = sc.Scan(
		&entryId,
		&entryStudyPlaceId,
		&groupId,
		&roomId,
		&subjectId,
		&teacherId,
		&lesson.StartTime,
		&lesson.EndTime,
		&lesson.DayIndex,
		&lesson.LessonIndex,
		&lesson.PrimaryColor,
		&lesson.SecondaryColor,
	); err != nil {
		return
	}

	lesson.ID = uuid.UUID(entryId)
	lesson.StudyPlaceId = uuid.UUID(entryStudyPlaceId)
	lesson.GroupId = uuid.UUID(groupId)
	lesson.RoomId = uuid.UUID(roomId)
	lesson.SubjectId = uuid.UUID(subjectId)
	lesson.TeacherId = uuid.UUID(teacherId)

	return
}
