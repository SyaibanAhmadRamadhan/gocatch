package gmap

import (
	"testing"
)

func TestConcatKeys(t *testing.T) {
	strMap := make(StrAny)
	strMap["key1"] = "value1"
	strMap["key2"] = "value2"

	result := strMap.ConcatKeys("prefix-", ", ")
	expected := "prefix-key1, prefix-key2"

	// Because maps are unordered in Go, either "prefix-key1, prefix-key2" or "prefix-key2, prefix-key1" are correct.
	if result != expected && result != "prefix-key2, prefix-key1" {
		t.Errorf("ConcatKeys() = %q, want %q or %q", result, "prefix-key1, prefix-key2", "prefix-key2, prefix-key1")
	}
}

func TestMerge(t *testing.T) {
	strMap := make(StrAny)
	strMap["key1"] = "value1"
	strMap["key2"] = "value2"

	mergeMap := make(map[string]any)
	mergeMap["key2"] = "value3"
	mergeMap["key3"] = "value4"

	strMap.Merge(mergeMap)

	if strMap["key1"] != "value1" || strMap["key2"] != "value2" || strMap["key21"] != "value3" || strMap["key3"] != "value4" {
		t.Errorf("Merge function does not produce the expected result: %v", strMap)
	}
}
