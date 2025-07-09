window.BENCHMARK_DATA = {
  "lastUpdate": 1752088471228,
  "repoUrl": "https://github.com/aasanchez/ocpp16messages",
  "entries": {
    "Go Benchmark on 1.23)": [
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
        "date": 1752087982786,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizeConfirmation",
            "value": 97.3,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "245650542 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation - ns/op",
            "value": 97.3,
            "unit": "ns/op",
            "extra": "245650542 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "245650542 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "245650542 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest",
            "value": 19.02,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest - ns/op",
            "value": 19.02,
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
            "value": 95.06,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "252609284 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus - ns/op",
            "value": 95.06,
            "unit": "ns/op",
            "extra": "252609284 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "252609284 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "252609284 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue",
            "value": 0.6254,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue - ns/op",
            "value": 0.6254,
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
            "value": 0.312,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValidate - ns/op",
            "value": 0.312,
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
            "value": 0.3114,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValue - ns/op",
            "value": 0.3114,
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
            "value": 99.08,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "242127946 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo - ns/op",
            "value": 99.08,
            "unit": "ns/op",
            "extra": "242127946 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "242127946 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "242127946 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue",
            "value": 2.802,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue - ns/op",
            "value": 2.802,
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
            "value": 0.3119,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTokenValue - ns/op",
            "value": 0.3119,
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
            "value": 432.1,
            "unit": "ns/op\t     240 B/op\t       4 allocs/op",
            "extra": "55868989 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate - ns/op",
            "value": 432.1,
            "unit": "ns/op",
            "extra": "55868989 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "55868989 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "55868989 times\n4 procs"
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
        "date": 1752088470821,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkAuthorizeConfirmation",
            "value": 97.79,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "246789626 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation - ns/op",
            "value": 97.79,
            "unit": "ns/op",
            "extra": "246789626 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "246789626 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeConfirmation - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "246789626 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest",
            "value": 18.92,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizeRequest - ns/op",
            "value": 18.92,
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
            "value": 94.97,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "251062203 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus - ns/op",
            "value": 94.97,
            "unit": "ns/op",
            "extra": "251062203 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "251062203 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatus - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "251062203 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue",
            "value": 0.6242,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkAuthorizationStatusValue - ns/op",
            "value": 0.6242,
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
            "value": 0.3122,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkConfirmationPayloadValue - ns/op",
            "value": 0.3122,
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
            "value": 99.26,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "242242665 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo - ns/op",
            "value": 99.26,
            "unit": "ns/op",
            "extra": "242242665 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "242242665 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfo - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "242242665 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue",
            "value": 2.799,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTagInfoValue - ns/op",
            "value": 2.799,
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
            "value": 0.3116,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdToken - ns/op",
            "value": 0.3116,
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
            "value": 0.311,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "1000000000 times\n4 procs"
          },
          {
            "name": "BenchmarkIdTokenValue - ns/op",
            "value": 0.311,
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
            "value": 428,
            "unit": "ns/op\t     240 B/op\t       4 allocs/op",
            "extra": "56307513 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate - ns/op",
            "value": 428,
            "unit": "ns/op",
            "extra": "56307513 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "56307513 times\n4 procs"
          },
          {
            "name": "BenchmarkRequestPayloadValidate - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "56307513 times\n4 procs"
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
      }
    ]
  }
}