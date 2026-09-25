window.BENCHMARK_DATA = {
  "lastUpdate": 1790324829002,
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
          "id": "fe90cb3afc87b29a87c4771cdaa6a22f1aa54524",
          "message": "chore(deps): update github/codeql-action action to v4.38.1 (#21)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-09-18T22:05:14Z",
          "url": "https://github.com/moov-io/pamspr/commit/fe90cb3afc87b29a87c4771cdaa6a22f1aa54524"
        },
        "date": 1789803825268,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 207040,
            "unit": "ns/op\t  359672 B/op\t     431 allocs/op",
            "extra": "5619 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 207040,
            "unit": "ns/op",
            "extra": "5619 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359672,
            "unit": "B/op",
            "extra": "5619 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5619 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 199184,
            "unit": "ns/op\t  359670 B/op\t     431 allocs/op",
            "extra": "5787 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 199184,
            "unit": "ns/op",
            "extra": "5787 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359670,
            "unit": "B/op",
            "extra": "5787 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5787 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1911039,
            "unit": "ns/op\t 3023346 B/op\t    4042 allocs/op",
            "extra": "596 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1911039,
            "unit": "ns/op",
            "extra": "596 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023346,
            "unit": "B/op",
            "extra": "596 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "596 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1912042,
            "unit": "ns/op\t 3023334 B/op\t    4042 allocs/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1912042,
            "unit": "ns/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023334,
            "unit": "B/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 20350396,
            "unit": "ns/op\t42777517 B/op\t   40137 allocs/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 20350396,
            "unit": "ns/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777517,
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
            "value": 20626604,
            "unit": "ns/op\t42777538 B/op\t   40137 allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 20626604,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777538,
            "unit": "B/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "49 times\n4 procs"
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
          "id": "fe90cb3afc87b29a87c4771cdaa6a22f1aa54524",
          "message": "chore(deps): update github/codeql-action action to v4.38.1 (#21)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-09-18T22:05:14Z",
          "url": "https://github.com/moov-io/pamspr/commit/fe90cb3afc87b29a87c4771cdaa6a22f1aa54524"
        },
        "date": 1789891937648,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 276895,
            "unit": "ns/op\t  359682 B/op\t     431 allocs/op",
            "extra": "3972 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 276895,
            "unit": "ns/op",
            "extra": "3972 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359682,
            "unit": "B/op",
            "extra": "3972 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "3972 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 197955,
            "unit": "ns/op\t  359668 B/op\t     431 allocs/op",
            "extra": "6109 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 197955,
            "unit": "ns/op",
            "extra": "6109 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359668,
            "unit": "B/op",
            "extra": "6109 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6109 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 2301052,
            "unit": "ns/op\t 3023366 B/op\t    4043 allocs/op",
            "extra": "537 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 2301052,
            "unit": "ns/op",
            "extra": "537 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023366,
            "unit": "B/op",
            "extra": "537 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "537 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 2315644,
            "unit": "ns/op\t 3023353 B/op\t    4043 allocs/op",
            "extra": "484 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 2315644,
            "unit": "ns/op",
            "extra": "484 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023353,
            "unit": "B/op",
            "extra": "484 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "484 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 21026628,
            "unit": "ns/op\t42777535 B/op\t   40137 allocs/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 21026628,
            "unit": "ns/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777535,
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
            "value": 21488916,
            "unit": "ns/op\t42777520 B/op\t   40137 allocs/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 21488916,
            "unit": "ns/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777520,
            "unit": "B/op",
            "extra": "57 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "57 times\n4 procs"
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
          "id": "fe90cb3afc87b29a87c4771cdaa6a22f1aa54524",
          "message": "chore(deps): update github/codeql-action action to v4.38.1 (#21)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-09-18T22:05:14Z",
          "url": "https://github.com/moov-io/pamspr/commit/fe90cb3afc87b29a87c4771cdaa6a22f1aa54524"
        },
        "date": 1789979327292,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 179902,
            "unit": "ns/op\t  359669 B/op\t     431 allocs/op",
            "extra": "5710 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 179902,
            "unit": "ns/op",
            "extra": "5710 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359669,
            "unit": "B/op",
            "extra": "5710 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5710 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 174582,
            "unit": "ns/op\t  359669 B/op\t     431 allocs/op",
            "extra": "6750 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 174582,
            "unit": "ns/op",
            "extra": "6750 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359669,
            "unit": "B/op",
            "extra": "6750 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6750 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1907868,
            "unit": "ns/op\t 3023358 B/op\t    4043 allocs/op",
            "extra": "655 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1907868,
            "unit": "ns/op",
            "extra": "655 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023358,
            "unit": "B/op",
            "extra": "655 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "655 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1943535,
            "unit": "ns/op\t 3023363 B/op\t    4043 allocs/op",
            "extra": "595 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1943535,
            "unit": "ns/op",
            "extra": "595 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023363,
            "unit": "B/op",
            "extra": "595 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "595 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 19093514,
            "unit": "ns/op\t42777556 B/op\t   40137 allocs/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 19093514,
            "unit": "ns/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777556,
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
            "value": 18936932,
            "unit": "ns/op\t42777585 B/op\t   40138 allocs/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 18936932,
            "unit": "ns/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777585,
            "unit": "B/op",
            "extra": "62 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "62 times\n4 procs"
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
          "id": "fe90cb3afc87b29a87c4771cdaa6a22f1aa54524",
          "message": "chore(deps): update github/codeql-action action to v4.38.1 (#21)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-09-18T22:05:14Z",
          "url": "https://github.com/moov-io/pamspr/commit/fe90cb3afc87b29a87c4771cdaa6a22f1aa54524"
        },
        "date": 1790064585143,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 161481,
            "unit": "ns/op\t  359673 B/op\t     431 allocs/op",
            "extra": "6385 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 161481,
            "unit": "ns/op",
            "extra": "6385 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359673,
            "unit": "B/op",
            "extra": "6385 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6385 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 158464,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "6968 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 158464,
            "unit": "ns/op",
            "extra": "6968 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "6968 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6968 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1634585,
            "unit": "ns/op\t 3023415 B/op\t    4043 allocs/op",
            "extra": "735 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1634585,
            "unit": "ns/op",
            "extra": "735 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023415,
            "unit": "B/op",
            "extra": "735 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "735 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1625410,
            "unit": "ns/op\t 3023415 B/op\t    4043 allocs/op",
            "extra": "692 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1625410,
            "unit": "ns/op",
            "extra": "692 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023415,
            "unit": "B/op",
            "extra": "692 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "692 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 15986049,
            "unit": "ns/op\t42777556 B/op\t   40137 allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 15986049,
            "unit": "ns/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777556,
            "unit": "B/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "75 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 15816464,
            "unit": "ns/op\t42777572 B/op\t   40138 allocs/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 15816464,
            "unit": "ns/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777572,
            "unit": "B/op",
            "extra": "74 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "74 times\n4 procs"
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
          "id": "fe90cb3afc87b29a87c4771cdaa6a22f1aa54524",
          "message": "chore(deps): update github/codeql-action action to v4.38.1 (#21)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-09-18T22:05:14Z",
          "url": "https://github.com/moov-io/pamspr/commit/fe90cb3afc87b29a87c4771cdaa6a22f1aa54524"
        },
        "date": 1790151111899,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 198087,
            "unit": "ns/op\t  359672 B/op\t     431 allocs/op",
            "extra": "5484 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 198087,
            "unit": "ns/op",
            "extra": "5484 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359672,
            "unit": "B/op",
            "extra": "5484 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5484 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 195883,
            "unit": "ns/op\t  359669 B/op\t     431 allocs/op",
            "extra": "5436 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 195883,
            "unit": "ns/op",
            "extra": "5436 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359669,
            "unit": "B/op",
            "extra": "5436 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5436 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 2181697,
            "unit": "ns/op\t 3023368 B/op\t    4043 allocs/op",
            "extra": "548 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 2181697,
            "unit": "ns/op",
            "extra": "548 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023368,
            "unit": "B/op",
            "extra": "548 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "548 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 2149753,
            "unit": "ns/op\t 3023371 B/op\t    4043 allocs/op",
            "extra": "567 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 2149753,
            "unit": "ns/op",
            "extra": "567 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023371,
            "unit": "B/op",
            "extra": "567 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "567 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 21238609,
            "unit": "ns/op\t42777552 B/op\t   40137 allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 21238609,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777552,
            "unit": "B/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 21454112,
            "unit": "ns/op\t42777554 B/op\t   40137 allocs/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 21454112,
            "unit": "ns/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777554,
            "unit": "B/op",
            "extra": "48 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "48 times\n4 procs"
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
          "id": "fe90cb3afc87b29a87c4771cdaa6a22f1aa54524",
          "message": "chore(deps): update github/codeql-action action to v4.38.1 (#21)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-09-18T22:05:14Z",
          "url": "https://github.com/moov-io/pamspr/commit/fe90cb3afc87b29a87c4771cdaa6a22f1aa54524"
        },
        "date": 1790236959358,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 142721,
            "unit": "ns/op\t  359673 B/op\t     431 allocs/op",
            "extra": "8631 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 142721,
            "unit": "ns/op",
            "extra": "8631 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359673,
            "unit": "B/op",
            "extra": "8631 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "8631 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 140194,
            "unit": "ns/op\t  359672 B/op\t     431 allocs/op",
            "extra": "7989 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 140194,
            "unit": "ns/op",
            "extra": "7989 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359672,
            "unit": "B/op",
            "extra": "7989 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "7989 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1389308,
            "unit": "ns/op\t 3023403 B/op\t    4043 allocs/op",
            "extra": "830 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1389308,
            "unit": "ns/op",
            "extra": "830 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023403,
            "unit": "B/op",
            "extra": "830 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "830 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1478482,
            "unit": "ns/op\t 3023396 B/op\t    4043 allocs/op",
            "extra": "837 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1478482,
            "unit": "ns/op",
            "extra": "837 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023396,
            "unit": "B/op",
            "extra": "837 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "837 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 13992108,
            "unit": "ns/op\t42777560 B/op\t   40137 allocs/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 13992108,
            "unit": "ns/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777560,
            "unit": "B/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 13779563,
            "unit": "ns/op\t42777529 B/op\t   40137 allocs/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 13779563,
            "unit": "ns/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777529,
            "unit": "B/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "88 times\n4 procs"
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
          "id": "8c5ae66370eec5da293a0094a8fd5e126de6cfaf",
          "message": "chore(deps): update github/codeql-action action to v4.38.2 (#22)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-09-24T22:40:33Z",
          "url": "https://github.com/moov-io/pamspr/commit/8c5ae66370eec5da293a0094a8fd5e126de6cfaf"
        },
        "date": 1790324827566,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 191166,
            "unit": "ns/op\t  359672 B/op\t     431 allocs/op",
            "extra": "6290 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 191166,
            "unit": "ns/op",
            "extra": "6290 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359672,
            "unit": "B/op",
            "extra": "6290 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6290 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 181607,
            "unit": "ns/op\t  359670 B/op\t     431 allocs/op",
            "extra": "5788 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 181607,
            "unit": "ns/op",
            "extra": "5788 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359670,
            "unit": "B/op",
            "extra": "5788 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5788 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1881631,
            "unit": "ns/op\t 3023404 B/op\t    4043 allocs/op",
            "extra": "630 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1881631,
            "unit": "ns/op",
            "extra": "630 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023404,
            "unit": "B/op",
            "extra": "630 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "630 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1896557,
            "unit": "ns/op\t 3023415 B/op\t    4043 allocs/op",
            "extra": "638 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1896557,
            "unit": "ns/op",
            "extra": "638 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023415,
            "unit": "B/op",
            "extra": "638 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "638 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 20469557,
            "unit": "ns/op\t42777599 B/op\t   40138 allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 20469557,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777599,
            "unit": "B/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 20345103,
            "unit": "ns/op\t42777603 B/op\t   40138 allocs/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 20345103,
            "unit": "ns/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777603,
            "unit": "B/op",
            "extra": "55 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "55 times\n4 procs"
          }
        ]
      }
    ]
  }
}