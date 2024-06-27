package safe_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/fasibio/safe"
	"github.com/stretchr/testify/assert"
)

func TestResultExample(t *testing.T) {
	exp := assert.New(t)
	result := safe.NewResult(somePossibleFailedFunc())
	exp.True(result.IsOk())
	if result.IsOk() {
		log.Println("is no error")
		//happy path
	} else {
		// not ok
		err := result.Err()
		log.Println(err)
	}
}

func TestCombined(t *testing.T) {
	//or combined

	result := safe.NewResult(somePossibleFailedOptinalFunc(false))
	i := -1
	value := result.OkOrDefault(safe.Some(&i)).SomeOrDefault(safe.GetPtr(-1))
	assert.Equal(t, safe.GetPtr(0), value)

	result = safe.NewResult(somePossibleFailedOptinalFunc(true))
	value = result.OkOrDefault(safe.Some(&i)).SomeOrDefault(safe.GetPtr(-1))
	assert.Equal(t, safe.GetPtr(-1), value)
}

func somePossibleFailedOptinalFunc(shoudErr bool) (safe.Option[int], error) {
	if shoudErr {
		return safe.None[int](), fmt.Errorf("Some error")
	}
	return safe.Some(safe.GetPtr(0)), nil
}

func somePossibleFailedFunc() (int, error) {
	return 0, nil
}

func TestResult(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		err      error
		shouldOk bool
	}{
		{
			name:     "Simple happy test",
			value:    1,
			err:      nil,
			shouldOk: true,
		},
		{
			name:     "is error",
			value:    nil,
			err:      fmt.Errorf("Some error"),
			shouldOk: false,
		},
		{
			name:     "error is nil but value as well so not ok ",
			value:    nil,
			err:      nil,
			shouldOk: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			exp := assert.New(t)

			res := safe.NewResult(test.value, test.err)
			exp.Equal(test.shouldOk, res.IsOk())
			exp.Equal(!test.shouldOk, res.IsErr())
			if test.shouldOk {
				exp.Equal(test.value, res.Ok())
			} else {
				exp.Error(res.Err())
			}

		})
	}
}
