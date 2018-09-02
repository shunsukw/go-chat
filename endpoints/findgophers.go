package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shunsukw/go-chat/common"
	"github.com/shunsukw/go-chat/common/authenticate"
)

// FindGophers ...
func FindGophers(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gsSession, err := authenticate.SessionStore.Get(r, "gopherface-session")
		if err != nil {
			log.Print(err)
			return
		}

		uuid := gsSession.Values["uuid"].(string)
		var serchTerm string

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		err := json.Unmarshal(reqBody, &serchTerm)
		if err != nil {
			log.Print("Encountered error when attempting to unmarshal JSON: ", err)
		}

		gophers, err := env.DB.FindGophers(uuid, searchTerm)

		if err != nil {
			log.Print(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gophers)
	})
}
