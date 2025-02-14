package safe

import (
	"encoding/json"
	"errors"
)

var ErrValueIsNil = errors.New("Value is nil")

// Wrapper around potenzial nil Values. Force nil checks before T can be used.
type Option[T any] struct {
	value *T
}

// Some creates an option of type T could nil.
func Some[T any](value *T) Option[T] {
	if value == nil {
		return None[T]()
	}
	return Option[T]{value: value}
}

// None create a None Option by given T.
func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

// IsSome returns true if the Option contains a value.
func (o Option[T]) IsSome() bool {
	return o.value != nil
}

// IsNone returns true if the Option is empty.
func (o Option[T]) IsNone() bool {
	return !o.IsSome()
}

// Some returns the value and true if Option is not empty, otherwise nil and false.
func (o Option[T]) Some() (*T, bool) {

	if o.IsSome() {
		return o.value, true
	}
	return nil, false
}

// Unwrap returns the contained value without a nil check (use with caution).
func (o Option[T]) Unwrap() *T {
	return o.value
}

// SomeOrDefault returns the value if present, otherwise returns the provided default value.
func (o Option[T]) SomeOrDefault(value *T) *T {
	if result, ok := o.Some(); ok {
		return result
	}
	return value
}

// SomeOrDefaultFn returns the value if present, otherwise calls the provided function to generate a default.
func (o Option[T]) SomeOrDefaultFn(fn func() *T) *T {
	if result, ok := o.Some(); ok {
		return result
	}
	return fn()
}

// SomeOrError returns the value if present, otherwise returns the provided error.
func (o Option[T]) SomeOrError(e error) (*T, error) {
	if v, ok := o.Some(); ok {
		return v, nil
	}
	return nil, e
}

// SomeAndThen calls the provided function if Option contains a value.
func (o Option[T]) SomeAndThen(fn func(value *T)) {
	if v, ok := o.Some(); ok {
		fn(v)
	}
}

// NoneAndThen calls the provided function if Option is empty.
func (o Option[T]) NoneAndThen(fn func()) {
	if o.IsNone() {
		fn()
	}
}

// CopyOrDefault returns a copy of the value if present, otherwise returns the default value.
func (o Option[T]) CopyOrDefault(defaultValue T) T {
	if v, ok := o.Some(); ok {
		return *v
	}
	return defaultValue
}

// CopySome returns a copy of the value and true if present, otherwise a zero-value and false.
func (o Option[T]) CopySome() (T, bool) {
	if v, ok := o.Some(); ok {
		return *v, true
	}
	var d T
	return d, false
}

// Transforms an Option[T] to an Option[P] using the provided function.
func SomeAndMap[T any, P any](o Option[T], fn func(*T) Option[P]) Option[P] {
	if v, ok := o.Some(); ok {
		return fn(v)
	}
	return None[P]()
}

var _ json.Marshaler = (*Option[any])(nil)
var _ json.Unmarshaler = (*Option[any])(nil)

func (e *Option[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		e.value = nil
		return nil
	}
	var v T
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	e.value = &v
	return nil
}

func (e *Option[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.value)
}
