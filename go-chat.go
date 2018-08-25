package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shunsukw/go-chat/handlers"
)

const (
	WEBSERVERPORT = ":8443"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("")

	http.ListenAndServe(WEBSERVERPORT, r)
}
