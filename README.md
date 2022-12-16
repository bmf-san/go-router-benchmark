# go-router-benchmark
Compare the performance of routers built with golang.

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

Since [net/http#ServeMux](https://pkg.go.dev/net/http#ServeMux) does not have the capability to support path param route, only the static route test case is comparable.

# Motivation
I have implemented a router called [bmf-san/goblin](https://github.com/bmf-san/goblin), and I created this repository to compare it with other routers and get hints on how to improve [bmf-san/goblin](https://github.com/bmf-san/goblin).

Another reason is that [julienschmidt/go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) seems to have stopped being maintained, so I wanted to have a benchmarker for the router that I could manage myself.

# Benchmark test
This benchmark tests is not a perfect comparison of HTTP Router performance.
The reasons are as follows

- Not practical to cover all test cases because each HTTP Router has different specifications
- Performance may be unfairly evaluated depending on the routing test cases defined, since each HTTP Router has its own strengths and weaknesses depending on its data structures and algorithms.

Although the benchmark test is based on a specific case, it is possible to see some trends in performance differences.

Performance measurements will be made on the routing process of the HTTP Router.
More specifically, we will test the `ServeHTTP` function of [http#Handler](https://pkg.go.dev/net/http#Handler).

We do not measure the performance of the process that defines the routing of the HTTP Router.
The process of defining routing is the process of registering data for the routing process.
For example, the code for net/http is as follows.

```go
package main

import (
	"fmt"
	"net/http" )
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler) // here
	ListenAndServe(":8080", mux)
}
```

We believe that the part that handles routing is called more often than the part that defines routing, and thus accounts for the majority of the HTTP Router's performance.

## Static route
A static route is a route without variable parameters such as `/foo/bar`.

Static route measures performance by benchmarking routing tests with the following three input values.

- `/`
- `/foo`
- `/foo/bar/baz/qux/quux`
- `/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred`

## Path parameter route
A path parameter route is a route with variable parameters such as `/foo/:bar`.

In path parameter route, we perform benchmark routing tests with the following three input values to measure the performance.

- `/foo/:bar`
- `/foo/:bar/:baz/:qux/:quux/:corge`
- `/foo/:bar/:baz/:qux/:quux/:corge/:grault/:garply/:waldo/:fred/:plugh`

Since different HTTP Routers have different ways of expressing parameters, there are several cases where another symbol is used in addition to `:`.

# How to run benchmark test
`make test-benchmark`

# Results
```sh
go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/go-router-benchmark
cpu: VirtualApple @ 2.50GHz
BenchmarkStaticRoutesRootServeMux-8                             23305570                50.96 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes1ServeMux-8                                21346630                56.29 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes5ServeMux-8                                13507384                88.60 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes10ServeMux-8                                8871856               135.2 ns/op             0 B/op          0 allocs/op
BenchmarkStaticRoutesRootGoblin-8                               31467156                37.69 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes1Goblin-8                                  16274073                71.43 ns/op           16 B/op          1 allocs/op
BenchmarkStaticRoutes5Goblin-8                                   5900212               203.9 ns/op            80 B/op          1 allocs/op
BenchmarkStaticRoutes10Goblin-8                                  3205656               374.0 ns/op           160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes1Goblin-8                           1936185               619.5 ns/op           408 B/op          6 allocs/op
BenchmarkPathParamColonRoutes5Goblin-8                            520683              2237 ns/op             964 B/op         13 allocs/op
BenchmarkPathParamColonRoutes10Goblin-8                           276434              4292 ns/op            1611 B/op         19 allocs/op
BenchmarkStaticRoutesRootHTTPRouter-8                           79534504                13.29 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes1HTTPRouter-8                              79775299                13.40 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes5HTTPRouter-8                              79895688                13.38 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes10HTTPRouter-8                             80145372                13.40 ns/op            0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1HTTPRouter-8                      24690871                46.86 ns/op           32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes5HTTPRouter-8                       9880882               119.9 ns/op           160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes10HTTPRouter-8                      5949777               200.5 ns/op           320 B/op          1 allocs/op
BenchmarkStaticRoutesRootChi-8                                   5480019               216.8 ns/op           304 B/op          2 allocs/op
BenchmarkStaticRoutes1Chi-8                                      5508302               214.6 ns/op           304 B/op          2 allocs/op
BenchmarkStaticRoutes5Chi-8                                      5643144               211.5 ns/op           304 B/op          2 allocs/op
BenchmarkStaticRoutes10Chi-8                                     5650141               212.2 ns/op           304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes1Chi-8                            4443690               269.6 ns/op           304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes5Chi-8                            2747660               443.0 ns/op           304 B/op          2 allocs/op
BenchmarkPathParamBracketRoutes10Chi-8                           1867706               639.0 ns/op           304 B/op          2 allocs/op
BenchmarkStaticRoutesRootGin-8                                  33840350                35.03 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes1Gin-8                                     33891363                34.97 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes5Gin-8                                     33447457                35.87 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes10Gin-8                                    32442332                36.17 ns/op            0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Gin-8                             28137997                40.66 ns/op            0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5Gin-8                             15604334                76.89 ns/op            0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10Gin-8                             9713322               123.6 ns/op             0 B/op          0 allocs/op
BenchmarkStaticRoutesRootBunRouter-8                            62378167                19.04 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes1BunRouter-8                               54281038                21.52 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes5BunRouter-8                               54312976                21.70 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes10BunRouter-8                              54036307                21.85 ns/op            0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1BunRouter-8                       36516248                32.37 ns/op            0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5BunRouter-8                        8669384               138.0 ns/op             0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10BunRouter-8                       4173550               284.7 ns/op             0 B/op          0 allocs/op
BenchmarkStaticRoutesRootHTTPTreeMux-8                           6744598               179.6 ns/op           328 B/op          3 allocs/op
BenchmarkStaticRoutes1HTTPTreeMux-8                              6489350               180.7 ns/op           328 B/op          3 allocs/op
BenchmarkStaticRoutes5HTTPTreeMux-8                              5359826               234.5 ns/op           328 B/op          3 allocs/op
BenchmarkStaticRoutes10HTTPTreeMux-8                             4188642               286.9 ns/op           328 B/op          3 allocs/op
BenchmarkPathParamColonRoutes1HTTPTreeMux-8                      3014743               378.0 ns/op           680 B/op          6 allocs/op
BenchmarkPathParamColonRoutes5HTTPTreeMux-8                      1498790               787.8 ns/op           904 B/op          9 allocs/op
BenchmarkPathParamColonRoutes10HTTPTreeMux-8                      805834              1450 ns/op            1742 B/op         11 allocs/op
BenchmarkStaticRoutesRootBeegoMux-8                             23302629                50.27 ns/op           32 B/op          1 allocs/op
BenchmarkStaticRoutes1BeegoMux-8                                16677880                71.03 ns/op           32 B/op          1 allocs/op
BenchmarkStaticRoutes5BeegoMux-8                                 1000000              1077 ns/op              32 B/op          1 allocs/op
BenchmarkStaticRoutes10BeegoMux-8                                 577743              2026 ns/op              32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes1BeegoMux-8                         3329392               366.7 ns/op           672 B/op          5 allocs/op
BenchmarkPathParamColonRoutes5BeegoMux-8                          813811              1439 ns/op             672 B/op          5 allocs/op
BenchmarkPathParamColonRoutes10BeegoMux-8                         344250              3417 ns/op            1254 B/op          6 allocs/op
BenchmarkStaticRoutesRootGorillaMux-8                            2082201               571.2 ns/op           720 B/op          7 allocs/op
BenchmarkStaticRoutes1GorillaMux-8                               2088121               575.9 ns/op           720 B/op          7 allocs/op
BenchmarkStaticRoutes5GorillaMux-8                               1977076               612.0 ns/op           720 B/op          7 allocs/op
BenchmarkStaticRoutes10GorillaMux-8                              1843194               648.8 ns/op           720 B/op          7 allocs/op
BenchmarkPathParamBracketRoutes1GorillaMux-8                     1366618               874.8 ns/op          1024 B/op          8 allocs/op
BenchmarkPathParamBracketRoutes5GorillaMux-8                      516780              2217 ns/op            1088 B/op          8 allocs/op
BenchmarkPathParamBracketRoutes10GorillaMux-8                     242192              4879 ns/op            1751 B/op          9 allocs/op
BenchmarkStaticRoutesRootBon-8                                  84426278                13.91 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes1Bon-8                                     84835378                13.92 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes5Bon-8                                     84868131                13.90 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes10Bon-8                                    81009922                14.74 ns/op            0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Bon-8                              6376364               186.0 ns/op           304 B/op          2 allocs/op
BenchmarkPathParamColonRoutes5Bon-8                              4626661               258.9 ns/op           304 B/op          2 allocs/op
BenchmarkPathParamColonRoutes10Bon-8                             3383512               354.7 ns/op           304 B/op          2 allocs/op
BenchmarkStaticRoutesRootDenco-8                                86391010                13.73 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes1Denco-8                                   86784367                13.73 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes5Denco-8                                   86470388                13.72 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes10Denco-8                                  86911906                13.72 ns/op            0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Denco-8                           19862026                59.49 ns/op           32 B/op          1 allocs/op
BenchmarkPathParamColonRoutes5Denco-8                            8485465               141.3 ns/op           160 B/op          1 allocs/op
BenchmarkPathParamColonRoutes10Denco-8                           5047719               237.2 ns/op           320 B/op          1 allocs/op
BenchmarkStaticRoutesRootEcho-8                                 41954373                28.08 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes1Echo-8                                    36967816                32.18 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes5Echo-8                                    23895873                49.96 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes10Echo-8                                   12395061                96.61 ns/op            0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes1Echo-8                            30886804                39.32 ns/op            0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes5Echo-8                            12091669                98.75 ns/op            0 B/op          0 allocs/op
BenchmarkPathParamColonRoutes10Echo-8                            6830480               176.0 ns/op             0 B/op          0 allocs/op
BenchmarkStaticRoutesRootGocraftWeb-8                            1259409               943.2 ns/op           288 B/op          6 allocs/op
BenchmarkStaticRoutes1GocraftWeb-8                               1252980               955.4 ns/op           288 B/op          6 allocs/op
BenchmarkStaticRoutes5GocraftWeb-8                               1000000              1107 ns/op             352 B/op          6 allocs/op
BenchmarkStaticRoutes10GocraftWeb-8                               935240              1217 ns/op             432 B/op          6 allocs/op
BenchmarkPathParamColonRoutes1GocraftWeb-8                        966060              1165 ns/op             656 B/op          9 allocs/op
BenchmarkPathParamColonRoutes5GocraftWeb-8                        756178              1545 ns/op             944 B/op         12 allocs/op
BenchmarkPathParamColonRoutes10GocraftWeb-8                       518744              2281 ns/op            1862 B/op         14 allocs/op
BenchmarkStaticRoutesRootGorouter-8                             23695804                50.08 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes1Gorouter-8                                26121121                44.53 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes5Gorouter-8                                11303084               102.8 ns/op             0 B/op          0 allocs/op
BenchmarkStaticRoutes10Gorouter-8                                7328988               163.8 ns/op             0 B/op          0 allocs/op
BenchmarkPathParamBracketRoutes1Gorouter-8                       4579246               267.3 ns/op           360 B/op          4 allocs/op
BenchmarkPathParamBracketRoutes5Gorouter-8                       2987978               426.6 ns/op           488 B/op          4 allocs/op
BenchmarkPathParamBracketRoutes10Gorouter-8                      2060394               550.8 ns/op           648 B/op          4 allocs/op
BenchmarkStaticRoutesRootOzzoRouting-8                          28247007                41.33 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes1OzzoRouting-8                             28425942                42.13 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes5OzzoRouting-8                             23588305                50.34 ns/op            0 B/op          0 allocs/op
BenchmarkStaticRoutes10OzzoRouting-8                            18309638                65.64 ns/op            0 B/op          0 allocs/op
BenchmarkPathParamInequalitySignRoutes1OzzoRouting-8            23027325                51.88 ns/op            0 B/op          0 allocs/op
BenchmarkPathParamInequalitySignRoutes5OzzoRouting-8            11985876                99.66 ns/op            0 B/op          0 allocs/op
BenchmarkPathParamInequalitySignRoutes10OzzoRouting-8            7554290               159.3 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/go-router-benchmark  140.014s
```

# Contribution
We are always accepting issues, pull requests, and other requests and questions.

We look forward to your contribution！

If you have an HTTP Router or test case you would like to add, please send us an Issue or Pull Request.

# License
This project is licensed under the terms of the MIT license.

## Author
- [@bmf-san](https://twitter.com/bmf_san)
- [bmf-tech.com](http://bmf-tech.com/)
