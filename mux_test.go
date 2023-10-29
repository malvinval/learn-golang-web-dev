package learngowebdev

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMux(t *testing.T) {
	mux := http.NewServeMux()

	// implementasi handler secara langsung di parameter kedua
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // mux.HandleFunc
		fmt.Fprint(w, "/")
	})

	// implementasi handler yang dideklarasi terlebih dahulu (bertipe http.HandlerFunc)
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
