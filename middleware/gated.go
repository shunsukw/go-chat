package middleware

import (
	"log"
	"net/http"

	"github.com/shunsukw/go-chat/common/authenticate"
)

// GatedContentHandler ...
func GatedContentHandler(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		shouldRedirectToLogin := false

		secureCookieMap, err := authenticate.ReadSecureCookieValues(w, r)
		if err != nil {
			log.Print(err)
		}

		if _, ok := secureCookieMap["sid"]; ok == true {
			gfSession, err := authenticate.SessionStore.Get(r, "gopherface-session")
			if err != nil {
				log.Print(err)
				return
			}

			if gfSession.Values["SessionID"] == secureCookieMap["sid"] && gfSession.Values["username"] == secureCookieMap["sid"] {
				next(w, r)
			} else {
				shouldRedirectToLogin = true
			}
		} else {
			shouldRedirectToLogin = true
		}

		if shouldRedirectToLogin == true {
			http.Redirect(w, r, "/login", 302)
		}
	})
}
