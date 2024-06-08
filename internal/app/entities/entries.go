package entities

import (
	"github.com/google/uuid"
)

type UniqueEntry struct {
	Id           string
	StudyPlaceId uuid.UUID
	TeacherId    uuid.UUID
	SubjectId    uuid.UUID
	GroupId      uuid.UUID
}
