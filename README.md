# Akeyless Sheller

Akeyless Sheller is a Golang library designed to shell out to the Akeyless CLI to retrieve a token. This token can then be used in multiple tools such as CLI helpers and Terraform.

## Prerequisites

Before you begin, ensure you have met the following requirements:
* You have installed the latest version of Go.
* You have a `<Windows/Linux/Mac>` machine. State which OS is supported or works with the software.
* You have read `<guide/link/documentation_related_to_project>`.

## Installation

To install Akeyless Sheller, follow these steps:

1. Clone the repository
2. Navigate to the cloned repository
3. Run `go build`

## Usage

To use Akeyless Sheller, follow these steps:

1. Import it into your Golang project.
2. Run the `GetTokenFromAkeylessCommandLine` function. This function requires the `AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND` environment variable to be set. If it is not set, or if it is set to a path that does not point to an executable file, the function will return an error. The value of this variable should be something like:

```sh
akeyless auth --access-id p-jgk2szbi1vwd --access-type saml --json --jq-expression '.token'
```

or

```sh
/path/to/akeyless auth --access-id p-jgk2szbi1vwd --access-type saml --json --jq-expression '.token'
```

## Testing

To run tests, execute the following command in the project directory:

```sh
go test ./...
```

## Contributing to Akeyless Sheller

To contribute to Akeyless Sheller, follow these steps:

1. Fork the repository.
2. Create a new branch: `git checkout -b '<branch_name>'`.
3. Make your changes and commit them: `git commit -m '<commit_message>'`
4. Push to the original branch: `git push origin '<project_name>/<location>'`
5. Create the pull request.

## License

Akeyless Sheller is licensed under the Apache 2.0 License.
