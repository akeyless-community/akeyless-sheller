package sheller

import (
	"errors"
	"fmt"
	"github.com/spf13/afero"
	"path/filepath"
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
// cliPath: Path to the Akeyless CLI executable
// profile: Name of the Akeyless CLI profile to use
// akeylessPath: Path to the .akeyless directory
// expiryBuffer: Buffer time before token expiry to trigger re-authentication
// debug: Debug flag to enable or disable debug logging
func NewConfig(cliPath, profile, akeylessPath string, expiryBuffer time.Duration, debug bool) *Config {
	return &Config{
		CLIPath:      cliPath,
		Profile:      profile,
		AkeylessPath: akeylessPath,
		ExpiryBuffer: expiryBuffer,
		Debug:        debug,
	}
}

// NewConfigWithDefaults creates a new Config instance with default values.
// It pulls the CLIPath from the system path and uses the "default" CLI profile.
func NewConfigWithDefaults() *Config {
	homeDir, err := afero.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Found Home Directory:", homeDir)
	akeylessHomeDir := filepath.Join(homeDir, ".akeyless")

	return NewConfig("", "default", akeylessHomeDir, 0, false)
}

// LoadConfigFromEnv loads configuration options from environment variables.
// It updates the provided config object with values from the environment.
func LoadConfigFromEnv(config *Config) {
	cliPath := afero.Getenv("AKEYLESS_SHELLER_CLI_PATH")
	if cliPath != "" {
		config.CLIPath = cliPath
	}
	profile := os.Getenv("AKEYLESS_SHELLER_PROFILE")
	if profile != "" {
		config.Profile = profile
	}
	akeylessPath := os.Getenv("AKEYLESS_SHELLER_HOME_DIRECTORY_PATH")
	if akeylessPath != "" {
		config.AkeylessPath = akeylessPath
	}
	expiryBufferStr := os.Getenv("AKEYLESS_SHELLER_EXPIRY_BUFFER")
	if expiryBufferStr != "" {
		expiryBuffer, err := time.ParseDuration(expiryBufferStr)
		if err == nil {
			config.ExpiryBuffer = expiryBuffer
		}
	}
	if config.ExpiryBuffer == 0 {
		config.ExpiryBuffer = DEFAULT_EXPIRY_BUFFER
	}

	debugStr := os.Getenv("AKEYLESS_SHELLER_DEBUG")
	if debugStr != "" {
		config.Debug = true
	}
}

// ValidateConfig validates the provided configuration.
// It checks if the CLIPath, Profile, and AkeylessPath are set correctly and if the files and directories they point to exist and are accessible.
func ValidateConfig(config *Config) error {
	if config.CLIPath == "" {
		akeylessFound := which.Which("akeyless")
		if akeylessFound == "" {
			return errors.New("the CLIPath is not set and akeyless is not in the system path")
		} else {
			config.CLIPath = akeylessFound
		}
	}

	if config.Profile == "" {
		cliProfileExistsError := ValidateAkeylessCliProfileExists(config.AkeylessPath, "default")
		if cliProfileExistsError == nil {
			config.Profile = "default"
		} else {
			return errors.New("the Akeyless CLI Profile name to use is not set and the default profile does not exist")
		}
	}

	// Check if the CLIPath is an executable file
	fileInfo, err := afero.Stat(config.CLIPath)
	if err != nil {
		return err
	}
	if (fileInfo.Mode() & 0111) == 0 {
		return errors.New("the CLIPath does not lead to an executable file")
	}

	// Check if the AkeylessPath property is not empty
	if config.AkeylessPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error:", err)
		}

		if config.Debug {
			fmt.Println("**DEBUG** Home Directory:", homeDir)
		}

		akeylessHomeDir := filepath.Join(homeDir, ".akeyless")

	if err := ValidateAkeylessHomeDirectoryExists(akeylessHomeDir, "default"); err != nil {
		return err
	}
	if config.Debug {
		fmt.Println("**DEBUG** Akeyless Home Directory exists")
	}
	config.AkeylessPath = akeylessHomeDir
	}

	// Check if the AkeylessPath property leads to an existing directory
	// Check if the profiles subdirectory exists inside the AkeylessPath directory
	// Check if the profile file exists and is readable
	if err := ValidateAkeylessCliProfileExists(config.AkeylessPath, config.Profile); err != nil {
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

func ValidateAkeylessHomeDirectoryExists(akeylessHomeDir string, profileName string) error {
	akeylessPathInfo, err := afero.Stat(akeylessHomeDir)
	if err != nil {
		return err
	}
	if !akeylessPathInfo.IsDir() {
		return errors.New("the AkeylessPath does not lead to a directory")
	}

	profilesDirPath := filepath.Join(akeylessHomeDir, "profiles")
	profilesDirInfo, err := afero.Stat(profilesDirPath)
	if err != nil {
		return err
	}
	if !profilesDirInfo.IsDir() {
		return errors.New("the profiles subdirectory does not exist inside the AkeylessPath directory meaning that the AkeylessPath is likely not a valid Akeyless home directory")
	}
	
	return nil
}

func ValidateAkeylessCliProfileExists(profilesDirPath string, profileName string) error {
	profileFilePath := filepath.Join(profilesDirPath, profileName + ".toml")
	_, err := afero.Stat(profileFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("the profile file does not exist")
		}
		if os.IsPermission(err) {
			return errors.New("the profile file is not readable")
		}
		return err
	}
	return nil
}

// InitializeLibrary initializes the Sheller library with the provided configuration.
// It loads the configuration from environment variables and validates it.
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
// It returns the retrieved token or an error if something went wrong.
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
