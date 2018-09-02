package endpoints

import (
	"log"
	"net/http"

	"github.com/shunsukw/go-chat/common"
	"github.com/shunsukw/go-chat/common/authenticate"
)

func FetchPostEndpoints(env *common.Env) http.HandleFunc {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		gsSession, err := authenticate.SessionStore.Get(r, "gopherface-session")
		if err != nil {
			log.Print(err)
			return
		}

		uuid := gsSession.Values["uuid"].(string)

		posts, err := env.DB.FetchPosts(uuid)
		if err != nil {
			// Error 処理

		}
	})
}
