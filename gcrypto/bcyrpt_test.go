package gcrypto

import (
	"testing"
)

func TestHash(t *testing.T) {
	tests := []struct {
		name, password string
	}{
		{"Short password", "1234"},
		{"32 character password", "abcdefghabcdefghabcdefghabcdefgh"},
		{"64 character password", "abcdefghabcdefghabcdefghabcdefghabcdefghabcdefghabcdefghabcdefgh"},
		{"Empty password", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := Hash(tt.password)
			if err != nil {
				t.Errorf("Hash() error = %v", err)
				return
			}
			if len(hash) == 0 {
				t.Errorf("Fail Hash(). Got an empty string")
			}
		})
	}
}

// TestSuite for CompareHash function
func TestCompareHash(t *testing.T) {
	tests := []struct {
		name, password string
	}{
		{"With valid password", "1234"},
		{"With invalid password", "notmypassword"},
		{"With empty password", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, _ := Hash(tt.password)
			if CompareHash(tt.password, hash) != true {
				t.Errorf("expected CompareHash() to return true, got false")
			}
		})
	}
}
