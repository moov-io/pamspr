window.BENCHMARK_DATA = {
  "lastUpdate": 1793090062598,
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
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash@ela.city"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ac6978bc3aa15c0b4aaff0803defabafdbe2d07e",
          "message": "Check the ABA routing number in ACHPayment.Validate (#23)\n\nSchedule validation calls this method. A nine-digit number with a bad check digit was accepted even though Validator.validateRoutingNumber rejects it.",
          "timestamp": "2026-09-25T14:22:57Z",
          "url": "https://github.com/moov-io/pamspr/commit/ac6978bc3aa15c0b4aaff0803defabafdbe2d07e"
        },
        "date": 1790410443060,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 89551,
            "unit": "ns/op\t  359672 B/op\t     431 allocs/op",
            "extra": "13800 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 89551,
            "unit": "ns/op",
            "extra": "13800 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359672,
            "unit": "B/op",
            "extra": "13800 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "13800 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 89742,
            "unit": "ns/op\t  359672 B/op\t     431 allocs/op",
            "extra": "13171 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 89742,
            "unit": "ns/op",
            "extra": "13171 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359672,
            "unit": "B/op",
            "extra": "13171 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "13171 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 899220,
            "unit": "ns/op\t 3023355 B/op\t    4042 allocs/op",
            "extra": "1358 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 899220,
            "unit": "ns/op",
            "extra": "1358 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023355,
            "unit": "B/op",
            "extra": "1358 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "1358 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 958474,
            "unit": "ns/op\t 3023365 B/op\t    4043 allocs/op",
            "extra": "1255 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 958474,
            "unit": "ns/op",
            "extra": "1255 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023365,
            "unit": "B/op",
            "extra": "1255 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "1255 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 8996789,
            "unit": "ns/op\t42777611 B/op\t   40138 allocs/op",
            "extra": "132 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 8996789,
            "unit": "ns/op",
            "extra": "132 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777611,
            "unit": "B/op",
            "extra": "132 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "132 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 8978783,
            "unit": "ns/op\t42777677 B/op\t   40139 allocs/op",
            "extra": "128 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 8978783,
            "unit": "ns/op",
            "extra": "128 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777677,
            "unit": "B/op",
            "extra": "128 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40139,
            "unit": "allocs/op",
            "extra": "128 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash@ela.city"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ac6978bc3aa15c0b4aaff0803defabafdbe2d07e",
          "message": "Check the ABA routing number in ACHPayment.Validate (#23)\n\nSchedule validation calls this method. A nine-digit number with a bad check digit was accepted even though Validator.validateRoutingNumber rejects it.",
          "timestamp": "2026-09-25T14:22:57Z",
          "url": "https://github.com/moov-io/pamspr/commit/ac6978bc3aa15c0b4aaff0803defabafdbe2d07e"
        },
        "date": 1790499064569,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 85068,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "13939 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 85068,
            "unit": "ns/op",
            "extra": "13939 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "13939 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "13939 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 84377,
            "unit": "ns/op\t  359670 B/op\t     431 allocs/op",
            "extra": "13606 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 84377,
            "unit": "ns/op",
            "extra": "13606 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359670,
            "unit": "B/op",
            "extra": "13606 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "13606 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 910591,
            "unit": "ns/op\t 3023361 B/op\t    4043 allocs/op",
            "extra": "1342 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 910591,
            "unit": "ns/op",
            "extra": "1342 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023361,
            "unit": "B/op",
            "extra": "1342 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "1342 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 956797,
            "unit": "ns/op\t 3023371 B/op\t    4043 allocs/op",
            "extra": "1303 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 956797,
            "unit": "ns/op",
            "extra": "1303 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023371,
            "unit": "B/op",
            "extra": "1303 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "1303 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 8690074,
            "unit": "ns/op\t42777579 B/op\t   40138 allocs/op",
            "extra": "138 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 8690074,
            "unit": "ns/op",
            "extra": "138 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777579,
            "unit": "B/op",
            "extra": "138 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "138 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 8594173,
            "unit": "ns/op\t42777561 B/op\t   40138 allocs/op",
            "extra": "140 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 8594173,
            "unit": "ns/op",
            "extra": "140 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777561,
            "unit": "B/op",
            "extra": "140 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "140 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash@ela.city"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ac6978bc3aa15c0b4aaff0803defabafdbe2d07e",
          "message": "Check the ABA routing number in ACHPayment.Validate (#23)\n\nSchedule validation calls this method. A nine-digit number with a bad check digit was accepted even though Validator.validateRoutingNumber rejects it.",
          "timestamp": "2026-09-25T14:22:57Z",
          "url": "https://github.com/moov-io/pamspr/commit/ac6978bc3aa15c0b4aaff0803defabafdbe2d07e"
        },
        "date": 1790586917354,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 87932,
            "unit": "ns/op\t  359673 B/op\t     431 allocs/op",
            "extra": "13930 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 87932,
            "unit": "ns/op",
            "extra": "13930 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359673,
            "unit": "B/op",
            "extra": "13930 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "13930 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 82807,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "14154 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 82807,
            "unit": "ns/op",
            "extra": "14154 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "14154 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "14154 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 837499,
            "unit": "ns/op\t 3023353 B/op\t    4042 allocs/op",
            "extra": "1414 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 837499,
            "unit": "ns/op",
            "extra": "1414 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023353,
            "unit": "B/op",
            "extra": "1414 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "1414 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 871786,
            "unit": "ns/op\t 3023367 B/op\t    4043 allocs/op",
            "extra": "1452 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 871786,
            "unit": "ns/op",
            "extra": "1452 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023367,
            "unit": "B/op",
            "extra": "1452 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "1452 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 7946399,
            "unit": "ns/op\t42777597 B/op\t   40138 allocs/op",
            "extra": "151 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 7946399,
            "unit": "ns/op",
            "extra": "151 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777597,
            "unit": "B/op",
            "extra": "151 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "151 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 8021077,
            "unit": "ns/op\t42777589 B/op\t   40138 allocs/op",
            "extra": "148 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 8021077,
            "unit": "ns/op",
            "extra": "148 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777589,
            "unit": "B/op",
            "extra": "148 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "148 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash@ela.city"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ac6978bc3aa15c0b4aaff0803defabafdbe2d07e",
          "message": "Check the ABA routing number in ACHPayment.Validate (#23)\n\nSchedule validation calls this method. A nine-digit number with a bad check digit was accepted even though Validator.validateRoutingNumber rejects it.",
          "timestamp": "2026-09-25T14:22:57Z",
          "url": "https://github.com/moov-io/pamspr/commit/ac6978bc3aa15c0b4aaff0803defabafdbe2d07e"
        },
        "date": 1790673781787,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 160647,
            "unit": "ns/op\t  359670 B/op\t     431 allocs/op",
            "extra": "7406 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 160647,
            "unit": "ns/op",
            "extra": "7406 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359670,
            "unit": "B/op",
            "extra": "7406 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "7406 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 166396,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "7226 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 166396,
            "unit": "ns/op",
            "extra": "7226 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "7226 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "7226 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1621202,
            "unit": "ns/op\t 3023406 B/op\t    4043 allocs/op",
            "extra": "682 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1621202,
            "unit": "ns/op",
            "extra": "682 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023406,
            "unit": "B/op",
            "extra": "682 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "682 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1641924,
            "unit": "ns/op\t 3023419 B/op\t    4043 allocs/op",
            "extra": "700 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1641924,
            "unit": "ns/op",
            "extra": "700 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023419,
            "unit": "B/op",
            "extra": "700 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "700 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 17401460,
            "unit": "ns/op\t42777539 B/op\t   40137 allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 17401460,
            "unit": "ns/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777539,
            "unit": "B/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "69 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 17024513,
            "unit": "ns/op\t42777575 B/op\t   40138 allocs/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 17024513,
            "unit": "ns/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777575,
            "unit": "B/op",
            "extra": "68 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "68 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash@ela.city"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ac6978bc3aa15c0b4aaff0803defabafdbe2d07e",
          "message": "Check the ABA routing number in ACHPayment.Validate (#23)\n\nSchedule validation calls this method. A nine-digit number with a bad check digit was accepted even though Validator.validateRoutingNumber rejects it.",
          "timestamp": "2026-09-25T14:22:57Z",
          "url": "https://github.com/moov-io/pamspr/commit/ac6978bc3aa15c0b4aaff0803defabafdbe2d07e"
        },
        "date": 1790759601417,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 126826,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "7994 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 126826,
            "unit": "ns/op",
            "extra": "7994 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "7994 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "7994 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 126538,
            "unit": "ns/op\t  359670 B/op\t     431 allocs/op",
            "extra": "8773 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 126538,
            "unit": "ns/op",
            "extra": "8773 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359670,
            "unit": "B/op",
            "extra": "8773 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "8773 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1314482,
            "unit": "ns/op\t 3023352 B/op\t    4043 allocs/op",
            "extra": "841 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1314482,
            "unit": "ns/op",
            "extra": "841 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023352,
            "unit": "B/op",
            "extra": "841 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "841 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1443258,
            "unit": "ns/op\t 3023349 B/op\t    4042 allocs/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1443258,
            "unit": "ns/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023349,
            "unit": "B/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 13359818,
            "unit": "ns/op\t42777587 B/op\t   40138 allocs/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 13359818,
            "unit": "ns/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777587,
            "unit": "B/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 13756824,
            "unit": "ns/op\t42777595 B/op\t   40138 allocs/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 13756824,
            "unit": "ns/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777595,
            "unit": "B/op",
            "extra": "88 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "88 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash@ela.city"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ac6978bc3aa15c0b4aaff0803defabafdbe2d07e",
          "message": "Check the ABA routing number in ACHPayment.Validate (#23)\n\nSchedule validation calls this method. A nine-digit number with a bad check digit was accepted even though Validator.validateRoutingNumber rejects it.",
          "timestamp": "2026-09-25T14:22:57Z",
          "url": "https://github.com/moov-io/pamspr/commit/ac6978bc3aa15c0b4aaff0803defabafdbe2d07e"
        },
        "date": 1790847658124,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 93813,
            "unit": "ns/op\t  359670 B/op\t     431 allocs/op",
            "extra": "12684 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 93813,
            "unit": "ns/op",
            "extra": "12684 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359670,
            "unit": "B/op",
            "extra": "12684 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "12684 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 92255,
            "unit": "ns/op\t  359673 B/op\t     431 allocs/op",
            "extra": "13148 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 92255,
            "unit": "ns/op",
            "extra": "13148 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359673,
            "unit": "B/op",
            "extra": "13148 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "13148 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 941507,
            "unit": "ns/op\t 3023359 B/op\t    4043 allocs/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 941507,
            "unit": "ns/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023359,
            "unit": "B/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 916892,
            "unit": "ns/op\t 3023350 B/op\t    4042 allocs/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 916892,
            "unit": "ns/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023350,
            "unit": "B/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 9971521,
            "unit": "ns/op\t42777597 B/op\t   40138 allocs/op",
            "extra": "117 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 9971521,
            "unit": "ns/op",
            "extra": "117 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777597,
            "unit": "B/op",
            "extra": "117 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "117 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 10100917,
            "unit": "ns/op\t42777612 B/op\t   40138 allocs/op",
            "extra": "117 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 10100917,
            "unit": "ns/op",
            "extra": "117 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777612,
            "unit": "B/op",
            "extra": "117 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "117 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash@ela.city"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ac6978bc3aa15c0b4aaff0803defabafdbe2d07e",
          "message": "Check the ABA routing number in ACHPayment.Validate (#23)\n\nSchedule validation calls this method. A nine-digit number with a bad check digit was accepted even though Validator.validateRoutingNumber rejects it.",
          "timestamp": "2026-09-25T14:22:57Z",
          "url": "https://github.com/moov-io/pamspr/commit/ac6978bc3aa15c0b4aaff0803defabafdbe2d07e"
        },
        "date": 1790932546461,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 191743,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "5953 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 191743,
            "unit": "ns/op",
            "extra": "5953 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "5953 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5953 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 190454,
            "unit": "ns/op\t  359668 B/op\t     431 allocs/op",
            "extra": "5767 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 190454,
            "unit": "ns/op",
            "extra": "5767 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359668,
            "unit": "B/op",
            "extra": "5767 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5767 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1957440,
            "unit": "ns/op\t 3023374 B/op\t    4043 allocs/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1957440,
            "unit": "ns/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023374,
            "unit": "B/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "592 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1958113,
            "unit": "ns/op\t 3023371 B/op\t    4043 allocs/op",
            "extra": "595 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1958113,
            "unit": "ns/op",
            "extra": "595 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023371,
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
            "value": 19986765,
            "unit": "ns/op\t42777552 B/op\t   40137 allocs/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 19986765,
            "unit": "ns/op",
            "extra": "51 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777552,
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
            "value": 20014102,
            "unit": "ns/op\t42777524 B/op\t   40137 allocs/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 20014102,
            "unit": "ns/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777524,
            "unit": "B/op",
            "extra": "50 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "50 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash@ela.city"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ac6978bc3aa15c0b4aaff0803defabafdbe2d07e",
          "message": "Check the ABA routing number in ACHPayment.Validate (#23)\n\nSchedule validation calls this method. A nine-digit number with a bad check digit was accepted even though Validator.validateRoutingNumber rejects it.",
          "timestamp": "2026-09-25T14:22:57Z",
          "url": "https://github.com/moov-io/pamspr/commit/ac6978bc3aa15c0b4aaff0803defabafdbe2d07e"
        },
        "date": 1791017238396,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 128450,
            "unit": "ns/op\t  359670 B/op\t     431 allocs/op",
            "extra": "8845 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 128450,
            "unit": "ns/op",
            "extra": "8845 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359670,
            "unit": "B/op",
            "extra": "8845 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "8845 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 126115,
            "unit": "ns/op\t  359669 B/op\t     431 allocs/op",
            "extra": "8101 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 126115,
            "unit": "ns/op",
            "extra": "8101 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359669,
            "unit": "B/op",
            "extra": "8101 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "8101 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1407854,
            "unit": "ns/op\t 3023348 B/op\t    4042 allocs/op",
            "extra": "832 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1407854,
            "unit": "ns/op",
            "extra": "832 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023348,
            "unit": "B/op",
            "extra": "832 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "832 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1408433,
            "unit": "ns/op\t 3023350 B/op\t    4042 allocs/op",
            "extra": "885 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1408433,
            "unit": "ns/op",
            "extra": "885 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023350,
            "unit": "B/op",
            "extra": "885 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "885 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 13491484,
            "unit": "ns/op\t42777545 B/op\t   40137 allocs/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 13491484,
            "unit": "ns/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777545,
            "unit": "B/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "76 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 13577911,
            "unit": "ns/op\t42777547 B/op\t   40137 allocs/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 13577911,
            "unit": "ns/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777547,
            "unit": "B/op",
            "extra": "87 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "87 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash@ela.city"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ac6978bc3aa15c0b4aaff0803defabafdbe2d07e",
          "message": "Check the ABA routing number in ACHPayment.Validate (#23)\n\nSchedule validation calls this method. A nine-digit number with a bad check digit was accepted even though Validator.validateRoutingNumber rejects it.",
          "timestamp": "2026-09-25T14:22:57Z",
          "url": "https://github.com/moov-io/pamspr/commit/ac6978bc3aa15c0b4aaff0803defabafdbe2d07e"
        },
        "date": 1791105330364,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 129724,
            "unit": "ns/op\t  359672 B/op\t     431 allocs/op",
            "extra": "9153 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 129724,
            "unit": "ns/op",
            "extra": "9153 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359672,
            "unit": "B/op",
            "extra": "9153 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "9153 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 131137,
            "unit": "ns/op\t  359672 B/op\t     431 allocs/op",
            "extra": "8492 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 131137,
            "unit": "ns/op",
            "extra": "8492 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359672,
            "unit": "B/op",
            "extra": "8492 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "8492 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1406903,
            "unit": "ns/op\t 3023385 B/op\t    4043 allocs/op",
            "extra": "836 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1406903,
            "unit": "ns/op",
            "extra": "836 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023385,
            "unit": "B/op",
            "extra": "836 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "836 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1336959,
            "unit": "ns/op\t 3023357 B/op\t    4042 allocs/op",
            "extra": "841 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1336959,
            "unit": "ns/op",
            "extra": "841 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023357,
            "unit": "B/op",
            "extra": "841 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "841 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 14339538,
            "unit": "ns/op\t42777559 B/op\t   40137 allocs/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 14339538,
            "unit": "ns/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777559,
            "unit": "B/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "86 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 14401940,
            "unit": "ns/op\t42777578 B/op\t   40138 allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 14401940,
            "unit": "ns/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777578,
            "unit": "B/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "80 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash@ela.city"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ac6978bc3aa15c0b4aaff0803defabafdbe2d07e",
          "message": "Check the ABA routing number in ACHPayment.Validate (#23)\n\nSchedule validation calls this method. A nine-digit number with a bad check digit was accepted even though Validator.validateRoutingNumber rejects it.",
          "timestamp": "2026-09-25T14:22:57Z",
          "url": "https://github.com/moov-io/pamspr/commit/ac6978bc3aa15c0b4aaff0803defabafdbe2d07e"
        },
        "date": 1791194065393,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 96227,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "12094 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 96227,
            "unit": "ns/op",
            "extra": "12094 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "12094 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "12094 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 85150,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "13669 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 85150,
            "unit": "ns/op",
            "extra": "13669 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "13669 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "13669 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 890782,
            "unit": "ns/op\t 3023357 B/op\t    4043 allocs/op",
            "extra": "1342 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 890782,
            "unit": "ns/op",
            "extra": "1342 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023357,
            "unit": "B/op",
            "extra": "1342 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "1342 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 929161,
            "unit": "ns/op\t 3023373 B/op\t    4043 allocs/op",
            "extra": "1306 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 929161,
            "unit": "ns/op",
            "extra": "1306 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023373,
            "unit": "B/op",
            "extra": "1306 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "1306 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 9576548,
            "unit": "ns/op\t42777620 B/op\t   40138 allocs/op",
            "extra": "130 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 9576548,
            "unit": "ns/op",
            "extra": "130 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777620,
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
            "value": 8999291,
            "unit": "ns/op\t42777607 B/op\t   40138 allocs/op",
            "extra": "134 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 8999291,
            "unit": "ns/op",
            "extra": "134 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777607,
            "unit": "B/op",
            "extra": "134 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "134 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash@ela.city"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ac6978bc3aa15c0b4aaff0803defabafdbe2d07e",
          "message": "Check the ABA routing number in ACHPayment.Validate (#23)\n\nSchedule validation calls this method. A nine-digit number with a bad check digit was accepted even though Validator.validateRoutingNumber rejects it.",
          "timestamp": "2026-09-25T14:22:57Z",
          "url": "https://github.com/moov-io/pamspr/commit/ac6978bc3aa15c0b4aaff0803defabafdbe2d07e"
        },
        "date": 1791279683727,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 90017,
            "unit": "ns/op\t  359672 B/op\t     431 allocs/op",
            "extra": "12416 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 90017,
            "unit": "ns/op",
            "extra": "12416 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359672,
            "unit": "B/op",
            "extra": "12416 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "12416 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 93348,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "13117 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 93348,
            "unit": "ns/op",
            "extra": "13117 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "13117 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "13117 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 953510,
            "unit": "ns/op\t 3023364 B/op\t    4043 allocs/op",
            "extra": "1060 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 953510,
            "unit": "ns/op",
            "extra": "1060 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023364,
            "unit": "B/op",
            "extra": "1060 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "1060 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 964208,
            "unit": "ns/op\t 3023369 B/op\t    4043 allocs/op",
            "extra": "1246 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 964208,
            "unit": "ns/op",
            "extra": "1246 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023369,
            "unit": "B/op",
            "extra": "1246 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "1246 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 9804954,
            "unit": "ns/op\t42777578 B/op\t   40138 allocs/op",
            "extra": "133 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 9804954,
            "unit": "ns/op",
            "extra": "133 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777578,
            "unit": "B/op",
            "extra": "133 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "133 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 9214489,
            "unit": "ns/op\t42777599 B/op\t   40138 allocs/op",
            "extra": "126 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 9214489,
            "unit": "ns/op",
            "extra": "126 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777599,
            "unit": "B/op",
            "extra": "126 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "126 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash@ela.city"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ac6978bc3aa15c0b4aaff0803defabafdbe2d07e",
          "message": "Check the ABA routing number in ACHPayment.Validate (#23)\n\nSchedule validation calls this method. A nine-digit number with a bad check digit was accepted even though Validator.validateRoutingNumber rejects it.",
          "timestamp": "2026-09-25T14:22:57Z",
          "url": "https://github.com/moov-io/pamspr/commit/ac6978bc3aa15c0b4aaff0803defabafdbe2d07e"
        },
        "date": 1791365950847,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 94893,
            "unit": "ns/op\t  359672 B/op\t     431 allocs/op",
            "extra": "12837 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 94893,
            "unit": "ns/op",
            "extra": "12837 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359672,
            "unit": "B/op",
            "extra": "12837 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "12837 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 87542,
            "unit": "ns/op\t  359672 B/op\t     431 allocs/op",
            "extra": "12998 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 87542,
            "unit": "ns/op",
            "extra": "12998 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359672,
            "unit": "B/op",
            "extra": "12998 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "12998 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 975101,
            "unit": "ns/op\t 3023353 B/op\t    4042 allocs/op",
            "extra": "1328 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 975101,
            "unit": "ns/op",
            "extra": "1328 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023353,
            "unit": "B/op",
            "extra": "1328 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "1328 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 968985,
            "unit": "ns/op\t 3023350 B/op\t    4042 allocs/op",
            "extra": "1275 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 968985,
            "unit": "ns/op",
            "extra": "1275 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023350,
            "unit": "B/op",
            "extra": "1275 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "1275 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 9741146,
            "unit": "ns/op\t42777580 B/op\t   40138 allocs/op",
            "extra": "128 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 9741146,
            "unit": "ns/op",
            "extra": "128 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777580,
            "unit": "B/op",
            "extra": "128 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "128 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 9502554,
            "unit": "ns/op\t42777616 B/op\t   40138 allocs/op",
            "extra": "126 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 9502554,
            "unit": "ns/op",
            "extra": "126 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777616,
            "unit": "B/op",
            "extra": "126 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "126 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "Sash",
            "username": "SashaMIT",
            "email": "sash@ela.city"
          },
          "committer": {
            "name": "GitHub",
            "username": "web-flow",
            "email": "noreply@github.com"
          },
          "id": "ac6978bc3aa15c0b4aaff0803defabafdbe2d07e",
          "message": "Check the ABA routing number in ACHPayment.Validate (#23)\n\nSchedule validation calls this method. A nine-digit number with a bad check digit was accepted even though Validator.validateRoutingNumber rejects it.",
          "timestamp": "2026-09-25T14:22:57Z",
          "url": "https://github.com/moov-io/pamspr/commit/ac6978bc3aa15c0b4aaff0803defabafdbe2d07e"
        },
        "date": 1791452977074,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 205134,
            "unit": "ns/op\t  359672 B/op\t     431 allocs/op",
            "extra": "5774 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 205134,
            "unit": "ns/op",
            "extra": "5774 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359672,
            "unit": "B/op",
            "extra": "5774 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5774 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 192412,
            "unit": "ns/op\t  359670 B/op\t     431 allocs/op",
            "extra": "6238 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 192412,
            "unit": "ns/op",
            "extra": "6238 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359670,
            "unit": "B/op",
            "extra": "6238 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6238 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 2006502,
            "unit": "ns/op\t 3023361 B/op\t    4043 allocs/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 2006502,
            "unit": "ns/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023361,
            "unit": "B/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "601 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1986158,
            "unit": "ns/op\t 3023365 B/op\t    4043 allocs/op",
            "extra": "577 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1986158,
            "unit": "ns/op",
            "extra": "577 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023365,
            "unit": "B/op",
            "extra": "577 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "577 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 20884952,
            "unit": "ns/op\t42777543 B/op\t   40137 allocs/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 20884952,
            "unit": "ns/op",
            "extra": "49 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777543,
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
            "value": 20775513,
            "unit": "ns/op\t42777551 B/op\t   40137 allocs/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 20775513,
            "unit": "ns/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777551,
            "unit": "B/op",
            "extra": "56 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "56 times\n4 procs"
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
          "id": "e54ff1ac13679fde2ca14fbeb22bc006900ac38c",
          "message": "chore(deps): update github/codeql-action action to v4.38.3 (#24)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-08T20:01:27Z",
          "url": "https://github.com/moov-io/pamspr/commit/e54ff1ac13679fde2ca14fbeb22bc006900ac38c"
        },
        "date": 1791539647604,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 162720,
            "unit": "ns/op\t  359672 B/op\t     431 allocs/op",
            "extra": "6415 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 162720,
            "unit": "ns/op",
            "extra": "6415 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359672,
            "unit": "B/op",
            "extra": "6415 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6415 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 161879,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "6922 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 161879,
            "unit": "ns/op",
            "extra": "6922 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "6922 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6922 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1686388,
            "unit": "ns/op\t 3023414 B/op\t    4043 allocs/op",
            "extra": "702 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1686388,
            "unit": "ns/op",
            "extra": "702 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023414,
            "unit": "B/op",
            "extra": "702 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "702 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1692842,
            "unit": "ns/op\t 3023415 B/op\t    4043 allocs/op",
            "extra": "634 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1692842,
            "unit": "ns/op",
            "extra": "634 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023415,
            "unit": "B/op",
            "extra": "634 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "634 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 15932682,
            "unit": "ns/op\t42777584 B/op\t   40138 allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 15932682,
            "unit": "ns/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777584,
            "unit": "B/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "78 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 15971604,
            "unit": "ns/op\t42777586 B/op\t   40138 allocs/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 15971604,
            "unit": "ns/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777586,
            "unit": "B/op",
            "extra": "63 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "63 times\n4 procs"
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
          "id": "e54ff1ac13679fde2ca14fbeb22bc006900ac38c",
          "message": "chore(deps): update github/codeql-action action to v4.38.3 (#24)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-08T20:01:27Z",
          "url": "https://github.com/moov-io/pamspr/commit/e54ff1ac13679fde2ca14fbeb22bc006900ac38c"
        },
        "date": 1791623892975,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 229580,
            "unit": "ns/op\t  359670 B/op\t     431 allocs/op",
            "extra": "5511 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 229580,
            "unit": "ns/op",
            "extra": "5511 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359670,
            "unit": "B/op",
            "extra": "5511 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5511 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 219988,
            "unit": "ns/op\t  359670 B/op\t     431 allocs/op",
            "extra": "4993 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 219988,
            "unit": "ns/op",
            "extra": "4993 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359670,
            "unit": "B/op",
            "extra": "4993 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "4993 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 2443297,
            "unit": "ns/op\t 3023424 B/op\t    4043 allocs/op",
            "extra": "496 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 2443297,
            "unit": "ns/op",
            "extra": "496 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023424,
            "unit": "B/op",
            "extra": "496 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "496 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 2339024,
            "unit": "ns/op\t 3023410 B/op\t    4043 allocs/op",
            "extra": "463 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 2339024,
            "unit": "ns/op",
            "extra": "463 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023410,
            "unit": "B/op",
            "extra": "463 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "463 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 24503605,
            "unit": "ns/op\t42777638 B/op\t   40138 allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 24503605,
            "unit": "ns/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777638,
            "unit": "B/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "42 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 25031585,
            "unit": "ns/op\t42777630 B/op\t   40138 allocs/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 25031585,
            "unit": "ns/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777630,
            "unit": "B/op",
            "extra": "46 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "46 times\n4 procs"
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
          "id": "e54ff1ac13679fde2ca14fbeb22bc006900ac38c",
          "message": "chore(deps): update github/codeql-action action to v4.38.3 (#24)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-08T20:01:27Z",
          "url": "https://github.com/moov-io/pamspr/commit/e54ff1ac13679fde2ca14fbeb22bc006900ac38c"
        },
        "date": 1791711300207,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 157979,
            "unit": "ns/op\t  359673 B/op\t     431 allocs/op",
            "extra": "7500 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 157979,
            "unit": "ns/op",
            "extra": "7500 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359673,
            "unit": "B/op",
            "extra": "7500 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "7500 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 153827,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "6836 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 153827,
            "unit": "ns/op",
            "extra": "6836 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "6836 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6836 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1513684,
            "unit": "ns/op\t 3023410 B/op\t    4043 allocs/op",
            "extra": "766 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1513684,
            "unit": "ns/op",
            "extra": "766 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023410,
            "unit": "B/op",
            "extra": "766 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "766 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1570602,
            "unit": "ns/op\t 3023411 B/op\t    4043 allocs/op",
            "extra": "736 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1570602,
            "unit": "ns/op",
            "extra": "736 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023411,
            "unit": "B/op",
            "extra": "736 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "736 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 15173951,
            "unit": "ns/op\t42777533 B/op\t   40137 allocs/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 15173951,
            "unit": "ns/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777533,
            "unit": "B/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "79 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 15101642,
            "unit": "ns/op\t42777543 B/op\t   40137 allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 15101642,
            "unit": "ns/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777543,
            "unit": "B/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "80 times\n4 procs"
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
          "id": "e54ff1ac13679fde2ca14fbeb22bc006900ac38c",
          "message": "chore(deps): update github/codeql-action action to v4.38.3 (#24)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-08T20:01:27Z",
          "url": "https://github.com/moov-io/pamspr/commit/e54ff1ac13679fde2ca14fbeb22bc006900ac38c"
        },
        "date": 1791799552807,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 145041,
            "unit": "ns/op\t  359671 B/op\t     431 allocs/op",
            "extra": "7948 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 145041,
            "unit": "ns/op",
            "extra": "7948 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359671,
            "unit": "B/op",
            "extra": "7948 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "7948 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 137189,
            "unit": "ns/op\t  359669 B/op\t     431 allocs/op",
            "extra": "8001 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 137189,
            "unit": "ns/op",
            "extra": "8001 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359669,
            "unit": "B/op",
            "extra": "8001 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "8001 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1437571,
            "unit": "ns/op\t 3023372 B/op\t    4043 allocs/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1437571,
            "unit": "ns/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023372,
            "unit": "B/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "807 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1451003,
            "unit": "ns/op\t 3023364 B/op\t    4043 allocs/op",
            "extra": "786 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1451003,
            "unit": "ns/op",
            "extra": "786 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023364,
            "unit": "B/op",
            "extra": "786 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "786 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 14321358,
            "unit": "ns/op\t42777563 B/op\t   40137 allocs/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 14321358,
            "unit": "ns/op",
            "extra": "82 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777563,
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
            "value": 14317588,
            "unit": "ns/op\t42777549 B/op\t   40137 allocs/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 14317588,
            "unit": "ns/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777549,
            "unit": "B/op",
            "extra": "70 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "70 times\n4 procs"
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
          "id": "e54ff1ac13679fde2ca14fbeb22bc006900ac38c",
          "message": "chore(deps): update github/codeql-action action to v4.38.3 (#24)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-08T20:01:27Z",
          "url": "https://github.com/moov-io/pamspr/commit/e54ff1ac13679fde2ca14fbeb22bc006900ac38c"
        },
        "date": 1791885522240,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 141743,
            "unit": "ns/op\t  359670 B/op\t     431 allocs/op",
            "extra": "8012 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 141743,
            "unit": "ns/op",
            "extra": "8012 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359670,
            "unit": "B/op",
            "extra": "8012 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "8012 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 138297,
            "unit": "ns/op\t  359670 B/op\t     431 allocs/op",
            "extra": "8925 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 138297,
            "unit": "ns/op",
            "extra": "8925 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359670,
            "unit": "B/op",
            "extra": "8925 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "8925 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1463330,
            "unit": "ns/op\t 3023346 B/op\t    4042 allocs/op",
            "extra": "753 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1463330,
            "unit": "ns/op",
            "extra": "753 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023346,
            "unit": "B/op",
            "extra": "753 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "753 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1384159,
            "unit": "ns/op\t 3023341 B/op\t    4042 allocs/op",
            "extra": "874 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1384159,
            "unit": "ns/op",
            "extra": "874 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023341,
            "unit": "B/op",
            "extra": "874 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "874 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 14750468,
            "unit": "ns/op\t42777576 B/op\t   40138 allocs/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 14750468,
            "unit": "ns/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777576,
            "unit": "B/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "81 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 14972129,
            "unit": "ns/op\t42777589 B/op\t   40138 allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 14972129,
            "unit": "ns/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777589,
            "unit": "B/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "80 times\n4 procs"
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
          "id": "e54ff1ac13679fde2ca14fbeb22bc006900ac38c",
          "message": "chore(deps): update github/codeql-action action to v4.38.3 (#24)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-08T20:01:27Z",
          "url": "https://github.com/moov-io/pamspr/commit/e54ff1ac13679fde2ca14fbeb22bc006900ac38c"
        },
        "date": 1791965897231,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 157524,
            "unit": "ns/op\t  359673 B/op\t     431 allocs/op",
            "extra": "8091 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 157524,
            "unit": "ns/op",
            "extra": "8091 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359673,
            "unit": "B/op",
            "extra": "8091 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "8091 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 149390,
            "unit": "ns/op\t  359672 B/op\t     431 allocs/op",
            "extra": "7440 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 149390,
            "unit": "ns/op",
            "extra": "7440 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359672,
            "unit": "B/op",
            "extra": "7440 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "7440 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1505723,
            "unit": "ns/op\t 3023355 B/op\t    4043 allocs/op",
            "extra": "823 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1505723,
            "unit": "ns/op",
            "extra": "823 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023355,
            "unit": "B/op",
            "extra": "823 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "823 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1530433,
            "unit": "ns/op\t 3023358 B/op\t    4043 allocs/op",
            "extra": "721 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1530433,
            "unit": "ns/op",
            "extra": "721 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023358,
            "unit": "B/op",
            "extra": "721 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4043,
            "unit": "allocs/op",
            "extra": "721 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 14924644,
            "unit": "ns/op\t42777590 B/op\t   40138 allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 14924644,
            "unit": "ns/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777590,
            "unit": "B/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "80 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 15372327,
            "unit": "ns/op\t42777647 B/op\t   40138 allocs/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 15372327,
            "unit": "ns/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777647,
            "unit": "B/op",
            "extra": "72 times\n4 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40138,
            "unit": "allocs/op",
            "extra": "72 times\n4 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "a2676c9f0252b58b027058fa70dfd4bba4b5f2c8",
          "message": "ci: store median of 6 bench samples at 2 CPUs\n\nRun go test -count=6 -cpu=2 and collapse each benchmark to a median\nline before github-action-benchmark records it. Raise alert-threshold\nto 150% and pin ubuntu-24.04.",
          "timestamp": "2026-10-14T14:53:19Z",
          "url": "https://github.com/moov-io/pamspr/commit/a2676c9f0252b58b027058fa70dfd4bba4b5f2c8"
        },
        "date": 1791989815471,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 187842,
            "unit": "ns/op\t359635 B/op\t431 allocs/op",
            "extra": "6046 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 187842,
            "unit": "ns/op",
            "extra": "6046 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359635,
            "unit": "B/op",
            "extra": "6046 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6046 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 187233,
            "unit": "ns/op\t359635 B/op\t431 allocs/op",
            "extra": "6117 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 187233,
            "unit": "ns/op",
            "extra": "6117 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359635,
            "unit": "B/op",
            "extra": "6117 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6117 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1956047,
            "unit": "ns/op\t3023012 B/op\t4041 allocs/op",
            "extra": "621 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1956047,
            "unit": "ns/op",
            "extra": "621 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023012,
            "unit": "B/op",
            "extra": "621 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "621 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1944921,
            "unit": "ns/op\t3023008 B/op\t4041 allocs/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1944921,
            "unit": "ns/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023008,
            "unit": "B/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 20337922,
            "unit": "ns/op\t42777251 B/op\t40137 allocs/op",
            "extra": "50 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 20337922,
            "unit": "ns/op",
            "extra": "50 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777251,
            "unit": "B/op",
            "extra": "50 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "50 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 20386435,
            "unit": "ns/op\t42777230 B/op\t40137 allocs/op",
            "extra": "50 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 20386435,
            "unit": "ns/op",
            "extra": "50 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777230,
            "unit": "B/op",
            "extra": "50 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "50 times\n2 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "f12b47f963103fe2f9a318c8c725ca7f0308082c",
          "message": "ci: keep bench.yml on ubuntu-24.04\n\nRenovate automerged ubuntu-26.04 onto the benchmark workflow. Pin\n24.04 again and ignore that package in bench.yml so nightly samples\nstay comparable.",
          "timestamp": "2026-10-14T14:59:48Z",
          "url": "https://github.com/moov-io/pamspr/commit/f12b47f963103fe2f9a318c8c725ca7f0308082c"
        },
        "date": 1791990602909,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 191187,
            "unit": "ns/op\t359635 B/op\t431 allocs/op",
            "extra": "6026 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 191187,
            "unit": "ns/op",
            "extra": "6026 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359635,
            "unit": "B/op",
            "extra": "6026 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6026 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 196926,
            "unit": "ns/op\t359635 B/op\t431 allocs/op",
            "extra": "6003 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 196926,
            "unit": "ns/op",
            "extra": "6003 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359635,
            "unit": "B/op",
            "extra": "6003 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6003 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 2031147,
            "unit": "ns/op\t3023013 B/op\t4041 allocs/op",
            "extra": "580 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 2031147,
            "unit": "ns/op",
            "extra": "580 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023013,
            "unit": "B/op",
            "extra": "580 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "580 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1911136,
            "unit": "ns/op\t3023008 B/op\t4041 allocs/op",
            "extra": "621 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1911136,
            "unit": "ns/op",
            "extra": "621 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023008,
            "unit": "B/op",
            "extra": "621 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "621 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 21240713,
            "unit": "ns/op\t42777242 B/op\t40137 allocs/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 21240713,
            "unit": "ns/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777242,
            "unit": "B/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 21608376,
            "unit": "ns/op\t42777234 B/op\t40137 allocs/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 21608376,
            "unit": "ns/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777234,
            "unit": "B/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "55 times\n2 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "f12b47f963103fe2f9a318c8c725ca7f0308082c",
          "message": "ci: keep bench.yml on ubuntu-24.04\n\nRenovate automerged ubuntu-26.04 onto the benchmark workflow. Pin\n24.04 again and ignore that package in bench.yml so nightly samples\nstay comparable.",
          "timestamp": "2026-10-14T14:59:48Z",
          "url": "https://github.com/moov-io/pamspr/commit/f12b47f963103fe2f9a318c8c725ca7f0308082c"
        },
        "date": 1791991752987,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 213457.5,
            "unit": "ns/op\t359633 B/op\t431 allocs/op",
            "extra": "5370 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 213457.5,
            "unit": "ns/op",
            "extra": "5370 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359633,
            "unit": "B/op",
            "extra": "5370 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5370 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 209987.5,
            "unit": "ns/op\t359633 B/op\t431 allocs/op",
            "extra": "5462 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 209987.5,
            "unit": "ns/op",
            "extra": "5462 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359633,
            "unit": "B/op",
            "extra": "5462 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5462 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 2150073,
            "unit": "ns/op\t3023041 B/op\t4042 allocs/op",
            "extra": "550 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 2150073,
            "unit": "ns/op",
            "extra": "550 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023041,
            "unit": "B/op",
            "extra": "550 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "550 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 2163410.5,
            "unit": "ns/op\t3023033 B/op\t4042 allocs/op",
            "extra": "552 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 2163410.5,
            "unit": "ns/op",
            "extra": "552 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023033,
            "unit": "B/op",
            "extra": "552 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "552 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 25609884.5,
            "unit": "ns/op\t42777182 B/op\t40136 allocs/op",
            "extra": "46 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 25609884.5,
            "unit": "ns/op",
            "extra": "46 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777182,
            "unit": "B/op",
            "extra": "46 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40136,
            "unit": "allocs/op",
            "extra": "46 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 25799433,
            "unit": "ns/op\t42777187 B/op\t40136 allocs/op",
            "extra": "46 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 25799433,
            "unit": "ns/op",
            "extra": "46 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777187,
            "unit": "B/op",
            "extra": "46 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40136,
            "unit": "allocs/op",
            "extra": "46 times\n2 procs"
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
            "name": "Adam Shannon",
            "username": "adamdecaf",
            "email": "adamkshannon@gmail.com"
          },
          "id": "f12b47f963103fe2f9a318c8c725ca7f0308082c",
          "message": "ci: keep bench.yml on ubuntu-24.04\n\nRenovate automerged ubuntu-26.04 onto the benchmark workflow. Pin\n24.04 again and ignore that package in bench.yml so nightly samples\nstay comparable.",
          "timestamp": "2026-10-14T14:59:48Z",
          "url": "https://github.com/moov-io/pamspr/commit/f12b47f963103fe2f9a318c8c725ca7f0308082c"
        },
        "date": 1792050954862,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 199514,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "5874 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 199514,
            "unit": "ns/op",
            "extra": "5874 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "5874 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5874 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 195069,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "5724 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 195069,
            "unit": "ns/op",
            "extra": "5724 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "5724 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5724 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 2110509.5,
            "unit": "ns/op\t3023002 B/op\t4041 allocs/op",
            "extra": "565 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 2110509.5,
            "unit": "ns/op",
            "extra": "565 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023002,
            "unit": "B/op",
            "extra": "565 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "565 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 2125948,
            "unit": "ns/op\t3023008 B/op\t4041 allocs/op",
            "extra": "549 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 2125948,
            "unit": "ns/op",
            "extra": "549 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023008,
            "unit": "B/op",
            "extra": "549 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "549 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 21483605.5,
            "unit": "ns/op\t42777256 B/op\t40137 allocs/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 21483605.5,
            "unit": "ns/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777256,
            "unit": "B/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 21669874,
            "unit": "ns/op\t42777258 B/op\t40137 allocs/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 21669874,
            "unit": "ns/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777258,
            "unit": "B/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "55 times\n2 procs"
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
          "id": "daeee281b4989a97eee6b5f13615294036b96a74",
          "message": "chore(deps): update github/codeql-action action to v4.39.0 (#26)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-15T11:47:02Z",
          "url": "https://github.com/moov-io/pamspr/commit/daeee281b4989a97eee6b5f13615294036b96a74"
        },
        "date": 1792136929515,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 186567,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "6409 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 186567,
            "unit": "ns/op",
            "extra": "6409 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "6409 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6409 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 186254.5,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "6213 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 186254.5,
            "unit": "ns/op",
            "extra": "6213 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "6213 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6213 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1973254.5,
            "unit": "ns/op\t3023012 B/op\t4041 allocs/op",
            "extra": "604 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1973254.5,
            "unit": "ns/op",
            "extra": "604 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023012,
            "unit": "B/op",
            "extra": "604 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "604 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1980403.5,
            "unit": "ns/op\t3023013 B/op\t4041 allocs/op",
            "extra": "614 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1980403.5,
            "unit": "ns/op",
            "extra": "614 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023013,
            "unit": "B/op",
            "extra": "614 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "614 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 20082821.5,
            "unit": "ns/op\t42777254 B/op\t40137 allocs/op",
            "extra": "51 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 20082821.5,
            "unit": "ns/op",
            "extra": "51 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777254,
            "unit": "B/op",
            "extra": "51 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "51 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 20312858.5,
            "unit": "ns/op\t42777206 B/op\t40136 allocs/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 20312858.5,
            "unit": "ns/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777206,
            "unit": "B/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40136,
            "unit": "allocs/op",
            "extra": "58 times\n2 procs"
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
          "id": "daeee281b4989a97eee6b5f13615294036b96a74",
          "message": "chore(deps): update github/codeql-action action to v4.39.0 (#26)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-15T11:47:02Z",
          "url": "https://github.com/moov-io/pamspr/commit/daeee281b4989a97eee6b5f13615294036b96a74"
        },
        "date": 1792221974159,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 186429.5,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "5938 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 186429.5,
            "unit": "ns/op",
            "extra": "5938 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "5938 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5938 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 186787.5,
            "unit": "ns/op\t359635 B/op\t431 allocs/op",
            "extra": "6061 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 186787.5,
            "unit": "ns/op",
            "extra": "6061 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359635,
            "unit": "B/op",
            "extra": "6061 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6061 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 2024251,
            "unit": "ns/op\t3023008 B/op\t4041 allocs/op",
            "extra": "602 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 2024251,
            "unit": "ns/op",
            "extra": "602 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023008,
            "unit": "B/op",
            "extra": "602 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "602 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1984463.5,
            "unit": "ns/op\t3023006 B/op\t4041 allocs/op",
            "extra": "600 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1984463.5,
            "unit": "ns/op",
            "extra": "600 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023006,
            "unit": "B/op",
            "extra": "600 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "600 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 20507492.5,
            "unit": "ns/op\t42777251 B/op\t40137 allocs/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 20507492.5,
            "unit": "ns/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777251,
            "unit": "B/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 20590686,
            "unit": "ns/op\t42777220 B/op\t40136 allocs/op",
            "extra": "56 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 20590686,
            "unit": "ns/op",
            "extra": "56 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777220,
            "unit": "B/op",
            "extra": "56 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40136,
            "unit": "allocs/op",
            "extra": "56 times\n2 procs"
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
          "id": "daeee281b4989a97eee6b5f13615294036b96a74",
          "message": "chore(deps): update github/codeql-action action to v4.39.0 (#26)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-15T11:47:02Z",
          "url": "https://github.com/moov-io/pamspr/commit/daeee281b4989a97eee6b5f13615294036b96a74"
        },
        "date": 1792309282123,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 190914,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "6010 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 190914,
            "unit": "ns/op",
            "extra": "6010 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "6010 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6010 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 186824,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "6153 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 186824,
            "unit": "ns/op",
            "extra": "6153 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "6153 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6153 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1942714.5,
            "unit": "ns/op\t3023014 B/op\t4041 allocs/op",
            "extra": "613 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1942714.5,
            "unit": "ns/op",
            "extra": "613 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023014,
            "unit": "B/op",
            "extra": "613 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "613 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1966191.5,
            "unit": "ns/op\t3023012 B/op\t4041 allocs/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1966191.5,
            "unit": "ns/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023012,
            "unit": "B/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 19840503.5,
            "unit": "ns/op\t42777269 B/op\t40137 allocs/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 19840503.5,
            "unit": "ns/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777269,
            "unit": "B/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 19784735,
            "unit": "ns/op\t42777266 B/op\t40137 allocs/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 19784735,
            "unit": "ns/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777266,
            "unit": "B/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "58 times\n2 procs"
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
          "id": "daeee281b4989a97eee6b5f13615294036b96a74",
          "message": "chore(deps): update github/codeql-action action to v4.39.0 (#26)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-15T11:47:02Z",
          "url": "https://github.com/moov-io/pamspr/commit/daeee281b4989a97eee6b5f13615294036b96a74"
        },
        "date": 1792398215253,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 185990,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "6236 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 185990,
            "unit": "ns/op",
            "extra": "6236 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "6236 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6236 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 187544,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "6470 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 187544,
            "unit": "ns/op",
            "extra": "6470 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "6470 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6470 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 2048023.5,
            "unit": "ns/op\t3023021 B/op\t4042 allocs/op",
            "extra": "590 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 2048023.5,
            "unit": "ns/op",
            "extra": "590 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023021,
            "unit": "B/op",
            "extra": "590 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "590 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 2007477,
            "unit": "ns/op\t3023017 B/op\t4041 allocs/op",
            "extra": "591 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 2007477,
            "unit": "ns/op",
            "extra": "591 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023017,
            "unit": "B/op",
            "extra": "591 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "591 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 20246128,
            "unit": "ns/op\t42777252 B/op\t40137 allocs/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 20246128,
            "unit": "ns/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777252,
            "unit": "B/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 20165726.5,
            "unit": "ns/op\t42777220 B/op\t40136 allocs/op",
            "extra": "57 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 20165726.5,
            "unit": "ns/op",
            "extra": "57 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777220,
            "unit": "B/op",
            "extra": "57 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40136,
            "unit": "allocs/op",
            "extra": "57 times\n2 procs"
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
          "id": "daeee281b4989a97eee6b5f13615294036b96a74",
          "message": "chore(deps): update github/codeql-action action to v4.39.0 (#26)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-15T11:47:02Z",
          "url": "https://github.com/moov-io/pamspr/commit/daeee281b4989a97eee6b5f13615294036b96a74"
        },
        "date": 1792483635462,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 190258.5,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "5955 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 190258.5,
            "unit": "ns/op",
            "extra": "5955 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "5955 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5955 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 187500,
            "unit": "ns/op\t359635 B/op\t431 allocs/op",
            "extra": "6186 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 187500,
            "unit": "ns/op",
            "extra": "6186 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359635,
            "unit": "B/op",
            "extra": "6186 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6186 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1959261.5,
            "unit": "ns/op\t3023016 B/op\t4041 allocs/op",
            "extra": "606 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1959261.5,
            "unit": "ns/op",
            "extra": "606 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023016,
            "unit": "B/op",
            "extra": "606 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "606 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1924456.5,
            "unit": "ns/op\t3023004 B/op\t4041 allocs/op",
            "extra": "612 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1924456.5,
            "unit": "ns/op",
            "extra": "612 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023004,
            "unit": "B/op",
            "extra": "612 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "612 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 20465705,
            "unit": "ns/op\t42777230 B/op\t40136 allocs/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 20465705,
            "unit": "ns/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777230,
            "unit": "B/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40136,
            "unit": "allocs/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 20518221.5,
            "unit": "ns/op\t42777232 B/op\t40137 allocs/op",
            "extra": "51 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 20518221.5,
            "unit": "ns/op",
            "extra": "51 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777232,
            "unit": "B/op",
            "extra": "51 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "51 times\n2 procs"
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
          "id": "daeee281b4989a97eee6b5f13615294036b96a74",
          "message": "chore(deps): update github/codeql-action action to v4.39.0 (#26)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-15T11:47:02Z",
          "url": "https://github.com/moov-io/pamspr/commit/daeee281b4989a97eee6b5f13615294036b96a74"
        },
        "date": 1792570367196,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 182322,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "5965 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 182322,
            "unit": "ns/op",
            "extra": "5965 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "5965 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5965 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 182250,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "6075 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 182250,
            "unit": "ns/op",
            "extra": "6075 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "6075 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6075 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1882249.5,
            "unit": "ns/op\t3023028 B/op\t4042 allocs/op",
            "extra": "643 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1882249.5,
            "unit": "ns/op",
            "extra": "643 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023028,
            "unit": "B/op",
            "extra": "643 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "643 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1894941.5,
            "unit": "ns/op\t3023025 B/op\t4042 allocs/op",
            "extra": "632 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1894941.5,
            "unit": "ns/op",
            "extra": "632 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023025,
            "unit": "B/op",
            "extra": "632 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "632 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 23225376.5,
            "unit": "ns/op\t42777244 B/op\t40137 allocs/op",
            "extra": "49 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 23225376.5,
            "unit": "ns/op",
            "extra": "49 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777244,
            "unit": "B/op",
            "extra": "49 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "49 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 23330062.5,
            "unit": "ns/op\t42777209 B/op\t40136 allocs/op",
            "extra": "49 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 23330062.5,
            "unit": "ns/op",
            "extra": "49 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777209,
            "unit": "B/op",
            "extra": "49 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40136,
            "unit": "allocs/op",
            "extra": "49 times\n2 procs"
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
          "id": "f6e2bb0a059473ca64faaa91b4c1a0e5e40530e0",
          "message": "chore(deps): update github/codeql-action action to v4.39.1 (#27)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-21T19:03:16Z",
          "url": "https://github.com/moov-io/pamspr/commit/f6e2bb0a059473ca64faaa91b4c1a0e5e40530e0"
        },
        "date": 1792656859731,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 93167.5,
            "unit": "ns/op\t359635 B/op\t431 allocs/op",
            "extra": "12786 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 93167.5,
            "unit": "ns/op",
            "extra": "12786 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359635,
            "unit": "B/op",
            "extra": "12786 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "12786 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 93734.5,
            "unit": "ns/op\t359636 B/op\t431 allocs/op",
            "extra": "12849 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 93734.5,
            "unit": "ns/op",
            "extra": "12849 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359636,
            "unit": "B/op",
            "extra": "12849 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "12849 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 964571,
            "unit": "ns/op\t3023015 B/op\t4041 allocs/op",
            "extra": "1255 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 964571,
            "unit": "ns/op",
            "extra": "1255 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023015,
            "unit": "B/op",
            "extra": "1255 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "1255 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 973459,
            "unit": "ns/op\t3023018 B/op\t4042 allocs/op",
            "extra": "1261 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 973459,
            "unit": "ns/op",
            "extra": "1261 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023018,
            "unit": "B/op",
            "extra": "1261 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4042,
            "unit": "allocs/op",
            "extra": "1261 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 10486395.5,
            "unit": "ns/op\t42777259 B/op\t40137 allocs/op",
            "extra": "116 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 10486395.5,
            "unit": "ns/op",
            "extra": "116 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777259,
            "unit": "B/op",
            "extra": "116 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "116 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 10152854.5,
            "unit": "ns/op\t42777253 B/op\t40137 allocs/op",
            "extra": "117 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 10152854.5,
            "unit": "ns/op",
            "extra": "117 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777253,
            "unit": "B/op",
            "extra": "117 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "117 times\n2 procs"
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
          "id": "f6e2bb0a059473ca64faaa91b4c1a0e5e40530e0",
          "message": "chore(deps): update github/codeql-action action to v4.39.1 (#27)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-21T19:03:16Z",
          "url": "https://github.com/moov-io/pamspr/commit/f6e2bb0a059473ca64faaa91b4c1a0e5e40530e0"
        },
        "date": 1792743113407,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 133607,
            "unit": "ns/op\t359635 B/op\t431 allocs/op",
            "extra": "8518 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 133607,
            "unit": "ns/op",
            "extra": "8518 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359635,
            "unit": "B/op",
            "extra": "8518 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "8518 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 132747,
            "unit": "ns/op\t359635 B/op\t431 allocs/op",
            "extra": "9070 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 132747,
            "unit": "ns/op",
            "extra": "9070 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359635,
            "unit": "B/op",
            "extra": "9070 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "9070 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1331006,
            "unit": "ns/op\t3023004 B/op\t4041 allocs/op",
            "extra": "883 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1331006,
            "unit": "ns/op",
            "extra": "883 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023004,
            "unit": "B/op",
            "extra": "883 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "883 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1304481.5,
            "unit": "ns/op\t3022994 B/op\t4041 allocs/op",
            "extra": "925 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1304481.5,
            "unit": "ns/op",
            "extra": "925 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3022994,
            "unit": "B/op",
            "extra": "925 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "925 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 13850197,
            "unit": "ns/op\t42777242 B/op\t40137 allocs/op",
            "extra": "75 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 13850197,
            "unit": "ns/op",
            "extra": "75 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777242,
            "unit": "B/op",
            "extra": "75 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "75 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 14066917.5,
            "unit": "ns/op\t42777221 B/op\t40136 allocs/op",
            "extra": "85 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 14066917.5,
            "unit": "ns/op",
            "extra": "85 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777221,
            "unit": "B/op",
            "extra": "85 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40136,
            "unit": "allocs/op",
            "extra": "85 times\n2 procs"
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
          "id": "f6e2bb0a059473ca64faaa91b4c1a0e5e40530e0",
          "message": "chore(deps): update github/codeql-action action to v4.39.1 (#27)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-21T19:03:16Z",
          "url": "https://github.com/moov-io/pamspr/commit/f6e2bb0a059473ca64faaa91b4c1a0e5e40530e0"
        },
        "date": 1792828192085,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 188147.5,
            "unit": "ns/op\t359635 B/op\t431 allocs/op",
            "extra": "5998 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 188147.5,
            "unit": "ns/op",
            "extra": "5998 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359635,
            "unit": "B/op",
            "extra": "5998 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5998 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 187497.5,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "6048 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 187497.5,
            "unit": "ns/op",
            "extra": "6048 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "6048 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6048 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1933374,
            "unit": "ns/op\t3023014 B/op\t4041 allocs/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1933374,
            "unit": "ns/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023014,
            "unit": "B/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1927715,
            "unit": "ns/op\t3023011 B/op\t4041 allocs/op",
            "extra": "626 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1927715,
            "unit": "ns/op",
            "extra": "626 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023011,
            "unit": "B/op",
            "extra": "626 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "626 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 20045285,
            "unit": "ns/op\t42777215 B/op\t40136 allocs/op",
            "extra": "51 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 20045285,
            "unit": "ns/op",
            "extra": "51 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777215,
            "unit": "B/op",
            "extra": "51 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40136,
            "unit": "allocs/op",
            "extra": "51 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 20345235,
            "unit": "ns/op\t42777212 B/op\t40136 allocs/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 20345235,
            "unit": "ns/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777212,
            "unit": "B/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40136,
            "unit": "allocs/op",
            "extra": "58 times\n2 procs"
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
          "id": "f6e2bb0a059473ca64faaa91b4c1a0e5e40530e0",
          "message": "chore(deps): update github/codeql-action action to v4.39.1 (#27)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-21T19:03:16Z",
          "url": "https://github.com/moov-io/pamspr/commit/f6e2bb0a059473ca64faaa91b4c1a0e5e40530e0"
        },
        "date": 1792915320246,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 189658.5,
            "unit": "ns/op\t359635 B/op\t431 allocs/op",
            "extra": "6091 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 189658.5,
            "unit": "ns/op",
            "extra": "6091 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359635,
            "unit": "B/op",
            "extra": "6091 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6091 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 190490.5,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "6033 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 190490.5,
            "unit": "ns/op",
            "extra": "6033 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "6033 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6033 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1957456.5,
            "unit": "ns/op\t3023006 B/op\t4041 allocs/op",
            "extra": "607 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1957456.5,
            "unit": "ns/op",
            "extra": "607 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023006,
            "unit": "B/op",
            "extra": "607 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "607 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1954693.5,
            "unit": "ns/op\t3023006 B/op\t4041 allocs/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1954693.5,
            "unit": "ns/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023006,
            "unit": "B/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "618 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 20707907,
            "unit": "ns/op\t42777226 B/op\t40137 allocs/op",
            "extra": "57 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 20707907,
            "unit": "ns/op",
            "extra": "57 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777226,
            "unit": "B/op",
            "extra": "57 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "57 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 20384573,
            "unit": "ns/op\t42777240 B/op\t40137 allocs/op",
            "extra": "56 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 20384573,
            "unit": "ns/op",
            "extra": "56 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777240,
            "unit": "B/op",
            "extra": "56 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "56 times\n2 procs"
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
          "id": "f6e2bb0a059473ca64faaa91b4c1a0e5e40530e0",
          "message": "chore(deps): update github/codeql-action action to v4.39.1 (#27)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-21T19:03:16Z",
          "url": "https://github.com/moov-io/pamspr/commit/f6e2bb0a059473ca64faaa91b4c1a0e5e40530e0"
        },
        "date": 1793003024556,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 187881.5,
            "unit": "ns/op\t359635 B/op\t431 allocs/op",
            "extra": "5889 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 187881.5,
            "unit": "ns/op",
            "extra": "5889 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359635,
            "unit": "B/op",
            "extra": "5889 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5889 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 187673.5,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "6061 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 187673.5,
            "unit": "ns/op",
            "extra": "6061 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "6061 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6061 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1998159,
            "unit": "ns/op\t3023013 B/op\t4041 allocs/op",
            "extra": "598 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1998159,
            "unit": "ns/op",
            "extra": "598 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023013,
            "unit": "B/op",
            "extra": "598 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "598 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 2018023.5,
            "unit": "ns/op\t3023011 B/op\t4041 allocs/op",
            "extra": "607 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 2018023.5,
            "unit": "ns/op",
            "extra": "607 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023011,
            "unit": "B/op",
            "extra": "607 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "607 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 20897161,
            "unit": "ns/op\t42777250 B/op\t40137 allocs/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 20897161,
            "unit": "ns/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777250,
            "unit": "B/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "55 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 21109214.5,
            "unit": "ns/op\t42777248 B/op\t40137 allocs/op",
            "extra": "49 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 21109214.5,
            "unit": "ns/op",
            "extra": "49 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777248,
            "unit": "B/op",
            "extra": "49 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "49 times\n2 procs"
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
          "id": "f6e2bb0a059473ca64faaa91b4c1a0e5e40530e0",
          "message": "chore(deps): update github/codeql-action action to v4.39.1 (#27)\n\nCo-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>",
          "timestamp": "2026-10-21T19:03:16Z",
          "url": "https://github.com/moov-io/pamspr/commit/f6e2bb0a059473ca64faaa91b4c1a0e5e40530e0"
        },
        "date": 1793090061893,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments",
            "value": 187857,
            "unit": "ns/op\t359634 B/op\t431 allocs/op",
            "extra": "5966 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - ns/op",
            "value": 187857,
            "unit": "ns/op",
            "extra": "5966 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - B/op",
            "value": 359634,
            "unit": "B/op",
            "extra": "5966 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "5966 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments",
            "value": 191244,
            "unit": "ns/op\t359635 B/op\t431 allocs/op",
            "extra": "6280 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - ns/op",
            "value": 191244,
            "unit": "ns/op",
            "extra": "6280 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - B/op",
            "value": 359635,
            "unit": "B/op",
            "extra": "6280 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_100_payments - allocs/op",
            "value": 431,
            "unit": "allocs/op",
            "extra": "6280 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments",
            "value": 1951995.5,
            "unit": "ns/op\t3023014 B/op\t4041 allocs/op",
            "extra": "607 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - ns/op",
            "value": 1951995.5,
            "unit": "ns/op",
            "extra": "607 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - B/op",
            "value": 3023014,
            "unit": "B/op",
            "extra": "607 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "607 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments",
            "value": 1982836,
            "unit": "ns/op\t3023014 B/op\t4041 allocs/op",
            "extra": "621 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - ns/op",
            "value": 1982836,
            "unit": "ns/op",
            "extra": "621 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - B/op",
            "value": 3023014,
            "unit": "B/op",
            "extra": "621 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_1000_payments - allocs/op",
            "value": 4041,
            "unit": "allocs/op",
            "extra": "621 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments",
            "value": 20401417,
            "unit": "ns/op\t42777236 B/op\t40137 allocs/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - ns/op",
            "value": 20401417,
            "unit": "ns/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - B/op",
            "value": 42777236,
            "unit": "B/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Traditional_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "58 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments",
            "value": 20252135.5,
            "unit": "ns/op\t42777248 B/op\t40137 allocs/op",
            "extra": "51 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - ns/op",
            "value": 20252135.5,
            "unit": "ns/op",
            "extra": "51 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - B/op",
            "value": 42777248,
            "unit": "B/op",
            "extra": "51 times\n2 procs"
          },
          {
            "name": "BenchmarkStreamingWriter_vs_TraditionalWriter/Streaming_Writer_10000_payments - allocs/op",
            "value": 40137,
            "unit": "allocs/op",
            "extra": "51 times\n2 procs"
          }
        ]
      }
    ]
  }
}