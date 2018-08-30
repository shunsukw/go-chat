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
	})
}
