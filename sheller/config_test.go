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
	config := NewConfigWithDefaults()

	if config.Profile != "default" {
		t.Errorf("Expected Profile to be 'default', but got %s", config.Profile)
	}
	if config.ExpiryBuffer != 0 {
		t.Errorf("Expected ExpiryBuffer to be 0, but got %s", config.ExpiryBuffer)
	}
	if config.Debug != false {
		t.Errorf("Expected Debug to be false, but got %v", config.Debug)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("AKEYLESS_SHELLER_CLI_PATH", "/path/to/cli")
	os.Setenv("AKEYLESS_SHELLER_PROFILE", "testProfile")
	os.Setenv("AKEYLESS_SHELLER_HOME_DIRECTORY_PATH", "/path/to/akeyless")
	os.Setenv("AKEYLESS_SHELLER_EXPIRY_BUFFER", "10m")
	os.Setenv("AKEYLESS_SHELLER_DEBUG", "true")

	config := NewConfig("", "", "", 0, false)
	LoadConfigFromEnv(config)

	if config.CLIPath != "/path/to/cli" {
		t.Errorf("Expected CLIPath to be '/path/to/cli', but got %s", config.CLIPath)
	}
	if config.Profile != "testProfile" {
		t.Errorf("Expected Profile to be 'testProfile', but got %s", config.Profile)
	}
	if config.AkeylessPath != "/path/to/akeyless" {
		t.Errorf("Expected AkeylessPath to be '/path/to/akeyless', but got %s", config.AkeylessPath)
	}
	if config.ExpiryBuffer != 10*time.Minute {
		t.Errorf("Expected ExpiryBuffer to be 10m, but got %s", config.ExpiryBuffer)
	}
	if config.Debug != true {
		t.Errorf("Expected Debug to be true, but got %v", config.Debug)
	}
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
