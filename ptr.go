package safe

func Ptr[T any](v T) *T {
	return &v
}

func SomePtr[T any](v T) Option[T] {
	return Some(Ptr(v))
}
