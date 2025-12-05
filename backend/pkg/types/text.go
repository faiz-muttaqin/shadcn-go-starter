package types

type Text string

func (t Text) String() string  { return string(t) }
func (t Text) Kind() FieldType { return FieldText }

// optional
func (t Text) IsEmpty() bool { return len(t) == 0 }
