package handlers

import (
	"log"
	"net/http"

	"github.com/shunsukw/go-chat/common"
	"github.com/shunsukw/go-chat/models"
	"github.com/shunsukw/go-chat/validationkit"
)

// SignupForm struct ...
type SignupForm struct {
	PageTitle  string
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

// SignupHandler ...
func SignupHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := SignupForm{}
		s.FieldNames = []string{"username", "firstName", "lastName", "email"}
		s.Fields = make(map[string]string)
		s.Errors = make(map[string]string)
		s.PageTitle = "Sign Up"

		switch r.Method {
		case "GET":
			DisplaySignupForm(w, r, &s)
		case "POST":
			ValidateSignupForm(w, r, &s, env)
		default:
			DisplaySignupForm(w, r, &s)
		}
	})
}

// DisplaySignupForm ...
func DisplaySignupForm(w http.ResponseWriter, r *http.Request, s *SignupForm) {
	RenderTemplate(w, WebAppRoot+"/templates/signupform.html", s)
}

// -------------------------------------------------

// ValidateSignupForm ...
func ValidateSignupForm(w http.ResponseWriter, r *http.Request, s *SignupForm, e *common.Env) {
	PopulateFormFields(r, s)

	if r.FormValue("username") == "" {
		s.Errors["usernameError"] = "The username field is required."
	}

	if r.FormValue("firstName") == "" {
		s.Errors["firstNameError"] = "The first name field is required."
	}

	if r.FormValue("lastName") == "" {
		s.Errors["lastNameError"] = "The last name field is required."
	}

	if r.FormValue("email") == "" {
		s.Errors["emailError"] = "The e-mail address field is required."
	}

	if r.FormValue("password") == "" {
		s.Errors["passwordError"] = "The password field is required."
	}

	if r.FormValue("confirmPassword") == "" {
		s.Errors["confirmPasswordError"] = "The confirm password field is required."
	}

	if validationkit.CheckUsernameSyntax(r.FormValue("username")) == false {
		usernameErrorMessage := "The username entered has an improper syntax."
		if _, ok := s.Errors["usernameError"]; ok {
			s.Errors["usernameError"] += " " + usernameErrorMessage
		} else {
			s.Errors["usernameError"] = usernameErrorMessage
		}
	}

	if validationkit.CheckEmailSyntax(r.FormValue("email")) == false {
		emailErrorMessage := "The e-mail address entered has an improper syntax."
		if _, ok := s.Errors["usernameError"]; ok {
			s.Errors["emailError"] += " " + emailErrorMessage
		} else {
			s.Errors["emailError"] = emailErrorMessage
		}
	}

	if r.FormValue("password") != r.FormValue("confirmPassword") {
		s.Errors["confirmPasswordError"] = "The password and confirm pasword fields do not match."
	}

	if len(s.Errors) > 0 {
		DisplaySignupForm(w, r, s)
	} else {
		ProcessSignupForm(w, r, s, e)
	}
}

// PopulateFormFields ...
func PopulateFormFields(r *http.Request, s *SignupForm) {
	for _, fieldname := range s.FieldNames {
		s.Fields[fieldname] = r.FormValue(fieldname)
	}
}

// ProcessSignupForm ...
func ProcessSignupForm(w http.ResponseWriter, r *http.Request, s *SignupForm, env *common.Env) {
	u := models.NewUser(r.FormValue("username"), r.FormValue("firstName"), r.FormValue("lastName"), r.FormValue("email"), r.FormValue("password"))
	err := env.DB.CreateUser(u)
	if err != nil {
		log.Print(err)
	}

	// Display form confirmation message
	DisplayConfirmation(w, r, s)
}

// --------------------------------------------------

// DisplayConfirmation ...
func DisplayConfirmation(w http.ResponseWriter, r *http.Request, s *SignupForm) {
	RenderTemplate(w, WebAppRoot+"/templates/signupconfirmation.html", s)
}
