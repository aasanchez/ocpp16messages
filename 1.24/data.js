window.BENCHMARK_DATA = {
  "lastUpdate": 1752089342768,
  "repoUrl": "https://github.com/aasanchez/ocpp16messages",
  "entries": {
    "Go Benchmark on 1.24)": [
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
          "id": "e12752e5813e98274ff2cbb518ebd8e0dfd19e3a",
          "message": "chore(benchmark.yml): add permissions for deployments and contents to enable GitHub Pages updates",
          "timestamp": "2025-07-09T19:01:33Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/e12752e5813e98274ff2cbb518ebd8e0dfd19e3a"
        },
        "date": 1752087977501,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizeConfirmation",
            "value": 99.48,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "241339702 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation - ns/op",
            "value": 99.48,
            "unit": "ns/op",
            "extra": "241339702 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "241339702 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "241339702 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest",
            "value": 18.87,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest - ns/op",
            "value": 18.87,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus",
            "value": 94.98,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "251932801 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus - ns/op",
            "value": 94.98,
            "unit": "ns/op",
            "extra": "251932801 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "251932801 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "251932801 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue",
            "value": 0.3115,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue - ns/op",
            "value": 0.3115,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValidate",
            "value": 0.3117,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValidate - ns/op",
            "value": 0.3117,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValidate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValidate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValue",
            "value": 0.3119,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValue - ns/op",
            "value": 0.3119,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValue - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo",
            "value": 98.65,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "243239420 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo - ns/op",
            "value": 98.65,
            "unit": "ns/op",
            "extra": "243239420 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "243239420 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "243239420 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue",
            "value": 2.804,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue - ns/op",
            "value": 2.804,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken",
            "value": 0.3112,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken - ns/op",
            "value": 0.3112,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTokenValue",
            "value": 0.3111,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTokenValue - ns/op",
            "value": 0.3111,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTokenValue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTokenValue - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate",
            "value": 406.4,
            "unit": "ns/op\t     240 B/op\t       4 allocs/op",
            "extra": "59504017 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate - ns/op",
            "value": 406.4,
            "unit": "ns/op",
            "extra": "59504017 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "59504017 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "59504017 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValue",
            "value": 0.3117,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValue - ns/op",
            "value": 0.3117,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValue - allocs/op",
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
          "id": "f5baa987a60024680acff18d1a87937c77e84e64",
          "message": "chore(benchmark.yml): remove unnecessary permissions for deployments and contents to streamline workflow configuration",
          "timestamp": "2025-07-09T19:11:22Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/f5baa987a60024680acff18d1a87937c77e84e64"
        },
        "date": 1752088465010,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizeConfirmation",
            "value": 97.72,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "245377395 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation - ns/op",
            "value": 97.72,
            "unit": "ns/op",
            "extra": "245377395 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "245377395 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "245377395 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest",
            "value": 18.85,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest - ns/op",
            "value": 18.85,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus",
            "value": 94.21,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "255420224 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus - ns/op",
            "value": 94.21,
            "unit": "ns/op",
            "extra": "255420224 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "255420224 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "255420224 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue",
            "value": 0.3113,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue - ns/op",
            "value": 0.3113,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValidate",
            "value": 0.3118,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValidate - ns/op",
            "value": 0.3118,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValidate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValidate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValue",
            "value": 0.311,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValue - ns/op",
            "value": 0.311,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValue - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo",
            "value": 96.91,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "247855917 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo - ns/op",
            "value": 96.91,
            "unit": "ns/op",
            "extra": "247855917 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "247855917 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "247855917 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue",
            "value": 2.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue - ns/op",
            "value": 2.8,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken",
            "value": 0.3117,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken - ns/op",
            "value": 0.3117,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTokenValue",
            "value": 0.3115,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTokenValue - ns/op",
            "value": 0.3115,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTokenValue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTokenValue - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate",
            "value": 408.2,
            "unit": "ns/op\t     240 B/op\t       4 allocs/op",
            "extra": "58640216 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate - ns/op",
            "value": 408.2,
            "unit": "ns/op",
            "extra": "58640216 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "58640216 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "58640216 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValue",
            "value": 0.3117,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValue - ns/op",
            "value": 0.3117,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValue - allocs/op",
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
          "id": "2c1d74e41097325660c9ab51c389b5dd7e6cbcbc",
          "message": "chore(lint.yml): rename analysis jobs for clarity and consistency in workflow configuration",
          "timestamp": "2025-07-09T19:24:29Z",
          "url": "https://github.com/aasanchez/ocpp16messages/commit/2c1d74e41097325660c9ab51c389b5dd7e6cbcbc"
        },
        "date": 1752089342418,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizeConfirmation",
            "value": 96.71,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "247841038 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation - ns/op",
            "value": 96.71,
            "unit": "ns/op",
            "extra": "247841038 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "247841038 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "247841038 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest",
            "value": 18.81,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest - ns/op",
            "value": 18.81,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus",
            "value": 94.92,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "252908029 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus - ns/op",
            "value": 94.92,
            "unit": "ns/op",
            "extra": "252908029 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "252908029 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "252908029 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue",
            "value": 0.3116,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue - ns/op",
            "value": 0.3116,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValidate",
            "value": 0.311,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValidate - ns/op",
            "value": 0.311,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValidate - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValidate - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValue",
            "value": 0.3119,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValue - ns/op",
            "value": 0.3119,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValue - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo",
            "value": 98.7,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "242557795 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo - ns/op",
            "value": 98.7,
            "unit": "ns/op",
            "extra": "242557795 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "242557795 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "242557795 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue",
            "value": 2.821,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue - ns/op",
            "value": 2.821,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken",
            "value": 0.3115,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken - ns/op",
            "value": 0.3115,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTokenValue",
            "value": 0.3113,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTokenValue - ns/op",
            "value": 0.3113,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTokenValue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTokenValue - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate",
            "value": 407.7,
            "unit": "ns/op\t     240 B/op\t       4 allocs/op",
            "extra": "58923726 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate - ns/op",
            "value": 407.7,
            "unit": "ns/op",
            "extra": "58923726 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "58923726 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "58923726 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValue",
            "value": 0.3115,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValue - ns/op",
            "value": 0.3115,
            "unit": "ns/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValue - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValue - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "1000000000 times\n4 procs"
          }
        ]
      }
    ]
  }
}