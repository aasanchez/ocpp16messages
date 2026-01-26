# Release Checklist (v1.x)

- [ ] Update `CHANGELOG.md` with the new version entry.
- [ ] Verify `go.mod` and CI workflows target the same Go version (current:
      1.24.6).
- [ ] Run quality gates locally:
  - [ ] `make test`
  - [ ] `make test-all`
  - [ ] `make test-race`
  - [ ] `make test-fuzz` (short fuzz budget)
  - [ ] `make test-bench` (smoke)
- [ ] Ensure README/AGENTS reflect status and commands.
- [ ] Tag the release (e.g., `v1.0.0`) and push tags.
- [ ] Create GitHub release notes from `CHANGELOG.md` and attach artifacts if
      desired.
- [ ] (Optional) Trigger nightly workflow run to confirm opt-in suites on CI.
