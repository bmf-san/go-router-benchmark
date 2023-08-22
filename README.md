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
- [jba/muxpatterns](https://github.com/jba/muxpatterns)

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
|       time        | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux          | 32706376           | 31260028        | 15832058        | 7264929          |
| goblin            | 76439406           | 40379964        | 9580180         | 4323399          |
| httprouter        | 144948184          | 134442093       | 129799516       | 100000000        |
| chi               | 8070156            | 8052058         | 8008406         | 8048611          |
| gin               | 44602617           | 44147605        | 42466225        | 41878116         |
| bunrouter         | 122551147          | 83394621        | 80952087        | 75950764         |
| httptreemux       | 8886540            | 8690398         | 6828529         | 5250801          |
| beegomux          | 31958985           | 25187836        | 8400084         | 4969017          |
| gorillamux        | 3404859            | 3336543         | 3059030         | 2562878          |
| bon               | 100000000          | 100000000       | 100000000       | 100000000        |
| denco             | 133318143          | 134562303       | 133360999       | 134754518        |
| echo              | 64878891           | 59950042        | 39227169        | 23970236         |
| gocraftweb        | 1608999            | 1608612         | 1421674         | 1213682          |
| gorouter          | 71320846           | 56132800        | 29620791        | 18047248         |
| ozzorouting       | 47726209           | 46019254        | 35856304        | 28017253         |
| techbook13-sample | 11669716           | 8451075         | 3286712         | 1664376          |

![time.png](/images/static-routes/time.png)

[Graph - time](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=800028423&format=interactive)

### nsop
|       nsop        | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux          | 37.13              | 38.26           | 75.65           | 148.7            |
| goblin            | 15.66              | 29.72           | 125.2           | 277.2            |
| httprouter        | 8.288              | 8.987           | 9.248           | 10.22            |
| chi               | 148.7              | 148.2           | 148.6           | 148.3            |
| gin               | 26.91              | 27.15           | 28.24           | 28.71            |
| bunrouter         | 9.844              | 14.36           | 14.85           | 15.8             |
| httptreemux       | 133.7              | 138.4           | 176.3           | 229.9            |
| beegomux          | 37.7               | 47.5            | 143.3           | 241.6            |
| gorillamux        | 352.3              | 358.8           | 393.1           | 469.6            |
| bon               | 10.6               | 10.62           | 10.57           | 10.52            |
| denco             | 8.944              | 9.059           | 8.999           | 8.963            |
| echo              | 18.51              | 20.12           | 30.67           | 50.5             |
| gocraftweb        | 743.6              | 744             | 849.2           | 986.7            |
| gorouter          | 16.54              | 21.49           | 40.53           | 66.6             |
| ozzorouting       | 25.36              | 26.12           | 33.56           | 42.92            |
| techbook13-sample | 100.3              | 140.5           | 365.4           | 723              |

![nsop.png](/images/static-routes/nsop.png)

[Graph - nsop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1691114342&format=interactive)

### bop
|        bop        | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux          | 0                  | 0               | 0               | 0                |
| goblin            | 0                  | 0               | 0               | 0                |
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
|       time        | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------------- | ------------------ | ------------------ | ------------------- |
| goblin            | 5651378            | 3350758            | 2237424             |
| httprouter        | 35218974           | 13397467           | 8215023             |
| chi               | 7278495            | 4290862            | 2868873             |
| gin               | 37004343           | 19501255           | 12229501            |
| bunrouter         | 62347647           | 10800812           | 5265159             |
| httptreemux       | 4364602            | 2175886            | 1000000             |
| beegomux          | 4552228            | 2861100            | 1420195             |
| gorillamux        | 2252773            | 1000000            | 513360              |
| bon               | 9118044            | 6031597            | 4243491             |
| denco             | 25824546           | 10170320           | 6332478             |
| echo              | 45946130           | 15379792           | 8352284             |
| gocraftweb        | 1315434            | 969732             | 693129              |
| gorouter          | 7179020            | 4748841            | 3296766             |
| ozzorouting       | 37090459           | 16726933           | 10019931            |
| techbook13-sample | 3875658            | 1223982            | 651006              |

![time.png](/images/pathparam-routes/time.png)

[Graph - time](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1039813866&format=interactive)

### nsop
|       nsop        | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------------- | ------------------ | ------------------ | ------------------- |
| goblin            | 194                | 341.2              | 529.2               |
| httprouter        | 33.85              | 89.47              | 146.9               |
| chi               | 164.3              | 280.9              | 421.4               |
| gin               | 32.5               | 60.58              | 97.68               |
| bunrouter         | 19.26              | 110.7              | 228.4               |
| httptreemux       | 276.8              | 553.1              | 1111                |
| beegomux          | 262.1              | 420.6              | 845                 |
| gorillamux        | 530.7              | 1092               | 2320                |
| bon               | 131.9              | 198.9              | 281.8               |
| denco             | 46.26              | 118.3              | 189.8               |
| echo              | 26.14              | 78.09              | 143.7               |
| gocraftweb        | 914.5              | 1196               | 1694                |
| gorouter          | 168.9              | 253.2              | 362.9               |
| ozzorouting       | 32.27              | 71.76              | 119.8               |
| techbook13-sample | 308.6              | 982.1              | 1794                |

![nsop.png](/images/pathparam-routes/nsop.png)

[Graph - nsop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1534246873&format=interactive)

### bop
|        bop        | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------------- | ------------------ | ------------------ | ------------------- |
| goblin            | 328                | 328                | 328                 |
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
| goblin            | 3                  | 3                  | 3                   |
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
