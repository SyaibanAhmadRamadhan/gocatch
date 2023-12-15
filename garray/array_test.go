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

func TestFilterDifferentElem(t *testing.T) {
	tables := []struct {
		source     []string
		refference []string
		expec      []string
	}{
		{
			source:     []string{"a", "b", "c", "d"},
			refference: []string{"a", "x", "c", "y"},
			expec:      []string{"a", "c"},
		},
		{
			source:     []string{"a", "b", "y", "d"},
			refference: []string{"z", "y"},
			expec:      []string{"y"},
		},
		{
			source:     []string{"a", "b"},
			refference: []string{"a", "x", "c", "y"},
			expec:      []string{"a"},
		},
	}

	for _, table := range tables {
		res := FilterDifferentElem(table.source, table.refference)
		if !SlicesMatch(res, table.expec) {
			t.Errorf("expec and actual is not same. actual: %v, expected: %v", res, table.expec)
		}

	}
}
