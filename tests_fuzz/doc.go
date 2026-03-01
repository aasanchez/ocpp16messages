// Package fuzz contains fuzz tests for this module.
//
// Fuzzers in this directory are guarded by the `fuzz` build tag to avoid
// running during normal `go test ./...` runs. Run them via:
//
//	go test -tags=fuzz ./fuzz
//
// Or use the project helper target:
//
//	make test-fuzz
package fuzz
