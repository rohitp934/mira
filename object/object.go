package object

import "fmt"

const (
	INTEGER_TYPE = "INTEGER"
	BOOL_TYPE    = "BOOL"
	NULL_TYPE    = "NULL"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGER_TYPE
}

type Bool struct {
	Value bool
}

func (i *Bool) Inspect() string {
	return fmt.Sprintf("%t", i.Value)
}

func (i *Bool) Type() ObjectType {
	return BOOL_TYPE
}

type Null struct{}

func (i *Null) Inspect() string {
	return "null"
}

func (i *Null) Type() ObjectType {
	return NULL_TYPE
}
