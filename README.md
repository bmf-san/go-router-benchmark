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
- [gocraft/web](https://github.com/gocraft/web)
- [vardius/gorouter](https://github.com/vardius/gorouter)
- [go-ozzo/ozzo-routing](https://github.com/go-ozzo/ozzo-routing)

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
BenchmarkStaticRoutesRootGoblin-8                               37810608                31.62 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes1Goblin-8                                  17207332                68.55 ns/op             16 B/op          1 allocs/op
BenchmarkStaticRoutes5Goblin-8                                   5786544               208.3 ns/op      80 B/op          1 allocs/op
BenchmarkStaticRoutes10Goblin-8                                  2992412               400.3 ns/op     160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes1Goblin-8                           1907245               627.5 ns/op     409 B/op          6 allocs/op
BenchmarkPathParamColonRoutes5Goblin-8                            525704              2269 ns/op      965 B/op          13 allocs/op
BenchmarkPathParamColonRoutes10Goblin-8                           275090              4399 ns/op     1608 B/op          19 allocs/op
BenchmarkStaticRoutesRootHTTPRouter-8                           100000000               10.46 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes1HTTPRouter-8                              100000000               10.33 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes5HTTPRouter-8                              100000000               10.44 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes10HTTPRouter-8                             100000000               10.36 ns/op              0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1HTTPRouter-8                      27414811                43.20 ns/op             32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes5HTTPRouter-8                      10295874               114.4 ns/op     160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes10HTTPRouter-8                      6092434               197.3 ns/op     320 B/op          1 allocs/op
BenchmarkStaticRoutesRootChi-8                                   5429904               220.0 ns/op     304 B/op          2 allocs/op
BenchmarkStaticRoutes1Chi-8                                      5290057               220.6 ns/op     304 B/op          2 allocs/op
BenchmarkStaticRoutes5Chi-8                                      5434046               219.3 ns/op     304 B/op          2 allocs/op
BenchmarkStaticRoutes10Chi-8                                     5505698               212.7 ns/op     304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes1Chi-8                            4445071               268.9 ns/op     304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes5Chi-8                            2730680               437.0 ns/op     304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes10Chi-8                           1888065               636.3 ns/op     304 B/op          2 allocs/op
BenchmarkStaticRoutesRootGin-8                                  32862794                36.33 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes1Gin-8                                     33070339                36.07 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes5Gin-8                                     32586816                36.17 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes10Gin-8                                    32143107                37.20 ns/op              0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Gin-8                             27500701                42.62 ns/op              0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5Gin-8                             15258900                78.41 ns/op              0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10Gin-8                             9522596               127.7 ns/op       0 B/op          0 allocs/op
BenchmarkStaticRoutesRootBunRouter-8                            64215045                18.42 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes1BunRouter-8                               56828720                21.02 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes5BunRouter-8                               56755030                20.99 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes10BunRouter-8                              56908671                20.84 ns/op              0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1BunRouter-8                       35162946                33.80 ns/op              0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5BunRouter-8                        8425648               140.7 ns/op       0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10BunRouter-8                       4257427               283.7 ns/op       0 B/op          0 allocs/op
BenchmarkStaticRoutesRootHTTPTreeMux-8                           6690556               179.2 ns/op     328 B/op          3 allocs/op
BenchmarkStaticRoutes1HTTPTreeMux-8                              6562962               183.0 ns/op     328 B/op          3 allocs/op
BenchmarkStaticRoutes5HTTPTreeMux-8                              5362312               224.5 ns/op     328 B/op          3 allocs/op
BenchmarkStaticRoutes10HTTPTreeMux-8                             4219711               284.1 ns/op     328 B/op          3 allocs/op
BenchmarkPathParamColonRoutes1HTTPTreeMux-8                      3174328               377.0 ns/op     680 B/op          6 allocs/op
BenchmarkPathParamColonRoutes5HTTPTreeMux-8                      1545889               757.7 ns/op     904 B/op          9 allocs/op
BenchmarkPathParamColonRoutes10HTTPTreeMux-8                      807580              1474 ns/op     1742 B/op          11 allocs/op
BenchmarkStaticRoutesRootBeegoMux-8                             22258030                52.52 ns/op             32 B/op          1 allocs/op
BenchmarkStaticRoutes1BeegoMux-8                                15893838                75.32 ns/op             32 B/op          1 allocs/op
BenchmarkStaticRoutes5BeegoMux-8                                 1000000              1071 ns/op       32 B/op           1 allocs/op
BenchmarkStaticRoutes10BeegoMux-8                                 592300              2029 ns/op       32 B/op           1 allocs/op
BenchmarkPathParamColonRoutes1BeegoMux-8                         3219261               373.9 ns/op     672 B/op          5 allocs/op
BenchmarkPathParamColonRoutes5BeegoMux-8                          824432              1420 ns/op      672 B/op           5 allocs/op
BenchmarkPathParamColonRoutes10BeegoMux-8                         354176              3350 ns/op     1254 B/op           6 allocs/op
BenchmarkStaticRoutesRootGorillaMux-8                            2076849               574.2 ns/op     720 B/op          7 allocs/op
BenchmarkStaticRoutes1GorillaMux-8                               2066920               579.6 ns/op     720 B/op          7 allocs/op
BenchmarkStaticRoutes5GorillaMux-8                               1967010               610.1 ns/op     720 B/op          7 allocs/op
BenchmarkStaticRoutes10GorillaMux-8                              1824106               658.2 ns/op     720 B/op          7 allocs/op
BenchmarkPathParamBracketRoutes1GorillaMux-8                     1339666               895.5 ns/op    1024 B/op          8 allocs/op
BenchmarkPathParamBracketRoutes5GorillaMux-8                      526460              2262 ns/op     1088 B/op           8 allocs/op
BenchmarkPathParamBracketRoutes10GorillaMux-8                     238946              5008 ns/op     1751 B/op           9 allocs/op
BenchmarkStaticRoutesRootBon-8                                  88309966                13.51 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes1Bon-8                                     87832992                13.51 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes5Bon-8                                     86620446                13.48 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes10Bon-8                                    87585646                14.03 ns/op              0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Bon-8                              6392168               186.3 ns/op     304 B/op          2 allocs/op
BenchmarkPathParamColonRoutes5Bon-8                              4628920               259.3 ns/op     304 B/op          2 allocs/op
BenchmarkPathParamColonRoutes10Bon-8                             3381514               355.9 ns/op     304 B/op          2 allocs/op
BenchmarkStaticRoutesRootDenco-8                                79634784                14.46 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes1Denco-8                                   79188319                14.38 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes5Denco-8                                   79437754                14.36 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes10Denco-8                                  74332171                15.32 ns/op              0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Denco-8                           19618006                60.16 ns/op             32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes5Denco-8                            8545141               140.5 ns/op     160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes10Denco-8                           5004351               239.7 ns/op     320 B/op          1 allocs/op
BenchmarkStaticRoutesRootEcho-8                                 33267644                35.72 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes1Echo-8                                    30976101                38.24 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes5Echo-8                                    21207642                56.36 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes10Echo-8                                   12649693                94.39 ns/op              0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Echo-8                            26137147                45.25 ns/op              0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5Echo-8                            11194768               106.2 ns/op       0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10Echo-8                            6530691               183.1 ns/op       0 B/op          0 allocs/op
BenchmarkStaticRoutesRootGocraftWeb-8                            1276323               934.8 ns/op     288 B/op          6 allocs/op
BenchmarkStaticRoutes1GocraftWeb-8                               1246320               949.0 ns/op     288 B/op          6 allocs/op
BenchmarkStaticRoutes5GocraftWeb-8                               1000000              1075 ns/op      352 B/op           6 allocs/op
BenchmarkStaticRoutes10GocraftWeb-8                               967153              1210 ns/op      432 B/op           6 allocs/op
BenchmarkPathParamColonRoutes1GocraftWeb-8                       1000000              1154 ns/op      656 B/op           9 allocs/op
BenchmarkPathParamColonRoutes5GocraftWeb-8                        794268              1542 ns/op      944 B/op          12 allocs/op
BenchmarkPathParamColonRoutes10GocraftWeb-8                       519505              2243 ns/op     1862 B/op          14 allocs/op
BenchmarkStaticRoutesRootGorouter-8                             41112081                29.07 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes1Gorouter-8                                33128236                36.39 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes5Gorouter-8                                18102231                66.23 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes10Gorouter-8                               10874481               110.1 ns/op       0 B/op          0 allocs/op
BenchmarkPathParamBracketRoutes1Gorouter-8                       4731598               253.9 ns/op     360 B/op          4 allocs/op
BenchmarkPathParamBracketRoutes5Gorouter-8                       2986797               401.4 ns/op     488 B/op          4 allocs/op
BenchmarkPathParamBracketRoutes10Gorouter-8                      2219754               539.8 ns/op     648 B/op          4 allocs/op
BenchmarkStaticRoutesRootOzzoRouting-8                          36044740                32.67 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes1OzzoRouting-8                             34365777                34.35 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes5OzzoRouting-8                             25027525                47.43 ns/op              0 B/op          0 allocs/op
BenchmarkStaticRoutes10OzzoRouting-8                            18702718                63.80 ns/op              0 B/op          0 allocs/op
BenchmarkPathParamInequalitySignRoutes1OzzoRouting-8            27764709                42.93 ns/op              0 B/op          0 allocs/op
BenchmarkPathParamInequalitySignRoutes5OzzoRouting-8            13905720                86.12 ns/op              0 B/op          0 allocs/op
BenchmarkPathParamInequalitySignRoutes10OzzoRouting-8            8533772               140.4 ns/op       0 B/op          0 allocs/op
PASS
ok      github.com/go-router-benchmark  134.328s
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
