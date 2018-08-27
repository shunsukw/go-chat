package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is a home handler")
}
