package safe_test

import (
	"log"
	"testing"

	"github.com/fasibio/safe"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	a string
}

func TestOptionalExample(t *testing.T) {
	exp := assert.New(t)

	opt := safe.Some(SomeFunctionWichReturnPointer())
	exp.True(opt.IsSome())
	if opt.IsSome() {
		// happy path
		value, _ := opt.Some()
		log.Println(value)
	} else {
		log.Println("is None")
		// is not set
	}

	//or

	if value, ok := opt.Some(); ok {
		// happy path
		log.Println(value)
	} else {
		log.Println("is None")
		// is not set
	}

}

func SomeFunctionWichReturnPointer() *int {
	result := 1
	return &result
}

func giveSomeUnsetPointerValue(shouldNil bool) TestStruct {
	var result TestStruct
	if !shouldNil {
		result.a = "bla"
	}
	return result
}

func TestOption(t *testing.T) {
	doublePointed := giveSomeUnsetPointerValue(true)
	tests := []struct {
		name       string
		value      interface{}
		shouldSome bool
	}{
		{
			name:       "simple happy path test",
			value:      1,
			shouldSome: true,
		},
		{
			name:       "simple unhappy path test",
			value:      nil,
			shouldSome: false,
		},
		{
			name:       "unhappy strange unset value",
			value:      giveSomeUnsetPointerValue(true),
			shouldSome: false,
		},
		{
			name:       "unhappy strange unset value double pointed",
			value:      &doublePointed,
			shouldSome: false,
		},
		{
			name:       "unset stuff use safe.None()",
			value:      safe.None[string](),
			shouldSome: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			exp := assert.New(t)
			option := safe.Some(&test.value)
			exp.Equal(test.shouldSome, option.IsSome())
			value, ok := option.Some()
			exp.Equal(test.shouldSome, ok)
			if ok {
				exp.Equal(test.value, *value)
			}

		})
	}

}
