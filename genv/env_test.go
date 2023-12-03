package genv

import (
	"os"
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
)

func TestInitialize(t *testing.T) {
	defer func() {
		err := os.Remove("./.env")
		gcommon.PanicIfError(err)

		err = os.Remove("./.env.override")
		gcommon.PanicIfError(err)
	}()

	// Create .env and .env.override files for testing
	_ = os.WriteFile("./.env", []byte("TEST_ENV=initial"), 0644)
	_ = os.WriteFile("./.env.override", []byte("TEST_ENV=override"), 0644)

	// Test godotEnv without override
	err := Initialize(DefaultEnvLib, false)
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	// Test godotEnv with override
	err = Initialize(DefaultEnvLib, true)
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	if os.Getenv("TEST_ENV") != "override" {
		t.Errorf("Expected TEST_ENV to be 'override', but got: %s", os.Getenv("TEST_ENV"))
	}
}
