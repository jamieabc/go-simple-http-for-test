package handlers

import (
	"io"
	"net/http"
)

func New() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/hello", helloHandler)
	return m
}

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = io.WriteString(w, "hello, world!\n")
}
