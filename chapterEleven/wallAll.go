package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

func (h notAllowedHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	handler(rw, r)
}

type notAllowedHandler struct{}

func main() {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(handler)
	notAllowed := notAllowedHandler{}

	r.MethodNotAllowedHandler = notAllowed

	// get methods

	getMux := r.Methods(http.MethodGet).Subrouter()
	getMux.HandleFunc("/time", handler)
	getMux.HandleFunc("/getall", handler)
	getMux.HandleFunc("/getid", handler)
	getMux.HandleFunc("/logged", handler)
	getMux.HandleFunc("/username/{id:[0-9]+}", handler)

	// put method
	putMux := r.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/update", handler)

	// post methods
	postMux := r.Methods(http.MethodPost).Subrouter()
	postMux.HandleFunc("/add", handler)
	postMux.HandleFunc("/login", handler)
	postMux.HandleFunc("/logout", handler)

	deleteMux := r.Methods(http.MethodDelete).Subrouter()
	deleteMux.HandleFunc("/username:{id:[0-9]+}", handler)

	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("Route:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		qT, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries template:", strings.Join(qT, ","))
		}
		qRegExps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries Regexs:", strings.Join(qRegExps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	http.Handle("/", r)
}
