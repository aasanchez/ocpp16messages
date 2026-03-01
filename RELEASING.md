# Releasing

This repository uses:

- Signed, annotated tags for releases (required; enforced in CI).
- Release notes extracted from `CHANGELOG.md`.

## Optional: signed commits

Signed commits are recommended but not required.

To enable signing commits locally:

```sh
git config --global commit.gpgsign true
git config --global user.signingkey <YOUR_KEY_ID>
```

## One-command release

Prerequisites:

- Working tree clean.
- On `main`, and `main` is pushed (matches `origin/main`).
- GPG configured for `git tag -s`.
- GitHub CLI authenticated (`gh auth login`).

Release steps:

1. Add the new version section to `CHANGELOG.md` (e.g., `## [1.0.4] - YYYY-MM-DD`)
   and push the commit to `main`.
2. Run:

```sh
bash scripts/release.sh vX.Y.Z "vX.Y.Z: <summary>"
```

This will:

- Run `make test-all`
- Create a signed tag `vX.Y.Z`
- Push the tag
- Create the GitHub release with notes extracted from `CHANGELOG.md`

For the manual checklist, see `RELEASE.md`.
