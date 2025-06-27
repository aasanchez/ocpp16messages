window.BENCHMARK_DATA = {
  "lastUpdate": 1751017136007,
  "repoUrl": "https://github.com/aasanchez/ocpp16messages",
  "entries": {
    "Go Benchmark (bootnotification - Go 1.24.4)": [
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
        "date": 1751017135322,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConfirmation",
            "value": 111.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "10739161 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation - ns/op",
            "value": 111.9,
            "unit": "ns/op",
            "extra": "10739161 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "10739161 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "10739161 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload",
            "value": 227,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "5154638 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - ns/op",
            "value": 227,
            "unit": "ns/op",
            "extra": "5154638 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "5154638 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "5154638 times\n4 procs"
          }
        ]
      }
    ]
  }
}