window.BENCHMARK_DATA = {
  "lastUpdate": 1751017142234,
  "repoUrl": "https://github.com/aasanchez/ocpp16messages",
  "entries": {
    "Go Benchmark (authorize - Go 1.24.4)": [
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
        "date": 1751017141471,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizeConfirmation_valid",
            "value": 215.2,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5611702 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - ns/op",
            "value": 215.2,
            "unit": "ns/op",
            "extra": "5611702 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5611702 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5611702 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus",
            "value": 770.1,
            "unit": "ns/op\t     368 B/op\t       7 allocs/op",
            "extra": "1534837 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - ns/op",
            "value": 770.1,
            "unit": "ns/op",
            "extra": "1534837 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1534837 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1534837 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate",
            "value": 1496,
            "unit": "ns/op\t     952 B/op\t      19 allocs/op",
            "extra": "706806 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - ns/op",
            "value": 1496,
            "unit": "ns/op",
            "extra": "706806 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - B/op",
            "value": 952,
            "unit": "B/op",
            "extra": "706806 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "706806 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag",
            "value": 837.4,
            "unit": "ns/op\t     504 B/op\t       7 allocs/op",
            "extra": "1428648 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - ns/op",
            "value": 837.4,
            "unit": "ns/op",
            "extra": "1428648 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - B/op",
            "value": 504,
            "unit": "B/op",
            "extra": "1428648 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1428648 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest",
            "value": 18.87,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "62917734 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest - ns/op",
            "value": 18.87,
            "unit": "ns/op",
            "extra": "62917734 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "62917734 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "62917734 times\n4 procs"
          }
        ]
      }
    ]
  }
}