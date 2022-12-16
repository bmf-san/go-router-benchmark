package main

import (
	"net/http"

	beegomux "github.com/beego/mux"
	"github.com/bmf-san/goblin"
	"github.com/dimfeld/httptreemux/v5"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	ozzorouting "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/gocraft/web"
	gorillamux "github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo"
	"github.com/naoina/denco"
	"github.com/nissy/bon"
	"github.com/uptrace/bunrouter"
	"github.com/vardius/gorouter/v4"
)

// route is a struct for route.
type route struct {
	path    string
	reqPath string
}

func loadServeMux(r route) http.Handler {
	router := http.NewServeMux()
	handler := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})
	router.HandleFunc(r.path, handler)
	return router
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

func loadDenco(r route) http.Handler {
	mux := denco.NewMux()
	h := func(w http.ResponseWriter, r *http.Request, _ denco.Params) {}
	handler, err := mux.Build([]denco.Handler{
		mux.GET(r.path, h),
	})
	if err != nil {
		panic(err)
	}
	return handler
}

func loadEcho(r route) http.Handler {
	e := echo.New()
	handler := func(_ echo.Context) error { return nil }
	e.GET(r.path, handler)

	return e
}

type GocraftWebContext struct{}

func (c *GocraftWebContext) gocraftwebHandler(_ web.ResponseWriter, _ *web.Request) {
}

func loadGocraftWeb(r route) http.Handler {
	router := web.New(GocraftWebContext{}).Get(r.path, (*GocraftWebContext).gocraftwebHandler)

	return router
}

func loadGorouter(r route) http.Handler {
	router := gorouter.New()
	handler := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})
	router.GET(r.path, handler)
	return router
}

func loadOzzoRouting(r route) http.Handler {
	router := ozzorouting.New()
	handler := func(_ *ozzorouting.Context) error { return nil }
	router.Get(r.path, handler)
	return router
}
