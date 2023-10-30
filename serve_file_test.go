package learngowebdev

import (
	"log"
	"net/http"
	"testing"
)

func TestServeFile(t *testing.T) {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: mux,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
