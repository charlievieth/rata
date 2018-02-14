package rata

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/gorilla/mux"
)

// Supported HTTP methods.
const (
	GET     = "GET"
	HEAD    = "HEAD"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	CONNECT = "CONNECT"
	OPTIONS = "OPTIONS"
	TRACE   = "TRACE"
)

// Handlers map route names to http.Handler objects.  Each Handler key must
// match a route Name in the Routes collection.
type Handlers map[string]http.Handler

func convertVar(p string) string {
	a := strings.Split(path.Clean(p), "/")
	for i, s := range a {
		if strings.HasPrefix(s, ":") {
			a[i] = "{" + s[1:] + "}"
		}
	}
	return strings.Join(a, "/")
}

// NewRouter combines a set of Routes with their corresponding Handlers to
// produce a http request multiplexer (AKA a "router").  If any route does
// not have a matching handler, an error occurs.
func NewRouter(routes Routes, handlers Handlers) (http.Handler, error) {
	m := mux.NewRouter()
	for _, route := range routes {
		handler, ok := handlers[route.Name]
		if !ok {
			return nil, fmt.Errorf("missing handler %s", route.Name)
		}

		switch method := strings.ToUpper(route.Method); method {
		case GET, HEAD, POST, PUT, PATCH, DELETE, CONNECT, OPTIONS, TRACE:
			m.Handle(convertVar(route.Path), handler).Methods(method)
		default:
			return nil, fmt.Errorf("invalid verb: %s", route.Method)
		}
	}

	return m, nil
}
