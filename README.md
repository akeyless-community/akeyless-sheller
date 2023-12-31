# Akeyless Sheller

This Go library streamlines the process of handling authentication and managing tokens when interacting with the Akeyless CLI (Command Line Interface). It automatically manages tokens by reusing valid ones and only prompts for re-authentication when needed, minimizing interruptions. This not only simplifies the user experience but also delegates authentication to the Akeyless CLI, ensuring top-notch security and compatibility with various authentication methods.

## Getting Started

1. **Install Library:**

```bash
go get github.com/akeyless-community/akeyless-sheller/sheller
```

2. **Environment Variables:**

The library can be configured using the following environment variables:

- `AKEYLESS_SHELLER_CLI_PATH`: Path to the Akeyless CLI executable
- `AKEYLESS_SHELLER_PROFILE`: Name of the Akeyless CLI profile to use
- `AKEYLESS_SHELLER_HOME_DIRECTORY_PATH`: Path to the .akeyless directory
- `AKEYLESS_SHELLER_EXPIRY_BUFFER`: Buffer time before token expiry to trigger re-authentication (in Go duration format, e.g., "10m" for 10 minutes)
- `AKEYLESS_SHELLER_DEBUG`: Debug flag to enable or disable debug logging (set to any value to enable)

## Sequence Diagram

The following sequence diagram provides an overview of how the `sheller` library operates:

```mermaid
sequenceDiagram
    participant Client as Client Code
    participant Sheller as Sheller Library
    participant ProfileManager as Sheller: Profile Manager
    participant TokenManager as Sheller: Token Manager
    participant CLI as Akeyless CLI

    Client->>Sheller: InitializeLibrary(config)
    Sheller->>ProfileManager: GetProfile(config.profileName)
    ProfileManager->>Sheller: Return profile
    Client->>Sheller: GetToken(profile)
    Sheller->>TokenManager: Check for existing token
    alt Token exists and is valid
        TokenManager->>Sheller: Return existing token
    else Token doesn't exist or is invalid
        TokenManager->>CLI: ShellOutForNewToken(profile)
        CLI->>TokenManager: Return new token
    end
    TokenManager->>Sheller: Return token
    Sheller->>Client: Return token
```

## Example Quickstart

### Prerequisites

- Akeyless CLI is installed, executable and accessible through the system path
- Akeyless CLI has a `default` profile configured
- The Akeyless home directory exists at `~/.akeyless/` (created through the Akeyless CLI initialization)

### Quickstart Code Example

The following code snippet provides a quickstart example of how to use the `sheller` library to obtain a token for a specified profile:

```go
package main

import (
    "fmt"
    "time"

    "github.com/akeyless-community/akeyless-sheller/sheller"
)

func main() {
    // Define the configuration
    config := sheller.NewConfigWithDefaults()

    // Initialize the sheller library and get the token
    token, err := sheller.InitializeAndGetToken(config)
    if err != nil {
        fmt.Printf("Failed to initialize and get token: %v\n", err)
        return
    }

    // Print the obtained token
    fmt.Printf("Obtained token: %v\n", token.Token)
}

```

## Example Full Implementation

### Full Implementation Explanation

The `main.go` file in the root directory serves as an example quickstart implementation of the `sheller` library. Below is a brief explanation of how the full example operates:

1. Define the configuration using `sheller.NewConfig`.
2. Initialize the `sheller` library using `sheller.InitializeLibrary`.
3. Load the specified profile using `sheller.GetProfile`.
4. Obtain a token for the specified profile using `sheller.GetToken`.
5. Print the obtained token to the console.

### Full Implementation Code Example

```go
package main

import (
    "fmt"
    "sheller"
    "time"

    "github.com/akeyless-community/akeyless-sheller/sheller"
)

func main() {
    // Define the configuration
    config := sheller.NewConfig(
        "", // will pull from path if akeyless CLI is on system path
        "default", // the name of the CLI profile to use
        "/Users/chrisgruel/.akeyless", // the path to the Akeyless home directory usually located at ~/.akeyless
        10*time.Minute, // if the token expiration is within this upcoming interval, a new token will be obtained
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

- `sheller/config.go`: Configuration Manager: Defines the configuration structure and provides a function to initialize the library.
- `sheller/profile.go`: Profile Manager: Provides functions to load and list Akeyless CLI profiles.
- `sheller/token.go`: Token Manager: Provides functions to check for existing tokens, shell out for new tokens, and retrieve tokens for specified profiles.

## Testing

To run the provided tests, use the following command:

```bash
go test ./...
```

## License

This project is licensed under the terms of the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
