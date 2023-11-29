package Jtype

import (
	"testing"
)

func name() string {
	return "rama"
}

func TestTernary(t *testing.T) {
	nameRes := Ternary(name() == "", name(), "no name")
	t.Log(nameRes)
}
