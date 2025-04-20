# Contributing to TGLog

Thank you for considering contributing to TGLog! This document outlines the guidelines for contributing to the project.

## Getting Started

1. **Fork the repository** and clone your fork
2. **Make sure the tests pass** on your machine:
   ```
   go test ./...
   ```
3. **Create a new branch** for your changes:
   ```
   git checkout -b your-branch-name
   ```

## Making Changes

### Code Style

Follow the standard Go style guidelines:
- Format your code with `go fmt`
- Run `go vet` to check for common errors
- Follow naming conventions in existing code

### Security

Since this package deals with potentially sensitive information:
- Never commit tokens, credentials, or sensitive data in your code
- Don't lower security standards for convenience
- Report security issues privately, not in public issues

### Commits

- Use clear, descriptive commit messages
- Keep commits focused on single changes
- Reference issue numbers in commit messages

### Tests

- Add tests for new features
- Ensure existing tests pass
- Include both unit tests and integration tests when appropriate

## Pull Requests

1. Update the README.md with details of changes if appropriate
2. Update the documentation for any new features or changes
3. The PR should work with Go 1.18 and newer
4. Ensure all tests pass and there are no linting errors

## Review Process

Once you submit a PR:
1. Maintainers will review your code
2. You may need to make changes before your PR is accepted
3. Once approved, your PR will be merged

## Code of Conduct

- Be respectful and inclusive in your communications
- Focus on the technical merits of contributions
- Help others learn and grow

## Questions?

If you have any questions about contributing, feel free to open an issue for clarification. 