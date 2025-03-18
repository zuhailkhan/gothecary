package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	port := 3001

	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, Chi")
	})

	fmt.Printf("Server listening on Port: %v", port)
	http.ListenAndServe(":3001", r)
}
