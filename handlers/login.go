package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shunsukw/go-chat/common/utility"

	"github.com/shunsukw/go-chat/common"
	"github.com/shunsukw/go-chat/common/authenticate"
	"github.com/shunsukw/go-chat/validationkit"
)

type LoginForm struct {
	PageTitle  string
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

func LoginHandler(e *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := LoginForm{}
		l.FieldNames = []string{"username"}
		l.Fields = make(map[string]string)
		l.Errors = make(map[string]string)
		l.PageTitle = "Log In"

		switch r.Method {
		case "GET":
			DisplayLoginForm(w, r, &l)
		case "POST":
			ValidateLoginForm(w, r, &l, e)
		default:
			DisplayLoginForm(w, r, &l)
		}
	})
}

// DisplayLoginForm ...
func DisplayLoginForm(w http.ResponseWriter, r *http.Request, l *LoginForm) {
	RenderTemplate(w, "./templates/loginform.html", l)
}

// -------------------------------------

// ValidateLoginForm ...
func ValidateLoginForm(w http.ResponseWriter, r *http.Request, l *LoginForm, e *common.Env) {
	PopulateLoginFormFields(r, l)

	if r.FormValue("username") == "" {
		l.Errors["usernameError"] = "The username field is required."
	}

	if r.FormValue("password") == "" {
		l.Errors["passwordError"] = "The password field is required."
	}

	if validationkit.CheckUsernameSyntax(r.FormValue("username")) == false {

		usernameErrorMessage := "The username entered has an improper syntax."
		if _, ok := l.Errors["usernameError"]; ok {
			l.Errors["usernameError"] += " " + usernameErrorMessage
		} else {
			l.Errors["usernameError"] = usernameErrorMessage
		}
	}

	if len(l.Errors) > 0 {
		DisplayLoginForm(w, r, l)
	} else {
		ProcessLoginForm(w, r, l, e)
	}
}

// PopulateLoginFormFields ...
func PopulateLoginFormFields(r *http.Request, l *LoginForm) {
	for _, fieldname := range l.FieldNames {
		l.Fields[fieldname] = r.FormValue(fieldname)
	}
}

// ProcessLoginForm ...
func ProcessLoginForm(w http.ResponseWriter, r *http.Request, l *LoginForm, e *common.Env) {
	authResult := authenticate.VerifyCredentials(e, r.FormValue("username"), r.FormValue("password"))
	fmt.Println("auth result: ", authResult)

	if authResult == true {
		sessionID := utility.GenerateUUID()
		fmt.Println("sessid: ", sessionID)

		u, err := e.DB.GetUser(r.FormValue("username"))
		if err != nil {
			log.Print("Encountered error when attempting to fetch user record: ", err)
			http.Redirect(w, r, "/login", 302)
			return
		}

		err = authenticate.CreateSecureCookie(u, sessionID, w, r)
		if err != nil {
			log.Print("Encountered error when attempting to create secure cookie: ", err)
			http.Redirect(w, r, "/login", 302)
			return
		}

		err = authenticate.CreateUserSession(u, sessionID, w, r)
		if err != nil {
			log.Print("Encountered error when attempting to create secure cookie: ", err)
			http.Redirect(w, r, "/login", 302)
			return
		}

		http.Redirect(w, r, "/feed", 302)
	} else {
		l.Errors["usernameError"] = "Invalid login."
		DisplayLoginForm(w, r, l)
	}
}
