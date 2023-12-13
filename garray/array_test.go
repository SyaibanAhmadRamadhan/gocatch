package garray

import (
	"fmt"
	"testing"
)

func TestAppendUniqueVal(t *testing.T) {
	name := []string{
		"rama",
	}
	name = AppendUniqueVal(name, "rama4", "rama", "rama2", "rama0")
	fmt.Println(name)

	name2, err := AppendUniqueValWithErr(name, "rama3", "rama123", "rama212345")
	fmt.Println(name2)
	fmt.Println(err)
}
