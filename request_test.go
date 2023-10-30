package learngowebdev

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequesr(t *testing.T) {
	// declare mux
	mux := http.NewServeMux()

	// memberikan informasi content length request
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.ContentLength)
	})

	// memberikan informasi request header
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Header)
	})

	addr := "127.0.0.1:8000"

	// declare server struct
	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// execute server
	err := server.ListenAndServe()

	// server execution error handling
	if err != nil {
		panic(err)
	}
}
