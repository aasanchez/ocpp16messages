window.BENCHMARK_DATA = {
  "lastUpdate": 1750967139443,
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
      },
      {
        "commit": {
          "author": {
            "name": "aasanchez",
            "username": "aasanchez"
          },
          "committer": {
            "name": "aasanchez",
            "username": "aasanchez"
          },
          "id": "717b5ab8813f94c108df8c3927457c08ba4a3ae1",
          "message": "Bootnotification confirmation",
          "timestamp": "2025-06-26T15:30:23Z",
          "url": "https://github.com/aasanchez/ocpp16messages/pull/2/commits/717b5ab8813f94c108df8c3927457c08ba4a3ae1"
        },
        "date": 1750953597741,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRequest_validPayload",
            "value": 226.3,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "5019027 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - ns/op",
            "value": 226.3,
            "unit": "ns/op",
            "extra": "5019027 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "5019027 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "5019027 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "aasanchez",
            "username": "aasanchez"
          },
          "committer": {
            "name": "aasanchez",
            "username": "aasanchez"
          },
          "id": "05be3e675a08fb78edaadf0de82c21f7fde1ca72",
          "message": "Bootnotification confirmation",
          "timestamp": "2025-06-26T15:30:23Z",
          "url": "https://github.com/aasanchez/ocpp16messages/pull/2/commits/05be3e675a08fb78edaadf0de82c21f7fde1ca72"
        },
        "date": 1750956496398,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRequest_validPayload",
            "value": 224.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "5187124 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - ns/op",
            "value": 224.5,
            "unit": "ns/op",
            "extra": "5187124 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "5187124 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "5187124 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "aasanchez",
            "username": "aasanchez"
          },
          "committer": {
            "name": "aasanchez",
            "username": "aasanchez"
          },
          "id": "7ea895ca2cd6175b5f4af6b3b7089e6903bbaf02",
          "message": "Bootnotification confirmation",
          "timestamp": "2025-06-26T15:30:23Z",
          "url": "https://github.com/aasanchez/ocpp16messages/pull/2/commits/7ea895ca2cd6175b5f4af6b3b7089e6903bbaf02"
        },
        "date": 1750965941451,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRequest_validPayload",
            "value": 226.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "5061588 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - ns/op",
            "value": 226.1,
            "unit": "ns/op",
            "extra": "5061588 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "5061588 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "5061588 times\n4 procs"
          }
        ]
      },
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
          "id": "539507f8840894520162bf6682a76188e21925b6",
          "message": "feat(bootnotification): add benchmarks for confirmation and payload validation, and refactor confirmation function to export it",
          "timestamp": "2025-06-26T19:37:56Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/539507f8840894520162bf6682a76188e21925b6"
        },
        "date": 1750966893328,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConfirmation_Valid",
            "value": 112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "10739318 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_Valid - ns/op",
            "value": 112,
            "unit": "ns/op",
            "extra": "10739318 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "10739318 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "10739318 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload",
            "value": 236.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "5000230 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - ns/op",
            "value": 236.9,
            "unit": "ns/op",
            "extra": "5000230 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "5000230 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "5000230 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "aasanchez@gmail.com",
            "name": "Alexis Sánchez",
            "username": "aasanchez"
          },
          "committer": {
            "email": "aasanchez@gmail.com",
            "name": "Alexis Sánchez",
            "username": "aasanchez"
          },
          "distinct": true,
          "id": "539507f8840894520162bf6682a76188e21925b6",
          "message": "feat(bootnotification): add benchmarks for confirmation and payload validation, and refactor confirmation function to export it",
          "timestamp": "2025-06-26T21:37:56+02:00",
          "tree_id": "9d6912b70f14d6a4f091a58c5a78de0298ecf65e",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/539507f8840894520162bf6682a76188e21925b6"
        },
        "date": 1750967139039,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkConfirmation_Valid",
            "value": 111.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "10707744 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_Valid - ns/op",
            "value": 111.4,
            "unit": "ns/op",
            "extra": "10707744 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "10707744 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "10707744 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload",
            "value": 228.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "5024966 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - ns/op",
            "value": 228.4,
            "unit": "ns/op",
            "extra": "5024966 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "5024966 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_validPayload - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "5024966 times\n4 procs"
          }
        ]
      }
    ]
  }
}