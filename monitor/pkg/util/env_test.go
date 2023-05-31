package util

import (
	"os"
	"testing"
)

func TestGetEnvWithDefault(t *testing.T) {
	if GetEnvWithDefault("NON_EXISTENT_KEY", "test") != "test" {
		t.Fatal("Default is not working")
	}

	_ = os.Setenv("KEY", "value")
	if GetEnvWithDefault("KEY", "foo") != "value" {
		t.Fatal("Retrieving env vars is not working as expected")
	}
}
