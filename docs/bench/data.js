window.BENCHMARK_DATA = {
  "lastUpdate": 1789717928428,
  "repoUrl": "https://github.com/moov-io/pamspr",
  "entries": {
    "moov-io/pamspr": [
      {
        "commit": {
          "author": {
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ff465a84d190895c44e5a38da112bafc9f38c017",
          "message": "ci: run streaming-writer Go benchmarks in this repository (#19)",
          "timestamp": "2026-09-14T18:43:40Z",
          "url": "https://github.com/moov-io/pamspr/commit/ff465a84d190895c44e5a38da112bafc9f38c017"
        },
        "date": 1789411512626,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 194251,
            "unit": "ns/op\t  359670 B/op\t     431 allocs/op",
            "extra": "5950 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 194251,
            "unit": "ns/op",
            "extra": "5950 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359670,
            "unit": "B/op",
            "extra": "5950 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5950 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 192998,
            "unit": "ns/op\t  359668 B/op\t     431 allocs/op",
            "extra": "5834 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 192998,
            "unit": "ns/op",
            "extra": "5834 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359668,
            "unit": "B/op",
            "extra": "5834 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5834 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 2051337,
            "unit": "ns/op\t 3023347 B/op\t    4042 allocs/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 2051337,
            "unit": "ns/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023347,
            "unit": "B/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 2097539,
            "unit": "ns/op\t 3023349 B/op\t    4042 allocs/op",
            "extra": "571 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 2097539,
            "unit": "ns/op",
            "extra": "571 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023349,
            "unit": "B/op",
            "extra": "571 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "571 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 21205757,
            "unit": "ns/op\t42777537 B/op\t   40137 allocs/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 21205757,
            "unit": "ns/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777537,
            "unit": "B/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 21545836,
            "unit": "ns/op\t42777596 B/op\t   40138 allocs/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 21545836,
            "unit": "ns/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777596,
            "unit": "B/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "56 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ff465a84d190895c44e5a38da112bafc9f38c017",
          "message": "ci: run streaming-writer Go benchmarks in this repository (#19)",
          "timestamp": "2026-09-14T18:43:40Z",
          "url": "https://github.com/moov-io/pamspr/commit/ff465a84d190895c44e5a38da112bafc9f38c017"
        },
        "date": 1789460300970,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 180156,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "5575 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 180156,
            "unit": "ns/op",
            "extra": "5575 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "5575 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5575 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 185510,
            "unit": "ns/op\t  359670 B/op\t     431 allocs/op",
            "extra": "6321 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 185510,
            "unit": "ns/op",
            "extra": "6321 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359670,
            "unit": "B/op",
            "extra": "6321 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6321 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1920328,
            "unit": "ns/op\t 3023352 B/op\t    4043 allocs/op",
            "extra": "620 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1920328,
            "unit": "ns/op",
            "extra": "620 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023352,
            "unit": "B/op",
            "extra": "620 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "620 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1947434,
            "unit": "ns/op\t 3023352 B/op\t    4042 allocs/op",
            "extra": "627 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1947434,
            "unit": "ns/op",
            "extra": "627 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023352,
            "unit": "B/op",
            "extra": "627 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "627 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 18594540,
            "unit": "ns/op\t42777619 B/op\t   40138 allocs/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 18594540,
            "unit": "ns/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777619,
            "unit": "B/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "54 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 18959784,
            "unit": "ns/op\t42777579 B/op\t   40138 allocs/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 18959784,
            "unit": "ns/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777579,
            "unit": "B/op",
            "extra": "60 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "60 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "91c7bb19d989df90cd5b9012f53bddad675ef936",
          "message": "chore(deps): update benchmark-action/github-action-benchmark action to v1.22.2 (#20)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-09-15T12:14:47Z",
          "url": "https://github.com/moov-io/pamspr/commit/91c7bb19d989df90cd5b9012f53bddad675ef936"
        },
        "date": 1789546285717,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 198098,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "6697 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 198098,
            "unit": "ns/op",
            "extra": "6697 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "6697 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6697 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 185071,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "5440 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 185071,
            "unit": "ns/op",
            "extra": "5440 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "5440 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5440 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1971511,
            "unit": "ns/op\t 3023356 B/op\t    4042 allocs/op",
            "extra": "621 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1971511,
            "unit": "ns/op",
            "extra": "621 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023356,
            "unit": "B/op",
            "extra": "621 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "621 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1927130,
            "unit": "ns/op\t 3023361 B/op\t    4043 allocs/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1927130,
            "unit": "ns/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023361,
            "unit": "B/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 20504580,
            "unit": "ns/op\t42777536 B/op\t   40137 allocs/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 20504580,
            "unit": "ns/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777536,
            "unit": "B/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 21353480,
            "unit": "ns/op\t42777552 B/op\t   40137 allocs/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 21353480,
            "unit": "ns/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777552,
            "unit": "B/op",
            "extra": "52 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "52 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "91c7bb19d989df90cd5b9012f53bddad675ef936",
          "message": "chore(deps): update benchmark-action/github-action-benchmark action to v1.22.2 (#20)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-09-15T12:14:47Z",
          "url": "https://github.com/moov-io/pamspr/commit/91c7bb19d989df90cd5b9012f53bddad675ef936"
        },
        "date": 1789633085564,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 96252,
            "unit": "ns/op\t  359672 B/op\t     431 allocs/op",
            "extra": "12459 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 96252,
            "unit": "ns/op",
            "extra": "12459 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359672,
            "unit": "B/op",
            "extra": "12459 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "12459 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 95783,
            "unit": "ns/op\t  359673 B/op\t     431 allocs/op",
            "extra": "12900 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 95783,
            "unit": "ns/op",
            "extra": "12900 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359673,
            "unit": "B/op",
            "extra": "12900 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "12900 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 942171,
            "unit": "ns/op\t 3023374 B/op\t    4043 allocs/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 942171,
            "unit": "ns/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023374,
            "unit": "B/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 963051,
            "unit": "ns/op\t 3023363 B/op\t    4043 allocs/op",
            "extra": "1278 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 963051,
            "unit": "ns/op",
            "extra": "1278 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023363,
            "unit": "B/op",
            "extra": "1278 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "1278 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 8353322,
            "unit": "ns/op\t42777575 B/op\t   40138 allocs/op",
            "extra": "130 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 8353322,
            "unit": "ns/op",
            "extra": "130 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777575,
            "unit": "B/op",
            "extra": "130 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "130 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 8798110,
            "unit": "ns/op\t42777591 B/op\t   40138 allocs/op",
            "extra": "141 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 8798110,
            "unit": "ns/op",
            "extra": "141 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777591,
            "unit": "B/op",
            "extra": "141 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "141 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "renovate[bot]",
            "username": "renovate[bot]",
            "email": "29139614+renovate[bot]@users.noreply.github.com"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "91c7bb19d989df90cd5b9012f53bddad675ef936",
          "message": "chore(deps): update benchmark-action/github-action-benchmark action to v1.22.2 (#20)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-09-15T12:14:47Z",
          "url": "https://github.com/moov-io/pamspr/commit/91c7bb19d989df90cd5b9012f53bddad675ef936"
        },
        "date": 1789717927361,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 201200,
            "unit": "ns/op\t  359669 B/op\t     431 allocs/op",
            "extra": "6121 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 201200,
            "unit": "ns/op",
            "extra": "6121 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359669,
            "unit": "B/op",
            "extra": "6121 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6121 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 202503,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "5710 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 202503,
            "unit": "ns/op",
            "extra": "5710 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "5710 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5710 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 2009302,
            "unit": "ns/op\t 3023349 B/op\t    4042 allocs/op",
            "extra": "595 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 2009302,
            "unit": "ns/op",
            "extra": "595 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023349,
            "unit": "B/op",
            "extra": "595 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "595 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 2017953,
            "unit": "ns/op\t 3023356 B/op\t    4043 allocs/op",
            "extra": "602 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 2017953,
            "unit": "ns/op",
            "extra": "602 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023356,
            "unit": "B/op",
            "extra": "602 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "602 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 20620764,
            "unit": "ns/op\t42777552 B/op\t   40137 allocs/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 20620764,
            "unit": "ns/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777552,
            "unit": "B/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 20380344,
            "unit": "ns/op\t42777569 B/op\t   40138 allocs/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 20380344,
            "unit": "ns/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777569,
            "unit": "B/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "57 times\n4 procs"
          }
        ]
      }
    ]
  }
}