package main

import (
	"flag"
	"net/http"

	"log"
)

func main() {

	addr := flag.String("addr", ":9000", "The addr where the server is listening")
	flag.Parse()

	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./images"))))
	log.Printf("Listening in %s ....", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
