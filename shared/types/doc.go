// Package sharedtypes provides small, validated
// value types used across OCPP 1.6 payloads.
// The goal is to keep message fields precise,
// safe, and predictable in both CP and CSMS.
//
// Why this matters for OCPP 1.6
//   - Deterministic parsing and formatting of
//     timestamps and bounded strings prevents
//     drift, overflow, and injection issues.
//   - Compact encodings reduce bandwidth and
//     improve logging and troubleshooting.
//   - Consistent errors make validation easy to
//     reason about across handlers and tests.
//
// Provided types
//
//   - DateTime: RFC3339 input, normalized to UTC.
//     String() returns RFC3339Nano in UTC for
//     stable logs, tests, and message signing.
//     Common in StartTransaction, MeterValues,
//     StatusNotification, and BootNotification.
//
//   - Integer: unsigned 16-bit (0..65535). Use
//     for connector indices, counters, and small
//     numeric identifiers seen in request flows.
//
//   - CiStringXX: bounded, printable ASCII only
//     (32..126). Families: 20, 25, 50, 255, 500.
//     Typical mappings include idTag (20), vendor
//     and model descriptors (25/50), and longer
//     diagnostic or vendor fields (255/500).
//
// Construction and behavior
//   - SetDateTime parses RFC3339 and normalizes
//     to UTC; errors wrap the parse failure.
//   - SetInteger parses base-10 into uint16;
//     rejects empty, negatives, overflow, non
//     decimal text. Never panics.
//   - SetCiStringXXType validates ASCII range and
//     length cap. Empty is allowed unless message
//     schema forbids it. All values are immutable
//     and safe to share across goroutines.
//
// Error conventions
//   - See errors.go for short, consistent formats.
//     Prefer ErrorWrapper when adding context:
//     fmt.Errorf(ErrorWrapper, err, input)
//     Keep labels protocol-aligned (e.g. "IdTag").
//
// Interop and safety notes
//   - DateTime always renders UTC for predictable
//     comparisons and signatures across parties.
//   - CiString types reject non printable input to
//     protect logs, user interfaces, and relays.
//   - Integer bounds avoid silent truncation and
//     clarify intent for compact numeric fields.
//
// Examples and tests
//   - See example_*.go files for usage patterns.
//   - Unit, fuzz, race, and benchmark suites live
//     under shared/types/{tests,fuzz,race,benchmark}.
//     They focus on atomic, adversarial cases and
//     invariants important to OCPP 1.6 reliability.
package sharedtypes
