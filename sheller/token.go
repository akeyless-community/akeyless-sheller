package sheller

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/pelletier/go-toml"
)

// Token holds the details of an authentication token.
type Token struct {
	AccessID string    `json:"access_id"`
	Token    string    `json:"token"`
	Expiry   time.Time `json:"expiry"`
}

// CheckForExistingToken checks for an existing valid token for the specified profile.
func CheckForExistingToken(profile *Profile, config *Config) (*Token, error) {
	tokenFilesPath := filepath.Join(config.AkeylessPath, ".tmp_creds")
	files, err := os.ReadDir(tokenFilesPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == "" { // Assuming token files have no extension
			token, err := ParseTokenFile(filepath.Join(tokenFilesPath, file.Name()))
			if err != nil {
				return nil, err
			}

			// Check if the token's AccessID matches the profile's AccessID and the token is not expired
			if token.AccessID == profile.AccessID && token.Expiry.After(time.Now().Add(config.ExpiryBuffer)) {
				return token, nil
			}
		}
	}

	return nil, errors.New("no valid token found")
}

// ParseTokenFile parses a token file and returns a Token struct.
func ParseTokenFile(path string) (*Token, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	token := &Token{}
	err = json.Unmarshal(data, token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// convertUnderscoresToHyphens converts underscores to hyphens in a string.
func convertUnderscoresToHyphens(s string) string {
	return strings.ReplaceAll(s, "_", "-")
}

// ShellOutForNewToken shells out to the Akeyless CLI to obtain a new token for the specified profile.
func ShellOutForNewToken(profile *Profile, config *Config) (*Token, error) {
	// Load the profile configuration file
	profilePath := filepath.Join(config.AkeylessPath, "profiles", fmt.Sprintf("%s.toml", profile.Name))
	profileConfig, err := toml.LoadFile(profilePath)
	if err != nil {
		return nil, err
	}

	// Construct the command string based on the profile and config
	cmdStrBuilder := strings.Builder{}
	fmt.Fprintf(&cmdStrBuilder, "%s auth", filepath.Join(config.CLIPath, "akeyless"))

	// Iterate through the profile configuration and append properties to the command string
	profileConfigTree := profileConfig.Get(profile.Name).(*toml.Tree)
	for _, key := range profileConfigTree.Keys() {
		value := profileConfigTree.Get(key).(string)
		cmdKey := convertUnderscoresToHyphens(key)
		fmt.Fprintf(&cmdStrBuilder, " --%s %s", cmdKey, value)
	}

	cmdStr := cmdStrBuilder.String()

	cmdParts := strings.Fields(cmdStr)

	// Check if the path points to an executable file
	if _, err := os.Stat(cmdParts[0]); os.IsNotExist(err) {
		return nil, errors.New("the path does not point to an executable file")
	}

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	tokenCode := strings.TrimSpace(string(output))

	token := &Token{
		AccessID: profile.AccessID,
		Token:    tokenCode,
		Expiry:   time.Now().Add(1 * time.Hour), // Assuming token expiry is 1 hour
	}

	return token, nil
}

// GetToken retrieves a token for the specified profile, either by reusing an existing valid token or by shelling out to the Akeyless CLI.
func GetToken(profile *Profile, config *Config) (*Token, error) {
	token, err := CheckForExistingToken(profile, config)
	if err == nil {
		return token, nil
	}

	// If no valid token found, shell out for a new token
	return ShellOutForNewToken(profile, config)
}
