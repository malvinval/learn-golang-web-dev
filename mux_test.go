package learngowebdev

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMux(t *testing.T) {
	mux := http.NewServeMux()

	// implementasi handler menggunakan anonymous function di parameter kedua
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // mux.HandleFunc
		fmt.Fprint(w, "/")
	})

	// implementasi handler menggunakan http.HandlerFunc yang dideklarasi terlebih dahulu
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "/about")
	}
	mux.Handle("/about", handler) // mux.Handle

	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
