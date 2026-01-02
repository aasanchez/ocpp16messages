# Security Policy

## Supported Versions

We release patches for security vulnerabilities in the following versions:

| Version | Supported          |
| ------- | ------------------ |
| 0.x.x   | :white_check_mark: |

**Note:** This project is currently in active development (pre-1.0). Security updates are provided for the latest version only.

## Reporting a Vulnerability

We take the security of this OCPP 1.6 library seriously. If you discover a security vulnerability, please report it responsibly.

### How to Report

**Please do NOT report security vulnerabilities through public GitHub issues.**

Instead, please report security vulnerabilities using one of the following methods:

1. **GitHub Security Advisories (Preferred)**
   - Navigate to the [Security Advisories](https://github.com/aasanchez/ocpp16messages/security/advisories) page
   - Click "Report a vulnerability"
   - Provide detailed information about the vulnerability

2. **Email**
   - Send details to the repository maintainer via GitHub email
   - Include "SECURITY" in the subject line

### What to Include

Please include the following information in your report:

- Type of vulnerability (e.g., input validation, authentication bypass, injection)
- Full paths of affected source files
- Location of the affected code (tag/branch/commit or direct URL)
- Step-by-step instructions to reproduce the issue
- Proof-of-concept or exploit code (if possible)
- Impact assessment and potential attack scenarios
- Any suggested fixes or mitigations

### Response Timeline

- **Initial Response:** Within 48 hours
- **Status Update:** Within 7 days with preliminary assessment
- **Fix Timeline:** Critical vulnerabilities will be addressed within 30 days
- **Public Disclosure:** After a fix is released, coordinated with the reporter

## Security Considerations for This Library

### Scope

This library provides OCPP 1.6 message type definitions and validation for Go applications. Security vulnerabilities in scope include:

#### In Scope

- **Input Validation Failures:** Bypass of length constraints, character encoding issues
- **Injection Vulnerabilities:** SQL injection, command injection, or similar attacks via message fields
- **Authentication Issues:** IDToken validation bypass or manipulation
- **Data Exposure:** Sensitive data leakage through error messages or logging
- **Protocol Implementation Flaws:** Violations of OCPP 1.6 spec that create security risks
- **Denial of Service:** Resource exhaustion through malformed messages
- **Type Confusion:** Vulnerabilities arising from type system bypasses

#### Out of Scope

- Network layer security (TLS/SSL configuration)
- WebSocket security (handled by transport layer)
- Application-level authorization logic (business logic in consuming applications)
- Vulnerabilities in third-party dependencies (report to respective projects)
- Issues requiring physical access to charging infrastructure

### Security Best Practices for Users

When using this library in your OCPP implementation:

1. **Always validate input:** Use the provided `Validate()` methods on all message types
2. **Handle errors securely:** Don't expose internal error details to untrusted parties
3. **Use type constructors:** Always use `New*()` constructors (e.g., `NewIDToken()`) rather than direct struct initialization
4. **Bounds checking:** The library enforces OCPP 1.6 string length limits - don't bypass them
5. **Sanitize logs:** Be careful when logging message contents that may contain sensitive IDTokens
6. **Keep dependencies updated:** Regularly update to the latest version for security patches
7. **Thread safety:** All types use value receivers and are designed for concurrent use - don't modify after creation

### Known Security Features

- ✅ Strict type validation using constructor pattern
- ✅ OCPP 1.6 length constraints enforced at construction time
- ✅ ASCII printable character validation for CiString types
- ✅ Empty value rejection for required fields
- ✅ Immutable types with private fields
- ✅ Thread-safe by design (value receivers, no mutation)
- ✅ Error wrapping for context without information disclosure

## Security Update Policy

- Security patches will be released as soon as possible after verification
- Critical vulnerabilities receive immediate attention
- All security updates will be documented in release notes with CVE identifiers when applicable
- GitHub Security Advisories will be published for all confirmed vulnerabilities

## Acknowledgments

We appreciate the security research community's efforts to responsibly disclose vulnerabilities. Contributors who report valid security issues will be acknowledged in:

- The security advisory (unless they prefer to remain anonymous)
- The CHANGELOG for the patched release
- This SECURITY.md file's acknowledgments section

Thank you for helping keep this project and its users secure.
