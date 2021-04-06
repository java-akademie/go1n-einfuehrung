package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "<h1>Welcome</h1>")
	fmt.Fprintln(w, "<a href='/static/test1.html'>test1</a>")
	fmt.Fprintln(w, "<br>")
	fmt.Fprintln(w, "<a href='/static/test2.html'>test2</a>")
}

func webserver4() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	fs := http.FileServer(http.Dir("static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Starting server on :8084")
	err := http.ListenAndServe(":8084", mux)
	log.Fatal(err)
}
