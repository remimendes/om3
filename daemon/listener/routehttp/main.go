/*
Package routehttp provides http mux

It defines routing for Opensvc listener daemons
*/
package routehttp

import (
	"context"
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"

	"github.com/opensvc/om3/daemon/api"
	"github.com/opensvc/om3/daemon/daemonapi"
)

type (
	T struct {
		mux *echo.Echo
	}
)

var (
	mwProm = echoprometheus.NewMiddleware("opensvc_api")
)

// New returns *T with log, rootDaemon
// it prepares middlewares and routes for Opensvc daemon listeners
// when enableUi is true swagger-ui is serverd from /ui
func New(ctx context.Context, enableUi bool) *T {
	e := echo.New()
	pprof.Register(e)
	e.Use(mwProm)
	e.GET("/metrics", echoprometheus.NewHandler())
	e.Use(daemonapi.LogMiddleware(ctx))
	e.Use(daemonapi.AuthMiddleware(ctx))
	e.Use(daemonapi.LogUserMiddleware(ctx))
	e.Use(daemonapi.LogRequestMiddleWare(ctx))
	api.RegisterHandlers(e, daemonapi.New(ctx))
	g := e.Group("/public/ui")
	if enableUi {
		g.Use(daemonapi.UiMiddleware(ctx))
	}

	return &T{mux: e}
}

// ServerHTTP implement http.Handler interface for T
func (t *T) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.mux.ServeHTTP(w, r)
}
