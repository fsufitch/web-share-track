package webtrack

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func InitRouter() {
	Router = mux.NewRouter().StrictSlash(false)
}

func AddRoute(name string, method string, url string, handler http.HandlerFunc) {
	Router.Methods(method).Path(url).Handler(handler).Name(name)
}
