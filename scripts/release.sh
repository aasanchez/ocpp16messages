#!/usr/bin/env bash
set -euo pipefail

usage() {
  cat >&2 <<'EOF'
usage: scripts/release.sh vX.Y.Z ["tag message"]

Runs release checks, creates a signed tag, pushes it, and publishes a GitHub
release with notes extracted from CHANGELOG.md.

Requirements:
- Clean git working tree
- On branch 'main' and up-to-date with origin/main
- GPG configured for signed tags (git tag -s)
- GitHub CLI authenticated (gh auth login)
EOF
  exit 2
}

if [[ $# -lt 1 || $# -gt 2 ]]; then
  usage
fi

tag="$1"
tag_message="${2:-Release ${tag}}"

if [[ ! "${tag}" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "invalid tag '${tag}': expected vX.Y.Z" >&2
  exit 2
fi

if ! git diff --quiet || ! git diff --cached --quiet; then
  echo "working tree is not clean; commit or stash changes first" >&2
  exit 1
fi

branch="$(git rev-parse --abbrev-ref HEAD)"
if [[ "${branch}" != "main" ]]; then
  echo "release must be run from branch 'main' (current: ${branch})" >&2
  exit 1
fi

git fetch origin --tags

if ! git show-ref --verify --quiet "refs/remotes/origin/main"; then
  echo "origin/main not found; is 'origin' configured?" >&2
  exit 1
fi

head_sha="$(git rev-parse HEAD)"
origin_main_sha="$(git rev-parse origin/main)"
if [[ "${head_sha}" != "${origin_main_sha}" ]]; then
  echo "main is not up-to-date with origin/main" >&2
  echo "HEAD:        ${head_sha}" >&2
  echo "origin/main: ${origin_main_sha}" >&2
  echo "push main before releasing" >&2
  exit 1
fi

if git show-ref --verify --quiet "refs/tags/${tag}"; then
  echo "tag '${tag}' already exists locally" >&2
  exit 1
fi

if git ls-remote --tags origin "${tag}" | rg -q "${tag}$"; then
  echo "tag '${tag}' already exists on origin" >&2
  exit 1
fi

tmp_notes="$(mktemp -t ocpp16messages-release-notes.XXXXXX)"
trap 'rm -f "${tmp_notes}"' EXIT

if ! bash scripts/release-notes.sh "${tag}" >"${tmp_notes}"; then
  echo "CHANGELOG.md section for '${tag}' not found" >&2
  exit 1
fi

echo "running quality gates (make test-all)..." >&2
make test-all

echo "creating signed tag '${tag}'..." >&2
git tag -s "${tag}" -m "${tag_message}"

echo "pushing tag '${tag}'..." >&2
git push origin "${tag}"

if gh release view "${tag}" >/dev/null 2>&1; then
  echo "GitHub release '${tag}' already exists; refusing to overwrite" >&2
  exit 1
fi

echo "publishing GitHub release '${tag}'..." >&2
gh release create "${tag}" \
  --title "${tag}" \
  --notes-file "${tmp_notes}" \
  --verify-tag

echo "done: released ${tag}" >&2

