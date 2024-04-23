# go-router-benchmark
Compare the performance of routers built with golang.

# Table of contents
- [go-router-benchmark](#go-router-benchmark)
- [Table of contents](#table-of-contents)
- [Listed routers](#listed-routers)
- [Motivation](#motivation)
- [Benchmark test](#benchmark-test)
	- [Static route](#static-route)
	- [Path parameter route](#path-parameter-route)
- [Run benchmark tests](#run-benchmark-tests)
- [Results](#results)
	- [Static routes](#static-routes)
		- [time](#time)
		- [nsop](#nsop)
		- [bop](#bop)
		- [allocs](#allocs)
	- [Pathparams routes](#pathparams-routes)
		- [time](#time-1)
		- [nsop](#nsop-1)
		- [bop](#bop-1)
		- [allocs](#allocs-1)
- [Conclusion](#conclusion)
- [Contribution](#contribution)
	- [Want to add an HTTP Router ?](#want-to-add-an-http-router-)
- [Contribution](#contribution-1)
- [Sponsor](#sponsor)
- [License](#license)
	- [Author](#author)

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

Benchmark system
- go version: go1.19
- goos: darwin
- goarch: amd64
- pkg: github.com/go-router-benchmark
- cpu: VirtualApple @ 2.50GHz

## Static routes
### time
| time              | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux          | 27831720           | 19891950        | 4838466         | 2508039          |
| goblin            | 78151288           | 39773182        | 10053289        | 4709226          |
| httprouter        | 144303050          | 134171820       | 129710535       | 100000000        |
| chi               | 7614604            | 7730403         | 7558288         | 6086882          |
| gin               | 44050909           | 43401141        | 43176404        | 41443738         |
| bunrouter         | 100000000          | 79699796        | 76368669        | 71506602         |
| httptreemux       | 8800170            | 8486769         | 6520071         | 4769065          |
| beegomux          | 31233393           | 26120126        | 8418674         | 4674050          |
| gorillamux        | 3210363            | 3138418         | 2914868         | 2262861          |
| bon               | 100000000          | 100000000       | 100000000       | 97284812         |
| denco             | 134409880          | 134490919       | 133387552       | 131848746        |
| echo              | 65374460           | 60613204        | 39926578        | 23414919         |
| gocraftweb        | 1686422            | 1644936         | 1460192         | 1208860          |
| gorouter          | 89735841           | 64869247        | 31833826        | 18472964         |
| ozzorouting       | 47574130           | 45542164        | 35507553        | 27385978         |
| techbook13-sample | 10733499           | 7743865         | 3166988         | 1523804          |

![time.png](/images/static-routes/time.png)

[Graph - time](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=800028423&format=interactive)

### nsop
| nsop              | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux          | 43.09              | 60.18           | 226.1           | 457.7            |
| goblin            | 15.34              | 30.23           | 119.7           | 265.5            |
| httprouter        | 8.328              | 8.931           | 9.291           | 10.58            |
| chi               | 160                | 159.4           | 159.5           | 183.6            |
| gin               | 27.66              | 27.68           | 27.81           | 29.56            |
| bunrouter         | 10.63              | 15.24           | 15.72           | 18.61            |
| httptreemux       | 137.1              | 140.7           | 183.2           | 270.6            |
| beegomux          | 38.22              | 46.08           | 143.6           | 251.2            |
| gorillamux        | 374.9              | 380.8           | 411.7           | 550.8            |
| bon               | 10.68              | 10.46           | 10.7            | 10.93            |
| denco             | 9.027              | 8.986           | 8.962           | 9.076            |
| echo              | 18.54              | 19.73           | 30.14           | 50.64            |
| gocraftweb        | 709.5              | 718             | 826             | 1024             |
| gorouter          | 13.7               | 18.46           | 37.49           | 71.49            |
| ozzorouting       | 25.4               | 26.31           | 33.01           | 43.9             |
| techbook13-sample | 115.3              | 154.8           | 379.8           | 779.6            |

![nsop.png](/images/static-routes/nsop.png)

[Graph - nsop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1691114342&format=interactive)

### bop
| bop               | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux          | 0                  | 0               | 0               | 0                |
| goblin            | 0                  | 0               | 0               | 0                |
| httprouter        | 0                  | 0               | 0               | 0                |
| chi               | 336                | 336             | 336             | 336              |
| gin               | 0                  | 0               | 0               | 0                |
| bunrouter         | 0                  | 0               | 0               | 0                |
| httptreemux       | 360                | 360             | 360             | 360              |
| beegomux          | 32                 | 32              | 32              | 32               |
| gorillamux        | 784                | 784             | 784             | 784              |
| bon               | 0                  | 0               | 0               | 0                |
| denco             | 0                  | 0               | 0               | 0                |
| echo              | 0                  | 0               | 0               | 0                |
| gocraftweb        | 288                | 288             | 352             | 432              |
| gorouter          | 0                  | 0               | 0               | 0                |
| ozzorouting       | 0                  | 0               | 0               | 0                |
| techbook13-sample | 336                | 340             | 464             | 904              |

![bop.png](/images/static-routes/bop.png)

[Graph - bop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=675738282&format=interactive)

### allocs
| allocs            | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux          | 0                  | 0               | 0               | 0                |
| goblin            | 0                  | 0               | 0               | 0                |
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
| time              | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------------- | ------------------ | ------------------ | ------------------- |
| servemux          | 9879615            | 3095632            | 1716556             |
| goblin            | 5543427            | 3288781            | 2148358             |
| httprouter        | 33200109           | 12520192           | 7242321             |
| chi               | 4124253            | 2560972            | 1262286             |
| gin               | 35934104           | 18361316           | 10710009            |
| bunrouter         | 59392094           | 10587997           | 5214822             |
| httptreemux       | 3764035            | 2022978            | 1000000             |
| beegomux          | 3953678            | 2586265            | 1457932             |
| gorillamux        | 2004242            | 961486             | 521152              |
| bon               | 7565557            | 5322081            | 4093870             |
| denco             | 24679846           | 9708924            | 6409720             |
| echo              | 46360610           | 15237498           | 8387311             |
| gocraftweb        | 1287634            | 960009             | 717404              |
| gorouter          | 6209989            | 4295180            | 3322506             |
| ozzorouting       | 35132318           | 15635671           | 9333058             |
| techbook13-sample | 3436351            | 1000000            | 647994              |

![time.png](/images/pathparam-routes/time.png)

[Graph - time](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1039813866&format=interactive)

### nsop
| nsop              | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------------- | ------------------ | ------------------ | ------------------- |
| servemux          | 120.8              | 383.8              | 697.1               |
| goblin            | 217.7              | 367.2              | 558                 |
| httprouter        | 36.3               | 96.27              | 155.5               |
| chi               | 304.8              | 468.1              | 961.9               |
| gin               | 33.44              | 64.27              | 104.1               |
| bunrouter         | 20.3               | 115.1              | 229.9               |
| httptreemux       | 316.6              | 599                | 1080                |
| beegomux          | 316.7              | 466.9              | 826                 |
| gorillamux        | 609.5              | 1215               | 2327                |
| bon               | 159.2              | 226.9              | 293.9               |
| denco             | 49.59              | 123.2              | 187                 |
| echo              | 26.17              | 78.84              | 143.4               |
| gocraftweb        | 936.6              | 1228               | 1654                |
| gorouter          | 196.1              | 277.9              | 364.7               |
| ozzorouting       | 33.99              | 77.03              | 129                 |
| techbook13-sample | 355.3              | 1045               | 1813                |

![nsop.png](/images/pathparam-routes/nsop.png)

[Graph - nsop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1534246873&format=interactive)

### bop
| bop               | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------------- | ------------------ | ------------------ | ------------------- |
| servemux          | 16                 | 240                | 496                 |
| goblin            | 360                | 360                | 360                 |
| httprouter        | 32                 | 160                | 320                 |
| chi               | 672                | 672                | 1254                |
| gin               | 0                  | 0                  | 0                   |
| bunrouter         | 0                  | 0                  | 0                   |
| httptreemux       | 712                | 936                | 1774                |
| beegomux          | 704                | 704                | 1286                |
| gorillamux        | 1088               | 1152               | 1815                |
| bon               | 336                | 336                | 336                 |
| denco             | 32                 | 160                | 320                 |
| echo              | 0                  | 0                  | 0                   |
| gocraftweb        | 656                | 944                | 1862                |
| gorouter          | 392                | 520                | 680                 |
| ozzorouting       | 0                  | 0                  | 0                   |
| techbook13-sample | 464                | 1000               | 1824                |

![bop.png](/images/pathparam-routes/bop.png)

[Graph - bop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=73824357&format=interactive)

### allocs
| allocs            | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------------- | ------------------ | ------------------ | ------------------- |
| servemux          | 1                  | 4                  | 5                   |
| goblin            | 3                  | 3                  | 3                   |
| httprouter        | 1                  | 1                  | 1                   |
| chi               | 4                  | 4                  | 5                   |
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

# Contribution
Issues and Pull Requests are always welcome.

We would be happy to receive your contributions.

Please review the following documents before making a contribution.

[CODE_OF_CONDUCT](https://github.com/bmf-san/go-router-benchmark/blob/master/.github/CODE_OF_CONDUCT.md)
[CONTRIBUTING](https://github.com/bmf-san/go-router-benchmark/blob/master/.github/CONTRIBUTING.md)

# Sponsor
If you like it, I would be happy to have you sponsor it!

[GitHub Sponsors - bmf-san](https://github.com/sponsors/bmf-san)

Or I would be happy to get a STAR.

It motivates me to keep up with ongoing maintenance :D

# License
Based on the MIT License.

[LICENSE](https://github.com/bmf-san/goblin/blob/master/LICENSE)

## Author
[bmf-san](https://github.com/bmf-san)

- Email
  - bmf.infomation@gmail.com
- Blog
  - [bmf-tech.com](http://bmf-tech.com)
- Twitter
  - [bmf-san](https://twitter.com/bmf-san)
