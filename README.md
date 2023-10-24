# Akeyless Sheller

Akeyless Sheller is a Golang library designed to shell out to the Akeyless CLI to retrieve a token. This token can then be used in multiple tools such as CLI helpers and Terraform.

## Installation

To install Akeyless Sheller, you need to import it into your Golang project.

## Usage

To use Akeyless Sheller, run the `GetTokenFromAkeylessCommandLine` function. This function requires the `AKEYLESS_CLI_AUTHENTICATION_TOKEN_COMMAND` environment variable to be set. If it is not set, the function will return an error. The value of this variable should be something like:

```
akeyless auth --access-id p-jgk2szbi1vwd --access-type saml --json --jq-expression '.token'
```

or

```
/path/to/akeyless auth --access-id p-jgk2szbi1vwd --access-type saml --json --jq-expression '.token'
```

## License

Akeyless Sheller is licensed under the Apache 2.0 License.
