package entities

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
)

type Column string

var (
	ColumnSubject Column = "subject_id"
	ColumnRoom    Column = "room_id"
	ColumnGroup   Column = "group_id"
	ColumnTeacher Column = "teacher_id"
)

var (
	ColumnsMap = map[string]Column{
		"subject": ColumnSubject,
		"room":    ColumnRoom,
		"group":   ColumnGroup,
		"teacher": ColumnTeacher,
	}
)

func ColumnFromString(str string) (col Column, ok bool) {
	col, ok = ColumnsMap[str]
	return
}

func (c Column) String() (out string) {
	switch c {
	case ColumnSubject:
		out = "subject"
	case ColumnRoom:
		out = "room"
	case ColumnGroup:
		out = "group"
	case ColumnTeacher:
		out = "teacher"
	default:
		out = string(c)
	}

	return
}

func (c Column) Name(types models.TypesModels, id uuid.UUID) (out string) {
	switch c {
	case ColumnSubject:
		out = types.SubjectsIds[id].Name
	case ColumnRoom:
		out = types.RoomsIds[id].Name
	case ColumnGroup:
		out = types.GroupsIds[id].Name
	case ColumnTeacher:
		out = types.TeachersIds[id].Name
	default:
		out = id.String()
	}

	return
}
