package types

import "time"

type Date time.Time

func (d Date) String() string {
	return time.Time(d).Format("2006-01-02")
}
func (d Date) Kind() FieldType { return FieldDate }
