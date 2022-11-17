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

# How to run benchmark test
`make test-benchmark`

## Commands for running benchmark test
`make test-benchmark`

## Results
```
go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/go-router-benchmark
cpu: Intel(R) Xeon(R) Platinum 8124M CPU @ 3.00GHz
BenchmarkStaticRoutesRootGoblin-36                	27648112	        43.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticRoutes1Goblin-36                   	 9083848	       128.1 ns/op	      16 B/op	       1 allocs/op
BenchmarkStaticRoutes5Goblin-36                   	 3331718	       353.7 ns/op	      80 B/op	       1 allocs/op
BenchmarkStaticRoutes10Goblin-36                  	 1874312	       577.7 ns/op	     160 B/op	       1 allocs/op
BenchmarkPathParamColonRoutes1Goblin-36           	 1241379	       991.9 ns/op	     409 B/op	       6 allocs/op
BenchmarkPathParamColonRoutes5Goblin-36           	  388048	      3391 ns/op	     966 B/op	      13 allocs/op
BenchmarkPathParamColonRoutes10Goblin-36          	  187884	      5993 ns/op	    1613 B/op	      19 allocs/op
BenchmarkStaticRoutesRootHTTPRouter-36            	90323350	        13.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticRoutes1HTTPRouter-36               	89596381	        13.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticRoutes5HTTPRouter-36               	86710161	        13.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticRoutes10HTTPRouter-36              	87476588	        13.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkPathParamColonRoutes1HTTPRouter-36       	15888105	        72.22 ns/op	      32 B/op	       1 allocs/op
BenchmarkPathParamColonRoutes5HTTPRouter-36       	 6374383	       203.8 ns/op	     160 B/op	       1 allocs/op
BenchmarkPathParamColonRoutes10HTTPRouter-36      	 3814080	       330.8 ns/op	     320 B/op	       1 allocs/op
BenchmarkStaticRoutesRootChi-36                   	 3499848	       332.7 ns/op	     304 B/op	       2 allocs/op
BenchmarkStaticRoutes1Chi-36                      	 3483327	       349.8 ns/op	     304 B/op	       2 allocs/op
BenchmarkStaticRoutes5Chi-36                      	 3650314	       329.0 ns/op	     304 B/op	       2 allocs/op
BenchmarkStaticRoutes10Chi-36                     	 3711537	       327.3 ns/op	     304 B/op	       2 allocs/op
BenchmarkPathParamBracketRoutes1Chi-36            	 3403273	       377.5 ns/op	     304 B/op	       2 allocs/op
BenchmarkPathParamBracketRoutes5Chi-36            	 1919821	       618.9 ns/op	     304 B/op	       2 allocs/op
BenchmarkPathParamBracketRoutes10Chi-36           	 1407243	       810.2 ns/op	     304 B/op	       2 allocs/op
BenchmarkStaticRoutesRootGin-36                   	24909394	        47.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticRoutes1Gin-36                      	25048821	        48.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticRoutes5Gin-36                      	24995527	        49.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticRoutes10Gin-36                     	21957733	        52.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkPathParamColonRoutes1Gin-36              	18214605	        58.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkPathParamColonRoutes5Gin-36              	10900795	       118.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkPathParamColonRoutes10Gin-36             	 6629413	       190.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticRoutesRootBunRouter-36             	54979360	        22.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticRoutes1BunRouter-36                	41570702	        27.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticRoutes5BunRouter-36                	41219000	        28.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticRoutes10BunRouter-36               	40128004	        31.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkPathParamColonRoutes1BunRouter-36        	35211588	        33.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkPathParamColonRoutes5BunRouter-36        	10227812	       117.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkPathParamColonRoutes10BunRouter-36       	 5063023	       248.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticRoutesRootHTTPTreeMux-36           	 4521781	       286.4 ns/op	     328 B/op	       3 allocs/op
BenchmarkStaticRoutes1HTTPTreeMux-36              	 4056072	       293.4 ns/op	     328 B/op	       3 allocs/op
BenchmarkStaticRoutes5HTTPTreeMux-36              	 3465921	       367.0 ns/op	     328 B/op	       3 allocs/op
BenchmarkStaticRoutes10HTTPTreeMux-36             	 2574319	       450.1 ns/op	     328 B/op	       3 allocs/op
BenchmarkPathParamColonRoutes1HTTPTreeMux-36      	 2202633	       540.6 ns/op	     680 B/op	       6 allocs/op
BenchmarkPathParamColonRoutes5HTTPTreeMux-36      	 1000000	      1106 ns/op	     904 B/op	       9 allocs/op
BenchmarkPathParamColonRoutes10HTTPTreeMux-36     	  572616	      2174 ns/op	    1742 B/op	      11 allocs/op
BenchmarkStaticRoutesRootBeegoMux-36              	13203564	        81.36 ns/op	      32 B/op	       1 allocs/op
BenchmarkStaticRoutes1BeegoMux-36                 	14425970	        96.34 ns/op	      32 B/op	       1 allocs/op
BenchmarkStaticRoutes5BeegoMux-36                 	 6515413	       172.7 ns/op	      32 B/op	       1 allocs/op
BenchmarkStaticRoutes10BeegoMux-36                	 4286169	       282.6 ns/op	      32 B/op	       1 allocs/op
BenchmarkPathParamColonRoutes1BeegoMux-36         	 2392867	       499.6 ns/op	     672 B/op	       5 allocs/op
BenchmarkPathParamColonRoutes5BeegoMux-36         	 1725937	       697.6 ns/op	     672 B/op	       5 allocs/op
BenchmarkPathParamColonRoutes10BeegoMux-36        	  710916	      1486 ns/op	    1254 B/op	       6 allocs/op
BenchmarkStaticRoutesRootGorillaMux-36            	 1636922	       729.3 ns/op	     721 B/op	       7 allocs/op
BenchmarkStaticRoutes1GorillaMux-36               	 1612135	       779.3 ns/op	     721 B/op	       7 allocs/op
BenchmarkStaticRoutes5GorillaMux-36               	 1408734	       881.1 ns/op	     721 B/op	       7 allocs/op
BenchmarkStaticRoutes10GorillaMux-36              	 1243224	       956.0 ns/op	     721 B/op	       7 allocs/op
BenchmarkPathParamBracketRoutes1GorillaMux-36     	 1000000	      1165 ns/op	    1026 B/op	       8 allocs/op
BenchmarkPathParamBracketRoutes5GorillaMux-36     	  477472	      2191 ns/op	    1090 B/op	       8 allocs/op
BenchmarkPathParamBracketRoutes10GorillaMux-36    	  287336	      4409 ns/op	    1754 B/op	       9 allocs/op
PASS
ok  	github.com/go-router-benchmark	82.889s
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