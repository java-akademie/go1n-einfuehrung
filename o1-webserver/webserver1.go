package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func webserver1() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	log.Println("Starting server on :8081")
	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}

func handler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	len, err := w.Write([]byte("webserver1 - all path<br>\n"))
	_, _ = len, err
	len, err = fmt.Fprintf(w, "%v<br>\n", time.Now())
	_, _ = len, err
	// for writing, we can use
	// either w.Write or fmt.Fprintf, Fprint, Fprintln
	// len and err are not neccessary
	// for readability of the source text in Browser use Fprintln
	fmt.Fprint(w, "hello world<br>")
	fmt.Fprintln(w, "hello world<br>")
	w.Write([]byte("hello world<br>\n"))
}
