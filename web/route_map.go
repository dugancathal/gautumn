package web

import "net/http"

type HandlerMap map[string]http.Handler

// Maps handlers to routes.
// May be called multiple times
//
// Often used like so:
// web.MapRoutes(web.HandlerMap{
// 	 "/my/path": controller.InjectedIndex()
// })
func MapRoutes(routeMap HandlerMap) http.Handler {
	return &Mux{routeMap: routeMap}
}
