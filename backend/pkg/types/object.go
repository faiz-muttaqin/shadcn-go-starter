package types

type Object[T any] struct {
	Value T
}

func (o Object[T]) String() string  { return "" } // usually not a string
func (o Object[T]) Kind() FieldType { return FieldObject }

func (o Object[T]) Get() T { return o.Value }
