#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 1 ]]; then
  echo "usage: $0 <tag-or-version>" >&2
  exit 2
fi

tag="$1"
section="$tag"

if [[ "$tag" =~ ^v0\. ]]; then
  section="$tag"
elif [[ "$tag" =~ ^v[0-9] ]]; then
  section="${tag#v}"
fi

if [[ ! -f CHANGELOG.md ]]; then
  echo "CHANGELOG.md not found" >&2
  exit 1
fi

awk -v section="$section" '
  BEGIN { in_section = 0; printed = 0 }
  $0 ~ "^## \\[" section "\\]" {
    in_section = 1
    printed = 1
  }
  in_section && $0 ~ "^## \\[" && $0 !~ "^## \\[" section "\\]" {
    exit
  }
  in_section { print }
  END {
    if (printed == 0) {
      exit 1
    }
  }
' CHANGELOG.md
