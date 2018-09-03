package handlers

import (
	"net/http"

	"github.com/shunsukw/go-chat/common"
	"github.com/shunsukw/go-chat/forms"
	"go.isomorphicgo.org/go/isokit"
)

// FeedHandler ...
func FeedHandler(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		formParams := isokit.FormParams{ResponseWriter: w, Request: r}
		p := forms.NewSocialMediaPostForm(&formParams)
		p.PageTitle = "Feed"
		env.TemplateSet.Render("feed_page", &isokit.RenderParams{Writer: w, Data: p})
	})
}
