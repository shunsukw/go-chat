package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shunsukw/go-chat/common"
	"github.com/shunsukw/go-chat/common/authenticate"
)

// FetchPostsEndpoint ...
func FetchPostsEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gsSession, err := authenticate.SessionStore.Get(r, "gopherface-session")
		if err != nil {
			log.Print(err)
			return
		}

		uuid := gsSession.Values["uuid"].(string)

		posts, err := env.DB.FetchPosts(uuid)
		if err != nil {
			// Error 処理
			log.Print(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	})
}
