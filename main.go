package main

import (
	"fmt"

	"github.com/akeyless-community/akeyless-sheller/sheller"
)

func main() {
	// Define the configuration
	config := sheller.NewConfigWithDefaults()

	token, err := sheller.InitializeAndGetToken(config)
	if err != nil {
		fmt.Printf("Failed to initialize and get token: %v\n", err)
		return
	}

	// Print the obtained token
	fmt.Printf("Obtained token: %v\n", token.Token)
}
