package handlers

import "net/http"

// ProfileHandler ...
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("profile"))
}
