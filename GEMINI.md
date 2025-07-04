# Gemini CLI Instructions

This document provides instructions for the Gemini CLI to effectively interact with the `ocpp16messages` project.

## About the Project

This project is a Go library for creating and parsing Open Charge Point Protocol (OCPP) 1.6 messages. It is designed to be performant, type-safe, and easy to use.

## Key Technologies

- **Language:** Go
- **Linting:** `golangci-lint`
- **Testing:** Go's built-in testing framework
- **Code Quality:** SonarQube

## Project Conventions

### Directory Structure

- `messages/`: Contains the implementation for each OCPP message type. Each message has its own sub-package (e.g., `messages/authorize`).
- `shared/`: Contains shared types and utilities used across different messages.
- `tests/`: Subdirectories named `tests` (e.g., `messages/authorize/types/tests`) contain blackbox tests (`_external_test.go`).

### Testing

- **Whitebox Tests:** Standard unit tests are located in the same package as the code they test and have a `_test.go` suffix.
- **Blackbox Tests:** Integration-style tests are located in `tests` subdirectories and have an `_external_test.go` suffix. These are excluded from the main test runs.

### Code Style

Code is formatted using `gofmt`. Use the `make format` command to ensure consistent styling.

## Development Workflow & Commands

The `Makefile` contains targets for common development tasks.

- **Run all linters:**

  ```bash
  make lint
  ```

- **Run all tests (whitebox and blackbox) and generate coverage:**

  ```bash
  make test
  ```

- **Run only whitebox unit tests:**

  ```bash
  go test $(go list -mod=readonly ./... | grep -v '/tests')
  ```

- **Format the code:**

  ```bash
  make format
  ```

- **Run benchmark tests:**

  ```bash
  make benchmark
  ```

- **Generate and view HTML test coverage report:**

  ```bash
  make test-coverage-html
  ```

## Agent Instructions

- Before committing, always run `make lint` and `make test` to ensure the code is clean and all tests pass.
- When adding new features or fixing bugs, please add corresponding unit tests.
- When adding a new OCPP message, follow the existing directory structure and request/confirmation pattern.
- Adhere strictly to the project's code style and conventions.
