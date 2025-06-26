window.BENCHMARK_DATA = {
  "lastUpdate": 1750965961105,
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
          "id": "afe35890a08c7b935f1abb0241c8954cbb3b7ff7",
          "message": "chore(benchmark.yml): rename benchmark job and add authorize types benchmarks job for better clarity and organization",
          "timestamp": "2025-06-20T14:54:44Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/afe35890a08c7b935f1abb0241c8954cbb3b7ff7"
        },
        "date": 1750431416520,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizeConfirmation_valid",
            "value": 213.5,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5646331 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - ns/op",
            "value": 213.5,
            "unit": "ns/op",
            "extra": "5646331 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5646331 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5646331 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus",
            "value": 754.9,
            "unit": "ns/op\t     368 B/op\t       7 allocs/op",
            "extra": "1588610 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - ns/op",
            "value": 754.9,
            "unit": "ns/op",
            "extra": "1588610 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1588610 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1588610 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate",
            "value": 1471,
            "unit": "ns/op\t     952 B/op\t      19 allocs/op",
            "extra": "726823 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - ns/op",
            "value": 1471,
            "unit": "ns/op",
            "extra": "726823 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - B/op",
            "value": 952,
            "unit": "B/op",
            "extra": "726823 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "726823 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag",
            "value": 814.5,
            "unit": "ns/op\t     504 B/op\t       7 allocs/op",
            "extra": "1472732 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - ns/op",
            "value": 814.5,
            "unit": "ns/op",
            "extra": "1472732 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - B/op",
            "value": 504,
            "unit": "B/op",
            "extra": "1472732 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1472732 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid",
            "value": 19.02,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "61895824 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - ns/op",
            "value": 19.02,
            "unit": "ns/op",
            "extra": "61895824 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "61895824 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "61895824 times\n4 procs"
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
          "id": "1cb2bcb95adc26e49b8c1656438a61290f8db851",
          "message": "chore(golangci.yml): replace disabled linter godot with nilnil for better configuration clarity\nchore(ci_string_test.go): update comment for linting to clarify intentional control character test",
          "timestamp": "2025-06-25T12:23:52Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/1cb2bcb95adc26e49b8c1656438a61290f8db851"
        },
        "date": 1750861829797,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizeConfirmation_valid",
            "value": 215,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5598908 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - ns/op",
            "value": 215,
            "unit": "ns/op",
            "extra": "5598908 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5598908 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5598908 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus",
            "value": 780.2,
            "unit": "ns/op\t     368 B/op\t       7 allocs/op",
            "extra": "1539604 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - ns/op",
            "value": 780.2,
            "unit": "ns/op",
            "extra": "1539604 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1539604 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1539604 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate",
            "value": 1502,
            "unit": "ns/op\t     952 B/op\t      19 allocs/op",
            "extra": "729690 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - ns/op",
            "value": 1502,
            "unit": "ns/op",
            "extra": "729690 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - B/op",
            "value": 952,
            "unit": "B/op",
            "extra": "729690 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "729690 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag",
            "value": 843.3,
            "unit": "ns/op\t     504 B/op\t       7 allocs/op",
            "extra": "1365523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - ns/op",
            "value": 843.3,
            "unit": "ns/op",
            "extra": "1365523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - B/op",
            "value": 504,
            "unit": "B/op",
            "extra": "1365523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1365523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid",
            "value": 18.97,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "63539367 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - ns/op",
            "value": 18.97,
            "unit": "ns/op",
            "extra": "63539367 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "63539367 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "63539367 times\n4 procs"
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
          "id": "c0788140c64012d48e15b39087268e56da865a57",
          "message": "feat(benchmark): add Go benchmarks for Authorize and Authorize Types with GitHub Actions workflow",
          "timestamp": "2025-06-26T13:55:20Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/c0788140c64012d48e15b39087268e56da865a57"
        },
        "date": 1750947180961,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizeConfirmation_valid",
            "value": 211.5,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5667571 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - ns/op",
            "value": 211.5,
            "unit": "ns/op",
            "extra": "5667571 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5667571 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5667571 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus",
            "value": 753.9,
            "unit": "ns/op\t     368 B/op\t       7 allocs/op",
            "extra": "1600281 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - ns/op",
            "value": 753.9,
            "unit": "ns/op",
            "extra": "1600281 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1600281 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1600281 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate",
            "value": 1469,
            "unit": "ns/op\t     952 B/op\t      19 allocs/op",
            "extra": "722002 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - ns/op",
            "value": 1469,
            "unit": "ns/op",
            "extra": "722002 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - B/op",
            "value": 952,
            "unit": "B/op",
            "extra": "722002 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "722002 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag",
            "value": 815.7,
            "unit": "ns/op\t     504 B/op\t       7 allocs/op",
            "extra": "1476302 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - ns/op",
            "value": 815.7,
            "unit": "ns/op",
            "extra": "1476302 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - B/op",
            "value": 504,
            "unit": "B/op",
            "extra": "1476302 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1476302 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid",
            "value": 19.07,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "61791744 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - ns/op",
            "value": 19.07,
            "unit": "ns/op",
            "extra": "61791744 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "61791744 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "61791744 times\n4 procs"
          }
        ]
      }
    ],
    "Go Benchmark (authorize)": [
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
        "date": 1750949730113,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizeConfirmation_valid",
            "value": 212.7,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5648577 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - ns/op",
            "value": 212.7,
            "unit": "ns/op",
            "extra": "5648577 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5648577 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5648577 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus",
            "value": 754.4,
            "unit": "ns/op\t     368 B/op\t       7 allocs/op",
            "extra": "1541922 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - ns/op",
            "value": 754.4,
            "unit": "ns/op",
            "extra": "1541922 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1541922 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1541922 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate",
            "value": 1463,
            "unit": "ns/op\t     952 B/op\t      19 allocs/op",
            "extra": "818427 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - ns/op",
            "value": 1463,
            "unit": "ns/op",
            "extra": "818427 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - B/op",
            "value": 952,
            "unit": "B/op",
            "extra": "818427 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "818427 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag",
            "value": 857.9,
            "unit": "ns/op\t     504 B/op\t       7 allocs/op",
            "extra": "1459365 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - ns/op",
            "value": 857.9,
            "unit": "ns/op",
            "extra": "1459365 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - B/op",
            "value": 504,
            "unit": "B/op",
            "extra": "1459365 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1459365 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid",
            "value": 18.85,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "63665256 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - ns/op",
            "value": 18.85,
            "unit": "ns/op",
            "extra": "63665256 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "63665256 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "63665256 times\n4 procs"
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
        "date": 1750953615796,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizeConfirmation_valid",
            "value": 211.7,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5665897 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - ns/op",
            "value": 211.7,
            "unit": "ns/op",
            "extra": "5665897 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5665897 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5665897 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus",
            "value": 758.8,
            "unit": "ns/op\t     368 B/op\t       7 allocs/op",
            "extra": "1598991 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - ns/op",
            "value": 758.8,
            "unit": "ns/op",
            "extra": "1598991 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1598991 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1598991 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate",
            "value": 1532,
            "unit": "ns/op\t     952 B/op\t      19 allocs/op",
            "extra": "697696 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - ns/op",
            "value": 1532,
            "unit": "ns/op",
            "extra": "697696 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - B/op",
            "value": 952,
            "unit": "B/op",
            "extra": "697696 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "697696 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag",
            "value": 823.4,
            "unit": "ns/op\t     504 B/op\t       7 allocs/op",
            "extra": "1454518 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - ns/op",
            "value": 823.4,
            "unit": "ns/op",
            "extra": "1454518 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - B/op",
            "value": 504,
            "unit": "B/op",
            "extra": "1454518 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1454518 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid",
            "value": 18.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "61823480 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - ns/op",
            "value": 18.9,
            "unit": "ns/op",
            "extra": "61823480 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "61823480 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "61823480 times\n4 procs"
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
        "date": 1750956524583,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizeConfirmation_valid",
            "value": 216.2,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5567722 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - ns/op",
            "value": 216.2,
            "unit": "ns/op",
            "extra": "5567722 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5567722 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5567722 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus",
            "value": 761.5,
            "unit": "ns/op\t     368 B/op\t       7 allocs/op",
            "extra": "1589493 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - ns/op",
            "value": 761.5,
            "unit": "ns/op",
            "extra": "1589493 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1589493 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1589493 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate",
            "value": 1486,
            "unit": "ns/op\t     952 B/op\t      19 allocs/op",
            "extra": "718497 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - ns/op",
            "value": 1486,
            "unit": "ns/op",
            "extra": "718497 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - B/op",
            "value": 952,
            "unit": "B/op",
            "extra": "718497 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "718497 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag",
            "value": 826.7,
            "unit": "ns/op\t     504 B/op\t       7 allocs/op",
            "extra": "1412055 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - ns/op",
            "value": 826.7,
            "unit": "ns/op",
            "extra": "1412055 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - B/op",
            "value": 504,
            "unit": "B/op",
            "extra": "1412055 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1412055 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid",
            "value": 18.89,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "63855681 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - ns/op",
            "value": 18.89,
            "unit": "ns/op",
            "extra": "63855681 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "63855681 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "63855681 times\n4 procs"
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
        "date": 1750965960324,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizeConfirmation_valid",
            "value": 216.2,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5572759 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - ns/op",
            "value": 216.2,
            "unit": "ns/op",
            "extra": "5572759 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5572759 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_valid - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5572759 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus",
            "value": 763.3,
            "unit": "ns/op\t     368 B/op\t       7 allocs/op",
            "extra": "1565108 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - ns/op",
            "value": 763.3,
            "unit": "ns/op",
            "extra": "1565108 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1565108 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidStatus - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1565108 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate",
            "value": 1518,
            "unit": "ns/op\t     952 B/op\t      19 allocs/op",
            "extra": "722031 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - ns/op",
            "value": 1518,
            "unit": "ns/op",
            "extra": "722031 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - B/op",
            "value": 952,
            "unit": "B/op",
            "extra": "722031 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidExpiryDate - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "722031 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag",
            "value": 834.3,
            "unit": "ns/op\t     504 B/op\t       7 allocs/op",
            "extra": "1439898 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - ns/op",
            "value": 834.3,
            "unit": "ns/op",
            "extra": "1439898 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - B/op",
            "value": 504,
            "unit": "B/op",
            "extra": "1439898 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation_invalidParentIdTag - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1439898 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid",
            "value": 18.86,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "60538550 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - ns/op",
            "value": 18.86,
            "unit": "ns/op",
            "extra": "60538550 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "60538550 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "60538550 times\n4 procs"
          }
        ]
      }
    ]
  }
}