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
BenchmarkStaticRoutesRootGoblin-8                       38528376                30.89 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Goblin-8                          15491332                68.47 ns/op       16 B/op          1 allocs/op
BenchmarkStaticRoutes5Goblin-8                           5684391               208.1 ns/op        80 B/op          1 allocs/op
BenchmarkStaticRoutes10Goblin-8                          3004182               382.0 ns/op       160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes1Goblin-8                   1936789               627.4 ns/op       408 B/op          6 allocs/op
BenchmarkPathParamColonRoutes5Goblin-8                    517453              2279 ns/op         962 B/op         13 allocs/op
BenchmarkPathParamColonRoutes10Goblin-8                   280765              4308 ns/op        1610 B/op         19 allocs/op
BenchmarkStaticRoutesRootHTTPRouter-8                   100000000               10.84 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1HTTPRouter-8                      100000000               10.91 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5HTTPRouter-8                      100000000               10.81 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10HTTPRouter-8                     100000000               10.85 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1HTTPRouter-8              27327199                43.68 ns/op       32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes5HTTPRouter-8              10117357               116.5 ns/op       160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes10HTTPRouter-8              5942580               197.1 ns/op       320 B/op          1 allocs/op
BenchmarkStaticRoutesRootChi-8                           5652294               211.9 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutes1Chi-8                              5520016               234.4 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutes5Chi-8                              5517648               210.7 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutes10Chi-8                             5701597               210.4 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes1Chi-8                    3728184               275.8 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes5Chi-8                    2747817               459.0 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes10Chi-8                   1906172               653.1 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutesRootGin-8                          33641361                37.05 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Gin-8                             28879765                42.11 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5Gin-8                             28876406                39.51 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10Gin-8                            32357768                38.72 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Gin-8                     28526402                42.37 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5Gin-8                     15369319                77.93 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10Gin-8                     9566335               125.9 ns/op         0 B/op          0 allocs/op
BenchmarkStaticRoutesRootBunRouter-8                    63883444                18.59 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1BunRouter-8                       56544206                21.27 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5BunRouter-8                       55984617                21.21 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10BunRouter-8                      54595706                21.46 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1BunRouter-8               34981257                34.66 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5BunRouter-8                8501227               141.1 ns/op         0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10BunRouter-8               4157802               286.1 ns/op         0 B/op          0 allocs/op
BenchmarkStaticRoutesRootHTTPTreeMux-8                   6705519               182.0 ns/op       328 B/op          3 allocs/op
BenchmarkStaticRoutes1HTTPTreeMux-8                      6616117               193.0 ns/op       328 B/op          3 allocs/op
BenchmarkStaticRoutes5HTTPTreeMux-8                      5051487               230.9 ns/op       328 B/op          3 allocs/op
BenchmarkStaticRoutes10HTTPTreeMux-8                     4081179               294.3 ns/op       328 B/op          3 allocs/op
BenchmarkPathParamColonRoutes1HTTPTreeMux-8              3141382               391.8 ns/op       680 B/op          6 allocs/op
BenchmarkPathParamColonRoutes5HTTPTreeMux-8              1567244               769.4 ns/op       904 B/op          9 allocs/op
BenchmarkPathParamColonRoutes10HTTPTreeMux-8              798372              1487 ns/op        1742 B/op         11 allocs/op
BenchmarkStaticRoutesRootBeegoMux-8                     23698027                50.38 ns/op       32 B/op          1 allocs/op
BenchmarkStaticRoutes1BeegoMux-8                        16610826                72.54 ns/op       32 B/op          1 allocs/op
BenchmarkStaticRoutes5BeegoMux-8                         1000000              1075 ns/op          32 B/op          1 allocs/op
BenchmarkStaticRoutes10BeegoMux-8                         580822              2031 ns/op          32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes1BeegoMux-8                 3242056               371.4 ns/op       672 B/op          5 allocs/op
BenchmarkPathParamColonRoutes5BeegoMux-8                  804933              1428 ns/op         672 B/op          5 allocs/op
BenchmarkPathParamColonRoutes10BeegoMux-8                 355087              3402 ns/op        1254 B/op          6 allocs/op
BenchmarkStaticRoutesRootGorillaMux-8                    2029521               598.0 ns/op           720 B/op            7 allocs/op
BenchmarkStaticRoutes1GorillaMux-8                       2039371               578.6 ns/op           720 B/op            7 allocs/op
BenchmarkStaticRoutes5GorillaMux-8                       1978711               608.3 ns/op           720 B/op            7 allocs/op
BenchmarkStaticRoutes10GorillaMux-8                      1841149               656.1 ns/op           720 B/op            7 allocs/op
BenchmarkPathParamBracketRoutes1GorillaMux-8             1358829               888.5 ns/op          1024 B/op            8 allocs/op
BenchmarkPathParamBracketRoutes5GorillaMux-8              505826              3040 ns/op            1088 B/op            8 allocs/op
BenchmarkPathParamBracketRoutes10GorillaMux-8             202527              7443 ns/op            1750 B/op            9 allocs/op
BenchmarkStaticRoutesRootBon-8                          76577618                17.33 ns/op            0 B/op            0 allocs/op
BenchmarkStaticRoutes1Bon-8                             68924748                17.90 ns/op            0 B/op            0 allocs/op
BenchmarkStaticRoutes5Bon-8                             73622643                19.08 ns/op            0 B/op            0 allocs/op
BenchmarkStaticRoutes10Bon-8                            60186073                18.00 ns/op            0 B/op            0 allocs/op
BenchmarkPathParamColonRoutes1Bon-8                      6375943               189.7 ns/op           304 B/op            2 allocs/op
BenchmarkPathParamColonRoutes5Bon-8                      4367514               273.7 ns/op           304 B/op            2 allocs/op
BenchmarkPathParamColonRoutes10Bon-8                     3294938               359.5 ns/op           304 B/op            2 allocs/op
BenchmarkStaticRoutesRootDenco-8                        85436871                13.50 ns/op            0 B/op            0 allocs/op
BenchmarkStaticRoutes1Denco-8                           83122398                13.61 ns/op            0 B/op            0 allocs/op
BenchmarkStaticRoutes5Denco-8                           87743889                13.61 ns/op            0 B/op            0 allocs/op
BenchmarkStaticRoutes10Denco-8                          84731299                14.53 ns/op            0 B/op            0 allocs/op
BenchmarkPathParamColonRoutes1Denco-8                   19951202                60.73 ns/op           32 B/op            1 allocs/op
BenchmarkPathParamColonRoutes5Denco-8                    8575537               139.2 ns/op           160 B/op            1 allocs/op
BenchmarkPathParamColonRoutes10Denco-8                   4918402               237.5 ns/op           320 B/op            1 allocs/op
BenchmarkStaticRoutesRootEcho-8                         39526016                29.69 ns/op            0 B/op            0 allocs/op
BenchmarkStaticRoutes1Echo-8                            36468574                32.15 ns/op            0 B/op            0 allocs/op
BenchmarkStaticRoutes5Echo-8                            23207970                52.20 ns/op            0 B/op            0 allocs/op
BenchmarkStaticRoutes10Echo-8                           13377528                90.66 ns/op            0 B/op            0 allocs/op
BenchmarkPathParamColonRoutes1Echo-8                    28206118                43.29 ns/op            0 B/op            0 allocs/op
BenchmarkPathParamColonRoutes5Echo-8                    10763610               110.0 ns/op             0 B/op            0 allocs/op
BenchmarkPathParamColonRoutes10Echo-8                    6026521               197.8 ns/op             0 B/op            0 allocs/op
BenchmarkStaticRoutesRootGocraftWeb-8                    1228579               938.3 ns/op           288 B/op            6 allocs/op
BenchmarkStaticRoutes1GocraftWeb-8                       1000000              1002 ns/op             288 B/op            6 allocs/op
BenchmarkStaticRoutes5GocraftWeb-8                       1000000              1104 ns/op             352 B/op            6 allocs/op
BenchmarkStaticRoutes10GocraftWeb-8                       953325              1268 ns/op             432 B/op            6 allocs/op
BenchmarkPathParamColonRoutes1GocraftWeb-8                930950              1177 ns/op             656 B/op            9 allocs/op
BenchmarkPathParamColonRoutes5GocraftWeb-8                798331              1538 ns/op             944 B/op           12 allocs/op
BenchmarkPathParamColonRoutes10GocraftWeb-8               515178              2408 ns/op            1862 B/op           14 allocs/op
BenchmarkStaticRoutesRootGorouter-8                     27833481                42.80 ns/op            0 B/op            0 allocs/op
BenchmarkStaticRoutes1Gorouter-8                        31875154                37.92 ns/op            0 B/op            0 allocs/op
BenchmarkStaticRoutes5Gorouter-8                        17747660                66.06 ns/op            0 B/op            0 allocs/op
BenchmarkStaticRoutes10Gorouter-8                       11151841               110.5 ns/op             0 B/op            0 allocs/op
BenchmarkPathParamBracketRoutes1Gorouter-8               4313761               295.1 ns/op           360 B/op            4 allocs/op
BenchmarkPathParamBracketRoutes5Gorouter-8               2973651               430.5 ns/op           488 B/op            4 allocs/op
BenchmarkPathParamBracketRoutes10Gorouter-8              2228737               536.9 ns/op           648 B/op            4 allocs/op
PASS
ok      github.com/go-router-benchmark  125.479s
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