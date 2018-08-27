package handlers

import "net/http"

// FriendsHandler ...
func FriendsHandler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["PageTitle"] = "Friends"
	RenderGatedTemplate(w, "./templates/friends.html", m)
}
