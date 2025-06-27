window.BENCHMARK_DATA = {
  "lastUpdate": 1751017139188,
  "repoUrl": "https://github.com/aasanchez/ocpp16messages",
  "entries": {
    "Go Benchmark (bootnotification-types - Go 1.23.10)": [
      {
        "commit": {
          "author": {
            "name": "Alexis Sánchez",
            "username": "aasanchez",
            "email": "aasanchez@gmail.com"
          },
          "committer": {
            "name": "Alexis Sánchez",
            "username": "aasanchez",
            "email": "aasanchez@gmail.com"
          },
          "id": "3ccab0be6434c00488be9c30f804a4a1330c8e72",
          "message": "chore(benchmark.yml): update Go version matrix and adjust benchmark output paths for better organization",
          "timestamp": "2025-06-27T09:37:20Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/3ccab0be6434c00488be9c30f804a4a1330c8e72"
        },
        "date": 1751017138848,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConfirmationPayload_Validate",
            "value": 3.151,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "383371963 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate - ns/op",
            "value": 3.151,
            "unit": "ns/op",
            "extra": "383371963 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "383371963 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "383371963 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MissingCurrentTime",
            "value": 163.8,
            "unit": "ns/op\t      96 B/op\t       2 allocs/op",
            "extra": "7381299 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MissingCurrentTime - ns/op",
            "value": 163.8,
            "unit": "ns/op",
            "extra": "7381299 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MissingCurrentTime - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "7381299 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MissingCurrentTime - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7381299 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MissingInterval",
            "value": 167.8,
            "unit": "ns/op\t      96 B/op\t       2 allocs/op",
            "extra": "7410144 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MissingInterval - ns/op",
            "value": 167.8,
            "unit": "ns/op",
            "extra": "7410144 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MissingInterval - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "7410144 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MissingInterval - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7410144 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MissingStatus",
            "value": 162.2,
            "unit": "ns/op\t      96 B/op\t       2 allocs/op",
            "extra": "7435394 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MissingStatus - ns/op",
            "value": 162.2,
            "unit": "ns/op",
            "extra": "7435394 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MissingStatus - B/op",
            "value": 96,
            "unit": "B/op",
            "extra": "7435394 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MissingStatus - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7435394 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MixedParallel",
            "value": 66.7,
            "unit": "ns/op\t      72 B/op\t       1 allocs/op",
            "extra": "17476075 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MixedParallel - ns/op",
            "value": 66.7,
            "unit": "ns/op",
            "extra": "17476075 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MixedParallel - B/op",
            "value": 72,
            "unit": "B/op",
            "extra": "17476075 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_MixedParallel - allocs/op",
            "value": 1,
            "unit": "allocs/op",
            "extra": "17476075 times\n4 procs"
          },
          {
            "name": "BenchmarkRegistrationStatus_Accepted",
            "value": 62.92,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "19037056 times\n4 procs"
          },
          {
            "name": "BenchmarkRegistrationStatus_Accepted - ns/op",
            "value": 62.92,
            "unit": "ns/op",
            "extra": "19037056 times\n4 procs"
          },
          {
            "name": "BenchmarkRegistrationStatus_Accepted - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "19037056 times\n4 procs"
          },
          {
            "name": "BenchmarkRegistrationStatus_Accepted - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "19037056 times\n4 procs"
          },
          {
            "name": "BenchmarkRegistrationStatus_Pending",
            "value": 63.17,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "18974527 times\n4 procs"
          },
          {
            "name": "BenchmarkRegistrationStatus_Pending - ns/op",
            "value": 63.17,
            "unit": "ns/op",
            "extra": "18974527 times\n4 procs"
          },
          {
            "name": "BenchmarkRegistrationStatus_Pending - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "18974527 times\n4 procs"
          },
          {
            "name": "BenchmarkRegistrationStatus_Pending - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "18974527 times\n4 procs"
          },
          {
            "name": "BenchmarkRegistrationStatus_Rejected",
            "value": 63.54,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "18893406 times\n4 procs"
          },
          {
            "name": "BenchmarkRegistrationStatus_Rejected - ns/op",
            "value": 63.54,
            "unit": "ns/op",
            "extra": "18893406 times\n4 procs"
          },
          {
            "name": "BenchmarkRegistrationStatus_Rejected - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "18893406 times\n4 procs"
          },
          {
            "name": "BenchmarkRegistrationStatus_Rejected - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "18893406 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate",
            "value": 4.989,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "240600288 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate - ns/op",
            "value": 4.989,
            "unit": "ns/op",
            "extra": "240600288 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "240600288 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "240600288 times\n4 procs"
          }
        ]
      }
    ]
  }
}