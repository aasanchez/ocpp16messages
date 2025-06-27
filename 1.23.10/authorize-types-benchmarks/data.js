window.BENCHMARK_DATA = {
  "lastUpdate": 1751017146517,
  "repoUrl": "https://github.com/aasanchez/ocpp16messages",
  "entries": {
    "Go Benchmark (authorize-types - Go 1.23.10)": [
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
        "date": 1751017145167,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 97.26,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "11422465 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 97.26,
            "unit": "ns/op",
            "extra": "11422465 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "11422465 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "11422465 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.3112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.3112,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal",
            "value": 0.622,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.622,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full",
            "value": 0.6229,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.6229,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal",
            "value": 0.6335,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6335,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full",
            "value": 0.6262,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6262,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal",
            "value": 98.75,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12135118 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 98.75,
            "unit": "ns/op",
            "extra": "12135118 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12135118 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12135118 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 195.1,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "6157288 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 195.1,
            "unit": "ns/op",
            "extra": "6157288 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "6157288 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6157288 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.808,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "428104521 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.808,
            "unit": "ns/op",
            "extra": "428104521 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "428104521 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "428104521 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 101.8,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11546571 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 101.8,
            "unit": "ns/op",
            "extra": "11546571 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11546571 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11546571 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create",
            "value": 16.33,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "71396910 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - ns/op",
            "value": 16.33,
            "unit": "ns/op",
            "extra": "71396910 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "71396910 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "71396910 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3113,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3113,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid",
            "value": 18.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65622438 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.1,
            "unit": "ns/op",
            "extra": "65622438 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65622438 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65622438 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 158.9,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7451072 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 158.9,
            "unit": "ns/op",
            "extra": "7451072 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7451072 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7451072 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.3114,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.3114,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
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
          "id": "3ccab0be6434c00488be9c30f804a4a1330c8e72",
          "message": "chore(benchmark.yml): update Go version matrix and adjust benchmark output paths for better organization",
          "timestamp": "2025-06-27T09:37:20Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/3ccab0be6434c00488be9c30f804a4a1330c8e72"
        },
        "date": 1751017145167,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 97.26,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "11422465 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 97.26,
            "unit": "ns/op",
            "extra": "11422465 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "11422465 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "11422465 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.3112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.3112,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal",
            "value": 0.622,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.622,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full",
            "value": 0.6229,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.6229,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal",
            "value": 0.6335,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6335,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full",
            "value": 0.6262,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6262,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal",
            "value": 98.75,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12135118 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 98.75,
            "unit": "ns/op",
            "extra": "12135118 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12135118 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12135118 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 195.1,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "6157288 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 195.1,
            "unit": "ns/op",
            "extra": "6157288 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "6157288 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6157288 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.808,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "428104521 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.808,
            "unit": "ns/op",
            "extra": "428104521 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "428104521 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "428104521 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 101.8,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11546571 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 101.8,
            "unit": "ns/op",
            "extra": "11546571 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11546571 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11546571 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create",
            "value": 16.33,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "71396910 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - ns/op",
            "value": 16.33,
            "unit": "ns/op",
            "extra": "71396910 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "71396910 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "71396910 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3113,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3113,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid",
            "value": 18.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65622438 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.1,
            "unit": "ns/op",
            "extra": "65622438 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65622438 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65622438 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 158.9,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7451072 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 158.9,
            "unit": "ns/op",
            "extra": "7451072 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7451072 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7451072 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.3114,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.3114,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          }
        ]
      }
    ]
  }
}