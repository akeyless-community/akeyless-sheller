# Akeyless Sheller

A Go library for simplifying the process of authenticating and managing tokens with the Akeyless CLI. This library handles token retrieval and management, ensuring tokens are reused when valid and only re-authenticating when necessary to reduce unnecessary user prompts.

## Directory Structure

```plaintext
.
├── LICENSE
├── README.md
├── akeyless-sheller
├── go.mod
├── go.sum
├── main.go           # Example implementation of the library
├── main_test.go
└── sheller
    ├── config.go
    ├── profile.go
    └── token.go
```

## Getting Started

1. **Clone the Repository:**

```bash
git clone https://github.com/your-username/akeyless-sheller.git
cd akeyless-sheller
```

2. **Install Dependencies:**

```bash
go mod tidy
```

3. **Run the Example:**

```bash
go run main.go
```

## Example Implementation

The `main.go` file in the root directory serves as an example implementation of the `sheller` library. Below is a brief explanation of how it operates:

1. Define the configuration using `sheller.NewConfig`.
2. Initialize the `sheller` library using `sheller.InitializeLibrary`.
3. Load the specified profile using `sheller.GetProfile`.
4. Obtain a token for the specified profile using `sheller.GetToken`.
5. Print the obtained token to the console.

```go
package main

import (
    "fmt"
    "sheller"
    "time"
)

func main() {
    // Define the configuration
    config := sheller.NewConfig(
        "/path/to/akeyless-cli",
        "default",
        "/Users/chrisgruel/.akeyless",
        10*time.Minute,
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
```

## Library Structure

- `sheller/config.go`: Defines the configuration structure and provides a function to initialize the library.
- `sheller/profile.go`: Provides functions to load and list Akeyless CLI profiles.
- `sheller/token.go`: Provides functions to check for existing tokens, shell out for new tokens, and retrieve tokens for specified profiles.

## Testing

To run the provided tests, use the following command:

```bash
go test ./...
```

## License

This project is licensed under the terms of the MIT license. See the [LICENSE](LICENSE) file for details.