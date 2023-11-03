package sheller

import (
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	cliPath := "/path/to/cli"
	profile := "testProfile"
	akeylessPath := "/path/to/akeyless"
	expiryBuffer := 10 * time.Minute
	debug := true

	config := NewConfig(cliPath, profile, akeylessPath, expiryBuffer, debug)

	if config.CLIPath != cliPath {
		t.Errorf("Expected CLIPath to be %s, but got %s", cliPath, config.CLIPath)
	}
	if config.Profile != profile {
		t.Errorf("Expected Profile to be %s, but got %s", profile, config.Profile)
	}
	if config.AkeylessPath != akeylessPath {
		t.Errorf("Expected AkeylessPath to be %s, but got %s", akeylessPath, config.AkeylessPath)
	}
	if config.ExpiryBuffer != expiryBuffer {
		t.Errorf("Expected ExpiryBuffer to be %s, but got %s", expiryBuffer, config.ExpiryBuffer)
	}
	if config.Debug != debug {
		t.Errorf("Expected Debug to be %v, but got %v", debug, config.Debug)
	}
}

func TestNewConfigWithDefaults(t *testing.T) {
	// Test cases to be added
}

func TestLoadConfigFromEnv(t *testing.T) {
	// Test cases to be added
}

func TestValidateConfig(t *testing.T) {
	// Test cases to be added
}

func TestValidateAkeylessHomeDirectoryExists(t *testing.T) {
	// Test cases to be added
}

func TestValidateAkeylessCliProfileExists(t *testing.T) {
	// Test cases to be added
}

func TestInitializeLibrary(t *testing.T) {
	// Test cases to be added
}

func TestInitializeAndGetToken(t *testing.T) {
	// Test cases to be added
}
