package entities

import (
	"github.com/google/uuid"
)

type Entry struct {
	StudyPlaceId uuid.UUID
	TeacherId    uuid.UUID
	SubjectId    uuid.UUID
	GroupId      uuid.UUID
}
