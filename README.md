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
BenchmarkStaticRoutesRootGoblin-8                       36163220                32.89 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Goblin-8                          15753888                75.84 ns/op       16 B/op          1 allocs/op
BenchmarkStaticRoutes5Goblin-8                           4802425               249.6 ns/op        80 B/op          1 allocs/op
BenchmarkStaticRoutes10Goblin-8                          2363377               502.4 ns/op       160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes1Goblin-8                   1893566               630.8 ns/op       408 B/op          6 allocs/op
BenchmarkPathParamColonRoutes5Goblin-8                    504063              2443 ns/op         963 B/op         13 allocs/op
BenchmarkPathParamColonRoutes10Goblin-8                   278032              4380 ns/op        1605 B/op         19 allocs/op
BenchmarkStaticRoutesRootHTTPRouter-8                   100000000               10.87 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1HTTPRouter-8                      100000000               10.88 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5HTTPRouter-8                      100000000               10.85 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10HTTPRouter-8                     100000000               10.89 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1HTTPRouter-8              26945830                44.09 ns/op       32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes5HTTPRouter-8              10610876               112.6 ns/op       160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes10HTTPRouter-8              6222957               200.6 ns/op       320 B/op          1 allocs/op
BenchmarkStaticRoutesRootChi-8                           5477976               212.6 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutes1Chi-8                              5719608               210.5 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutes5Chi-8                              5486629               212.1 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutes10Chi-8                             5801430               208.7 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes1Chi-8                    4544932               263.9 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes5Chi-8                    2791695               433.2 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes10Chi-8                   1913809               624.4 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutesRootGin-8                          34209475                34.82 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Gin-8                             34483542                34.57 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5Gin-8                             34404409                34.87 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10Gin-8                            33386659                35.71 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Gin-8                     29252803                40.99 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5Gin-8                     14913160                76.89 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10Gin-8                     9699278               123.6 ns/op         0 B/op          0 allocs/op
BenchmarkStaticRoutesRootBunRouter-8                    64171836                18.41 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1BunRouter-8                       57424509                21.10 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5BunRouter-8                       55781737                21.16 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10BunRouter-8                      55802922                21.17 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1BunRouter-8               35038243                33.47 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5BunRouter-8                8442925               141.7 ns/op         0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10BunRouter-8               4200292               285.7 ns/op         0 B/op          0 allocs/op
BenchmarkStaticRoutesRootHTTPTreeMux-8                   6899055               174.8 ns/op       328 B/op          3 allocs/op
BenchmarkStaticRoutes1HTTPTreeMux-8                      6769173               177.9 ns/op       328 B/op          3 allocs/op
BenchmarkStaticRoutes5HTTPTreeMux-8                      5450172               217.9 ns/op       328 B/op          3 allocs/op
BenchmarkStaticRoutes10HTTPTreeMux-8                     3912093               281.6 ns/op       328 B/op          3 allocs/op
BenchmarkPathParamColonRoutes1HTTPTreeMux-8              3217290               375.3 ns/op       680 B/op          6 allocs/op
BenchmarkPathParamColonRoutes5HTTPTreeMux-8              1585006               745.8 ns/op       904 B/op          9 allocs/op
BenchmarkPathParamColonRoutes10HTTPTreeMux-8              690555              1471 ns/op        1742 B/op         11 allocs/op
BenchmarkStaticRoutesRootBeegoMux-8                     24045706                49.55 ns/op       32 B/op          1 allocs/op
BenchmarkStaticRoutes1BeegoMux-8                        16360400                72.44 ns/op       32 B/op          1 allocs/op
BenchmarkStaticRoutes5BeegoMux-8                         1000000              1083 ns/op          32 B/op          1 allocs/op
BenchmarkStaticRoutes10BeegoMux-8                         586821              2058 ns/op          32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes1BeegoMux-8                 3270447               368.8 ns/op       672 B/op          5 allocs/op
BenchmarkPathParamColonRoutes5BeegoMux-8                  833185              1446 ns/op         672 B/op          5 allocs/op
BenchmarkPathParamColonRoutes10BeegoMux-8                 358380              3327 ns/op        1254 B/op          6 allocs/op
BenchmarkStaticRoutesRootGorillaMux-8                    2197866               545.9 ns/op       720 B/op          7 allocs/op
BenchmarkStaticRoutes1GorillaMux-8                       2179435               551.1 ns/op       720 B/op          7 allocs/op
BenchmarkStaticRoutes5GorillaMux-8                       2054035               578.6 ns/op       720 B/op          7 allocs/op
BenchmarkStaticRoutes10GorillaMux-8                      1896001               628.9 ns/op       720 B/op          7 allocs/op
BenchmarkPathParamBracketRoutes1GorillaMux-8             1402426               866.3 ns/op      1024 B/op          8 allocs/op
BenchmarkPathParamBracketRoutes5GorillaMux-8              547293              2171 ns/op        1088 B/op          8 allocs/op
BenchmarkPathParamBracketRoutes10GorillaMux-8             249920              4768 ns/op        1751 B/op          9 allocs/op
BenchmarkStaticRoutesRootBon-8                          79469324                13.78 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Bon-8                             78999559                13.87 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5Bon-8                             79557568                13.77 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10Bon-8                            74634987                14.70 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Bon-8                      6622338               181.8 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamColonRoutes5Bon-8                      4711281               257.6 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamColonRoutes10Bon-8                     3433544               348.8 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutesRootDenco-8                        88291287                13.58 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Denco-8                           88792144                14.79 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5Denco-8                           86751950                13.97 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10Denco-8                          79622017                14.10 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Denco-8                   18787734                61.45 ns/op       32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes5Denco-8                    8833496               140.1 ns/op       160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes10Denco-8                   4776807               236.4 ns/op       320 B/op          1 allocs/op
BenchmarkStaticRoutesRootEcho-8                         32899846                35.84 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Echo-8                            31087545                38.42 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5Echo-8                            20495331                56.90 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10Echo-8                           12564380                95.05 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Echo-8                    25496359                46.17 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5Echo-8                    11158305               106.9 ns/op         0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10Echo-8                    6530382               184.6 ns/op         0 B/op          0 allocs/op
BenchmarkStaticRoutesRootGocraftWeb-8                    1264578               928.7 ns/op       288 B/op          6 allocs/op
BenchmarkStaticRoutes1GocraftWeb-8                       1000000              1046 ns/op         288 B/op          6 allocs/op
BenchmarkStaticRoutes5GocraftWeb-8                        953575              1158 ns/op         352 B/op          6 allocs/op
BenchmarkStaticRoutes10GocraftWeb-8                       943454              1265 ns/op         432 B/op          6 allocs/op
BenchmarkPathParamColonRoutes1GocraftWeb-8                970850              1155 ns/op         656 B/op          9 allocs/op
BenchmarkPathParamColonRoutes5GocraftWeb-8                808851              1551 ns/op         944 B/op         12 allocs/op
BenchmarkPathParamColonRoutes10GocraftWeb-8               547400              2179 ns/op        1862 B/op         14 allocs/op
PASS
ok      github.com/go-router-benchmark  114.608s
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