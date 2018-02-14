package rata

import (
	"net/http"

	"github.com/gorilla/mux"
)

//  Param returns the parameter with the given name from the given request.
func Param(req *http.Request, paramName string) string {
	if v := mux.Vars(req); v != nil {
		return v[paramName]
	}
	return ""
}
