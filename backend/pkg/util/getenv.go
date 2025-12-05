package util

import (
	"os"
)

func Getenv[T ~int | ~int64 | ~int32 | ~int16 | ~int8 |
	~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 |
	~float32 | ~float64 |
	~complex64 | ~complex128 |
	~bool | ~string](key string, defaultValue ...T) T {
	if os.Getenv(key) != "" {
		return ConvertStringTo[T](os.Getenv(key))
	}
	if len(defaultValue) > 0 {
		dv := defaultValue[0]
		os.Setenv(key, ConvertToString(dv))
		return defaultValue[0]
	}
	return *new(T)
}
func Setenv[T ~int | ~int64 | ~int32 | ~int16 | ~int8 |
	~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 |
	~float32 | ~float64 |
	~complex64 | ~complex128 |
	~bool | ~string](key string, value T) error {

	if err := os.Setenv(key, ConvertToString(value)); err != nil {
		return err
	}
	return nil
}
