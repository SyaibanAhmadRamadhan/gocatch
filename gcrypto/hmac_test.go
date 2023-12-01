package gcrypto

import (
	"testing"
)

func TestGenerateHMAC256(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		key  []byte
	}{
		{"With data and key", []byte("data"), []byte("key")},
		{"Empty data", []byte(""), []byte("key")},
		{"Empty key", []byte("data"), []byte("")},
		{"Large data and key", []byte("longData"), []byte("longKey")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hmac := GenerateHMAC256(tt.data, tt.key)
			if len(hmac) != 64 { // The length of an SHA-256 HMAC in hexadecimal form should always be 64
				t.Errorf("expected len(GenerateHMAC256()) to be 64, got %d", len(hmac))
			}
		})
	}
}

func TestCompareHMAC256(t *testing.T) {
	tests := []struct {
		name          string
		data          []byte
		key           []byte
		hmacToCompare string
		want          bool
	}{
		{"Valid comparison", []byte("data"), []byte("key"), "", true},
		{"Invalid comparison", []byte("data"), []byte("key"), "wrongHmac", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computedHmac := GenerateHMAC256(tt.data, tt.key)
			if tt.hmacToCompare == "" {
				tt.hmacToCompare = computedHmac
			}
			if got := CompareHMAC256(computedHmac, tt.hmacToCompare); got != tt.want {
				t.Errorf("CompareHMAC256() = %v, want %v", got, tt.want)
			}
		})
	}
}
