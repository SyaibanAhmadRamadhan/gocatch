package garray

import (
	"fmt"
	"testing"
)

func TestAppendUniqueVal(t *testing.T) {
	name := make([]int64, 0)
	name = AppendUniqueVal(name, 1, 3, 5, 5)
	fmt.Println(name)

	name2, err := AppendUniqueValWithErr(name, 1, 3, 5, 5)
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
