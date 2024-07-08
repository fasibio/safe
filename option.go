package safe

import (
	"errors"
)

var ErrValueIsNil = errors.New("Value is nil")

type Option[T any] struct {
	value *T
}

func Some[T any](value *T) Option[T] {
	if value == nil {
		return None[T]()
	}
	return Option[T]{value: value}
}

func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

func (o Option[T]) IsSome() bool {
	return o.value != nil
}

func (o Option[T]) IsNone() bool {
	return !o.IsSome()
}

func (o Option[T]) Some() (*T, bool) {

	if o.IsSome() {
		return o.value, true
	}
	return nil, false
}

func (o Option[T]) Unwrap() *T {
	return o.value
}

func (o Option[T]) SomeOrDefault(value *T) *T {
	if result, ok := o.Some(); ok {
		return result
	}
	return value
}

func (o Option[T]) SomeOrDefaultFn(fn func() *T) *T {
	if result, ok := o.Some(); ok {
		return result
	}
	return fn()
}

func (o Option[T]) SomeAndThen(fn func(value *T)) {
	if v, ok := o.Some(); ok {
		fn(v)
	}
}

func (o Option[T]) NoneAndThen(fn func()) {
	if o.IsNone() {
		fn()
	}
}

func (o Option[T]) CopyOrDefault(defaultValue T) T {
	if v, ok := o.Some(); ok {
		return *v
	}
	return defaultValue
}

func (o Option[T]) CopySome() (T, bool) {
	if v, ok := o.Some(); ok {
		return *v, true
	}
	var d T
	return d, false
}
