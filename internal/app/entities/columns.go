package entities

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
