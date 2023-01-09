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

Benchmark system
- go version: go1.19
- goos: darwin
- goarch: amd64
- pkg: github.com/go-router-benchmark
- cpu: VirtualApple @ 2.50GHz

## Static routes
### time
| time | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| --- | --- | --- | --- | --- |
| servemux | 32733160 | 31851897 | 13437040 | 7253216 |
| goblin | 51884132 | 35242333 | 8901708 | 4096243 |
| httprouter | 145829186 | 134900438 | 129621186 | 100000000 |
| chi | 7254184 | 7252261 | 7210196 | 7284903 |
| gin | 42563881 | 41631550 | 41211974 | 40183560 |
| bunrouter | 100000000 | 70006395 | 68203778 | 64502257 |
| httptreemux | 8065448 | 7859656 | 6201728 | 4849598 |
| beegomux | 31691360 | 25884801 | 8645566 | 5056942 |
| gorillamux | 3162660 | 3094188 | 2834072 | 2450286 |
| bon | 100000000 | 100000000 | 100000000 | 100000000 |
| denco | 135744692 | 133846348 | 134140972 | 133246112 |
| echo | 64909012 | 60949671 | 39699386 | 24012085 |
| gocraftweb | 1596338 | 1578338 | 1392926 | 1000000 |
| gorouter | 71930586 | 52163613 | 27706906 | 16627858 |
| ozzorouting | 47753750 | 45524383 | 35422566 | 27887978 |
| techbook13-sample | 10074039 | 7664803 | 3122547 | 1641374 |

![time.png](/images/static-routes/time.png)

[Graph - time](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=800028423&format=interactive)

### nsop
| nsop | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| --- | --- | --- | --- | --- |
| servemux | 36.33 | 37.56 | 75.73 | 148 |
| goblin | 23.33 | 33.99 | 135 | 290.5 |
| httprouter | 8.646 | 8.909 | 9.254 | 10.77 |
| chi | 166.5 | 163 | 166.4 | 174.8 |
| gin | 28.58 | 28.73 | 29.05 | 29.96 |
| bunrouter | 11.9 | 17.02 | 17.63 | 18.62 |
| httptreemux | 149.5 | 154 | 192.6 | 242.9 |
| beegomux | 37.85 | 46.46 | 138.2 | 238.5 |
| gorillamux | 381.9 | 386.4 | 419.6 | 489.2 |
| bon | 10.51 | 10.63 | 10.71 | 10.47 |
| denco | 8.91 | 8.916 | 9.011 | 9.005 |
| echo | 18.73 | 19.92 | 30.15 | 50.25 |
| gocraftweb | 749.8 | 768.7 | 862.6 | 1007 |
| gorouter | 16.65 | 24 | 45.26 | 71.23 |
| ozzorouting | 24.97 | 26.07 | 33.39 | 43.09 |
| techbook13-sample | 119.1 | 157.9 | 385 | 728.6 |

![nsop.png](/images/static-routes/nsop.png)

[Graph - nsop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1691114342&format=interactive)

### bop
| bop | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| --- | --- | --- | --- | --- |
| servemux | 0 | 0 | 0 | 0 |
| goblin | 0 | 0 | 0 | 0 |
| httprouter | 0 | 0 | 0 | 0 |
| chi | 304 | 304 | 304 | 304 |
| gin | 0 | 0 | 0 | 0 |
| bunrouter | 0 | 0 | 0 | 0 |
| httptreemux | 328 | 328 | 328 | 328 |
| beegomux | 32 | 32 | 32 | 32 |
| gorillamux | 720 | 720 | 720 | 720 |
| bon | 0 | 0 | 0 | 0 |
| denco | 0 | 0 | 0 | 0 |
| echo | 0 | 0 | 0 | 0 |
| gocraftweb | 288 | 288 | 352 | 432 |
| gorouter | 0 | 0 | 0 | 0 |
| ozzorouting | 0 | 0 | 0 | 0 |
| techbook13-sample | 304 | 308 | 432 | 872 |

![bop.png](/images/static-routes/bop.png)

[Graph - bop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=675738282&format=interactive)

### allocs
| allocs | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| --- | --- | --- | --- | --- |
| servemux | 0 | 0 | 0 | 0 |
| goblin | 0 | 0 | 0 | 0 |
| httprouter | 0 | 0 | 0 | 0 |
| chi | 2 | 2 | 2 | 2 |
| gin | 0 | 0 | 0 | 0 |
| bunrouter | 0 | 0 | 0 | 0 |
| httptreemux | 3 | 3 | 3 | 3 |
| beegomux | 1 | 1 | 1 | 1 |
| gorillamux | 7 | 7 | 7 | 7 |
| bon | 0 | 0 | 0 | 0 |
| denco | 0 | 0 | 0 | 0 |
| echo | 0 | 0 | 0 | 0 |
| gocraftweb | 6 | 6 | 6 | 6 |
| gorouter | 0 | 0 | 0 | 0 |
| ozzorouting | 0 | 0 | 0 | 0 |
| techbook13-sample | 2 | 3 | 11 | 21 |

![allocs.png](/images/static-routes/allocs.png)

[Graph - allocs](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1054975173&format=interactive)

## Pathparams routes
### time
| time | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| --- | --- | --- | --- |
| goblin | 4575871 | 2049280 | 1160929 |
| httprouter | 34437032 | 13029846 | 8165713 |
| chi | 5880913 | 3872946 | 2636372 |
| gin | 36084889 | 19343385 | 12088365 |
| bunrouter | 45573008 | 10754092 | 5223387 |
| httptreemux | 3641581 | 2070312 | 1000000 |
| beegomux | 4127463 | 2739680 | 1373012 |
| gorillamux | 2109531 | 1000000 | 497953 |
| bon | 8041033 | 5534796 | 4008288 |
| denco | 26334654 | 9830152 | 6212643 |
| echo | 44213551 | 15313064 | 8329726 |
| gocraftweb | 1291144 | 970765 | 708202 |
| gorouter | 6460888 | 4317782 | 3121602 |
| ozzorouting | 36781983 | 16313232 | 9731100 |
| techbook13-sample | 3656685 | 1000000 | 634324 |

![time.png](/images/pathparam-routes/time.png)

[Graph - time](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1039813866&format=interactive)

### nsop
| nsop | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| --- | --- | --- | --- |
| goblin | 251.9 | 574 | 1013 |
| httprouter | 41.44 | 92.52 | 146.3 |
| chi | 200.8 | 309.1 | 455.1 |
| gin | 33.9 | 62.48 | 98.31 |
| bunrouter | 23.22 | 112.8 | 229.8 |
| httptreemux | 297.5 | 581.1 | 1160 |
| beegomux | 272.1 | 435.6 | 872.2 |
| gorillamux | 558.9 | 1127 | 2367 |
| bon | 154.5 | 217.1 | 300.1 |
| denco | 46.57 | 122.7 | 193.3 |
| echo | 26.49 | 78.45 | 144 |
| gocraftweb | 927.4 | 1232 | 1770 |
| gorouter | 188.4 | 278.1 | 382.5 |
| ozzorouting | 33.78 | 73.6 | 123.4 |
| techbook13-sample | 326.2 | 1008 | 1841 |

![nsop.png](/images/pathparam-routes/nsop.png)

[Graph - nsop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1534246873&format=interactive)

### bop
| bop | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| --- | --- | --- | --- |
| goblin | 328 | 412 | 494 |
| httprouter | 32 | 160 | 320 |
| chi | 304 | 304 | 304 |
| gin | 0 | 0 | 0 |
| bunrouter | 0 | 0 | 0 |
| httptreemux | 680 | 904 | 1742 |
| beegomux | 672 | 672 | 1254 |
| gorillamux | 1024 | 1088 | 1751 |
| bon | 304 | 304 | 304 |
| denco | 32 | 160 | 320 |
| echo | 0 | 0 | 0 |
| gocraftweb | 656 | 944 | 1862 |
| gorouter | 360 | 488 | 648 |
| ozzorouting | 0 | 0 | 0 |
| techbook13-sample | 432 | 968 | 1792 |

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

# License
This project is licensed under the terms of the MIT license.

## Author
- [@bmf-san](https://twitter.com/bmf_san)
- [bmf-tech.com](http://bmf-tech.com/)
