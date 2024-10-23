package object

type ObjectType int

const (
	NULL_OBJ ObjectType = iota
	INTEGER_OBJ
	BOOLEAN_OBJ
	RETURN_VALUE_OBJ
	ERROR_OBJ
)

var ObjectTypeStrings = map[ObjectType]string{
	NULL_OBJ:         "NULL",
	INTEGER_OBJ:      "INTEGER",
	BOOLEAN_OBJ:      "BOOLEAN",
	RETURN_VALUE_OBJ: "RETURN_VALUE",
	ERROR_OBJ:        "ERROR",
}

func (o ObjectType) String() string {
	return ObjectTypeStrings[o]
}
