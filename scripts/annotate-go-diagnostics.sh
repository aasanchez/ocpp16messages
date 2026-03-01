#!/usr/bin/env bash
set -euo pipefail

level="${1:-error}"

encode() {
  local value="$1"
  value="${value//'%'/'%25'}"
  value="${value//$'\n'/'%0A'}"
  value="${value//$'\r'/'%0D'}"
  printf '%s' "${value}"
}

while IFS= read -r raw; do
  if [[ -z "${raw}" ]]; then
    continue
  fi

  line="$(encode "${raw}")"

  if [[ "${raw}" =~ ^([^:]+):([0-9]+):([0-9]+):[[:space:]](.*)$ ]]; then
    file="${BASH_REMATCH[1]}"
    line_no="${BASH_REMATCH[2]}"
    col_no="${BASH_REMATCH[3]}"
    message="$(encode "${BASH_REMATCH[4]}")"
    echo "::${level} file=$(encode "${file}"),line=${line_no},col=${col_no}::${message}"
    continue
  fi

  if [[ "${raw}" =~ ^([^:]+):([0-9]+):[[:space:]](.*)$ ]]; then
    file="${BASH_REMATCH[1]}"
    line_no="${BASH_REMATCH[2]}"
    message="$(encode "${BASH_REMATCH[3]}")"
    echo "::${level} file=$(encode "${file}"),line=${line_no}::${message}"
    continue
  fi

  echo "::${level}::${line}"
done

