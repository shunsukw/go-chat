package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shunsukw/go-chat/common"
	"github.com/shunsukw/go-chat/common/authenticate"
	"github.com/shunsukw/go-chat/models"
)

// SaveUserProfileEndpoint ...
func SaveUserProfileEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gfSession, err := authenticate.SessionStore.Get(r, "gopherface-session")
		if err != nil {
			log.Print(err)
			return
		}
		uuid := gfSession.Values["uuid"]

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		var u models.UserProfile
		json.Unmarshal(reqBody, &u)

		env.DB.UpdateUserProfile(uuid.(string), u.About, u.Location, u.Interests)
	})
}
