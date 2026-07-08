package safe

import (
	"fmt"
	"testing"

	"github.com/moznion/go-optional"
)

type Complex struct {
	A int
}

func TestMoznion(t *testing.T) {
	res, err := optional.Some[*Complex](nil).Take()
	fmt.Println(res, err)

}
