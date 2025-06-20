window.BENCHMARK_DATA = {
  "lastUpdate": 1750430754623,
  "repoUrl": "https://github.com/aasanchez/ocpp16messages",
  "entries": {
    "Go Benchmark": [
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
          "id": "cb77e7d1c2bb277985eb30424ab4a3ced8845967",
          "message": "chore(golangci.yml): disable testpackage linter to streamline linting process",
          "timestamp": "2025-06-20T14:45:09Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/cb77e7d1c2bb277985eb30424ab4a3ced8845967"
        },
        "date": 1750430753979,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizeConfirmation_valid",
            "value": 216.1,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5672028 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - ns/op",
            "value": 216.1,
            "unit": "ns/op",
            "extra": "5672028 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5672028 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5672028 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus",
            "value": 756.4,
            "unit": "ns/op\t     368 B/op\t       7 allocs/op",
            "extra": "1595096 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - ns/op",
            "value": 756.4,
            "unit": "ns/op",
            "extra": "1595096 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1595096 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1595096 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate",
            "value": 1464,
            "unit": "ns/op\t     952 B/op\t      19 allocs/op",
            "extra": "710136 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - ns/op",
            "value": 1464,
            "unit": "ns/op",
            "extra": "710136 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - B/op",
            "value": 952,
            "unit": "B/op",
            "extra": "710136 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "710136 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag",
            "value": 809.1,
            "unit": "ns/op\t     504 B/op\t       7 allocs/op",
            "extra": "1479691 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - ns/op",
            "value": 809.1,
            "unit": "ns/op",
            "extra": "1479691 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - B/op",
            "value": 504,
            "unit": "B/op",
            "extra": "1479691 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1479691 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid",
            "value": 18.94,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "61627479 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - ns/op",
            "value": 18.94,
            "unit": "ns/op",
            "extra": "61627479 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "61627479 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "61627479 times\n4 procs"
          }
        ]
      }
    ]
  }
}