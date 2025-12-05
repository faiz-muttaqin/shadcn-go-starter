package types

type Avatar string

func (a Avatar) String() string  { return string(a) }
func (a Avatar) Kind() FieldType { return FieldAvatar }

func (a Avatar) IsURL() bool {
	return len(a) > 4 && (a[:4] == "http" || a[:5] == "https")
}
func (a Avatar) IsPath() bool {
	return len(a) > 0 && a[0] == '/'
}
