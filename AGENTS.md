# AGENTS.md

## Prerequist

- Go ≥ 1.24
- Tools used by Makefile: golangci-lint, staticcheck, ripgrep (rg), gci, gofumpt, golines, pkgsite, sonar-scanner
  (optional)

## Coding Style

- Never add external third party libraries or package, we must use ONLY and exclusively the golang standard library
- All type of tests, which can be unitest, fuzzy test, benchmark, or even race test, must be very atomic, targeting
  individual cases, very atomic and simple.
- Leverage and exten [shared/types/errors.go](shared/types/errors.go) to consolidate the message errors
- No string should be longer than 40 charecters, no mather the reason.
