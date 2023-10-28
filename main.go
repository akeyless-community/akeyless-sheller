package main

import (
	"fmt"
	"time"

	"github.com/akeyless-community/akeyless-sheller/sheller"
)

func main() {
	// // Define the configuration
	config := sheller.NewConfig(
		"",                            // CLIPath
		"default",                     // Profile
		"/Users/chrisgruel/.akeyless", // AkeylessPath
		10*time.Minute,                // ExpiryBuffer
	)

	// Initialize the sheller library
	err := sheller.InitializeLibrary(config)
	if err != nil {
		fmt.Printf("Failed to initialize sheller library: %v\n", err)
		return
	}

	// Load the specified profile
	profile, err := sheller.GetProfile(config.Profile, config)
	if err != nil {
		fmt.Printf("Failed to load profile: %v\n", err)
		return
	}

	// Get a token for the specified profile
	token, err := sheller.GetToken(profile, config)
	if err != nil {
		fmt.Printf("Failed to get token: %v\n", err)
		return
	}

	// Print the obtained token
	fmt.Printf("Obtained token: %v\n", token.Token)
}
