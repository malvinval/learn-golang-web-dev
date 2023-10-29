package server_test

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	// deklarasi Server
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	// jalankan Server dengan ListenAndServe()

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
