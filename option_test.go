package safe_test

import (
	"log"
	"testing"

	"github.com/fasibio/safe"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

type OptionTest[T any] struct {
	Name        string
	InputOption safe.Option[T]
	ShouldSome  bool
	ShouldValue T
}

func TestOption(t *testing.T) {
	tests := []OptionTest[int]{
		{
			Name:        "simple happy path test",
			InputOption: safe.Some(safe.Ptr(1)),
			ShouldSome:  true,
			ShouldValue: 1,
		},
		{
			Name:        "simple unhappy path test",
			InputOption: safe.None[int](),
			ShouldSome:  false,
		},
		{
			Name:        "simple unhappy path test",
			InputOption: safe.Some[int](nil),
			ShouldSome:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			exp := assert.New(t)
			option := test.InputOption
			isSome := option.IsSome()
			exp.Equal(test.ShouldSome, isSome)
			value, ok := option.Some()
			exp.Equal(test.ShouldSome, ok)
			if ok {
				exp.Equal(test.ShouldValue, *value)
			}

		})
	}
}

type OptionDefaultTest[T any] struct {
	OptionTest[T]
	DefaultValue T
}

func TestDefault(t *testing.T) {
	tests := []OptionDefaultTest[int]{
		{
			OptionTest: OptionTest[int]{
				Name:        "simple happy path test",
				InputOption: safe.Some(safe.Ptr(1)),
				ShouldSome:  true,
				ShouldValue: 1,
			},
			DefaultValue: 10,
		},
		{
			OptionTest: OptionTest[int]{
				Name:        "simple unhappy path test use default 1",
				InputOption: safe.None[int](),
				ShouldSome:  false,
				ShouldValue: 10,
			},
			DefaultValue: 10,
		},
		{
			OptionTest: OptionTest[int]{
				Name:        "simple unhappy path test  use default 2",
				InputOption: safe.Some[int](nil),
				ShouldSome:  false,
				ShouldValue: 10,
			},
			DefaultValue: 10,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			exp := assert.New(t)
			option := test.InputOption
			isSome := option.IsSome()
			exp.Equal(test.ShouldSome, isSome)
			value := option.SomeOrDefault(&test.DefaultValue)
			exp.Equal(test.ShouldValue, *value)
		})
	}
}

type MockSomeFn struct {
	mock.Mock
}

func (m *MockSomeFn) SomeOrDefaultFn() *int {
	args := m.Called()
	res := args.Int(0)
	return &res
}

func (m *MockSomeFn) SomeAndThenFn(v *int) {
	m.Called(v)

}

func (m *MockSomeFn) NoneAndThenFn() {
	m.Called()
}

func TestSomeOrDefaultFn(t *testing.T) {
	tests := []struct {
		name            string
		value           *int
		shouldDefaultFn bool
	}{
		{
			name:            "simple happy path test",
			value:           safe.Ptr(1),
			shouldDefaultFn: false,
		},
		{
			name:            "simple unhappy path test",
			value:           nil,
			shouldDefaultFn: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockFn := new(MockSomeFn)
			if test.shouldDefaultFn {
				mockFn.On("SomeOrDefaultFn").Return(10)
			}
			option := safe.Some(test.value)
			option.SomeOrDefaultFn(mockFn.SomeOrDefaultFn)
			if test.shouldDefaultFn {
				mockFn.AssertExpectations(t)
			}
		})
	}
}

func TestSomeAndThenFn(t *testing.T) {
	tests := []struct {
		name                string
		value               *int
		shouldSomeAndThanFn bool
	}{
		{
			name:                "simple happy path test",
			value:               safe.Ptr(1),
			shouldSomeAndThanFn: true,
		},
		{
			name:                "simple unhappy path test",
			value:               nil,
			shouldSomeAndThanFn: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockFn := new(MockSomeFn)
			if test.shouldSomeAndThanFn {
				mockFn.On("SomeAndThenFn", test.value)
			}
			option := safe.Some(test.value)
			option.SomeAndThen(mockFn.SomeAndThenFn)
			if test.shouldSomeAndThanFn {
				mockFn.AssertExpectations(t)
			}
		})
	}
}

func TestNoneAndThenFn(t *testing.T) {
	tests := []struct {
		name                string
		value               *int
		shouldNoneAndThanFn bool
	}{
		{
			name:                "simple happy path test",
			value:               safe.Ptr(1),
			shouldNoneAndThanFn: false,
		},
		{
			name:                "simple unhappy path test",
			value:               nil,
			shouldNoneAndThanFn: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockFn := new(MockSomeFn)
			if test.shouldNoneAndThanFn {
				mockFn.On("NoneAndThenFn")
			}
			option := safe.Some(test.value)
			option.NoneAndThen(mockFn.NoneAndThenFn)
			if test.shouldNoneAndThanFn {
				mockFn.AssertExpectations(t)
			}
		})
	}
}

func TestCopyOrDefault(t *testing.T) {
	tests := []struct {
		name                string
		value               *int
		defaultValue        int
		shouldDefaultReturn bool
	}{
		{
			name:                "simple happy path test",
			value:               safe.Ptr(1),
			defaultValue:        10,
			shouldDefaultReturn: false,
		},
		{
			name:                "simple unhappy path test",
			value:               nil,
			defaultValue:        10,
			shouldDefaultReturn: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)
			option := safe.Some(test.value)
			res := option.CopyOrDefault(test.defaultValue)
			if test.shouldDefaultReturn {
				assert.Equal(test.defaultValue, res)
				res = test.defaultValue - 2
				assert.NotEqual(test.defaultValue, res)
			} else {
				assert.Equal(*test.value, res)
				res = *test.value - 2
				assert.NotEqual(*test.value, res)
			}
		})
	}
}
