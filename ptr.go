package safe

// Ptr returns a pointer of T.
func Ptr[T any](v T) *T {
	return &v
}

// SomePtr returns an option of  pointer of T.
func SomePtr[T any](v T) Option[T] {
	return Some(Ptr(v))
}
