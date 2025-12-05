package types

type Password string

func (p Password) String() string  { return string(p) }
func (p Password) Kind() FieldType { return FieldPassword }

func (p Password) Length() int { return len(p) }
