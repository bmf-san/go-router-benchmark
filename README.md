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

# Run benchmark tests
|         Command          |                                                          Description                                                          |
| :----------------------- | :---------------------------------------------------------------------------------------------------------------------------- |
| test-benchmark           | Run benchmark tests.                                                                                                          |
| test-benchmark-static    | Run benchmark tests only static.                                                                                              |
| test-benchmark-pathparam | Run benchmark tests only pathparam.                                                                                           |
| test-benchmark-by-regexp | Run benchmark tests using regexp. ex. make test-benchmark-by-regexp EXP=Goblin, make test-benchmark-by-name EXP=StaticRoutes1 |

# Results
Benchmark results are published in a spreadsheet.

[<Public>go-router-benchmark](https://docs.google.com/spreadsheets/d/1DrDNGJXfquw_PED3-eMqWqh7qbCTVBWZtqaKPXtngxg/edit#gid=1913830192)

## Static routes
### time
|    time     | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux    | 24101668           | 20803911        | 12802896        | 8919028          |
| goblin      | 36447806           | 17144642        | 5840305         | 3155595          |
| httprouter  | 100000000          | 100000000       | 84698406        | 83383032         |
| chi         | 5529261            | 5442328         | 5233761         | 5342503          |
| gin         | 34119340           | 34670139        | 33786320        | 34107339         |
| bunrouter   | 64181419           | 54405626        | 54746812        | 55535950         |
| httptreemux | 6516063            | 6448729         | 5253651         | 4152639          |
| beegomux    | 23423946           | 16189789        | 1000000         | 586872           |
| gorillamux  | 2130843            | 2079057         | 1995946         | 1850852          |
| bon         | 76976530           | 78365212        | 77670530        | 69186360         |
| denco       | 78300441           | 71648206        | 71871530        | 71719041         |
| echo        | 42333267           | 35916672        | 24450126        | 11766090         |
| gocraftweb  | 1260608            | 1243606         | 1000000         | 903750           |
| gorouter    | 30632953           | 23144054        | 15452823        | 8576277          |
| ozzorouting | 36290733           | 33224583        | 26447349        | 19445872         |

![time.png](/images/static-routes/time.png)

[Graph - time](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=800028423&format=interactive)

### nsop
|    nsop     |           static-routes-root           | static-routes-root |           static-routes-1           | static-routes-1 |           static-routes-5           | static-routes-5 |           static-routes-10           | static-routes-10 |
| ----------- | -------------------------------------- | ------------------ | ----------------------------------- | --------------- | ----------------------------------- | --------------- | ------------------------------------ | ---------------- |
| servemux    | BenchmarkStaticRoutesRootServeMux-8    | 50                 | BenchmarkStaticRoutes1ServeMux-8    | 55.03           | BenchmarkStaticRoutes5ServeMux-8    | 96.2            | BenchmarkStaticRoutes10ServeMux-8    | 137.8            |
| goblin      | BenchmarkStaticRoutesRootGoblin-8      | 34.09              | BenchmarkStaticRoutes1Goblin-8      | 67.76           | BenchmarkStaticRoutes5Goblin-8      | 201.9           | BenchmarkStaticRoutes10Goblin-8      | 377.6            |
| httprouter  | BenchmarkStaticRoutesRootHTTPRouter-8  | 11.84              | BenchmarkStaticRoutes1HTTPRouter-8  | 10.63           | BenchmarkStaticRoutes5HTTPRouter-8  | 12.99           | BenchmarkStaticRoutes10HTTPRouter-8  | 12.97            |
| chi         | BenchmarkStaticRoutesRootChi-8         | 232.7              | BenchmarkStaticRoutes1Chi-8         | 243.1           | BenchmarkStaticRoutes5Chi-8         | 248.3           | BenchmarkStaticRoutes10Chi-8         | 216.9            |
| gin         | BenchmarkStaticRoutesRootGin-8         | 34.56              | BenchmarkStaticRoutes1Gin-8         | 37.7            | BenchmarkStaticRoutes5Gin-8         | 35.77           | BenchmarkStaticRoutes10Gin-8         | 35.13            |
| bunrouter   | BenchmarkStaticRoutesRootBunRouter-8   | 19.39              | BenchmarkStaticRoutes1BunRouter-8   | 22.53           | BenchmarkStaticRoutes5BunRouter-8   | 22.03           | BenchmarkStaticRoutes10BunRouter-8   | 22.01            |
| httptreemux | BenchmarkStaticRoutesRootHTTPTreeMux-8 | 206.6              | BenchmarkStaticRoutes1HTTPTreeMux-8 | 185.6           | BenchmarkStaticRoutes5HTTPTreeMux-8 | 239.6           | BenchmarkStaticRoutes10HTTPTreeMux-8 | 291.7            |
| beegomux    | BenchmarkStaticRoutesRootBeegoMux-8    | 50.37              | BenchmarkStaticRoutes1BeegoMux-8    | 73.47           | BenchmarkStaticRoutes5BeegoMux-8    | 1086            | BenchmarkStaticRoutes10BeegoMux-8    | 2049             |
| gorillamux  | BenchmarkStaticRoutesRootGorillaMux-8  | 560.8              | BenchmarkStaticRoutes1GorillaMux-8  | 575.7           | BenchmarkStaticRoutes5GorillaMux-8  | 642.5           | BenchmarkStaticRoutes10GorillaMux-8  | 691.8            |
| bon         | BenchmarkStaticRoutesRootBon-8         | 15.43              | BenchmarkStaticRoutes1Bon-8         | 15.63           | BenchmarkStaticRoutes5Bon-8         | 15.85           | BenchmarkStaticRoutes10Bon-8         | 16.81            |
| denco       | BenchmarkStaticRoutesRootDenco-8       | 14.9               | BenchmarkStaticRoutes1Denco-8       | 16.56           | BenchmarkStaticRoutes5Denco-8       | 16.77           | BenchmarkStaticRoutes10Denco-8       | 16.09            |
| echo        | BenchmarkStaticRoutesRootEcho-8        | 28.04              | BenchmarkStaticRoutes1Echo-8        | 32.84           | BenchmarkStaticRoutes5Echo-8        | 49.97           | BenchmarkStaticRoutes10Echo-8        | 97.44            |
| gocraftweb  | BenchmarkStaticRoutesRootGocraftWeb-8  | 923.2              | BenchmarkStaticRoutes1GocraftWeb-8  | 944.7           | BenchmarkStaticRoutes5GocraftWeb-8  | 1168            | BenchmarkStaticRoutes10GocraftWeb-8  | 1215             |
| gorouter    | BenchmarkStaticRoutesRootGorouter-8    | 38.78              | BenchmarkStaticRoutes1Gorouter-8    | 50.27           | BenchmarkStaticRoutes5Gorouter-8    | 75.96           | BenchmarkStaticRoutes10Gorouter-8    | 147.9            |
| ozzorouting | BenchmarkStaticRoutesRootOzzoRouting-8 | 32.8               | BenchmarkStaticRoutes1OzzoRouting-8 | 35.08           | BenchmarkStaticRoutes5OzzoRouting-8 | 47.33           | BenchmarkStaticRoutes10OzzoRouting-8 | 61.13            |

![nsop.png](/images/static-routes/nsop.png)

[Graph - nsop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1691114342&format=interactive)

### bop
|     bop     | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux    | 0                  | 0               | 0               | 0                |
| goblin      | 0                  | 16              | 80              | 160              |
| httprouter  | 0                  | 0               | 0               | 0                |
| chi         | 304                | 304             | 304             | 304              |
| gin         | 0                  | 0               | 0               | 0                |
| bunrouter   | 0                  | 0               | 0               | 0                |
| httptreemux | 328                | 328             | 328             | 328              |
| beegomux    | 32                 | 32              | 32              | 32               |
| gorillamux  | 720                | 720             | 720             | 720              |
| bon         | 0                  | 0               | 0               | 0                |
| denco       | 0                  | 0               | 0               | 0                |
| echo        | 0                  | 0               | 0               | 0                |
| gocraftweb  | 288                | 288             | 352             | 432              |
| gorouter    | 0                  | 0               | 0               | 0                |
| ozzorouting | 0                  | 0               | 0               | 0                |

![bop.png](/images/static-routes/bop.png)

[Graph - bop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=675738282&format=interactive)

### allocs
|   allocs    | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux    | 0                  | 0               | 0               | 0                |
| goblin      | 0                  | 1               | 1               | 1                |
| httprouter  | 0                  | 0               | 0               | 0                |
| chi         | 2                  | 2               | 2               | 2                |
| gin         | 0                  | 0               | 0               | 0                |
| bunrouter   | 0                  | 0               | 0               | 0                |
| httptreemux | 3                  | 3               | 3               | 3                |
| beegomux    | 1                  | 1               | 1               | 1                |
| gorillamux  | 7                  | 7               | 7               | 7                |
| bon         | 0                  | 0               | 0               | 0                |
| denco       | 0                  | 0               | 0               | 0                |
| echo        | 0                  | 0               | 0               | 0                |
| gocraftweb  | 6                  | 6               | 6               | 6                |
| gorouter    | 0                  | 0               | 0               | 0                |
| ozzorouting | 0                  | 0               | 0               | 0                |

![allocs.png](/images/static-routes/allocs.png)

[Graph - allocs](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1054975173&format=interactive)

## Pathparams routes
### time
|    time     | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------- | ------------------ | ------------------ | ------------------- |
| goblin      | 1873453            | 507033             | 260714              |
| httprouter  | 26311221           | 10168474           | 5970110             |
| chi         | 4382772            | 2648826            | 1849846             |
| gin         | 29532524           | 14766572           | 9653036             |
| bunrouter   | 35877165           | 8573814            | 4150300             |
| httptreemux | 3127624            | 1560525            | 784714              |
| beegomux    | 3125070            | 767952             | 342159              |
| gorillamux  | 1268805            | 489907             | 223706              |
| bon         | 6451303            | 4559571            | 3263553             |
| denco       | 19714711           | 8164539            | 4503146             |
| echo        | 30035384           | 12032884           | 6730594             |
| gocraftweb  | 930712             | 714529             | 452245              |
| gorouter    | 4722282            | 3125533            | 2309708             |
| ozzorouting | 23310682           | 12970110           | 7913998             |

![time.png](/images/pathparam-routes/time.png)

[Graph - time](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1039813866&format=interactive)

### nsop
|    nsop     | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------- | ------------------ | ------------------ | ------------------- |
| goblin      | 630.4              | 2330               | 4384                |
| httprouter  | 45.62              | 121.5              | 198.5               |
| chi         | 273.6              | 504.2              | 644.3               |
| gin         | 40.54              | 78.41              | 124                 |
| bunrouter   | 33.25              | 139.5              | 288.4               |
| httptreemux | 383.2              | 783                | 1515                |
| beegomux    | 379.4              | 1447               | 3392                |
| gorillamux  | 891.7              | 2287               | 5656                |
| bon         | 189.1              | 264.3              | 370.1               |
| denco       | 59.81              | 141.5              | 248                 |
| echo        | 39.41              | 106.7              | 192                 |
| gocraftweb  | 1148               | 1561               | 2489                |
| gorouter    | 263.5              | 378.7              | 529.1               |
| ozzorouting | 46.26              | 93.6               | 151                 |

![nsop.png](/images/pathparam-routes/nsop.png)

[Graph - nsop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1534246873&format=interactive)

### bop
|     bop     | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------- | ------------------ | ------------------ | ------------------- |
| goblin      | 409                | 961                | 1607                |
| httprouter  | 32                 | 160                | 320                 |
| chi         | 304                | 304                | 304                 |
| gin         | 0                  | 0                  | 0                   |
| bunrouter   | 0                  | 0                  | 0                   |
| httptreemux | 680                | 904                | 1742                |
| beegomux    | 672                | 672                | 1254                |
| gorillamux  | 1024               | 1088               | 1751                |
| bon         | 304                | 304                | 304                 |
| denco       | 32                 | 160                | 320                 |
| echo        | 0                  | 0                  | 0                   |
| gocraftweb  | 656                | 944                | 1862                |
| gorouter    | 360                | 488                | 648                 |
| ozzorouting | 0                  | 0                  | 0                   |

![bop.png](/images/pathparam-routes/bop.png)

[Graph - bop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=73824357&format=interactive)

### allocs
|   allocs    | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------- | ------------------ | ------------------ | ------------------- |
| goblin      | 6                  | 13                 | 19                  |
| httprouter  | 1                  | 1                  | 1                   |
| chi         | 2                  | 2                  | 2                   |
| gin         | 0                  | 0                  | 0                   |
| bunrouter   | 0                  | 0                  | 0                   |
| httptreemux | 6                  | 9                  | 11                  |
| beegomux    | 5                  | 5                  | 6                   |
| gorillamux  | 8                  | 8                  | 9                   |
| bon         | 2                  | 2                  | 2                   |
| denco       | 1                  | 1                  | 1                   |
| echo        | 0                  | 0                  | 0                   |
| gocraftweb  | 9                  | 12                 | 14                  |
| gorouter    | 4                  | 4                  | 4                   |
| ozzorouting | 0                  | 0                  | 0                   |

![allocs.png](/images/pathparam-routes/allocs.png)

[Graph - allocs](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=344550080&format=interactive)

# Conclusion
It can be seen that the HTTP Router with better performance has less performance degradation in each test case.
This is a clear trend that shows that the implementation is optimized.

The better performing HTTP Route seems to employ a more sophisticated tree structure.
Echo,gin,httprouter,bon,chi seem to adopt Radix tree (Patricia trie) and denco double array.

As for my implementation of goblin, it is a proprietary extension of the trie tree, which is not very well optimized, and I could clearly see that its performance is lower than the other HTTP Routers. (I will try my best to improve it...)

I don't think it is reasonable to assume that HTTP Router should not be adopted because of its seemingly poor performance in this benchmark result.

Some HTTP Routers are considered to be highly functional and easy to use, even if their performance is a little lower.

Some of the HTTP Routers listed here are no longer under development.

Since goblin has no plans to stop development for the time being, please try it if you like!

# Contribution
We are always accepting issues, pull requests, and other requests and questions.

We look forward to your contribution！

## Want to add an HTTP Router ?
If you have an HTTP Router or test case you would like to add, please send us an Issue or Pull Request.

If you are submitting an Issue, please tell us about the HTTP Router you would like to add.

If you are submitting a PullRequest, please add a test case.
Reporting of benchmark runs is done by the owner, so you do not need to update the run results.

If you have any questions, you can ask them in Issue.

# License
This project is licensed under the terms of the MIT license.

## Author
- [@bmf-san](https://twitter.com/bmf_san)
- [bmf-tech.com](http://bmf-tech.com/)
