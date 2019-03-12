package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func Init(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "** default go page **", html.EscapeString(req.URL.Path))
}

func FirstHandleFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!!\nURL path is %q", html.EscapeString(req.URL.Path))
}
func SecondHandleFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "World!!\nURL path is %q", req.URL.Path)
}
func Foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Foo\nURL path is %q", req.URL.Path)
}
func Bar(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Bar\nURL path is %q", req.URL.Path)
}

func main() {
	// root
	// first arg is "pattern string" but forward match
	http.HandleFunc("/", Init)
	// 1st
	http.HandleFunc("/hello", FirstHandleFunc)
	// 2nd
	http.HandleFunc("/world", SecondHandleFunc)
	// foo
	http.HandleFunc("/foo", Foo)
	// foobar
	http.HandleFunc("/foo/bar", Bar)

	// listen
	las := http.ListenAndServe(":8080", nil)

	log.Fatal(las)
}
