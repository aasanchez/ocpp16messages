window.BENCHMARK_DATA = {
  "lastUpdate": 1750371789719,
  "repoUrl": "https://github.com/aasanchez/ocpp16messages",
  "entries": {
    "Go Benchmark": [
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
          "id": "e7ad537f55a05bba85fc90acfe449e44436481c3",
          "message": "Fix golint",
          "timestamp": "2025-06-19T19:14:58Z",
          "url": "https://github.com/aasanchez/ocpp16messages/pull/1/commits/e7ad537f55a05bba85fc90acfe449e44436481c3"
        },
        "date": 1750369341574,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRequest_Valid",
            "value": 18.34,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "63439934 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_Valid - ns/op",
            "value": 18.34,
            "unit": "ns/op",
            "extra": "63439934 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "63439934 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "63439934 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 11.84,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 11.84,
            "unit": "ns/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.6222,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.6222,
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
            "value": 0.6326,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.6326,
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
            "value": 0.6224,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.6224,
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
            "value": 0.6241,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6241,
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
            "value": 0.6746,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6746,
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
            "value": 14.64,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "81047932 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 14.64,
            "unit": "ns/op",
            "extra": "81047932 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "81047932 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "81047932 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 113.5,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "10510140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 113.5,
            "unit": "ns/op",
            "extra": "10510140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "10510140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "10510140 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 3.112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "383430625 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 3.112,
            "unit": "ns/op",
            "extra": "383430625 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "383430625 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "383430625 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 103.1,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11547201 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 103.1,
            "unit": "ns/op",
            "extra": "11547201 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11547201 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11547201 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid",
            "value": 16.84,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "68737737 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - ns/op",
            "value": 16.84,
            "unit": "ns/op",
            "extra": "68737737 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "68737737 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "68737737 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3121,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3121,
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
            "value": 18.51,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65275467 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.51,
            "unit": "ns/op",
            "extra": "65275467 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65275467 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65275467 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 149.7,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7940484 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 149.7,
            "unit": "ns/op",
            "extra": "7940484 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7940484 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7940484 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value",
            "value": 0.3113,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Value - ns/op",
            "value": 0.3113,
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
          },
          {
            "name": "BenchmarkCiString20_Create",
            "value": 17.97,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "66334862 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Create - ns/op",
            "value": 17.97,
            "unit": "ns/op",
            "extra": "66334862 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "66334862 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "66334862 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate",
            "value": 16.76,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "72873009 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate - ns/op",
            "value": 16.76,
            "unit": "ns/op",
            "extra": "72873009 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "72873009 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "72873009 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value",
            "value": 0.311,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value - ns/op",
            "value": 0.311,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create",
            "value": 21.51,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "54585068 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create - ns/op",
            "value": 21.51,
            "unit": "ns/op",
            "extra": "54585068 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "54585068 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "54585068 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate",
            "value": 19.62,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "59197710 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate - ns/op",
            "value": 19.62,
            "unit": "ns/op",
            "extra": "59197710 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "59197710 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "59197710 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value",
            "value": 0.3113,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value - ns/op",
            "value": 0.3113,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create",
            "value": 37.03,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "32016890 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create - ns/op",
            "value": 37.03,
            "unit": "ns/op",
            "extra": "32016890 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "32016890 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "32016890 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate",
            "value": 36.51,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "34017477 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate - ns/op",
            "value": 36.51,
            "unit": "ns/op",
            "extra": "34017477 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "34017477 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "34017477 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value",
            "value": 0.3228,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value - ns/op",
            "value": 0.3228,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create",
            "value": 170.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "6963002 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create - ns/op",
            "value": 170.5,
            "unit": "ns/op",
            "extra": "6963002 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "6963002 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "6963002 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate",
            "value": 168.6,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "7043818 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate - ns/op",
            "value": 168.6,
            "unit": "ns/op",
            "extra": "7043818 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "7043818 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "7043818 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value",
            "value": 0.3112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value - ns/op",
            "value": 0.3112,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create",
            "value": 457.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "3688708 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create - ns/op",
            "value": 457.9,
            "unit": "ns/op",
            "extra": "3688708 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "3688708 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "3688708 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate",
            "value": 320.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "3739638 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate - ns/op",
            "value": 320.4,
            "unit": "ns/op",
            "extra": "3739638 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "3739638 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "3739638 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value",
            "value": 0.3112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value - ns/op",
            "value": 0.3112,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid",
            "value": 37.68,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "31870778 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid - ns/op",
            "value": 37.68,
            "unit": "ns/op",
            "extra": "31870778 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "31870778 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "31870778 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat",
            "value": 1021,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat - ns/op",
            "value": 1021,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat - B/op",
            "value": 560,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime",
            "value": 37.82,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "31818844 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime - ns/op",
            "value": 37.82,
            "unit": "ns/op",
            "extra": "31818844 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "31818844 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "31818844 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value",
            "value": 0.3112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value - ns/op",
            "value": 0.3112,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value - allocs/op",
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
          "id": "0cb7a8ec7fa19222f5cd07edc398722f3ee8fc72",
          "message": "Fix golint",
          "timestamp": "2025-06-19T19:14:58Z",
          "url": "https://github.com/aasanchez/ocpp16messages/pull/1/commits/0cb7a8ec7fa19222f5cd07edc398722f3ee8fc72"
        },
        "date": 1750371275614,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRequest_Valid",
            "value": 18.19,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65779970 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_Valid - ns/op",
            "value": 18.19,
            "unit": "ns/op",
            "extra": "65779970 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65779970 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65779970 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 11.84,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 11.84,
            "unit": "ns/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "100000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.6224,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.6224,
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
            "value": 0.6238,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.6238,
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
            "value": 0.6235,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.6235,
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
            "value": 0.6411,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6411,
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
            "value": 0.6734,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6734,
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
            "value": 14.75,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "81737888 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 14.75,
            "unit": "ns/op",
            "extra": "81737888 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "81737888 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "81737888 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 114.3,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "10516098 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 114.3,
            "unit": "ns/op",
            "extra": "10516098 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "10516098 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "10516098 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 3.112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "385231390 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 3.112,
            "unit": "ns/op",
            "extra": "385231390 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "385231390 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "385231390 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 103.8,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11263509 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 103.8,
            "unit": "ns/op",
            "extra": "11263509 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11263509 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11263509 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid",
            "value": 16.89,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "67537863 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - ns/op",
            "value": 16.89,
            "unit": "ns/op",
            "extra": "67537863 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "67537863 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "67537863 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3125,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3125,
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
            "value": 18.69,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "64704519 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.69,
            "unit": "ns/op",
            "extra": "64704519 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "64704519 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "64704519 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 153.3,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7753801 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 153.3,
            "unit": "ns/op",
            "extra": "7753801 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7753801 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7753801 times\n4 procs"
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
          },
          {
            "name": "BenchmarkCiString20_Create",
            "value": 19.12,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "60778807 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Create - ns/op",
            "value": 19.12,
            "unit": "ns/op",
            "extra": "60778807 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "60778807 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "60778807 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate",
            "value": 16.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "71447300 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate - ns/op",
            "value": 16.5,
            "unit": "ns/op",
            "extra": "71447300 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "71447300 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "71447300 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value",
            "value": 0.3117,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value - ns/op",
            "value": 0.3117,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create",
            "value": 21.65,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "55146979 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create - ns/op",
            "value": 21.65,
            "unit": "ns/op",
            "extra": "55146979 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "55146979 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "55146979 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate",
            "value": 19.89,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "58828552 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate - ns/op",
            "value": 19.89,
            "unit": "ns/op",
            "extra": "58828552 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "58828552 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "58828552 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value",
            "value": 0.3143,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value - ns/op",
            "value": 0.3143,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create",
            "value": 38.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "31777172 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create - ns/op",
            "value": 38.4,
            "unit": "ns/op",
            "extra": "31777172 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "31777172 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "31777172 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate",
            "value": 35.89,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "31943547 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate - ns/op",
            "value": 35.89,
            "unit": "ns/op",
            "extra": "31943547 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "31943547 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "31943547 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value",
            "value": 0.3117,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value - ns/op",
            "value": 0.3117,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create",
            "value": 171.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "7021532 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create - ns/op",
            "value": 171.1,
            "unit": "ns/op",
            "extra": "7021532 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "7021532 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "7021532 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate",
            "value": 169.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "7169670 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate - ns/op",
            "value": 169.1,
            "unit": "ns/op",
            "extra": "7169670 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "7169670 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "7169670 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value",
            "value": 0.3116,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value - ns/op",
            "value": 0.3116,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create",
            "value": 325,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "3678194 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create - ns/op",
            "value": 325,
            "unit": "ns/op",
            "extra": "3678194 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "3678194 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "3678194 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate",
            "value": 320.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "3743895 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate - ns/op",
            "value": 320.7,
            "unit": "ns/op",
            "extra": "3743895 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "3743895 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "3743895 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value",
            "value": 0.3117,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value - ns/op",
            "value": 0.3117,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid",
            "value": 37.68,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "31706937 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid - ns/op",
            "value": 37.68,
            "unit": "ns/op",
            "extra": "31706937 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "31706937 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "31706937 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat",
            "value": 1030,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat - ns/op",
            "value": 1030,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat - B/op",
            "value": 560,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime",
            "value": 37.99,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "31724616 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime - ns/op",
            "value": 37.99,
            "unit": "ns/op",
            "extra": "31724616 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "31724616 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "31724616 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value",
            "value": 0.3121,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value - ns/op",
            "value": 0.3121,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value - allocs/op",
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
          "id": "0cb7a8ec7fa19222f5cd07edc398722f3ee8fc72",
          "message": "test: improve test coverage and error handling in authorization tests and update parent IdTag for consistency",
          "timestamp": "2025-06-19T22:13:27Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/0cb7a8ec7fa19222f5cd07edc398722f3ee8fc72"
        },
        "date": 1750371414184,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRequest_Valid",
            "value": 18.47,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "66486399 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_Valid - ns/op",
            "value": 18.47,
            "unit": "ns/op",
            "extra": "66486399 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "66486399 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "66486399 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_valid",
            "value": 201.7,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5971390 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_valid - ns/op",
            "value": 201.7,
            "unit": "ns/op",
            "extra": "5971390 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_valid - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5971390 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_valid - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5971390 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidStatus",
            "value": 757.7,
            "unit": "ns/op\t     368 B/op\t       7 allocs/op",
            "extra": "1606448 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidStatus - ns/op",
            "value": 757.7,
            "unit": "ns/op",
            "extra": "1606448 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidStatus - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1606448 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidStatus - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1606448 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidExpiryDate",
            "value": 1490,
            "unit": "ns/op\t     952 B/op\t      19 allocs/op",
            "extra": "705643 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidExpiryDate - ns/op",
            "value": 1490,
            "unit": "ns/op",
            "extra": "705643 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidExpiryDate - B/op",
            "value": 952,
            "unit": "B/op",
            "extra": "705643 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidExpiryDate - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "705643 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidParentIdTag",
            "value": 826.8,
            "unit": "ns/op\t     504 B/op\t       7 allocs/op",
            "extra": "1451784 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidParentIdTag - ns/op",
            "value": 826.8,
            "unit": "ns/op",
            "extra": "1451784 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidParentIdTag - B/op",
            "value": 504,
            "unit": "B/op",
            "extra": "1451784 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidParentIdTag - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1451784 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 94.49,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12704803 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 94.49,
            "unit": "ns/op",
            "extra": "12704803 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12704803 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12704803 times\n4 procs"
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
            "value": 0.6242,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6242,
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
            "value": 0.625,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.625,
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
            "value": 98.83,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12109754 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 98.83,
            "unit": "ns/op",
            "extra": "12109754 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12109754 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12109754 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 202.2,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5871228 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 202.2,
            "unit": "ns/op",
            "extra": "5871228 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5871228 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5871228 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.805,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "401683441 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.805,
            "unit": "ns/op",
            "extra": "401683441 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "401683441 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "401683441 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 103.8,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11624793 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 103.8,
            "unit": "ns/op",
            "extra": "11624793 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11624793 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11624793 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid",
            "value": 16.44,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "70680793 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - ns/op",
            "value": 16.44,
            "unit": "ns/op",
            "extra": "70680793 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "70680793 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "70680793 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value",
            "value": 0.3172,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_Value - ns/op",
            "value": 0.3172,
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
            "value": 18.39,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "62876824 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.39,
            "unit": "ns/op",
            "extra": "62876824 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "62876824 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "62876824 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 162.7,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7774153 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 162.7,
            "unit": "ns/op",
            "extra": "7774153 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7774153 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7774153 times\n4 procs"
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
          },
          {
            "name": "BenchmarkCiString20_Create",
            "value": 18.03,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "64283570 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Create - ns/op",
            "value": 18.03,
            "unit": "ns/op",
            "extra": "64283570 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "64283570 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "64283570 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate",
            "value": 16.71,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "69095336 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate - ns/op",
            "value": 16.71,
            "unit": "ns/op",
            "extra": "69095336 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "69095336 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "69095336 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value",
            "value": 0.3131,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value - ns/op",
            "value": 0.3131,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create",
            "value": 21.83,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "51608514 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create - ns/op",
            "value": 21.83,
            "unit": "ns/op",
            "extra": "51608514 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "51608514 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "51608514 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate",
            "value": 20.27,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "57941768 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate - ns/op",
            "value": 20.27,
            "unit": "ns/op",
            "extra": "57941768 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "57941768 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "57941768 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value",
            "value": 0.3127,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value - ns/op",
            "value": 0.3127,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create",
            "value": 37.95,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "30384399 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create - ns/op",
            "value": 37.95,
            "unit": "ns/op",
            "extra": "30384399 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "30384399 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "30384399 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate",
            "value": 36.4,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "32133204 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate - ns/op",
            "value": 36.4,
            "unit": "ns/op",
            "extra": "32133204 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "32133204 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "32133204 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value",
            "value": 0.3116,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value - ns/op",
            "value": 0.3116,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create",
            "value": 177.5,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "6657043 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create - ns/op",
            "value": 177.5,
            "unit": "ns/op",
            "extra": "6657043 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "6657043 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "6657043 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate",
            "value": 178.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "6636571 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate - ns/op",
            "value": 178.1,
            "unit": "ns/op",
            "extra": "6636571 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "6636571 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "6636571 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value",
            "value": 0.3113,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value - ns/op",
            "value": 0.3113,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create",
            "value": 340.1,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "3417102 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create - ns/op",
            "value": 340.1,
            "unit": "ns/op",
            "extra": "3417102 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "3417102 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "3417102 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate",
            "value": 338.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "3484288 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate - ns/op",
            "value": 338.8,
            "unit": "ns/op",
            "extra": "3484288 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "3484288 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "3484288 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value",
            "value": 0.3154,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value - ns/op",
            "value": 0.3154,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid",
            "value": 37.37,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "31893102 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid - ns/op",
            "value": 37.37,
            "unit": "ns/op",
            "extra": "31893102 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "31893102 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "31893102 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat",
            "value": 1050,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat - ns/op",
            "value": 1050,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat - B/op",
            "value": 560,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime",
            "value": 37.44,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "31975972 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime - ns/op",
            "value": 37.44,
            "unit": "ns/op",
            "extra": "31975972 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "31975972 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "31975972 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value",
            "value": 0.3133,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value - ns/op",
            "value": 0.3133,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
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
          "id": "0cb7a8ec7fa19222f5cd07edc398722f3ee8fc72",
          "message": "test: improve test coverage and error handling in authorization tests and update parent IdTag for consistency",
          "timestamp": "2025-06-20T00:13:27+02:00",
          "tree_id": "c33a7f53d707e57bd3d77a07191151495b154747",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/0cb7a8ec7fa19222f5cd07edc398722f3ee8fc72"
        },
        "date": 1750371789436,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkRequest_Valid",
            "value": 18.51,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "62166775 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_Valid - ns/op",
            "value": 18.51,
            "unit": "ns/op",
            "extra": "62166775 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "62166775 times\n4 procs"
          },
          {
            "name": "BenchmarkRequest_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "62166775 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_valid",
            "value": 201.6,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "5907283 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_valid - ns/op",
            "value": 201.6,
            "unit": "ns/op",
            "extra": "5907283 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_valid - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "5907283 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_valid - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "5907283 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidStatus",
            "value": 740.4,
            "unit": "ns/op\t     368 B/op\t       7 allocs/op",
            "extra": "1617391 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidStatus - ns/op",
            "value": 740.4,
            "unit": "ns/op",
            "extra": "1617391 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidStatus - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1617391 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidStatus - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1617391 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidExpiryDate",
            "value": 1478,
            "unit": "ns/op\t     952 B/op\t      19 allocs/op",
            "extra": "713178 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidExpiryDate - ns/op",
            "value": 1478,
            "unit": "ns/op",
            "extra": "713178 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidExpiryDate - B/op",
            "value": 952,
            "unit": "B/op",
            "extra": "713178 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidExpiryDate - allocs/op",
            "value": 19,
            "unit": "allocs/op",
            "extra": "713178 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidParentIdTag",
            "value": 812.7,
            "unit": "ns/op\t     504 B/op\t       7 allocs/op",
            "extra": "1447214 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidParentIdTag - ns/op",
            "value": 812.7,
            "unit": "ns/op",
            "extra": "1447214 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidParentIdTag - B/op",
            "value": 504,
            "unit": "B/op",
            "extra": "1447214 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmation_invalidParentIdTag - allocs/op",
            "value": 7,
            "unit": "allocs/op",
            "extra": "1447214 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create",
            "value": 94.71,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12623238 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - ns/op",
            "value": 94.71,
            "unit": "ns/op",
            "extra": "12623238 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12623238 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12623238 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value",
            "value": 0.3109,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus_Value - ns/op",
            "value": 0.3109,
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
            "value": 0.3891,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_minimal - ns/op",
            "value": 0.3891,
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
            "value": 0.3898,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Validate_full - ns/op",
            "value": 0.3898,
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
            "value": 0.6236,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_minimal - ns/op",
            "value": 0.6236,
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
            "value": 0.6295,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayload_Value_full - ns/op",
            "value": 0.6295,
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
            "value": 98.49,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "12181664 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - ns/op",
            "value": 98.49,
            "unit": "ns/op",
            "extra": "12181664 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "12181664 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "12181664 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full",
            "value": 199.2,
            "unit": "ns/op\t      48 B/op\t       2 allocs/op",
            "extra": "6032977 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - ns/op",
            "value": 199.2,
            "unit": "ns/op",
            "extra": "6032977 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - B/op",
            "value": 48,
            "unit": "B/op",
            "extra": "6032977 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Create_Full - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "6032977 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal",
            "value": 2.801,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "425800629 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - ns/op",
            "value": 2.801,
            "unit": "ns/op",
            "extra": "425800629 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "425800629 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Minimal - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "425800629 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full",
            "value": 110,
            "unit": "ns/op\t      56 B/op\t       3 allocs/op",
            "extra": "11515135 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - ns/op",
            "value": 110,
            "unit": "ns/op",
            "extra": "11515135 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - B/op",
            "value": 56,
            "unit": "B/op",
            "extra": "11515135 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo_Value_Full - allocs/op",
            "value": 3,
            "unit": "allocs/op",
            "extra": "11515135 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid",
            "value": 16.51,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "68869495 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - ns/op",
            "value": 16.51,
            "unit": "ns/op",
            "extra": "68869495 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "68869495 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken_CreateValid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "68869495 times\n4 procs"
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
            "value": 18.31,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "65903097 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - ns/op",
            "value": 18.31,
            "unit": "ns/op",
            "extra": "65903097 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "65903097 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "65903097 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty",
            "value": 150.5,
            "unit": "ns/op\t      64 B/op\t       2 allocs/op",
            "extra": "7866352 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - ns/op",
            "value": 150.5,
            "unit": "ns/op",
            "extra": "7866352 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - B/op",
            "value": 64,
            "unit": "B/op",
            "extra": "7866352 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayload_Validate_empty - allocs/op",
            "value": 2,
            "unit": "allocs/op",
            "extra": "7866352 times\n4 procs"
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
          },
          {
            "name": "BenchmarkCiString20_Create",
            "value": 18.11,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "62940550 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Create - ns/op",
            "value": 18.11,
            "unit": "ns/op",
            "extra": "62940550 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "62940550 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "62940550 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate",
            "value": 16.6,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "73603213 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate - ns/op",
            "value": 16.6,
            "unit": "ns/op",
            "extra": "73603213 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "73603213 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "73603213 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value",
            "value": 0.3115,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value - ns/op",
            "value": 0.3115,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString20_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create",
            "value": 21.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "54842694 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create - ns/op",
            "value": 21.57,
            "unit": "ns/op",
            "extra": "54842694 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "54842694 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "54842694 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate",
            "value": 20.01,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "57779766 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate - ns/op",
            "value": 20.01,
            "unit": "ns/op",
            "extra": "57779766 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "57779766 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "57779766 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value",
            "value": 0.3115,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value - ns/op",
            "value": 0.3115,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString25_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create",
            "value": 37.51,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "31113440 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create - ns/op",
            "value": 37.51,
            "unit": "ns/op",
            "extra": "31113440 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "31113440 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "31113440 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate",
            "value": 36.08,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "32764143 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate - ns/op",
            "value": 36.08,
            "unit": "ns/op",
            "extra": "32764143 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "32764143 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "32764143 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value",
            "value": 0.3117,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value - ns/op",
            "value": 0.3117,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString50_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create",
            "value": 170,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "7123011 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create - ns/op",
            "value": 170,
            "unit": "ns/op",
            "extra": "7123011 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "7123011 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "7123011 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate",
            "value": 170,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "7096053 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate - ns/op",
            "value": 170,
            "unit": "ns/op",
            "extra": "7096053 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "7096053 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "7096053 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value",
            "value": 0.3118,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value - ns/op",
            "value": 0.3118,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString255_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create",
            "value": 323.9,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "3719566 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create - ns/op",
            "value": 323.9,
            "unit": "ns/op",
            "extra": "3719566 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "3719566 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Create - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "3719566 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate",
            "value": 323.2,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "3755959 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate - ns/op",
            "value": 323.2,
            "unit": "ns/op",
            "extra": "3755959 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "3755959 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Validate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "3755959 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value",
            "value": 0.3135,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value - ns/op",
            "value": 0.3135,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkCiString500_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid",
            "value": 37.37,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "31920452 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid - ns/op",
            "value": 37.37,
            "unit": "ns/op",
            "extra": "31920452 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "31920452 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_Valid - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "31920452 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat",
            "value": 1025,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat - ns/op",
            "value": 1025,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat - B/op",
            "value": 560,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_InvalidFormat - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime",
            "value": 37.38,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "32046552 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime - ns/op",
            "value": 37.38,
            "unit": "ns/op",
            "extra": "32046552 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "32046552 times\n4 procs"
          },
          {
            "name": "BenchmarkNewDateTime_ZeroTime - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "32046552 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value",
            "value": 0.3116,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value - ns/op",
            "value": 0.3116,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkDateTime_Value - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          }
        ]
      }
    ]
  }
}