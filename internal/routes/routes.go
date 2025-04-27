package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/data", apiDataHandler)

	return mux
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome to the homepage")
}

func apiDataHandler(writer http.ResponseWriter, request *http.Request) {
	data := "Some data from the API"
	fmt.Fprintln(writer, data)
}
