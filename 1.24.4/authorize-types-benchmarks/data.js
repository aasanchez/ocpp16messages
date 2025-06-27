window.BENCHMARK_DATA = {
  "lastUpdate": 1751017148972,
  "repoUrl": "https://github.com/aasanchez/ocpp16messages",
  "entries": {
    "Go Benchmark (authorize-types - Go 1.24.4)": [
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
        "date": 1751017143642,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 95.58,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 95.58,
            "unit": "ns/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.312,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.312,
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
            "value": 0.3897,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.3897,
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
            "value": 0.3906,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.3906,
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
            "value": 0.6232,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6232,
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
            "value": 0.6234,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6234,
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
            "value": 99.11,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 99.11,
            "unit": "ns/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 203.9,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 203.9,
            "unit": "ns/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.804,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.804,
            "unit": "ns/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 104.2,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 104.2,
            "unit": "ns/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create",
            "value": 16.65,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - ns/op",
            "value": 16.65,
            "unit": "ns/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3118,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3118,
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
            "value": 18.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.57,
            "unit": "ns/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 153.3,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 153.3,
            "unit": "ns/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.3112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.3112,
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
        "date": 1751017143642,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 95.58,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 95.58,
            "unit": "ns/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.312,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.312,
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
            "value": 0.3897,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.3897,
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
            "value": 0.3906,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.3906,
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
            "value": 0.6232,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6232,
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
            "value": 0.6234,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6234,
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
            "value": 99.11,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 99.11,
            "unit": "ns/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 203.9,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 203.9,
            "unit": "ns/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.804,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.804,
            "unit": "ns/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 104.2,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 104.2,
            "unit": "ns/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create",
            "value": 16.65,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - ns/op",
            "value": 16.65,
            "unit": "ns/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3118,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3118,
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
            "value": 18.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.57,
            "unit": "ns/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 153.3,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 153.3,
            "unit": "ns/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.3112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.3112,
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
        "date": 1751017143642,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 95.58,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 95.58,
            "unit": "ns/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.312,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.312,
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
            "value": 0.3897,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.3897,
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
            "value": 0.3906,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.3906,
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
            "value": 0.6232,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6232,
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
            "value": 0.6234,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6234,
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
            "value": 99.11,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 99.11,
            "unit": "ns/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 203.9,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 203.9,
            "unit": "ns/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.804,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.804,
            "unit": "ns/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 104.2,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 104.2,
            "unit": "ns/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create",
            "value": 16.65,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - ns/op",
            "value": 16.65,
            "unit": "ns/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3118,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3118,
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
            "value": 18.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.57,
            "unit": "ns/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 153.3,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 153.3,
            "unit": "ns/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.3112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.3112,
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
        "date": 1751017143642,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 95.58,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 95.58,
            "unit": "ns/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.312,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.312,
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
            "value": 0.3897,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.3897,
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
            "value": 0.3906,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.3906,
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
            "value": 0.6232,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6232,
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
            "value": 0.6234,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6234,
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
            "value": 99.11,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 99.11,
            "unit": "ns/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 203.9,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 203.9,
            "unit": "ns/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.804,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.804,
            "unit": "ns/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 104.2,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 104.2,
            "unit": "ns/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create",
            "value": 16.65,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - ns/op",
            "value": 16.65,
            "unit": "ns/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3118,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3118,
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
            "value": 18.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.57,
            "unit": "ns/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 153.3,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 153.3,
            "unit": "ns/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.3112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.3112,
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
        "date": 1751017143642,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 95.58,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 95.58,
            "unit": "ns/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12573523 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.312,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.312,
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
            "value": 0.3897,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.3897,
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
            "value": 0.3906,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.3906,
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
            "value": 0.6232,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6232,
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
            "value": 0.6234,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6234,
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
            "value": 99.11,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 99.11,
            "unit": "ns/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12115426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 203.9,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 203.9,
            "unit": "ns/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5940285 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.804,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.804,
            "unit": "ns/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "426908140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 104.2,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 104.2,
            "unit": "ns/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11424097 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create",
            "value": 16.65,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - ns/op",
            "value": 16.65,
            "unit": "ns/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "71078934 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3118,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3118,
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
            "value": 18.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.57,
            "unit": "ns/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65383785 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 153.3,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 153.3,
            "unit": "ns/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7833080 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.3112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.3112,
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