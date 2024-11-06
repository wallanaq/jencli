# JenCLI

JenCLI is a command-line interface (CLI) tool designed to streamline interactions with Jenkins and Keycloak, allowing authentication, session management, and automation tasks in a DevOps environment.

## Features

- **Login and Logout**: Authenticate using Keycloak and manage sessions.
- **Configuration Management**: Easily store and retrieve CLI configurations.
- **Integration with Jenkins**: Automate tasks and access Jenkins resources with ease.

## Installation

To install, clone the repository and build the binary.

```bash
git clone https://github.com/wallanaq/jencli.git
cd jencli
go build -o jencli
```

## Usage

### Authentication

- **Login**: Launches the Keycloak authentication flow.

```bash
jencli login
```

- **Logout**: Ends the Keycloak session for the authenticated user.

```bash
jencli logout
```

### Configuration

Configurations are stored locally on each operating system in standard directories:

Windows: `%APPDATA%\sicoobcli\db`
Linux/Mac: `~/.config/jencli/db`

## Development

To contribute or modify this CLI:

1. Fork the repository.
2. Make your changes and submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
