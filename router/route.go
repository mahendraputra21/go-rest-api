package router

import (
	"context"
	"go-rest-api/controller"
	"net/http"
	"regexp"
	"strings"
)

type ctxKey struct{}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

var routes = []route{
	newRoute("GET", "/product", controller.GETProductByIdController),
	newRoute("GET", "/product/brand", controller.GETProductByBrandId),
	newRoute("GET", "/order", controller.GetOrderDetailByIdControl),
	newRoute("POST", "/brand", controller.POSTBrandControl),
	newRoute("POST", "/product", controller.POSTProductController),
	newRoute("POST", "/order", controller.POSTOrderControl),
}

func Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}
