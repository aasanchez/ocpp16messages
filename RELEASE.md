# Release Checklist (v1.x)

- Prefer using the one-command script in `scripts/release.sh` (see
  `RELEASING.md`). Use this checklist when you need to do the steps manually.

- [ ] Update `CHANGELOG.md` with the new version entry.
- [ ] Verify supported Go versions:
  - [ ] Minimum: Go 1.25.8 (`go.mod`)
  - [ ] Latest: Go 1.26.1 (CI matrix)
- [ ] Run quality gates locally:
  - [ ] `make test`
  - [ ] `make test-all`
  - [ ] `make test-race`
  - [ ] `make test-fuzz` (short fuzz budget)
  - [ ] `make test-bench` (smoke)
- [ ] Ensure README/AGENTS reflect status and commands.
- [ ] Tag the release with a **signed tag** (annotated):
  - [ ] `git tag -s vX.Y.Z -m "vX.Y.Z: <summary>"`
  - [ ] `git push origin vX.Y.Z`
- [ ] Generate release notes from `CHANGELOG.md` and publish the GitHub release:
  - [ ] `bash scripts/release-notes.sh vX.Y.Z | gh release create vX.Y.Z`
        `--title "vX.Y.Z" --notes-file -`
- [ ] Verify the tag policy workflow passed for `vX.Y.Z` (signed tag + changelog
      entry).
