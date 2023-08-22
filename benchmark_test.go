package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testServeHTTP(b *testing.B, r route, router http.Handler) {
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, r.reqPath, nil)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		router.ServeHTTP(rec, req)
		if rec.Code != 200 {
			panic(fmt.Sprintf("Request failed. path: %v request path:%v", r.path, r.reqPath))
		}
	}
}

func benchmark(b *testing.B, r route, router http.Handler) {
	testServeHTTP(b, r, router)
}

// net/http#ServeMux
// https://pkg.go.dev/net/http#ServeMux
func BenchmarkStaticRoutesRootServeMux(b *testing.B) {
	router := loadServeMux(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1ServeMux(b *testing.B) {
	router := loadServeMux(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5ServeMux(b *testing.B) {
	router := loadServeMux(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10ServeMux(b *testing.B) {
	router := loadServeMux(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

// bmf-san/goblin
// https://github.com/bmf-san/goblin
func BenchmarkStaticRoutesRootGoblin(b *testing.B) {
	router := loadGoblin(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1Goblin(b *testing.B) {
	router := loadGoblin(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5Goblin(b *testing.B) {
	router := loadGoblin(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10Goblin(b *testing.B) {
	router := loadGoblin(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1ColonGoblin(b *testing.B) {
	router := loadGoblin(pathParamRoutes1Colon)
	benchmark(b, pathParamRoutes1Colon, router)
}

func BenchmarkPathParamRoutes5ColonGoblin(b *testing.B) {
	router := loadGoblin(pathParamRoutes5Colon)
	benchmark(b, pathParamRoutes5Colon, router)
}

func BenchmarkPathParamRoutes10ColonGoblin(b *testing.B) {
	router := loadGoblin(pathParamRoutes10Colon)
	benchmark(b, pathParamRoutes10Colon, router)
}

// julienschmidt/httprouter
// https://github.com/julienschmidt/httprouter
func BenchmarkStaticRoutesRootHTTPRouter(b *testing.B) {
	router := loadHTTPRouter(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1HTTPRouter(b *testing.B) {
	router := loadHTTPRouter(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5HTTPRouter(b *testing.B) {
	router := loadHTTPRouter(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10HTTPRouter(b *testing.B) {
	router := loadHTTPRouter(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1ColonHTTPRouter(b *testing.B) {
	router := loadHTTPRouter(pathParamRoutes1Colon)
	benchmark(b, pathParamRoutes1Colon, router)
}

func BenchmarkPathParamRoutes5ColonHTTPRouter(b *testing.B) {
	router := loadHTTPRouter(pathParamRoutes5Colon)
	benchmark(b, pathParamRoutes5Colon, router)
}

func BenchmarkPathParamRoutes10ColonHTTPRouter(b *testing.B) {
	router := loadHTTPRouter(pathParamRoutes10Colon)
	benchmark(b, pathParamRoutes10Colon, router)
}

// go-chi/chi
// https://github.com/go-chi/chi
func BenchmarkStaticRoutesRootChi(b *testing.B) {
	router := loadChi(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1Chi(b *testing.B) {
	router := loadChi(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5Chi(b *testing.B) {
	router := loadChi(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10Chi(b *testing.B) {
	router := loadChi(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1BracketChi(b *testing.B) {
	router := loadChi(pathParamRoutes1Bracket)
	benchmark(b, pathParamRoutes1Bracket, router)
}

func BenchmarkPathParamRoutes5BracketChi(b *testing.B) {
	router := loadChi(pathParamRoutes5Bracket)
	benchmark(b, pathParamRoutes5Bracket, router)
}

func BenchmarkPathParamRoutes10BracketChi(b *testing.B) {
	router := loadChi(pathParamRoutes10Bracket)
	benchmark(b, pathParamRoutes10Bracket, router)
}

// gin-gonic/gin
// https://github.com/gin-gonic/gin
func BenchmarkStaticRoutesRootGin(b *testing.B) {
	router := loadGin(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1Gin(b *testing.B) {
	router := loadGin(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5Gin(b *testing.B) {
	router := loadGin(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10Gin(b *testing.B) {
	router := loadGin(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1ColonGin(b *testing.B) {
	router := loadGin(pathParamRoutes1Colon)
	benchmark(b, pathParamRoutes1Colon, router)
}

func BenchmarkPathParamRoutes5ColonGin(b *testing.B) {
	router := loadGin(pathParamRoutes5Colon)
	benchmark(b, pathParamRoutes5Colon, router)
}

func BenchmarkPathParamRoutes10ColonGin(b *testing.B) {
	router := loadGin(pathParamRoutes10Colon)
	benchmark(b, pathParamRoutes10Colon, router)
}

// uptrace/bunrouter
// https://github.com/uptrace/bunrouter
func BenchmarkStaticRoutesRootBunRouter(b *testing.B) {
	router := loadBunRouter(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1BunRouter(b *testing.B) {
	router := loadBunRouter(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5BunRouter(b *testing.B) {
	router := loadBunRouter(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10BunRouter(b *testing.B) {
	router := loadBunRouter(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1ColonBunRouter(b *testing.B) {
	router := loadBunRouter(pathParamRoutes1Colon)
	benchmark(b, pathParamRoutes1Colon, router)
}

func BenchmarkPathParamRoutes5ColonBunRouter(b *testing.B) {
	router := loadBunRouter(pathParamRoutes5Colon)
	benchmark(b, pathParamRoutes5Colon, router)
}

func BenchmarkPathParamRoutes10ColonBunRouter(b *testing.B) {
	router := loadBunRouter(pathParamRoutes10Colon)
	benchmark(b, pathParamRoutes10Colon, router)
}

// dimfeld/httptreemux
// https://github.com/dimfeld/httptreemux
func BenchmarkStaticRoutesRootHTTPTreeMux(b *testing.B) {
	router := loadHTTPTreeMux(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1HTTPTreeMux(b *testing.B) {
	router := loadHTTPTreeMux(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5HTTPTreeMux(b *testing.B) {
	router := loadHTTPTreeMux(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10HTTPTreeMux(b *testing.B) {
	router := loadHTTPTreeMux(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1ColonHTTPTreeMux(b *testing.B) {
	router := loadHTTPTreeMux(pathParamRoutes1Colon)
	benchmark(b, pathParamRoutes1Colon, router)
}

func BenchmarkPathParamRoutes5ColonHTTPTreeMux(b *testing.B) {
	router := loadHTTPTreeMux(pathParamRoutes5Colon)
	benchmark(b, pathParamRoutes5Colon, router)
}

func BenchmarkPathParamRoutes10ColonHTTPTreeMux(b *testing.B) {
	router := loadHTTPTreeMux(pathParamRoutes10Colon)
	benchmark(b, pathParamRoutes10Colon, router)
}

// beego/mux
// https://github.com/beego/mux
func BenchmarkStaticRoutesRootBeegoMux(b *testing.B) {
	router := loadBeegoMux(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1BeegoMux(b *testing.B) {
	router := loadBeegoMux(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5BeegoMux(b *testing.B) {
	router := loadBeegoMux(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10BeegoMux(b *testing.B) {
	router := loadBeegoMux(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1ColonBeegoMux(b *testing.B) {
	router := loadBeegoMux(pathParamRoutes1Colon)
	benchmark(b, pathParamRoutes1Colon, router)
}

func BenchmarkPathParamRoutes5ColonBeegoMux(b *testing.B) {
	router := loadBeegoMux(pathParamRoutes5Colon)
	benchmark(b, pathParamRoutes5Colon, router)
}

func BenchmarkPathParamRoutes10ColonBeegoMux(b *testing.B) {
	router := loadBeegoMux(pathParamRoutes10Colon)
	benchmark(b, pathParamRoutes10Colon, router)
}

// gorilla/mux
// https://github.com/gorilla/mux
func BenchmarkStaticRoutesRootGorillaMux(b *testing.B) {
	router := loadGorillaMux(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1GorillaMux(b *testing.B) {
	router := loadGorillaMux(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5GorillaMux(b *testing.B) {
	router := loadGorillaMux(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10GorillaMux(b *testing.B) {
	router := loadGorillaMux(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1BracketGorillaMux(b *testing.B) {
	router := loadGorillaMux(pathParamRoutes1Bracket)
	benchmark(b, pathParamRoutes1Bracket, router)
}

func BenchmarkPathParamRoutes5BracketGorillaMux(b *testing.B) {
	router := loadGorillaMux(pathParamRoutes5Bracket)
	benchmark(b, pathParamRoutes5Bracket, router)
}

func BenchmarkPathParamRoutes10BracketGorillaMux(b *testing.B) {
	router := loadGorillaMux(pathParamRoutes10Bracket)
	benchmark(b, pathParamRoutes10Bracket, router)
}

// nissy/bon
// https://github.com/nissy/bon
func BenchmarkStaticRoutesRootBon(b *testing.B) {
	router := loadBon(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1Bon(b *testing.B) {
	router := loadBon(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5Bon(b *testing.B) {
	router := loadBon(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10Bon(b *testing.B) {
	router := loadBon(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1ColonBon(b *testing.B) {
	router := loadBon(pathParamRoutes1Colon)
	benchmark(b, pathParamRoutes1Colon, router)
}

func BenchmarkPathParamRoutes5ColonBon(b *testing.B) {
	router := loadBon(pathParamRoutes5Colon)
	benchmark(b, pathParamRoutes5Colon, router)
}

func BenchmarkPathParamRoutes10ColonBon(b *testing.B) {
	router := loadBon(pathParamRoutes10Colon)
	benchmark(b, pathParamRoutes10Colon, router)
}

// naoina/denco
// https://github.com/naoina/denco
func BenchmarkStaticRoutesRootDenco(b *testing.B) {
	router := loadDenco(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1Denco(b *testing.B) {
	router := loadDenco(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5Denco(b *testing.B) {
	router := loadDenco(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10Denco(b *testing.B) {
	router := loadDenco(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1ColonDenco(b *testing.B) {
	router := loadDenco(pathParamRoutes1Colon)
	benchmark(b, pathParamRoutes1Colon, router)
}

func BenchmarkPathParamRoutes5ColonDenco(b *testing.B) {
	router := loadDenco(pathParamRoutes5Colon)
	benchmark(b, pathParamRoutes5Colon, router)
}

func BenchmarkPathParamRoutes10ColonDenco(b *testing.B) {
	router := loadDenco(pathParamRoutes10Colon)
	benchmark(b, pathParamRoutes10Colon, router)
}

// labstack/echo
// https://github.com/labstack/echo
func BenchmarkStaticRoutesRootEcho(b *testing.B) {
	router := loadEcho(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1Echo(b *testing.B) {
	router := loadEcho(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5Echo(b *testing.B) {
	router := loadEcho(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10Echo(b *testing.B) {
	router := loadEcho(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1ColonEcho(b *testing.B) {
	router := loadEcho(pathParamRoutes1Colon)
	benchmark(b, pathParamRoutes1Colon, router)
}

func BenchmarkPathParamRoutes5ColonEcho(b *testing.B) {
	router := loadEcho(pathParamRoutes5Colon)
	benchmark(b, pathParamRoutes5Colon, router)
}

func BenchmarkPathParamRoutes10ColonEcho(b *testing.B) {
	router := loadEcho(pathParamRoutes10Colon)
	benchmark(b, pathParamRoutes10Colon, router)
}

// gocraft/web
// https://github.com/gocraft/web
func BenchmarkStaticRoutesRootGocraftWeb(b *testing.B) {
	router := loadGocraftWeb(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1GocraftWeb(b *testing.B) {
	router := loadGocraftWeb(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5GocraftWeb(b *testing.B) {
	router := loadGocraftWeb(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10GocraftWeb(b *testing.B) {
	router := loadGocraftWeb(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1ColonGocraftWeb(b *testing.B) {
	router := loadGocraftWeb(pathParamRoutes1Colon)
	benchmark(b, pathParamRoutes1Colon, router)
}

func BenchmarkPathParamRoutes5ColonGocraftWeb(b *testing.B) {
	router := loadGocraftWeb(pathParamRoutes5Colon)
	benchmark(b, pathParamRoutes5Colon, router)
}

func BenchmarkPathParamRoutes10ColonGocraftWeb(b *testing.B) {
	router := loadGocraftWeb(pathParamRoutes10Colon)
	benchmark(b, pathParamRoutes10Colon, router)
}

// vardius/gorouter
// https://github.com/vardius/gorouter
func BenchmarkStaticRoutesRootGorouter(b *testing.B) {
	router := loadGorouter(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1Gorouter(b *testing.B) {
	router := loadGorouter(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5Gorouter(b *testing.B) {
	router := loadGorouter(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10Gorouter(b *testing.B) {
	router := loadGorouter(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1BracketGorouter(b *testing.B) {
	router := loadGorouter(pathParamRoutes1Bracket)
	benchmark(b, pathParamRoutes1Bracket, router)
}

func BenchmarkPathParamRoutes5BracketGorouter(b *testing.B) {
	router := loadGorouter(pathParamRoutes5Bracket)
	benchmark(b, pathParamRoutes5Bracket, router)
}

func BenchmarkPathParamRoutes10BracketGorouter(b *testing.B) {
	router := loadGorouter(pathParamRoutes10Bracket)
	benchmark(b, pathParamRoutes10Bracket, router)
}

// go-ozzo/ozzo-routing
// https://github.com/go-ozzo/ozzo-routing
func BenchmarkStaticRoutesRootOzzoRouting(b *testing.B) {
	router := loadOzzoRouting(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1OzzoRouting(b *testing.B) {
	router := loadOzzoRouting(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5OzzoRouting(b *testing.B) {
	router := loadOzzoRouting(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10OzzoRouting(b *testing.B) {
	router := loadOzzoRouting(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1InequalitySignOzzoRouting(b *testing.B) {
	router := loadOzzoRouting(pathParamRoutes1InequalitySign)
	benchmark(b, pathParamRoutes1InequalitySign, router)
}

func BenchmarkPathParamRoutes5InequalitySignOzzoRouting(b *testing.B) {
	router := loadOzzoRouting(pathParamRoutes5InequalitySign)
	benchmark(b, pathParamRoutes5InequalitySign, router)
}

func BenchmarkPathParamRoutes10InequalitySignOzzoRouting(b *testing.B) {
	router := loadOzzoRouting(pathParamRoutes10InequalitySign)
	benchmark(b, pathParamRoutes10InequalitySign, router)
}

// n9te9 router
// https://github.com/lkeix/techbook13-sample
func BenchmarkStaticRoutesRootN9tE9Routing(b *testing.B) {
	router := loadN9tE9Routing(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1ON9tE9Routing(b *testing.B) {
	router := loadN9tE9Routing(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5ON9tE9Routing(b *testing.B) {
	router := loadN9tE9Routing(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10N9tE9Routing(b *testing.B) {
	router := loadN9tE9Routing(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1ColonN9tE9Routing(b *testing.B) {
	router := loadN9tE9Routing(pathParamRoutes1Colon)
	benchmark(b, pathParamRoutes1Colon, router)
}

func BenchmarkPathParamRoutes5ColonN9tE9Routing(b *testing.B) {
	router := loadN9tE9Routing(pathParamRoutes5Colon)
	benchmark(b, pathParamRoutes5Colon, router)
}

func BenchmarkPathParamRoutes10ColonN9tE9Routing(b *testing.B) {
	router := loadN9tE9Routing(pathParamRoutes10Colon)
	benchmark(b, pathParamRoutes10Colon, router)
}

// muxpatterns
// https://github.com/jba/muxpatterns
func BenchmarkStaticRoutesRootMuxPatterns(b *testing.B) {
	router := loadMuxPatterns(staticRoutesRoot)
	benchmark(b, staticRoutesRoot, router)
}

func BenchmarkStaticRoutes1OMuxPatterns(b *testing.B) {
	router := loadMuxPatterns(staticRoutes1)
	benchmark(b, staticRoutes1, router)
}

func BenchmarkStaticRoutes5OMuxPatterns(b *testing.B) {
	router := loadMuxPatterns(staticRoutes5)
	benchmark(b, staticRoutes5, router)
}

func BenchmarkStaticRoutes10MuxPatterns(b *testing.B) {
	router := loadMuxPatterns(staticRoutes10)
	benchmark(b, staticRoutes10, router)
}

func BenchmarkPathParamRoutes1BracketMuxPatterns(b *testing.B) {
	router := loadMuxPatterns(pathParamRoutes1Bracket)
	benchmark(b, pathParamRoutes1Bracket, router)
}

func BenchmarkPathParamRoutes5BracketMuxPatterns(b *testing.B) {
	router := loadMuxPatterns(pathParamRoutes5Bracket)
	benchmark(b, pathParamRoutes5Bracket, router)
}

func BenchmarkPathParamRoutes10BracketMuxPatterns(b *testing.B) {
	router := loadMuxPatterns(pathParamRoutes10Bracket)
	benchmark(b, pathParamRoutes10Bracket, router)
}
