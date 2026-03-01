# Roadmap

This document captures planned work that would require a breaking change and
therefore a major version bump.

## v2 (breaking): make messages truly immutable

Today, `ReqMessage` / `ConfMessage` structs generally have exported fields. They
are safe for concurrent use only when treated as read-only by consumers.

In v2, the goal is to make message values *structurally immutable*:

- Make message struct fields unexported.
- Provide read-only getters for all fields.
- Ensure constructors perform deep defensive copies for pointers/slices.
- Preserve JSON shape and tags (via struct tags and/or custom JSON marshalers)
  to keep wire compatibility.

Benefits:

- Stronger concurrency guarantees (no accidental consumer mutation).
- Clearer invariants (messages remain valid after construction).
- Fewer race/aliasing footguns for downstream code.
