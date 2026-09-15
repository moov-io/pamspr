window.BENCHMARK_DATA = {
  "lastUpdate": 1789460301459,
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
      }
    ]
  }
}