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
- [lkeix/techbook13-sample](https://github.com/lkeix/techbook13-sample)

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
|       time        | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux          | 24301910           | 22053468        | 13324357        | 8851803          |
| goblin            | 32296879           | 16738813        | 5753088         | 3111172          |
| httprouter        | 100000000          | 100000000       | 100000000       | 72498970         |
| chi               | 5396652            | 5350285         | 5353856         | 5415325          |
| gin               | 34933861           | 34088810        | 34136852        | 33966028         |
| bunrouter         | 63478486           | 54812665        | 53564055        | 54345159         |
| httptreemux       | 6669231            | 6219157         | 5278312         | 4300488          |
| beegomux          | 22320199           | 15369320        | 1000000         | 577272           |
| gorillamux        | 1807042            | 2104210         | 1904696         | 1869037          |
| bon               | 72425132           | 56830177        | 59573305        | 58364338         |
| denco             | 90249313           | 92561344        | 89325312        | 73905086         |
| echo              | 41742093           | 36207878        | 23962478        | 12379764         |
| gocraftweb        | 1284613            | 1262863         | 1000000         | 889360           |
| gorouter          | 21622920           | 28592134        | 15582778        | 9636147          |
| ozzorouting       | 31406931           | 34989970        | 24825552        | 19431296         |
| techbook13-sample | 8176849            | 6349896         | 2684418         | 1384840          |

![time.png](/images/static-routes/time.png)

[Graph - time](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=800028423&format=interactive)

### nsop
|       nsop        | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux          | 50.44              | 54.97           | 89.81           | 135.2            |
| goblin            | 36.63              | 69.9            | 205.2           | 382.7            |
| httprouter        | 10.65              | 10.74           | 10.75           | 16.42            |
| chi               | 217.2              | 220.1           | 216.7           | 221.5            |
| gin               | 34.53              | 34.91           | 34.69           | 35.04            |
| bunrouter         | 18.77              | 21.78           | 22.41           | 22               |
| httptreemux       | 178.8              | 190.9           | 227.2           | 277.7            |
| beegomux          | 55.07              | 74.69           | 1080            | 2046             |
| gorillamux        | 595.7              | 572.8           | 626.5           | 643.3            |
| bon               | 15.75              | 20.17           | 18.87           | 19.16            |
| denco             | 14                 | 13.03           | 13.4            | 15.87            |
| echo              | 28.17              | 32.83           | 49.82           | 96.77            |
| gocraftweb        | 929.4              | 948.8           | 1078            | 1215             |
| gorouter          | 55.16              | 37.64           | 76.6            | 124.1            |
| ozzorouting       | 42.62              | 34.22           | 48.12           | 61.6             |
| techbook13-sample | 146.1              | 188.4           | 443.5           | 867.8            |

![nsop.png](/images/static-routes/nsop.png)

[Graph - nsop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1691114342&format=interactive)

### bop
|        bop        | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux          | 0                  | 0               | 0               | 0                |
| goblin            | 0                  | 16              | 80              | 160              |
| httprouter        | 0                  | 0               | 0               | 0                |
| chi               | 304                | 304             | 304             | 304              |
| gin               | 0                  | 0               | 0               | 0                |
| bunrouter         | 0                  | 0               | 0               | 0                |
| httptreemux       | 328                | 328             | 328             | 328              |
| beegomux          | 32                 | 32              | 32              | 32               |
| gorillamux        | 720                | 720             | 720             | 720              |
| bon               | 0                  | 0               | 0               | 0                |
| denco             | 0                  | 0               | 0               | 0                |
| echo              | 0                  | 0               | 0               | 0                |
| gocraftweb        | 288                | 288             | 352             | 432              |
| gorouter          | 0                  | 0               | 0               | 0                |
| ozzorouting       | 0                  | 0               | 0               | 0                |
| techbook13-sample | 304                | 308             | 432             | 872              |

![bop.png](/images/static-routes/bop.png)

[Graph - bop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=675738282&format=interactive)

### allocs
|      allocs       | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux          | 0                  | 0               | 0               | 0                |
| goblin            | 0                  | 1               | 1               | 1                |
| httprouter        | 0                  | 0               | 0               | 0                |
| chi               | 2                  | 2               | 2               | 2                |
| gin               | 0                  | 0               | 0               | 0                |
| bunrouter         | 0                  | 0               | 0               | 0                |
| httptreemux       | 3                  | 3               | 3               | 3                |
| beegomux          | 1                  | 1               | 1               | 1                |
| gorillamux        | 7                  | 7               | 7               | 7                |
| bon               | 0                  | 0               | 0               | 0                |
| denco             | 0                  | 0               | 0               | 0                |
| echo              | 0                  | 0               | 0               | 0                |
| gocraftweb        | 6                  | 6               | 6               | 6                |
| gorouter          | 0                  | 0               | 0               | 0                |
| ozzorouting       | 0                  | 0               | 0               | 0                |
| techbook13-sample | 2                  | 3               | 11              | 21               |

![allocs.png](/images/static-routes/allocs.png)

[Graph - allocs](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1054975173&format=interactive)

## Pathparams routes
### time
|       time        | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------------- | ------------------ | ------------------ | ------------------- |
| goblin            | 1802690            | 492392             | 252274              |
| httprouter        | 25775940           | 10057874           | 6060843             |
| chi               | 4337922            | 2687157            | 1772881             |
| gin               | 29479381           | 15714673           | 9586220             |
| bunrouter         | 37098772           | 8479642            | 3747968             |
| httptreemux       | 2610324            | 1550306            | 706356              |
| beegomux          | 3177818            | 797472             | 343969              |
| gorillamux        | 1364386            | 470180             | 223627              |
| bon               | 6639216            | 4486780            | 3285571             |
| denco             | 20093167           | 8503317            | 4988640             |
| echo              | 30667137           | 12028713           | 6721176             |
| gocraftweb        | 921375             | 734821             | 466641              |
| gorouter          | 4678617            | 3038450            | 2136946             |
| ozzorouting       | 27126000           | 12228037           | 7923040             |
| techbook13-sample | 3019774            | 917042             | 522897              |

![time.png](/images/pathparam-routes/time.png)

[Graph - time](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1039813866&format=interactive)

### nsop
|       nsop        | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------------- | ------------------ | ------------------ | ------------------- |
| goblin            | 652.4              | 2341               | 4504                |
| httprouter        | 45.73              | 117.4              | 204.2               |
| chi               | 276.4              | 442.8              | 677.6               |
| gin               | 40.21              | 76.39              | 124.3               |
| bunrouter         | 32.52              | 141.1              | 317.2               |
| httptreemux       | 399.7              | 778.5              | 1518                |
| beegomux          | 377.2              | 1446               | 3398                |
| gorillamux        | 850.3              | 2423               | 5264                |
| bon               | 186.5              | 269.6              | 364.4               |
| denco             | 60.47              | 139.4              | 238.7               |
| echo              | 39.36              | 99.6               | 175.7               |
| gocraftweb        | 1181               | 1540               | 2280                |
| gorouter          | 256.4              | 393                | 557.6               |
| ozzorouting       | 43.66              | 99.52              | 150.4               |
| techbook13-sample | 380.7              | 1154               | 2150                |

![nsop.png](/images/pathparam-routes/nsop.png)

[Graph - nsop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1534246873&format=interactive)

### bop
|        bop        | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------------- | ------------------ | ------------------ | ------------------- |
| goblin            | 409                | 962                | 1608                |
| httprouter        | 32                 | 160                | 320                 |
| chi               | 304                | 304                | 304                 |
| gin               | 0                  | 0                  | 0                   |
| bunrouter         | 0                  | 0                  | 0                   |
| httptreemux       | 680                | 904                | 1742                |
| beegomux          | 672                | 672                | 1254                |
| gorillamux        | 1024               | 1088               | 1751                |
| bon               | 304                | 304                | 304                 |
| denco             | 32                 | 160                | 320                 |
| echo              | 0                  | 0                  | 0                   |
| gocraftweb        | 656                | 944                | 1862                |
| gorouter          | 360                | 488                | 648                 |
| ozzorouting       | 0                  | 0                  | 0                   |
| techbook13-sample | 432                | 968                | 1792                |

![bop.png](/images/pathparam-routes/bop.png)

[Graph - bop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=73824357&format=interactive)

### allocs
|      allocs       | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------------- | ------------------ | ------------------ | ------------------- |
| goblin            | 6                  | 13                 | 19                  |
| httprouter        | 1                  | 1                  | 1                   |
| chi               | 2                  | 2                  | 2                   |
| gin               | 0                  | 0                  | 0                   |
| bunrouter         | 0                  | 0                  | 0                   |
| httptreemux       | 6                  | 9                  | 11                  |
| beegomux          | 5                  | 5                  | 6                   |
| gorillamux        | 8                  | 8                  | 9                   |
| bon               | 2                  | 2                  | 2                   |
| denco             | 1                  | 1                  | 1                   |
| echo              | 0                  | 0                  | 0                   |
| gocraftweb        | 9                  | 12                 | 14                  |
| gorouter          | 4                  | 4                  | 4                   |
| ozzorouting       | 0                  | 0                  | 0                   |
| techbook13-sample | 10                 | 33                 | 59                  |

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
