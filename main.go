package main

import (
	"fmt"
	"html"
	"net/http"
)

func TestHandleFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!!\nURL path is %q", html.EscapeString(req.URL.Path))
}

func main() {
	// 1st
	http.HandleFunc("/Hello", TestHandleFunc)
	// 2nd
	http.ListenAndServe(":8080", nil)
}
