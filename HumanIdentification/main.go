package main

import (
	"Equinox/goops"
	"net/http"
)

var log goops.GoLogger

func init() {

	log = goops.New()
}

func main() {

	router := newRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
