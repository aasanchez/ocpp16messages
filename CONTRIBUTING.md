# Contributing to OCPP 1.6 Messages for Go

Thank you for considering a contribution! We welcome improvements from the community. Please follow these guidelines for bug reports, new features, documentation, and code contributions.

## Getting Started

- Fork the repository and clone your fork locally.
- Install dependencies: `go mod tidy`
- Ensure you are using the Go version defined in `go.mod`

## Development Workflow

1. **Create a branch** for your fix/feature: `git checkout -b my-feature`
2. **Write clear, idiomatic code** following the patterns in `OpenCode.md`.
3. **Document all exported functions/types** using Go-style doc comments.
4. **Add/update unit tests** and (when relevant) benchmarks for all new or modified logic.
5. Run and pass all checks:
   - Tests: `make test` or `go test ./...`
   - Benchmarks (if added): `make benchmark`
   - Lint: `make lint` or `golangci-lint run ./...`
   - Format: `make format` or `gofmt -w .`
   - Markdown docs: validate with [markdownlint](https://github.com/DavidAnson/markdownlint)
6. **Commit with a clear, descriptive message**. Reference issues if applicable.

## Code Style & Quality

- All code (manual or generated) must pass `golangci-lint` and `staticcheck` using the repo configuration before PR/merge.
- Markdown docs (including README, OpenCode.md, etc.) must pass default markdownlint rules.
- Follow package, type, constant, and error naming conventions explained in `OpenCode.md`.
- Use table-driven and parallel (`t.Parallel()`) unit tests when possible.
- Keep PRs focused and as small/atomic as possible.

## Submitting a Pull Request

- Push your branch and open a PR against the `main` branch.
- Fill in the PR template describing what and why you changed.
- Ensure all checks in the PR pass: CI, codecov, Sonar, etc.
- Be responsive to review feedback and update as necessary.

## Issues

- Search existing issues before opening a new one.
- Describe the issue clearly; include OS, Go version, and steps to reproduce if reporting a bug.

## Code of Conduct

- Be respectful, inclusive, and constructive.
- Disagreement is fine—personal attacks are not.

---
Ready to contribute? Great! Feel free to open issues or pull requests—thank you for helping make this project better.
