# Compatibility Contract

This document defines the public compatibility guarantees for
`github.com/aasanchez/ocpp16messages`.

## Semantic Versioning (SemVer)

This project follows [SemVer](https://semver.org/).

- **MAJOR** (`v2.0.0`): breaking changes (see "Breaking changes" below).
- **MINOR** (`v1.1.0`): backward-compatible additions and improvements.
- **PATCH** (`v1.0.4`): backward-compatible bug fixes and documentation.

## Public API surface

For the purposes of SemVer, the "public API" includes:

- Exported identifiers (types, functions, methods, constants, variables).
- Package paths and directory structure (import paths are stable).
- JSON encoding/decoding behavior for exported message structs.
- Validation behavior of constructors (`Req`, `Conf`, `New*`) and error
  semantics (see below).

## Error contract

Error handling is part of the compatibility contract. Callers should:

- Use `errors.Is` for validation errors.
- Avoid depending on exact error strings.

### Stable guarantees

- Validation failures from public constructors wrap one of:
  - `types.ErrEmptyValue`
  - `types.ErrInvalidValue`
- Constructors may return `errors.Join(...)` to report multiple field errors;
  `errors.Is` will still match the underlying sentinel errors.

### Not guaranteed

- Exact error messages, formatting, and field prefixes.
- Ordering of joined errors.
- Presence or absence of additional context wrapping (as long as `errors.Is`
  behavior for the sentinels remains stable).

## Breaking changes (v1.x)

In `v1.x`, we treat the following as breaking changes and require a MAJOR bump:

- **API shape**
  - Removing or renaming exported identifiers.
  - Changing exported function/method signatures.
  - Changing exported struct fields (name, type, tags) in a way that breaks
    compilation or JSON behavior.
  - Moving packages (changing import paths).
- **Behavior / validation**
  - Tightening validation such that previously accepted inputs now return an
    error (including changes to optional vs required field semantics).
  - Changing the meaning of fields, enums, or defaulting behavior in a way that
    affects wire compatibility or application logic.
- **Error contracts**
  - Changing error wrapping such that `errors.Is` no longer matches the shared
    sentinel errors (`types.ErrEmptyValue`, `types.ErrInvalidValue`) for
    previously-invalid inputs.

The following are generally **not** considered breaking changes:

- Adding new exported identifiers.
- Fixing a bug that previously produced invalid values or violated the OCPP 1.6
  spec (even if it changes returned values for invalid inputs).
- Changing error messages, as long as error identity via `errors.Is` remains
  stable.

## Deprecation process

When we need to replace or retire a public API:

1. Mark the API as deprecated using a Go doc comment that starts with
   `Deprecated:` and points to the replacement.
2. Document the deprecation in `CHANGELOG.md` in the first release where it
   appears.
3. Keep deprecated APIs available until the next MAJOR release, unless the API
   was added and deprecated within the same MINOR release.

## Support window

For security support details, see `SECURITY.md` ("Supported Versions Policy").
