package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"
)

func webserver2() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/plain", plainHandler)
	mux.HandleFunc("/html", htmlHandler)
	mux.HandleFunc("/json", jsonHandler)
	mux.HandleFunc("/xml", xmlHandler)

	log.Println("Starting server on :8082")
	err := http.ListenAndServe(":8082", mux)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Home")
	fmt.Fprintf(w, "%v \n", time.Now())
}

func plainHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "<h1>plainHandler</h1>")
	fmt.Fprintln(w, "first line<br>")
	fmt.Fprintln(w, "second line<br>")
	fmt.Fprintln(w, "third line<br>")
}

func htmlHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h1>htmlHandler</h1>")
	fmt.Fprintln(w, "first line<br>")
	fmt.Fprintln(w, "second line<br>")
	fmt.Fprintln(w, "third line<br>")
}

type profile struct {
	Name    string
	Hobbies []string
}

var p = profile{"Alex", []string{"snowboarding", "programming"}}

func jsonHandler(w http.ResponseWriter, _ *http.Request) {
	js, err := json.MarshalIndent(p, "", "  ")
	//js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func xmlHandler(w http.ResponseWriter, _ *http.Request) {
	x, err := xml.MarshalIndent(p, "", "  ")
	//x, err := xml.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
