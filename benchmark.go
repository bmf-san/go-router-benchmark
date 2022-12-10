package main

import (
	"net/http"

	beegomux "github.com/beego/mux"
	"github.com/bmf-san/goblin"
	"github.com/dimfeld/httptreemux/v5"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	gorillamux "github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"github.com/nissy/bon"
	"github.com/uptrace/bunrouter"
)

// route is a struct for route.
type route struct {
	path    string
	reqPath string
}

var staticRoutesRoot = route{
	"/", "/",
}

var staticRoutes1 = route{
	"/foo", "/foo",
}

var staticRoutes5 = route{
	"/foo/bar/baz/qux/quux", "/foo/bar/baz/qux/quux",
}

var staticRoutes10 = route{
	"/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred", "/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred",
}

var pathParamColonRoutes1 = route{
	"/foo/:bar", "/foo/bar",
}

var pathParamColonRoutes5 = route{
	"/foo/:bar/:baz/:qux/:quux/:corge", "/foo/bar/baz/qux/quux/corge",
}

var pathParamColonRoutes10 = route{
	"/foo/:bar/:baz/:qux/:quux/:corge/:grault/:garply/:waldo/:fred/:plugh", "/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred/plugh",
}

var pathParamBracketRoutes1 = route{
	"/foo/{bar}", "/foo/bar",
}

var pathParamBracketRoutes5 = route{
	"/foo/{bar}/{baz}/{qux}/{quux}/{corge}", "/foo/bar/baz/qux/quux/corge",
}

var pathParamBracketRoutes10 = route{
	"/foo/{bar}/{baz}/{qux}/{quux}/{corge}/{grault}/{garply}/{waldo}/{fred}/{plugh}", "/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred/plugh",
}

func loadGoblin(r route) http.Handler {
	router := goblin.NewRouter()
	handler := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})
	router.Methods(http.MethodGet).Handler(r.path, handler)
	return router
}

func loadHTTPRouter(r route) http.Handler {
	router := httprouter.New()
	handler := func(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {}
	router.GET(r.path, handler)
	return router
}

func loadChi(r route) http.Handler {
	router := chi.NewRouter()
	handler := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})
	router.Get(r.path, handler)
	return router
}

func loadGin(r route) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	handler := func(_ *gin.Context) {}
	router.GET(r.path, handler)
	return router
}

func loadBunRouter(r route) http.Handler {
	router := bunrouter.New()
	handler := func(_ http.ResponseWriter, _ bunrouter.Request) error { return nil }
	router.GET(r.path, handler)
	return router
}

func loadHTTPTreeMux(r route) http.Handler {
	router := httptreemux.NewContextMux()
	handler := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})
	router.GET(r.path, handler)
	return router
}

func loadBeegoMux(r route) http.Handler {
	router := beegomux.New()
	handler := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})
	router.Get(r.path, handler)
	return router
}

func loadGorillaMux(r route) http.Handler {
	router := gorillamux.NewRouter()
	handler := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})
	router.HandleFunc(r.path, handler)
	return router
}

func loadBon(r route) http.Handler {
	router := bon.NewRouter()
	handler := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})
	router.Get(r.path, handler)
	return router
}
