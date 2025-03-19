package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	port := 3001

	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello from server")
	})

	r.Get("/hello/{name}", func(w http.ResponseWriter, req *http.Request) {
		name := req.URL.Path[7:]
		fmt.Fprintf(w, "name: %v", name)
	})

	r.Post("/helloPost", func(w http.ResponseWriter, req *http.Request) {
		data, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Bad Request", 400)
			return
		}

		var jsonData map[string]string
		err = json.Unmarshal(data, &jsonData)

		if err != nil {
			http.Error(w, "Bad JSON; Bad Request", 400)
			return
		}

		fmt.Printf("Recieved Data %+v\n", jsonData)

		fmt.Fprintf(w, "Data Recieved Successfully")
	})

	fmt.Printf("Server listening on Port: %v", port)
	http.ListenAndServe(":3001", r)
}
