package sheller

import (
	"errors"
	"os"
	"time"

	"github.com/hairyhenderson/go-which"
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
func LoadConfigFromEnv(config *Config) {
	cliPath := os.Getenv("SHELLER_CLI_PATH")
	if cliPath != "" {
		config.CLIPath = cliPath
	}
	profile := os.Getenv("SHELLER_PROFILE")
	if profile != "" {
		config.Profile = profile
	}
	akeylessPath := os.Getenv("SHELLER_AKEYLESS_HOME_DIRECTORY_PATH")
	if akeylessPath != "" {
		config.AkeylessPath = akeylessPath
	}
	expiryBufferStr := os.Getenv("SHELLER_EXPIRY_BUFFER")
	if expiryBufferStr != "" {
		expiryBuffer, err := time.ParseDuration(expiryBufferStr)
		if err == nil {
			config.ExpiryBuffer = expiryBuffer
		}
	}
}

// ValidateConfig validates the provided configuration.
func ValidateConfig(config *Config) error {
	if config.CLIPath == "" {
		akeylessFound := which.Which("akeyless")
		if akeylessFound == "" {
			return errors.New("the CLIPath is not set and akeyless is not in the system path")
		} else {
			config.CLIPath = akeylessFound
		}
	}

	// Check if the CLIPath is an executable file
	fileInfo, err := os.Stat(config.CLIPath)
	if err != nil {
		return err
	}
	if (fileInfo.Mode() & 0111) == 0 {
		return errors.New("the CLIPath does not lead to an executable file")
	}

	// TODO: Add more validations here.
	return nil
}

// InitializeLibrary initializes the Sheller library with the provided configuration.
func InitializeLibrary(config *Config) error {
	// Load configuration from environment variables
	LoadConfigFromEnv(config)

	// Validate the configuration.
	err := ValidateConfig(config)
	if err != nil {
		return err
	}

	return nil
}
