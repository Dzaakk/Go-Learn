package go_web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "Localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
