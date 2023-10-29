package learngowebdev

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	// deklarasi variabel handler bertipe http.HandlerFunc
	// `HandlerFunc` hanyalah sebuah user-defined data type yang sudah mengimplementasikan interface `Handler`.
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// w: response untuk client
		// r: request dari client
		fmt.Fprint(w, "Hello world!")
	}

	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
