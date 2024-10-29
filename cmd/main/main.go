package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"go-bookstore/pkg/routes"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
        path, err := route.GetPathTemplate()
        if err == nil {
            fmt.Println("Registered route:", path)
        }
        return nil
    })
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))
}