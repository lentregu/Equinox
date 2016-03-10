package main

import (
	"net/http"

	"github.com/lentregu/Equinox/goops"
)

func main() {

	router := newRouter()

	goops.Fatal(http.ListenAndServe(":8080", router))
}
