package main

import (
	"testing"
)

func TestGetTokenFromAkeylessCommandLine(t *testing.T) {
	// // Test case when AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND is not set
	// os.Unsetenv("AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND")
	// _, err := GetTokenFromAkeylessCommandLine()
	// if err == nil {
	// 	t.Errorf("Expected error when AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND is not set, got nil")
	// }

	// // Test case when AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND is set to a valid command
	// os.Setenv("AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND", "echo test")
	// token, err := GetTokenFromAkeylessCommandLine()
	// if err != nil || token != "test" {
	// 	t.Errorf("Expected token 'test', got '%s', error: %v", token, err)
	// }

	// // Test case when AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND is set to an invalid command
	// os.Setenv("AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND", "invalid_command")
	// _, err = GetTokenFromAkeylessCommandLine()
	// if err == nil {
	// 	t.Errorf("Expected error when AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND is set to an invalid command, got nil")
	// }

	// // Test case when AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND is set to a path that does not point to an executable file
	// os.Setenv("AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND", "/path/that/does/not/point/to/an/executable")
	// _, err = GetTokenFromAkeylessCommandLine()
	// if err == nil {
	// 	t.Errorf("Expected error when AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND is set to a path that does not point to an executable file, got nil")
	// }
}
