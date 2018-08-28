package authenticate

import (
	"log"
	"strings"

	"github.com/shunsukw/go-chat/common"
	"github.com/shunsukw/go-chat/common/utility"
)

// VerifyCredentials ...
func VerifyCredentials(env *common.Env, username string, password string) bool {
	u, err := env.DB.GetUser(username)
	if err != nil {
		log.Print(err)
	}

	if u == nil {
		return false
	}

	if strings.ToLower(username) == strings.ToLower(u.Username) && utility.SHA256OfString(password) == u.PasswordHash {
		log.Println("Successful login attempt from user: ", u.Username)
		return true
	} else {
		log.Println("Unsuccessful login attempt from user: ", u.Username)
		return false
	}
}
