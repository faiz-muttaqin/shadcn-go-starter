package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// TryConvertStringTo mencoba convert string ke tipe target.
// Kalau gagal, return zero value + error.
func TryConvertStringTo[T ~int | ~int64 | ~int32 | ~int16 | ~int8 |
	~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 |
	~float32 | ~float64 |
	~complex64 | ~complex128 |
	~bool | ~string](input string) (T, error) {

	var result any
	var err error

	switch any(*new(T)).(type) {
	case int:
		var v int
		v, err = strconv.Atoi(input)
		result = v
	case int8:
		var v int64
		v, err = strconv.ParseInt(input, 10, 8)
		result = int8(v)
	case int16:
		var v int64
		v, err = strconv.ParseInt(input, 10, 16)
		result = int16(v)
	case int32:
		var v int64
		v, err = strconv.ParseInt(input, 10, 32)
		result = int32(v)
	case int64:
		var v int64
		v, err = strconv.ParseInt(input, 10, 64)
		result = v

	case uint:
		var v uint64
		v, err = strconv.ParseUint(input, 10, 64)
		result = uint(v)
	case uint8:
		var v uint64
		v, err = strconv.ParseUint(input, 10, 8)
		result = uint8(v)
	case uint16:
		var v uint64
		v, err = strconv.ParseUint(input, 10, 16)
		result = uint16(v)
	case uint32:
		var v uint64
		v, err = strconv.ParseUint(input, 10, 32)
		result = uint32(v)
	case uint64:
		var v uint64
		v, err = strconv.ParseUint(input, 10, 64)
		result = v

	case float32:
		var v float64
		v, err = strconv.ParseFloat(input, 32)
		result = float32(v)
	case float64:
		var v float64
		v, err = strconv.ParseFloat(input, 64)
		result = v

	case complex64:
		var v complex128
		v, err = strconv.ParseComplex(input, 64)
		result = complex64(v)
	case complex128:
		var v complex128
		v, err = strconv.ParseComplex(input, 128)
		result = v
	case time.Duration:
		var v time.Duration
		v, err = time.ParseDuration(input)
		result = v
	case bool:
		lower := strings.ToLower(strings.TrimSpace(input))
		switch lower {
		case "true", "1", "-1":
			result = true
		case "false", "0", "":
			result = false
		default:
			err = fmt.Errorf("invalid bool: %s", input)
		}

	case string:
		result = input

	default:
		err = fmt.Errorf("unsupported type")
	}

	if err != nil {
		var zero T
		return zero, err
	}
	return result.(T), nil
}

// ConvertStringTo versi simple, tidak return error.
// Kalau gagal → log error, return zero value.
func ConvertStringTo[T ~int | ~int64 | ~int32 | ~int16 | ~int8 |
	~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 |
	~float32 | ~float64 |
	~complex64 | ~complex128 |
	~bool | ~string](input string, default_value ...T) T {

	v, err := TryConvertStringTo[T](input)
	if err != nil {
		logrus.Debugf("ConvertStringTo[%T]: cannot convert '%s': %v", v, input, err)
		if len(default_value) > 0 {
			return default_value[0]
		}
		var zero T
		return zero
	}
	return v
}

// ConvertToString mengubah berbagai tipe angka & bool ke string (base 10 untuk numerik)
func ConvertToString[T ~int | ~int64 | ~int32 | ~int16 | ~int8 |
	~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 |
	~float32 | ~float64 |
	~complex64 | ~complex128 |
	~bool | ~string](v T) string {

	switch val := any(v).(type) {
	case int, int64, int32, int16, int8:
		return fmt.Sprintf("%d", val)

	case uint, uint64, uint32, uint16, uint8:
		return fmt.Sprintf("%d", val)

	case float32:
		return fmt.Sprintf("%f", float64(val))
	case float64:
		return fmt.Sprintf("%f", val)

	case complex64, complex128:
		// pakai fmt biar output seperti "(1+2i)"
		return fmt.Sprint(val)

	case bool:
		// true/false jadi string langsung
		if val {
			return "true"
		}
		return "false"

	default:
		return fmt.Sprint(val)
	}
}
