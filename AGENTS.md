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

## Agent: Go Documentation Optimizer

### Purpose

Whenever I ask to “improve the documentation for a file,” this agent should automatically apply the following
instructions. The goal is to ensure all Go files in the project have **pkgsite-ready documentation** with domain context (OCPP 1.6), making them clear, professional, and highly usable.

Always write the file, if is possible

### Universal Rules for Documentation

- Always generate **pkgsite-style GoDoc comments** (top of file, package, and functions).
- Documentation must be:
  - Highly detailed
  - Precise and professional
  - Clear and idiomatic Go style
- Always frame explanations in the context of **OCPP 1.6** (Open Charge Point Protocol v1.6).
  - Explain why this type/test/function matters in OCPP.
  - Highlight robustness, validation, and prevention of malformed/malicious input.
  - Reference real OCPP 1.6 domains (transaction IDs, connector IDs, meter values, authorization IDs, etc.).
- Include at least three levels of comments:
  - Package-level comment — What this package/file is about and its role in OCPP.
  - Type/function-level comments — What it does, valid/invalid behavior, expected usage.
  - Contextual notes — Why this matters for OCPP 1.6 reliability, safety, and compliance.
- Output should be ready to copy into the source file without extra formatting.
- Do not just describe what the code does — connect it to why it matters in OCPP 1.6.

### Prompt Template (auto-applied)

```text
Please improve and optimize the Go documentation for this file.  
Follow the **Universal Rules for Documentation** in Agents.md.  

**File:** <replace-with-path>
```

### Usage Notes

- When I say:
  > “Improve the documentation for @shared/types/fuzz/datetime_fuzz_test.go”
  > The agent should automatically apply the rules above without me restating them.
- Works for any file path.
- Output must be directly usable in pkgsite.
