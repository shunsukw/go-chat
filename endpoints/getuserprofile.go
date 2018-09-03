package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shunsukw/go-chat/common"
	"github.com/shunsukw/go-chat/common/authenticate"
)

// GetUserProfileEndpoint ...
func GetUserProfileEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gfSession, err := authenticate.SessionStore.Get(r, "gopherface-session")
		if err != nil {
			log.Print(err)
			return
		}

		uuid := gfSession.Values["uuid"].(string)
		u, err := env.DB.GetUserProfile(uuid)
		if err != nil {
			log.Print(err)
		}

		u.Username = gfSession.Values["username"].(string)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
	})
}
