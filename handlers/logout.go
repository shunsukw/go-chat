package handlers

import (
	"net/http"

	"github.com/shunsukw/go-chat/common/authenticate"
)

// LogoutHandler ...
func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	authenticate.ExpireUserSession(w, r)
	authenticate.ExpireSecureCookie(w, r)
}
