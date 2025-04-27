package main

import (
	"fmt"
	"net/http"

	"github.com/mfrszpiotro/structure/internal/routes"
)

func main() {
	router := routes.NewRouter()
	server_address := ":8080"
	fmt.Printf("Server listening on http://localhost%s\n", server_address)
	err := http.ListenAndServe(server_address, router)
	if err != nil {
		panic(err)
	}
}
