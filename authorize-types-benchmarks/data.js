window.BENCHMARK_DATA = {
  "lastUpdate": 1750965962765,
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
        "date": 1750430769906,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 94.99,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "11838712 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 94.99,
            "unit": "ns/op",
            "extra": "11838712 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "11838712 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "11838712 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.3173,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.3173,
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
            "value": 0.3908,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.3908,
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
            "value": 0.4279,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.4279,
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
            "value": 0.6244,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6244,
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
            "value": 0.6273,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6273,
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
            "value": 100.3,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12105286 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 100.3,
            "unit": "ns/op",
            "extra": "12105286 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12105286 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12105286 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 203,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5900230 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 203,
            "unit": "ns/op",
            "extra": "5900230 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5900230 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5900230 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.814,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "424368100 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.814,
            "unit": "ns/op",
            "extra": "424368100 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "424368100 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "424368100 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 103.3,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11589694 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 103.3,
            "unit": "ns/op",
            "extra": "11589694 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11589694 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11589694 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid",
            "value": 16.86,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "69764500 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - ns/op",
            "value": 16.86,
            "unit": "ns/op",
            "extra": "69764500 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "69764500 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "69764500 times\n4 procs"
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
            "value": 18.46,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "63760989 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.46,
            "unit": "ns/op",
            "extra": "63760989 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "63760989 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "63760989 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 152.5,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7654905 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 152.5,
            "unit": "ns/op",
            "extra": "7654905 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7654905 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7654905 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.3117,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.3117,
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
          "id": "afe35890a08c7b935f1abb0241c8954cbb3b7ff7",
          "message": "chore(benchmark.yml): rename benchmark job and add authorize types benchmarks job for better clarity and organization",
          "timestamp": "2025-06-20T14:54:44Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/afe35890a08c7b935f1abb0241c8954cbb3b7ff7"
        },
        "date": 1750431424638,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 95.62,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12527424 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 95.62,
            "unit": "ns/op",
            "extra": "12527424 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12527424 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12527424 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.3119,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.3119,
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
            "value": 0.3896,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.3896,
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
            "value": 0.6238,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6238,
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
            "value": 0.6238,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6238,
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
            "value": 99.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12019338 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 99.9,
            "unit": "ns/op",
            "extra": "12019338 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12019338 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12019338 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 204.6,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5978290 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 204.6,
            "unit": "ns/op",
            "extra": "5978290 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5978290 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5978290 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.802,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "428747360 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.802,
            "unit": "ns/op",
            "extra": "428747360 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "428747360 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "428747360 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 102.3,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11527447 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 102.3,
            "unit": "ns/op",
            "extra": "11527447 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11527447 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11527447 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid",
            "value": 16.68,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "70895730 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - ns/op",
            "value": 16.68,
            "unit": "ns/op",
            "extra": "70895730 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "70895730 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "70895730 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3115,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3115,
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
            "value": 18.63,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "60443467 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.63,
            "unit": "ns/op",
            "extra": "60443467 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "60443467 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "60443467 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 151.7,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7859797 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 151.7,
            "unit": "ns/op",
            "extra": "7859797 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7859797 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7859797 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.3118,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.3118,
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
          "id": "1cb2bcb95adc26e49b8c1656438a61290f8db851",
          "message": "chore(golangci.yml): replace disabled linter godot with nilnil for better configuration clarity\nchore(ci_string_test.go): update comment for linting to clarify intentional control character test",
          "timestamp": "2025-06-25T12:23:52Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/1cb2bcb95adc26e49b8c1656438a61290f8db851"
        },
        "date": 1750861826807,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 94.71,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12666964 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 94.71,
            "unit": "ns/op",
            "extra": "12666964 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12666964 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12666964 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.3113,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.3113,
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
            "value": 0.3894,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.3894,
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
            "value": 1.285,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 1.285,
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
            "value": 0.6238,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6238,
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
            "value": 99.01,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12139734 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 99.01,
            "unit": "ns/op",
            "extra": "12139734 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12139734 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12139734 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 203.7,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5951372 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 203.7,
            "unit": "ns/op",
            "extra": "5951372 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5951372 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5951372 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.806,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "427912912 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.806,
            "unit": "ns/op",
            "extra": "427912912 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "427912912 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "427912912 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 102.4,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11569512 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 102.4,
            "unit": "ns/op",
            "extra": "11569512 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11569512 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11569512 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid",
            "value": 16.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "70031268 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - ns/op",
            "value": 16.8,
            "unit": "ns/op",
            "extra": "70031268 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "70031268 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "70031268 times\n4 procs"
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
            "value": 18.63,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "63805035 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.63,
            "unit": "ns/op",
            "extra": "63805035 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "63805035 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "63805035 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 150.6,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7837960 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 150.6,
            "unit": "ns/op",
            "extra": "7837960 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7837960 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7837960 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.3116,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.3116,
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
          "id": "c0788140c64012d48e15b39087268e56da865a57",
          "message": "feat(benchmark): add Go benchmarks for Authorize and Authorize Types with GitHub Actions workflow",
          "timestamp": "2025-06-26T13:55:20Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/c0788140c64012d48e15b39087268e56da865a57"
        },
        "date": 1750947190531,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 94.56,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12668322 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 94.56,
            "unit": "ns/op",
            "extra": "12668322 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12668322 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12668322 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.3113,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.3113,
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
            "value": 0.3894,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.3894,
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
            "value": 0.3894,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.3894,
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
            "value": 0.6229,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6229,
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
            "value": 0.6227,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6227,
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
            "value": 99.04,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12118537 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 99.04,
            "unit": "ns/op",
            "extra": "12118537 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12118537 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12118537 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 200.6,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5971179 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 200.6,
            "unit": "ns/op",
            "extra": "5971179 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5971179 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5971179 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.802,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "428484087 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.802,
            "unit": "ns/op",
            "extra": "428484087 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "428484087 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "428484087 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 104.9,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11160867 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 104.9,
            "unit": "ns/op",
            "extra": "11160867 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11160867 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11160867 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid",
            "value": 16.86,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "69232870 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - ns/op",
            "value": 16.86,
            "unit": "ns/op",
            "extra": "69232870 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "69232870 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "69232870 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3122,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3122,
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
            "value": 18.68,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "63771439 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.68,
            "unit": "ns/op",
            "extra": "63771439 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "63771439 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "63771439 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 151.9,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7782684 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 151.9,
            "unit": "ns/op",
            "extra": "7782684 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7782684 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7782684 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.3111,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.3111,
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
    ],
    "Go Benchmark (authorize-types)": [
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
        "date": 1750949731727,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 94.74,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "11578028 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 94.74,
            "unit": "ns/op",
            "extra": "11578028 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "11578028 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "11578028 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.3114,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.3114,
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
            "value": 0.3902,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.3902,
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
            "value": 0.3893,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.3893,
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
            "value": 0.6239,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6239,
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
            "value": 98.91,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12134386 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 98.91,
            "unit": "ns/op",
            "extra": "12134386 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12134386 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12134386 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 201.1,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5976116 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 201.1,
            "unit": "ns/op",
            "extra": "5976116 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5976116 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5976116 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.801,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "378932426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.801,
            "unit": "ns/op",
            "extra": "378932426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "378932426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "378932426 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 101.8,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11452298 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 101.8,
            "unit": "ns/op",
            "extra": "11452298 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11452298 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11452298 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid",
            "value": 16.83,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "68721889 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - ns/op",
            "value": 16.83,
            "unit": "ns/op",
            "extra": "68721889 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "68721889 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "68721889 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3123,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3123,
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
            "value": 19,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "63909963 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 19,
            "unit": "ns/op",
            "extra": "63909963 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "63909963 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "63909963 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 151.1,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7806561 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 151.1,
            "unit": "ns/op",
            "extra": "7806561 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7806561 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7806561 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.3117,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.3117,
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
        "date": 1750953618131,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 95.99,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12480692 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 95.99,
            "unit": "ns/op",
            "extra": "12480692 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12480692 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12480692 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.3115,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.3115,
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
            "value": 0.3892,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.3892,
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
            "value": 0.394,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.394,
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
            "value": 0.623,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.623,
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
            "value": 0.6278,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6278,
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
            "value": 100.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "11995543 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 100.1,
            "unit": "ns/op",
            "extra": "11995543 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "11995543 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "11995543 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 202.2,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5951278 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 202.2,
            "unit": "ns/op",
            "extra": "5951278 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5951278 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5951278 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.811,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "427655420 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.811,
            "unit": "ns/op",
            "extra": "427655420 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "427655420 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "427655420 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 102.1,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11587966 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 102.1,
            "unit": "ns/op",
            "extra": "11587966 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11587966 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11587966 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid",
            "value": 16.87,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "68909288 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - ns/op",
            "value": 16.87,
            "unit": "ns/op",
            "extra": "68909288 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "68909288 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "68909288 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3112,
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
            "value": 18.62,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "64308608 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.62,
            "unit": "ns/op",
            "extra": "64308608 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "64308608 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "64308608 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 156.5,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7846720 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 156.5,
            "unit": "ns/op",
            "extra": "7846720 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7846720 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7846720 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.3116,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.3116,
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
        "date": 1750956525904,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 94.86,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12630805 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 94.86,
            "unit": "ns/op",
            "extra": "12630805 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12630805 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12630805 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.3121,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.3121,
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
            "value": 0.3895,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.3895,
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
            "value": 0.3897,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.3897,
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
            "value": 0.6234,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6234,
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
            "value": 0.6243,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6243,
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
            "value": 99.46,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12072870 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 99.46,
            "unit": "ns/op",
            "extra": "12072870 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12072870 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12072870 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 200.3,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "6013615 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 200.3,
            "unit": "ns/op",
            "extra": "6013615 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "6013615 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6013615 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.804,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "404514414 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.804,
            "unit": "ns/op",
            "extra": "404514414 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "404514414 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "404514414 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 108.4,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11478375 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 108.4,
            "unit": "ns/op",
            "extra": "11478375 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11478375 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11478375 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid",
            "value": 16.94,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "72642474 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - ns/op",
            "value": 16.94,
            "unit": "ns/op",
            "extra": "72642474 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "72642474 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "72642474 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3148,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3148,
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
            "value": 18.55,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "62859460 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.55,
            "unit": "ns/op",
            "extra": "62859460 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "62859460 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "62859460 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 153.4,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7807510 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 153.4,
            "unit": "ns/op",
            "extra": "7807510 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7807510 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7807510 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.3174,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.3174,
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
        "date": 1750965962152,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 94.56,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12139102 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 94.56,
            "unit": "ns/op",
            "extra": "12139102 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12139102 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12139102 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.3142,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.3142,
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
            "value": 0.3903,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.3903,
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
            "value": 0.3909,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.3909,
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
            "value": 0.6274,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6274,
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
            "value": 0.6245,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6245,
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
            "value": 98.98,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12059820 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 98.98,
            "unit": "ns/op",
            "extra": "12059820 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12059820 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12059820 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 202.2,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5939020 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 202.2,
            "unit": "ns/op",
            "extra": "5939020 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5939020 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5939020 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.803,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "425902299 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.803,
            "unit": "ns/op",
            "extra": "425902299 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "425902299 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "425902299 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 104,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11426557 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 104,
            "unit": "ns/op",
            "extra": "11426557 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11426557 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11426557 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid",
            "value": 16.76,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "71131904 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - ns/op",
            "value": 16.76,
            "unit": "ns/op",
            "extra": "71131904 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "71131904 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "71131904 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3117,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3117,
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
            "value": 18.52,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65913508 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.52,
            "unit": "ns/op",
            "extra": "65913508 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65913508 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65913508 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 153.2,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7809650 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 153.2,
            "unit": "ns/op",
            "extra": "7809650 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7809650 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7809650 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.312,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.312,
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