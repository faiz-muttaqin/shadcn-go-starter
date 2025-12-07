package types

import (
	"database/sql"
	"html/template"
	"reflect"
	"time"
)

type FieldType string

const (
	FieldText     FieldType = "text"
	FieldEmail    FieldType = "email"
	FieldPhone    FieldType = "phone"
	FieldBadge    FieldType = "badge"
	FieldPassword FieldType = "password"
	FieldNumber   FieldType = "number"
	FieldBoolean  FieldType = "boolean"
	FieldTime     FieldType = "time"
	FieldDate     FieldType = "date"
	FieldDatetime FieldType = "datetime"
	FieldObject   FieldType = "object"
	FieldAvatar   FieldType = "avatar"
	FieldImage    FieldType = "image"
	FieldVideo    FieldType = "video"
	FieldAudio    FieldType = "audio"
	FieldFile     FieldType = "file"
	FieldDocument FieldType = "document"
	FieldArchive  FieldType = "archive"
	FieldMedia    FieldType = "media"
	FieldHTML     FieldType = "html"
	FieldCSS      FieldType = "css"
	FieldJS       FieldType = "js"
)

type Field interface {
	Kind() FieldType
	String() string
}

func DetectFieldType(t reflect.Type) FieldType {
	switch t {
	case reflect.TypeOf(Email("")):
		return FieldEmail
	case reflect.TypeOf(Phone("")):
		return FieldPhone
	case reflect.TypeOf(Avatar("")):
		return FieldAvatar
	case reflect.TypeOf(Image("")):
		return FieldImage
	case reflect.TypeOf(File("")):
		return FieldFile
	case reflect.TypeOf(Password("")):
		return FieldPassword
	case reflect.TypeOf(Badge("")):
		return FieldBadge
	case reflect.TypeOf(Image("")):
		return FieldImage
	case reflect.TypeOf(Video("")):
		return FieldVideo
	case reflect.TypeOf(Audio("")):
		return FieldAudio
	case reflect.TypeOf(Document("")):
		return FieldDocument
	case reflect.TypeOf(Archive("")):
		return FieldArchive
	case reflect.TypeOf(Media("")):
		return FieldMedia
	case reflect.TypeOf(File("")):
		return FieldFile
	case reflect.TypeOf(Datetime(time.Time{})):
		return FieldDatetime
	case reflect.TypeOf(Date(time.Time{})):
		return FieldDate
	case reflect.TypeOf(sql.NullTime{}):
		return FieldDatetime
	case reflect.TypeOf(time.Time{}):
		return FieldDatetime
	case reflect.TypeOf(Time(time.Time{})):
		return FieldTime
	case reflect.TypeOf(string("")):
		return FieldText
	case reflect.TypeOf(bool(false)):
		return FieldBoolean
	case reflect.TypeOf(HTML("")), reflect.TypeOf(template.HTML("")):
		return FieldHTML
	case reflect.TypeOf(CSS("")), reflect.TypeOf(template.CSS("")):
		return FieldCSS
	case reflect.TypeOf(JS("")), reflect.TypeOf(template.JS("")):
		return FieldJS
	case reflect.TypeOf(uint(0)), reflect.TypeOf(int(0)),
		reflect.TypeOf(uint8(0)), reflect.TypeOf(int8(0)),
		reflect.TypeOf(uint16(0)), reflect.TypeOf(int16(0)),
		reflect.TypeOf(uint32(0)), reflect.TypeOf(int32(0)),
		reflect.TypeOf(uint64(0)), reflect.TypeOf(int64(0)),
		reflect.TypeOf(float32(0)), reflect.TypeOf(float64(0)):
		return FieldNumber
	}
	return FieldText
}
