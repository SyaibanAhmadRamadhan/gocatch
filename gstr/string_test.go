package gstr

import (
	"reflect"
	"testing"
)

func TestGetLastSubstring(t *testing.T) {
	if got := GetLastSubstring("hello.world", "."); got != "world" {
		t.Errorf("GetLastSubstring() = %q, want %q", got, "world")
	}

	if got := GetLastSubstring("hello.world.com", "."); got != "com" {
		t.Errorf("GetLastSubstring() = %q, want %q", got, "com")
	}

	if got := GetLastSubstring("no.separator.here", "-"); got != "no.separator.here" {
		t.Errorf("GetLastSubstring() = %q, want %q", got, "no.separator.here")
	}
}

func TestLowercaseFirstChar(t *testing.T) {
	if got := LowercaseFirstChar("Hello"); got != "h" {
		t.Errorf("LowercaseFirstChar() = %q, want %q", got, "h")
	}

	if got := LowercaseFirstChar("HELLO"); got != "h" {
		t.Errorf("LowercaseFirstChar() = %q, want %q", got, "h")
	}
}

func TestToLowercase(t *testing.T) {
	if got := ToLowercase("HELLO"); got != "hello" {
		t.Errorf("ToLowercase() = %q, want %q", got, "hello")
	}

	if got := ToLowercase("HeLLo"); got != "hello" {
		t.Errorf("ToLowercase() = %q, want %q", got, "hello")
	}
}

func TestToUppercase(t *testing.T) {
	if got := ToUppercase("hello"); got != "HELLO" {
		t.Errorf("ToUppercase() = %q, want %q", got, "HELLO")
	}

	if got := ToUppercase("HeLLo"); got != "HELLO" {
		t.Errorf("ToUppercase() = %q, want %q", got, "HELLO")
	}
}

func TestSplitByWhiteSpace(t *testing.T) {
	result := SplitByWhiteSpace("Hello, World! This is a test string.")
	expected := []string{"Hello,", "World!", "This", "is", "a", "test", "string."}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestRepeatString(t *testing.T) {
	result := RepeatString("abc", 3)
	expected := "abcabcabc"
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestTrimSpace(t *testing.T) {
	result := TrimSpace(`   Hello, World!    `)
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCountSubstring(t *testing.T) {
	result := CountSubstring("Hello, World! Hello again!", "Hello")
	expected := 2
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
