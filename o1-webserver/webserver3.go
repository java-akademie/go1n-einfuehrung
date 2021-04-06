package main

import (
	"fmt"
	"log"
	"net/http"
)

func webserver3() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", pathAndQuery)

	log.Println("Starting server on :8083")
	err := http.ListenAndServe(":8083", mux)
	log.Fatal(err)
}

func pathAndQuery(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	query := r.URL.Query()
	fmt.Fprintln(w, "pathAndQuery")
	fmt.Fprintln(w, "path: ", path)
	fmt.Fprintln(w, "queries:")
	for k, v := range query {
		fmt.Fprintf(w, "\t Key: %v \t Value(s): %v \n", k, v)
	}
}
