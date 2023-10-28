package sheller

import (
	"os"
	"time"
)

// Config holds the configuration options for the Sheller library.
type Config struct {
	CLIPath      string        // Path to the Akeyless CLI executable
	Profile      string        // Name of the Akeyless CLI profile to use
	AkeylessPath string        // Path to the .akeyless directory
	ExpiryBuffer time.Duration // Buffer time before token expiry to trigger re-authentication
}

// NewConfig creates a new Config instance with the provided parameters.
func NewConfig(cliPath, profile, akeylessPath string, expiryBuffer time.Duration) *Config {
	return &Config{
		CLIPath:      cliPath,
		Profile:      profile,
		AkeylessPath: akeylessPath,
		ExpiryBuffer: expiryBuffer,
	}
}

// LoadConfigFromEnv loads configuration options from environment variables.
func LoadConfigFromEnv() *Config {
	cliPath := os.Getenv("SHELLER_CLI_PATH")
	profile := os.Getenv("SHELLER_PROFILE")
	akeylessPath := os.Getenv("SHELLER_AKEYLESS_HOME_DIRECTORY_PATH")
	expiryBufferStr := os.Getenv("SHELLER_EXPIRY_BUFFER")

	// Convert expiryBufferStr to time.Duration, with a default value
	expiryBuffer, err := time.ParseDuration(expiryBufferStr)
	if err != nil {
		expiryBuffer = 10 * time.Minute // Default value of 10 minutes
	}

	return NewConfig(cliPath, profile, akeylessPath, expiryBuffer)
}

// InitializeLibrary initializes the Sheller library with the provided configuration.
func InitializeLibrary(config *Config) error {
	// TODO: Implement initialization logic based on the provided configuration.
	// This might include validating the configuration, setting up logging, etc.
	return nil
}
