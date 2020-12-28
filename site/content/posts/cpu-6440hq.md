---
title: "i5-6440HQ CPU benchmark"
date: 2018-03-11T22:08:31+01:00
draft: false
tags: ["linux", "benchmark", "cpu", "sysbench"]
---

CPU benchmark from Dell Precision 15 5510 Core-i5 i5-6440HQ
```bash
$ sysbench cpu run --time=5
sysbench 1.0.9 (using system LuaJIT 2.1.0-beta3)

Running the test with following options:
Number of threads: 1
Initializing random number generator from current time


Prime numbers limit: 10000

Initializing worker threads...

Threads started!

CPU speed:
    events per second:  1157.88

General statistics:
    total time:                          5.0008s
    total number of events:              5792

Latency (ms):
         min:                                  0.80
         avg:                                  0.86
         max:                                  3.74
         95th percentile:                      1.06
         sum:                               4998.26

Threads fairness:
    events (avg/stddev):           5792.0000/0.00
    execution time (avg/stddev):   4.9983/0.00
```
