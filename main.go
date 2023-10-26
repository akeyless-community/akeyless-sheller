package sheller

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

// GetTokenFromAkeylessCommandLine shells out to the Akeyless CLI to handle authentication and return a token.
func GetTokenFromAkeylessCommandLine() (string, error) {
	cmdStr := os.Getenv("AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND")
	if cmdStr == "" {
		return "", errors.New("the AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND environment variable is not set")
	}

	cmdParts := strings.Fields(cmdStr)

	// Check if the path points to an executable file
	if _, err := os.Stat(cmdParts[0]); os.IsNotExist(err) {
		return "", errors.New("the path does not point to an executable file")
	}

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	token := strings.TrimSpace(string(output))
	return token, nil
}
