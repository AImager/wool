package main

import (
	"io"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		_, _ = io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/test", helloHandler)
	_ = http.ListenAndServe(":8080", nil)
}
