package types

type Badge string

func (b Badge) String() string  { return string(b) }
func (b Badge) Kind() FieldType { return FieldBadge }

func (b Badge) IsValid(options ...string) bool {
	for _, opt := range options {
		if string(b) == opt {
			return true
		}
	}
	return false
}
