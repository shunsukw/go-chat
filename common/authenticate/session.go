package authenticate

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/shunsukw/go-chat/models"
)

// SessionStore ...
var SessionStore *sessions.FilesystemStore

// CreateUserSession ...
func CreateUserSession(u *models.User, sessionID string, w http.ResponseWriter, r *http.Request) error {
	gfSession, err := SessionStore.Get(r, "gopherface-session")
	if err != nil {
		log.Print(err)
	}

	gfSession.Values["sessionID"] = sessionID
	gfSession.Values["username"] = u.Username
	gfSession.Values["firstName"] = u.FirstName
	gfSession.Values["lastName"] = u.LastName
	gfSession.Values["email"] = u.Email

	err = gfSession.Save(r, w)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// ExpireUserSession ...
func ExpireUserSession(w http.ResponseWriter, r *http.Request) {
	gfSession, err := SessionStore.Get(r, "gopherface-session")

	if err != nil {
		log.Print(err)
	}

	gfSession.Options.MaxAge = -1
	gfSession.Save(r, w)
}

func init() {
	SessionStore = sessions.NewFilesystemStore("/tmp/gopherface-sessions", []byte(os.Getenv("GOPHERFACE_HASH_KEY")))
}
