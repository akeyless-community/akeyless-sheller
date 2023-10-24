package akeyless

import (
	"os"
	"os/exec"
	"strings"
)

// GetTokenFromAkeylessCommandLine shells out to the Akeyless CLI to handle authentication and return a token.
func GetTokenFromAkeylessCommandLine() (string, error) {
	cmdStr := os.Getenv("AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND")
	if cmdStr == "" {
		return "", errors.New("AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND environment variable is not set")
	}

	cmdParts := strings.Fields(cmdStr)
	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	token := strings.TrimSpace(string(output))
	return token, nil
}
