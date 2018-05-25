REST server
-----------

the http.HandleFunc method takes a route and a callback as parameters, but does not make the distinction between the various http request methods (e.g., GET, POST, PUT, DELETE). We would like to provide a method such as:

HandleFunc(method string, pattern string, handler func())

The following small program does just this by defining a RESTRouter which implements the Handler interface and provides a HandleFunc method.

A map with a route (method and route itself) as a key and a callback function as a value is used to dispatch a request such as:

curl -X POST localhost:8000/route1

to the appropriate method such as:

func route1POSTHandler(w http.ResponseWriter, req *http.Request) {
        io.WriteString(w, "route1 POST\n")
}

