window.BENCHMARK_DATA = {
  "lastUpdate": 1750369341873,
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
      }
    ]
  }
}