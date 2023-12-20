package gdb

import (
	"github.com/SyaibanAhmadRamadhan/gocatch/gtypedata/garray"
)

type PaginationParam struct {
	Limit  int64
	Offset int64
}

type FindByOneColumnParam struct {
	Column string
	Value  any
}

type OrderBy struct {
	Column      string
	IsAscending bool
}

type OrderByParams []OrderBy

func (p OrderByParams) FilterDifferent(refColumns []string) (res OrderByParams) {
	for _, v := range p {
		if garray.Contains(refColumns, v.Column) {
			res = append(res, v)
		}
	}

	return
}
