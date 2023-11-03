package sheller

import (
	"os"
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
	// Test case 1: Valid configuration
	config1 := NewConfig("/path/to/cli", "testProfile", "/path/to/akeyless", 10*time.Minute, true)
	err := ValidateConfig(config1)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Test case 2: Invalid CLIPath
	config2 := NewConfig("", "testProfile", "/path/to/akeyless", 10*time.Minute, true)
	err = ValidateConfig(config2)
	if err == nil {
		t.Errorf("Expected error, but got none")
	}

	// Test case 3: Invalid Profile
	config3 := NewConfig("/path/to/cli", "", "/path/to/akeyless", 10*time.Minute, true)
	err = ValidateConfig(config3)
	if err == nil {
		t.Errorf("Expected error, but got none")
	}

	// Test case 4: Invalid AkeylessPath
	config4 := NewConfig("/path/to/cli", "testProfile", "", 10*time.Minute, true)
	err = ValidateConfig(config4)
	if err == nil {
		t.Errorf("Expected error, but got none")
	}
}

func TestValidateAkeylessHomeDirectoryExists(t *testing.T) {
	// Test case 1: Valid Akeyless home directory
	akeylessHomeDir1 := "/path/to/valid/akeyless/home/directory"
	err := ValidateAkeylessHomeDirectoryExists(akeylessHomeDir1, "default")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Test case 2: Invalid Akeyless home directory (directory does not exist)
	akeylessHomeDir2 := "/path/to/nonexistent/directory"
	err = ValidateAkeylessHomeDirectoryExists(akeylessHomeDir2, "default")
	if err == nil {
		t.Errorf("Expected error, but got none")
	}

	// Test case 3: Invalid Akeyless home directory (not a directory)
	akeylessHomeDir3 := "/path/to/file/not/directory"
	err = ValidateAkeylessHomeDirectoryExists(akeylessHomeDir3, "default")
	if err == nil {
		t.Errorf("Expected error, but got none")
	}
}

func TestValidateAkeylessCliProfileExists(t *testing.T) {
	// Test case 1: Valid profile file
	profilesDir1 := "/path/to/valid/profiles/directory"
	profileName1 := "validProfile"
	err := ValidateAkeylessCliProfileExists(profilesDir1, profileName1)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Test case 2: Profile file does not exist
	profilesDir2 := "/path/to/valid/profiles/directory"
	profileName2 := "nonexistentProfile"
	err = ValidateAkeylessCliProfileExists(profilesDir2, profileName2)
	if err == nil {
		t.Errorf("Expected error, but got none")
	}

	// Test case 3: Profile file is not readable
	profilesDir3 := "/path/to/valid/profiles/directory"
	profileName3 := "unreadableProfile"
	err = ValidateAkeylessCliProfileExists(profilesDir3, profileName3)
	if err == nil {
		t.Errorf("Expected error, but got none")
	}
}

func TestInitializeLibrary(t *testing.T) {
	// Test cases to be added
}

func TestInitializeAndGetToken(t *testing.T) {
	// Test cases to be added
}
