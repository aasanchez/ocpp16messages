window.BENCHMARK_DATA = {
  "lastUpdate": 1750949716275,
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
          "id": "c0788140c64012d48e15b39087268e56da865a57",
          "message": "feat(benchmark): add Go benchmarks for Authorize and Authorize Types with GitHub Actions workflow",
          "timestamp": "2025-06-26T13:55:20Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/c0788140c64012d48e15b39087268e56da865a57"
        },
        "date": 1750947178458,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRequest_validPayload",
            "value": 227.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "5068480 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - ns/op",
            "value": 227.1,
            "unit": "ns/op",
            "extra": "5068480 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "5068480 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "5068480 times\n4 procs"
          }
        ]
      }
    ],
    "Go Benchmark (bootnotification)": [
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
          "id": "50c2ec222bc992f6f508a3fba62205d2d4f8f71b",
          "message": "test(ciString): add test for non-printable character handling in CiString to ensure proper error reporting",
          "timestamp": "2025-06-26T14:54:22Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/50c2ec222bc992f6f508a3fba62205d2d4f8f71b"
        },
        "date": 1750949715912,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRequest_validPayload",
            "value": 232.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "5221764 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - ns/op",
            "value": 232.5,
            "unit": "ns/op",
            "extra": "5221764 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "5221764 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "5221764 times\n4 procs"
          }
        ]
      }
    ]
  }
}