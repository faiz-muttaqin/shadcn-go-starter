package types

import "time"

type Datetime time.Time

func (dt Datetime) String() string {
	return time.Time(dt).Format(time.RFC3339)
}
func (dt Datetime) Kind() FieldType { return FieldDatetime }
