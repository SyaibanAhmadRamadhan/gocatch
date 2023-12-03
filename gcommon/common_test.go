package gcommon

import (
	"errors"
	"testing"
)

func TestPanicIfError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	PanicIfError(errors.New("test error"))
}
