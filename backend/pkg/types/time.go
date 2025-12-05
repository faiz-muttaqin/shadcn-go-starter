package types

import "time"

type Time time.Time

func (d Time) String() string {
	return time.Time(d).Format("2006-01-02")
}
func (d Time) Kind() FieldType { return FieldTime }
