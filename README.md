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
BenchmarkStaticRoutesRootGoblin-8                       34961384                31.50 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Goblin-8                          18314901                65.87 ns/op       16 B/op          1 allocs/op
BenchmarkStaticRoutes5Goblin-8                           6131110               198.9 ns/op        80 B/op          1 allocs/op
BenchmarkStaticRoutes10Goblin-8                          3216090               371.2 ns/op       160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes1Goblin-8                   1910696               623.6 ns/op       407 B/op          6 allocs/op
BenchmarkPathParamColonRoutes5Goblin-8                    525855              2219 ns/op         962 B/op         13 allocs/op
BenchmarkPathParamColonRoutes10Goblin-8                   279876              4271 ns/op        1606 B/op         19 allocs/op
BenchmarkStaticRoutesRootHTTPRouter-8                   89100366                12.02 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1HTTPRouter-8                      88427376                12.10 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5HTTPRouter-8                      88872706                12.13 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10HTTPRouter-8                     89740317                12.14 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1HTTPRouter-8              27192128                43.05 ns/op       32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes5HTTPRouter-8              10607023               114.6 ns/op       160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes10HTTPRouter-8              6149580               194.0 ns/op       320 B/op          1 allocs/op
BenchmarkStaticRoutesRootChi-8                           5561600               215.9 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutes1Chi-8                              5583362               210.8 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutes5Chi-8                              5828372               204.2 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutes10Chi-8                             5775816               204.1 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes1Chi-8                    4622605               261.1 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes5Chi-8                    2838205               419.0 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes10Chi-8                   1939826               618.1 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutesRootGin-8                          33886138                35.53 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Gin-8                             33251511                35.42 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5Gin-8                             33898982                35.57 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10Gin-8                            32448765                37.04 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Gin-8                     29109050                41.75 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5Gin-8                     14857444                77.65 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10Gin-8                     9559023               124.5 ns/op         0 B/op          0 allocs/op
BenchmarkStaticRoutesRootBunRouter-8                    65604094                19.51 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1BunRouter-8                       56080006                21.77 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5BunRouter-8                       56156770                21.64 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10BunRouter-8                      54875016                21.80 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1BunRouter-8               36486039                32.79 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5BunRouter-8                8551426               140.7 ns/op         0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10BunRouter-8               4207790               285.3 ns/op         0 B/op          0 allocs/op
BenchmarkStaticRoutesRootHTTPTreeMux-8                   6720936               182.7 ns/op       328 B/op          3 allocs/op
BenchmarkStaticRoutes1HTTPTreeMux-8                      5723260               191.5 ns/op       328 B/op          3 allocs/op
BenchmarkStaticRoutes5HTTPTreeMux-8                      5181651               244.7 ns/op       328 B/op          3 allocs/op
BenchmarkStaticRoutes10HTTPTreeMux-8                     4032684               298.3 ns/op       328 B/op          3 allocs/op
BenchmarkPathParamColonRoutes1HTTPTreeMux-8              3069298               389.2 ns/op       680 B/op          6 allocs/op
BenchmarkPathParamColonRoutes5HTTPTreeMux-8              1563354               764.9 ns/op       904 B/op          9 allocs/op
BenchmarkPathParamColonRoutes10HTTPTreeMux-8              766495              1481 ns/op        1742 B/op         11 allocs/op
BenchmarkStaticRoutesRootBeegoMux-8                     23192737                50.00 ns/op       32 B/op          1 allocs/op
BenchmarkStaticRoutes1BeegoMux-8                        16262485                73.00 ns/op       32 B/op          1 allocs/op
BenchmarkStaticRoutes5BeegoMux-8                         1000000              1107 ns/op          32 B/op          1 allocs/op
BenchmarkStaticRoutes10BeegoMux-8                         569458              2099 ns/op          32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes1BeegoMux-8                 3176863               375.3 ns/op       672 B/op          5 allocs/op
BenchmarkPathParamColonRoutes5BeegoMux-8                  789951              1436 ns/op         672 B/op          5 allocs/op
BenchmarkPathParamColonRoutes10BeegoMux-8                 355358              3308 ns/op        1254 B/op          6 allocs/op
BenchmarkStaticRoutesRootGorillaMux-8                    2150677               555.1 ns/op       720 B/op          7 allocs/op
BenchmarkStaticRoutes1GorillaMux-8                       2152135               556.5 ns/op       720 B/op          7 allocs/op
BenchmarkStaticRoutes5GorillaMux-8                       2052021               584.8 ns/op       720 B/op          7 allocs/op
BenchmarkStaticRoutes10GorillaMux-8                      1899133               631.0 ns/op       720 B/op          7 allocs/op
BenchmarkPathParamBracketRoutes1GorillaMux-8             1361359               890.2 ns/op      1024 B/op          8 allocs/op
BenchmarkPathParamBracketRoutes5GorillaMux-8              513170              2272 ns/op        1088 B/op          8 allocs/op
BenchmarkPathParamBracketRoutes10GorillaMux-8             237112              5039 ns/op        1751 B/op          9 allocs/op
BenchmarkStaticRoutesRootBon-8                          85570648                14.30 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Bon-8                             84627976                14.34 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5Bon-8                             83464470                14.25 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10Bon-8                            80067390                16.06 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Bon-8                      5898654               202.7 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamColonRoutes5Bon-8                      4388452               254.6 ns/op       304 B/op          2 allocs/op
BenchmarkPathParamColonRoutes10Bon-8                     3445801               349.6 ns/op       304 B/op          2 allocs/op
BenchmarkStaticRoutesRootDenco-8                        88671046                13.49 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes1Denco-8                           88327569                13.55 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes5Denco-8                           87896529                14.50 ns/op        0 B/op          0 allocs/op
BenchmarkStaticRoutes10Denco-8                          87464620                13.56 ns/op        0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Denco-8                   20827050                56.15 ns/op       32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes5Denco-8                    8694339               135.1 ns/op       160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes10Denco-8                   5201998               259.7 ns/op       320 B/op          1 allocs/op
PASS
ok      github.com/go-router-benchmark  98.065s
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