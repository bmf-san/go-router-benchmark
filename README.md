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
- [naoina/denco](https://github.com/naoina/denco)
- [labstack/echo](https://github.com/labstack/echo/v4)

# How to run benchmark test
`make test-benchmark`

## Commands for running benchmark test
`make test-benchmark`

## Results
```sh
go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/go-router-benchmark
cpu: VirtualApple @ 2.50GHz
BenchmarkStaticRoutesRootGoblin-8                       37808821                31.21 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Goblin-8                          17749246                66.14 ns/op       16 B/op          1 allocs/op
BenchmarkStaticRoutes5Goblin-8                           6124653               197.5 ns/op        80 B/op          1 allocs/op
BenchmarkStaticRoutes10Goblin-8                          3210261               371.8 ns/op       160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes1Goblin-8                   1905600               628.9 ns/op       407 B/op          6 allocs/op
BenchmarkPathParamColonRoutes5Goblin-8                    511957              2297 ns/op         963 B/op         13 allocs/op
BenchmarkPathParamColonRoutes10Goblin-8                   269656              4482 ns/op        1606 B/op         19 allocs/op
BenchmarkStaticRoutesRootHTTPRouter-8                   75625483                14.98 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1HTTPRouter-8                      78181624                15.00 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5HTTPRouter-8                      75628862                15.01 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10HTTPRouter-8                     78458509                15.01 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1HTTPRouter-8              24709384                48.00 ns/op       32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes5HTTPRouter-8              10060582               117.2 ns/op       160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes10HTTPRouter-8              6091068               199.3 ns/op       320 B/op          1 allocs/op
BenchmarkStaticRoutesRootChi-8                           5681046               202.7 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutes1Chi-8                              5938426               210.9 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutes5Chi-8                              5887243               203.3 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutes10Chi-8                             5959963               201.1 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes1Chi-8                    4591572               259.3 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes5Chi-8                    2871210               417.3 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes10Chi-8                   1950718               619.0 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutesRootGin-8                          33904330                34.81 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Gin-8                             34333412                34.67 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5Gin-8                             34166740                34.69 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10Gin-8                            33355299                35.71 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Gin-8                     29163256                40.77 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5Gin-8                     15656606                77.37 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10Gin-8                     9627939               129.1 ns/op         0 B/op          0 allocs/op
BenchmarkStaticRoutesRootBunRouter-8                    61770626                22.37 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1BunRouter-8                       54016642                25.25 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5BunRouter-8                       48271202                25.21 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10BunRouter-8                      54948931                24.84 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1BunRouter-8               35822632                34.42 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5BunRouter-8                8529741               139.6 ns/op         0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10BunRouter-8               4199608               283.7 ns/op         0 B/op          0 allocs/op
BenchmarkStaticRoutesRootHTTPTreeMux-8                   6656719               174.2 ns/op       328 B/op          3 allocs/op
BenchmarkStaticRoutes1HTTPTreeMux-8                      6584396               180.9 ns/op       328 B/op          3 allocs/op
BenchmarkStaticRoutes5HTTPTreeMux-8                      5399803               219.8 ns/op       328 B/op          3 allocs/op
BenchmarkStaticRoutes10HTTPTreeMux-8                     4297136               279.5 ns/op       328 B/op          3 allocs/op
BenchmarkPathParamColonRoutes1HTTPTreeMux-8              3125630               379.6 ns/op       680 B/op          6 allocs/op
BenchmarkPathParamColonRoutes5HTTPTreeMux-8              1569038               757.9 ns/op       904 B/op          9 allocs/op
BenchmarkPathParamColonRoutes10HTTPTreeMux-8              794005              1448 ns/op        1742 B/op         11 allocs/op
BenchmarkStaticRoutesRootBeegoMux-8                     22699148                52.43 ns/op       32 B/op          1 allocs/op
BenchmarkStaticRoutes1BeegoMux-8                        16063226                74.21 ns/op       32 B/op          1 allocs/op
BenchmarkStaticRoutes5BeegoMux-8                         1000000              1074 ns/op          32 B/op          1 allocs/op
BenchmarkStaticRoutes10BeegoMux-8                         591600              2044 ns/op          32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes1BeegoMux-8                 3240128               370.9 ns/op       672 B/op          5 allocs/op
BenchmarkPathParamColonRoutes5BeegoMux-8                  829568              1424 ns/op         672 B/op          5 allocs/op
BenchmarkPathParamColonRoutes10BeegoMux-8                 353080              3380 ns/op        1254 B/op          6 allocs/op
BenchmarkStaticRoutesRootGorillaMux-8                    2142933               547.4 ns/op       720 B/op          7 allocs/op
BenchmarkStaticRoutes1GorillaMux-8                       2158372               552.3 ns/op       720 B/op          7 allocs/op
BenchmarkStaticRoutes5GorillaMux-8                       2043475               596.7 ns/op       720 B/op          7 allocs/op
BenchmarkStaticRoutes10GorillaMux-8                      1882173               648.0 ns/op       720 B/op          7 allocs/op
BenchmarkPathParamBracketRoutes1GorillaMux-8             1384304               878.9 ns/op      1024 B/op          8 allocs/op
BenchmarkPathParamBracketRoutes5GorillaMux-8              508077              2309 ns/op        1088 B/op          8 allocs/op
BenchmarkPathParamBracketRoutes10GorillaMux-8             203055              5198 ns/op        1751 B/op          9 allocs/op
BenchmarkStaticRoutesRootBon-8                          70143076                17.09 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Bon-8                             64823545                17.74 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5Bon-8                             65827966                17.28 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10Bon-8                            65070090                17.91 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Bon-8                      6431658               185.0 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamColonRoutes5Bon-8                      4647512               254.1 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamColonRoutes10Bon-8                     3424208               355.0 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutesRootDenco-8                        80690576                14.43 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Denco-8                           81357084                14.40 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5Denco-8                           81186217                14.42 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10Denco-8                          74032372                15.41 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Denco-8                   19499986                59.84 ns/op       32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes5Denco-8                    8739303               136.3 ns/op       160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes10Denco-8                   5138895               232.4 ns/op       320 B/op          1 allocs/op
BenchmarkStaticRoutesRootEcho-8                         41008466                29.68 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Echo-8                            37285414                31.98 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5Echo-8                            23113537                51.02 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10Echo-8                           13364814                91.95 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Echo-8                    30489030                38.85 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5Echo-8                    11871279                98.74 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10Echo-8                    6792343               184.3 ns/op         0 B/op          0 allocs/op
PASS
ok      github.com/go-router-benchmark  107.344s
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