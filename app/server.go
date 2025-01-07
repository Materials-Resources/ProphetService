package app

import (
	"connectrpc.com/grpcreflect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

func (a *App) RegisterService(path string, handler http.Handler, serviceName string) {
	a.mux.Handle(path, handler)
	a.reflectionRoutes = append(a.reflectionRoutes, serviceName)
}

func (a *App) serveHttp() error {
	return http.ListenAndServe(":8082", h2c.NewHandler(a.mux, &http2.Server{}))
}

func (a *App) registerReflection() {
	reflector := grpcreflect.NewStaticReflector(a.reflectionRoutes...)

	a.mux.Handle(grpcreflect.NewHandlerV1(reflector))
	a.mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

}
