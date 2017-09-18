package web

import "net/http"

type RouteMap map[string]http.Handler

// Maps handlers to routes.
// May be called multiple times
//
// Often used like so:
// web.MapRoutes(web.RouteMap{
// 	 "/my/path": controller.InjectedIndex()
// })
func MapRoutes(routeMap RouteMap) {
	for route, handler := range routeMap {
		http.Handle(route, handler)
	}
}
