package main

import (
	"io"
	"net/http"
)

type route struct {
	method string
	route  string
}

type RESTRouter struct {
	routes map[route]func(http.ResponseWriter, *http.Request)
}

func (r RESTRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	route := route{req.Method, req.URL.Path}
	f := r.routes[route]
	if f == nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "404 Not Found\n")
	} else {
		f(w, req)
	}
}

func (r RESTRouter) HandleFunc(method string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	route := route{method, pattern}
	r.routes[route] = handler
}

func NewRouter() RESTRouter {
	r := RESTRouter{}
	r.routes = make(map[route]func(http.ResponseWriter, *http.Request))
	return r
}

func route1GETHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "route1 GET\n")
}

func route1POSTHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "route1 POST\n")
}

func main() {
	r := NewRouter()
	r.HandleFunc("GET", "/route1", route1GETHandler)
	r.HandleFunc("POST", "/route1", route1POSTHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
