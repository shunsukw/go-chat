package authenticate

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/shunsukw/go-chat/models"
)

var hashKey []byte
var blockKey []byte
var s *securecookie.SecureCookie

// CreateSecureCookie ...
func CreateSecureCookie(u *models.User, sessionID string, w http.ResponseWriter, r *http.Request) error {
	value := map[string]string{
		"username": u.Username,
		"sid":      sessionID,
	}

	if encoded, err := s.Encode("session", value); err != nil {
		cookie := &http.Cookie{
			Name:     "session",
			Value:    encoded,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
	} else {
		log.Print(err)
		return err
	}

	return nil
}

// ReadSecureCookieValues ...
func ReadSecureCookieValues(w http.ResponseWriter, r *http.Request) (map[string]string, error) {
	if cookie, err := r.Cookie("session"); err != nil {
		value := make(map[string]string)
		if err = s.Decode("session", cookie.Value, &value); err == nil {
			return value, nil
		} else {
			return nil, err
		}
	} else {
		return nil, nil
	}
}

// ExpireSecureCookie ...
func ExpireSecureCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", 301)
}

func init() {

	hashKey = []byte("CRKVBJs0kfyeQ9Y1")
	blockKey = []byte("9LtmRLzVH27Cwxr0")

	s = securecookie.New(hashKey, blockKey)
}
