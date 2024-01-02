package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {

	// func ParseFiles(filenames ...string) (*Template, error)
	tpl, _ = template.ParseFiles("index.html")

	// registers the handler function for the given pattern in the DefaultServeMux.
	http.HandleFunc("/hello", helloHandleFunc)
	http.HandleFunc("/about", aboutHandleFunc)
	http.HandleFunc("/", indexHandler)
	// ListenAndServe listens on the TCP network address addr and then calls Serve
	// with handler to handle requests on incoming connections
	http.ListenAndServe(":8080", nil) // entering nil implicitly uses DefaultServeMux
	// Serve mux is an HTTP request multiplexer. It matches the URL of each incoming
	// request against a list of registered patterns anc calls the handler for the
	// pattern that most closely matches the URL.
}

func helloHandleFunc(w http.ResponseWriter, r*http.Request) {
	//fmt.Fprint(w, "Hello, World!")
	fmt.Fprint(w, "r.URL.Path:, %s!", r.URL.Path)
}

func aboutHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to our Gopher powered website.")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}