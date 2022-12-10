# go-router-benchmark
Compare the performance of routers built with golang.

# Motivation
I have implemented a router called [bmf-san/goblin](https://github.com/bmf-san/goblin), and I created this repository to compare it with other routers and get hints on how to improve [bmf-san/goblin](https://github.com/bmf-san/goblin).

Another reason is that [this nice repository](https://github.com/julienschmidt/go-http-routing-benchmark) seems to have stopped being maintained, so I wanted to have a benchmark test for the router that I could manage myself.

# Benchmark test
This benchmark test is not a complete representation of router performance differences.

This is because it is difficult to prepare a perfect test case due to differences in router specifications and different data structures.

Benchmarks are obtained by narrowing down to specific functions.

# Listed routers
- [bmf-san/goblin](https://github.com/bmf-san/goblin)
- [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
- [go-chi/chi](https://github.com/go-chi/chi)
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [uptrace/bunrouter](https://github.com/uptrace/bunrouter)
- [dimfeld/httptreemux](https://github.com/dimfeld/httptreemux)
- [beego/mux](https://github.com/beego/mux)
- [gorilla/mux](https://github.com/gorilla/mux)
- [nissy/bon](https://github.com/nissy/bon)

# How to run benchmark test
`make test-benchmark`

## Commands for running benchmark test
`make test-benchmark`

## Results
```
go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/go-router-benchmark
cpu: VirtualApple @ 2.50GHz
BenchmarkStaticRoutesRootGoblin-8                       32912930                31.54 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Goblin-8                          18060241                65.13 ns/op       16 B/op          1 allocs/op
BenchmarkStaticRoutes5Goblin-8                           6144348               201.7 ns/op        80 B/op          1 allocs/op
BenchmarkStaticRoutes10Goblin-8                          3275200               370.7 ns/op       160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes1Goblin-8                   1943931               661.0 ns/op       408 B/op          6 allocs/op
BenchmarkPathParamColonRoutes5Goblin-8                    520279              2285 ns/op         961 B/op         13 allocs/op
BenchmarkPathParamColonRoutes10Goblin-8                   281046              4281 ns/op        1604 B/op         19 allocs/op
BenchmarkStaticRoutesRootHTTPRouter-8                   93391267                13.11 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1HTTPRouter-8                      93725287                13.23 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5HTTPRouter-8                      93085196                13.15 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10HTTPRouter-8                     93078272                13.07 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1HTTPRouter-8              26540427                44.05 ns/op       32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes5HTTPRouter-8              10597780               112.7 ns/op       160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes10HTTPRouter-8              6192882               192.0 ns/op       320 B/op          1 allocs/op
BenchmarkStaticRoutesRootChi-8                           5741954               206.5 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutes1Chi-8                              5165967               203.9 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutes5Chi-8                              5891289               203.6 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutes10Chi-8                             5941209               201.4 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes1Chi-8                    4625229               257.8 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes5Chi-8                    2852914               418.9 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes10Chi-8                   1961700               611.0 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutesRootGin-8                          33409586                35.45 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Gin-8                             33237926                35.91 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5Gin-8                             33254545                36.96 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10Gin-8                            31706293                37.12 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Gin-8                     27547839                43.19 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5Gin-8                     13762492                81.08 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10Gin-8                     9499948               125.4 ns/op         0 B/op          0 allocs/op
BenchmarkStaticRoutesRootBunRouter-8                    62109922                18.19 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1BunRouter-8                       56690571                20.72 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5BunRouter-8                       56904399                20.77 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10BunRouter-8                      56887316                21.16 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1BunRouter-8               37093228                32.23 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5BunRouter-8                8553456               138.3 ns/op         0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10BunRouter-8               4257510               283.4 ns/op         0 B/op          0 allocs/op
BenchmarkStaticRoutesRootHTTPTreeMux-8                   6947364               174.4 ns/op       328 B/op          3 allocs/op
BenchmarkStaticRoutes1HTTPTreeMux-8                      6733014               177.0 ns/op       328 B/op          3 allocs/op
BenchmarkStaticRoutes5HTTPTreeMux-8                      5447610               223.1 ns/op       328 B/op          3 allocs/op
BenchmarkStaticRoutes10HTTPTreeMux-8                     4118466               330.0 ns/op       328 B/op          3 allocs/op
BenchmarkPathParamColonRoutes1HTTPTreeMux-8              3018307               413.8 ns/op       680 B/op          6 allocs/op
BenchmarkPathParamColonRoutes5HTTPTreeMux-8              1581075               770.8 ns/op       904 B/op          9 allocs/op
BenchmarkPathParamColonRoutes10HTTPTreeMux-8              789046              1432 ns/op        1742 B/op         11 allocs/op
BenchmarkStaticRoutesRootBeegoMux-8                     24145880                48.34 ns/op       32 B/op          1 allocs/op
BenchmarkStaticRoutes1BeegoMux-8                        16852018                71.90 ns/op       32 B/op          1 allocs/op
BenchmarkStaticRoutes5BeegoMux-8                         1000000              1126 ns/op          32 B/op          1 allocs/op
BenchmarkStaticRoutes10BeegoMux-8                         540434              2057 ns/op          32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes1BeegoMux-8                 3222766               364.9 ns/op       672 B/op          5 allocs/op
BenchmarkPathParamColonRoutes5BeegoMux-8                  822138              1409 ns/op         672 B/op          5 allocs/op
BenchmarkPathParamColonRoutes10BeegoMux-8                 355040              3312 ns/op        1254 B/op          6 allocs/op
BenchmarkStaticRoutesRootGorillaMux-8                    2122923               564.8 ns/op       720 B/op          7 allocs/op
BenchmarkStaticRoutes1GorillaMux-8                       1990176               569.1 ns/op       720 B/op          7 allocs/op
BenchmarkStaticRoutes5GorillaMux-8                       1972658               592.2 ns/op       720 B/op          7 allocs/op
BenchmarkStaticRoutes10GorillaMux-8                      1887168               642.8 ns/op       720 B/op          7 allocs/op
BenchmarkPathParamBracketRoutes1GorillaMux-8             1374578               889.1 ns/op      1024 B/op          8 allocs/op
BenchmarkPathParamBracketRoutes5GorillaMux-8              506086              2260 ns/op        1088 B/op          8 allocs/op
BenchmarkPathParamBracketRoutes10GorillaMux-8             233474              4988 ns/op        1751 B/op          9 allocs/op
BenchmarkStaticRoutesRootBon-8                          86125429                14.62 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Bon-8                             84076539                14.57 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5Bon-8                             83559884                14.52 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10Bon-8                            79310877                15.62 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Bon-8                      6237304               293.3 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamColonRoutes5Bon-8                      3235410               480.8 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamColonRoutes10Bon-8                     3211866               374.8 ns/op       304 B/op          2 allocs/op
PASS
ok      github.com/go-router-benchmark  90.756s
```

# Contribution
We are always accepting issues, pull requests, and other requests and questions.

We look forward to your contribution！

# License
This project is licensed under the terms of the MIT license.

## Author
bmf - A Web Developer in Japan.

- [@bmf-san](https://twitter.com/bmf_san)
- [bmf-tech](http://bmf-tech.com/)