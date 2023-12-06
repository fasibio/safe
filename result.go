package safe

import "fmt"

type Result[T any] struct {
	value Option[T]
	err   error
}

func NewResult[T any](value T, err error) Result[T] {
	return Result[T]{
		value: NewOption[T](value),
		err:   err,
	}
}

func (r Result[T]) IsOk() bool {
	if r.err == nil {
		return r.value.IsSome()
	}
	return false
}

func (r Result[T]) IsErr() bool {
	return !r.IsOk()
}

func (r Result[T]) Err() error {
	if r.err != nil {
		return r.err
	}
	if r.value.IsNone() {
		return fmt.Errorf("Err was nil but no value as well")
	}
	return nil
}

func (r Result[T]) OkOrDefault(value T) T {
	if r.IsOk() {
		return r.Ok()
	}
	return value
}

func (r Result[T]) Ok() T {
	if r.IsOk() {
		value, _ := r.value.Some() // already checked at IsOk
		return *value
	}
	panic("was not ok")
}

func (r Result[T]) Unpack() (*T, error) {
	if r.IsOk() {
		value, _ := r.value.Some() // already checked at IsOk
		return value, nil
	}
	return nil, r.Err()
}
