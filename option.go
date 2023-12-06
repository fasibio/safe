package safe

import "fmt"

type Option[T any] struct {
	value *T
}

func NewOption[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (o Option[T]) IsSome() bool {
	if o.value == nil {
		return false
	}
	stringify := fmt.Sprintf("%v", *o.value)
	switch stringify {
	case "<nil>", "{<nil>}", "&{<nil>}", "{}", "&{}":
		return false
	}

	return true
}

func (o Option[T]) IsNone() bool {
	return !o.IsSome()
}

func (o Option[T]) Some() (*T, bool) {
	if o.IsNone() {
		return nil, false
	}
	return o.value, o.IsSome()
}

func (o Option[T]) SomeOrDefault(value T) T {
	if o.IsSome() {
		result, _ := o.Some()
		return *result
	}
	return value
}
