# safe

A lightweight Go library that provides the `Option[T]` struct, which enforces explicit handling of nil values by requiring unwrapping to access the actual value.

## Features
- Prevents accidental `nil` dereferencing.
- Provides utility methods to safely handle optional values.
- Implements JSON marshaling and unmarshaling.
- Inspired by Rust's `Option<T>` type.
- Can be used as a field within structs.

## Installation
```sh
go get github.com/fasibio/safe
```

## Usage

### Creating an Option
```go
package main

import (
	"fmt"
	"github.com/fasibio/safe"
)

func main() {
	value := 42
	opt := safe.Some(&value)
	none := safe.None[int]()

	fmt.Println("Is Some?", opt.IsSome()) // true
	fmt.Println("Is None?", none.IsNone()) // true
}
```

### Using `Option` as a Struct Field
```go
type Obj struct {
	V safe.Option[string]
	B string
}
```

### Unwrapping a Value
`Unwrap` is for use when you're sure the value is not `nil`.
```go
opt := safe.Some("Hello")
value := opt.Unwrap()
fmt.Println(*value)   // Output: Hello
```

### Handling `None` Gracefully
```go
opt := safe.None[string]()
value := opt.SomeOrDefault("default value")
fmt.Println(*value) // Output: default value
```

### Common Pattern for Using `Some`
```go
opt := safe.Some("GoLang")
if value, ok := opt.Some(); ok {
	fmt.Println("Value:", *value) // Output: Value: GoLang
}
```

### Applying Transformations with `SomeAndMap`
```go
opt := safe.Some(5)
result := safe.SomeAndMap(opt, func(v *int) safe.Option[string] {
	str := fmt.Sprintf("Number: %d", *v)
	return safe.Some(&str)
})

if v, ok := result.Some(); ok {
	fmt.Println(*v) // Output: Number: 5
}
```

### JSON Serialization
`Option` implements `json.Marshaler` and `json.Unmarshaler`, allowing it to be used in structs seamlessly.
```go
import (
	"encoding/json"
	"fmt"
)

type TestStruct struct {
	A int              `json:"a,omitempty"`
	B safe.Option[int] `json:"b,omitempty"`
	C string           `json:"c,omitempty"`
}

opt := safe.Some(TestStruct{
	A: 1,
	B: safe.Some(safe.Ptr(10)),
	C: "test",
})

data, _ := json.Marshal(opt)
fmt.Println(string(data))

var newOpt safe.Option[TestStruct]
json.Unmarshal(data, &newOpt)
fmt.Println(newOpt.IsSome()) // Output: true
```

## Running Tests
Unit tests are available to validate library functionality. Run tests using:
```sh
go test ./...
```

## License
Apache License 2.0

