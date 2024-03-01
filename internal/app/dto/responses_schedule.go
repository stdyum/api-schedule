package dto

import (
	"time"

	"github.com/google/uuid"
)

type ScheduleResponseDTO struct {
	Lessons []ScheduleLessonResponseDTO `json:"lessons"`
	Info    ScheduleInfoResponseDTO     `json:"info"`
}

type ScheduleLessonResponseDTO struct {
	ID             uuid.UUID                        `json:"id"`
	StudyPlaceId   uuid.UUID                        `json:"studyPlaceId"`
	Group          ScheduleLessonGroupResponseDTO   `json:"group"`
	Room           ScheduleLessonRoomResponseDTO    `json:"room"`
	Subject        ScheduleLessonSubjectResponseDTO `json:"subject"`
	Teacher        ScheduleLessonTeacherResponseDTO `json:"teacher"`
	StartTime      time.Time                        `json:"startTime"`
	EndTime        time.Time                        `json:"endTime"`
	LessonIndex    int                              `json:"lessonIndex"`
	PrimaryColor   string                           `json:"primaryColor"`
	SecondaryColor string                           `json:"secondaryColor"`
}

type ScheduleLessonGroupResponseDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ScheduleLessonRoomResponseDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ScheduleLessonSubjectResponseDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ScheduleLessonTeacherResponseDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ScheduleInfoResponseDTO struct {
	StudyPlaceId uuid.UUID `json:"studyPlaceId"`
	Column       string    `json:"column"`
	ColumnName   string    `json:"columnName"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
}