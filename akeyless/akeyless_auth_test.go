package akeyless

import (
	"os"
	"testing"
)

func TestGetTokenFromAkeylessCommandLine(t *testing.T) {
	// Test case when AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND is not set
	os.Unsetenv("AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND")
	_, err := GetTokenFromAkeylessCommandLine()
	if err == nil {
		t.Errorf("Expected error when AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND is not set, got nil")
	}
}
