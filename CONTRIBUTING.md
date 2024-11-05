# Contributing to Go REST API v2.0

First off, thank you for considering contributing to Go REST API! It's people like you that make this project such a great tool.

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the existing issues as you might find out that you don't need to create one. When you are creating a bug report, please include as many details as possible:

- **Use a clear and descriptive title**
- **Describe the exact steps to reproduce the problem**
- **Provide specific examples to demonstrate the steps**
- **Describe the behavior you observed and what behavior you expected**
- **Include screenshots if possible**
- **Include your environment details** (OS, Go version, etc.)

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, please include:

- **Use a clear and descriptive title**
- **Provide a detailed description of the suggested enhancement**
- **Provide specific examples to demonstrate the enhancement**
- **Explain why this enhancement would be useful**

### Pull Requests

1. Fork the repository
2. Create a new branch from `main`
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. Make your changes
4. Run tests and ensure they pass
   ```bash
   go test -v ./...
   ```
5. Format your code
   ```bash
   go fmt ./...
   ```
6. Commit your changes with a descriptive commit message
   ```bash
   git commit -m 'Add some amazing feature'
   ```
7. Push to your branch
   ```bash
   git push origin feature/amazing-feature
   ```
8. Open a Pull Request

### Coding Standards

- Follow Go best practices and idioms
- Write clear, readable code with comments where necessary
- Keep functions focused and small
- Use meaningful variable and function names
- Write tests for new functionality
- Ensure all tests pass before submitting PR
- Follow the existing code structure and patterns

### Commit Messages

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests liberally after the first line

### Go Code Style

- Run `go fmt` on your code
- Run `go vet` to catch common mistakes
- Follow the [Effective Go](https://golang.org/doc/effective_go.html) guidelines
- Use `gofmt` or `goimports` to format your code

## Project Structure

Please maintain the clean architecture structure:
- `cmd/`: Application entry points
- `internal/`: Private application code
  - `config/`: Configuration management
  - `database/`: Database connections
  - `models/`: Domain models
  - `repository/`: Data access layer
  - `service/`: Business logic layer
  - `handler/`: HTTP handlers
  - `middleware/`: HTTP middleware
  - `utils/`: Utility functions

## Testing

- Write unit tests for new functionality
- Ensure all tests pass before submitting
- Aim for meaningful test coverage
- Test edge cases and error conditions

## Documentation

- Update the README.md if needed
- Add comments to exported functions
- Update API documentation for new endpoints
- Include examples for new features

## Questions?

Feel free to open an issue with your question or contact the maintainers.

Thank you for contributing!
