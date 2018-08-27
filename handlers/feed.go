package handlers

import "net/http"

// FeedHandler ...
func FeedHandler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["PageTitle"] = "Feed"
	RenderGatedTemplate(w, "./templates/feed.html", m)
}
