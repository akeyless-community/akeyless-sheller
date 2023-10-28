package sheller

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/hairyhenderson/go-which"
)

var DEFAULT_EXPIRY_BUFFER = 10 * time.Minute

// Config holds the configuration options for the Sheller library.
type Config struct {
	CLIPath      string        // Path to the Akeyless CLI executable
	Profile      string        // Name of the Akeyless CLI profile to use
	AkeylessPath string        // Path to the .akeyless directory
	ExpiryBuffer time.Duration // Buffer time before token expiry to trigger re-authentication
	Debug        bool          // Debug flag to enable or disable debug logging
}

// NewConfig creates a new Config instance with the provided parameters.
func NewConfig(cliPath, profile, akeylessPath string, expiryBuffer time.Duration, debug bool) *Config {
	return &Config{
		CLIPath:      cliPath,
		Profile:      profile,
		AkeylessPath: akeylessPath,
		ExpiryBuffer: expiryBuffer,
		Debug:        debug,
	}
}

// NewConfigWithDefaults creates a new Config instance with default values like
// pulling the CLIPath from the system path and using the "default" CLI profile
func NewConfigWithDefaults() *Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Found Home Directory:", homeDir)
	akeylessHomeDir := homeDir + "/.akeyless"
	return NewConfig("", "default", akeylessHomeDir, 0, false)
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
	if config.ExpiryBuffer == 0 {
		config.ExpiryBuffer = DEFAULT_EXPIRY_BUFFER
	}

	debugStr := os.Getenv("SHELLER_DEBUG")
	if debugStr != "" {
		config.Debug = true
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

	// Check if the AkeylessPath property is not empty
	if config.AkeylessPath == "" {
		return errors.New("the AkeylessPath is not set")
	}

	// Check if the AkeylessPath property leads to an existing directory
	akeylessPathInfo, err := os.Stat(config.AkeylessPath)
	if err != nil {
		return err
	}
	if !akeylessPathInfo.IsDir() {
		return errors.New("the AkeylessPath does not lead to a directory")
	}

	// Check if the profiles subdirectory exists inside the AkeylessPath directory
	profilesDirPath := config.AkeylessPath + "/profiles"
	profilesDirInfo, err := os.Stat(profilesDirPath)
	if err != nil {
		return err
	}
	if !profilesDirInfo.IsDir() {
		return errors.New("the profiles subdirectory does not exist inside the AkeylessPath directory meaning that the AkeylessPath is likely not a valid Akeyless home directory")
	}

	// Check if the profile file exists and is readable
	profileFilePath := profilesDirPath + "/" + config.Profile + ".toml"
	_, err = os.Stat(profileFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("the profile file does not exist")
		}
		if os.IsPermission(err) {
			return errors.New("the profile file is not readable")
		}
		return err
	}

	if config.Debug {
		fmt.Println("**DEBUG** Loaded configuration:")
		fmt.Println("CLIPath:", config.CLIPath)
		fmt.Println("Profile:", config.Profile)
		fmt.Println("AkeylessPath:", config.AkeylessPath)
		fmt.Println("ExpiryBuffer:", config.ExpiryBuffer)
		fmt.Println("Debug:", config.Debug)
	}

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

// InitializeAndGetToken initializes the library, gets the profile, and retrieves the token.
func InitializeAndGetToken(config *Config) (*Token, error) {
	err := InitializeLibrary(config)
	if err != nil {
		return nil, err
	}
	profile, errProfile := GetProfile(config.Profile, config)
	if errProfile != nil {
		return nil, errProfile
	}
	token, err := GetToken(profile, config)
	if err != nil {
		return nil, err
	}
	return token, nil
}
