package gstruct

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	AField     string `tag:"AFieldTag"`
	BField     string `tag:"-"`
	AnotherTag string `anotherTag:"TagValue"`
	Ignore     string `tag:"ignore"`
}

func TestExtractStructTagsAndFields(t *testing.T) {
	var ts TestStruct
	m := ExtractStructTagsAndFields(ts, "", "tag")

	expected := map[string]string{
		"AField": "AFieldTag|string",
	}

	if !reflect.DeepEqual(m, expected) {
		t.Errorf("Expected: %v, got: %v", expected, m)
	}
}
