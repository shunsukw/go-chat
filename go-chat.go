package main

import (
	"log"
	"net/http"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/shunsukw/go-chat/common"
	"github.com/shunsukw/go-chat/common/datastore"
	"github.com/shunsukw/go-chat/endpoints"
	"github.com/shunsukw/go-chat/handlers"
	"github.com/shunsukw/go-chat/middleware"

	"go.isomorphicgo.org/go/isokit"
)

const (
	// WEBSERVERPORT ...
	WEBSERVERPORT = ":8443"
)

var WEBAPPROOT = "/Users/Shunsuke/dev/src/github.com/shunsukw/go-chat"

// TODO: slt / asyncq

func main() {
	db, err := datastore.NewDatastore(datastore.MYSQL, "gopherface:gopherface@/gopherfacedb")
	//db, err := datastore.NewDatastore(datastore.MONGODB, "localhost:27017")
	//db, err := datastore.NewDatastore(datastore.REDIS, "localhost:6379")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	env := common.Env{}
	isokit.TemplateFilesPath = WEBAPPROOT + "/templates"
	isokit.TemplateFileExtension = ".html"
	ts := isokit.NewTemplateSet()
	ts.GatherTemplates()
	env.TemplateSet = ts
	env.DB = db

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	// Signin Login
	r.Handle("/signup", handlers.SignupHandler(&env)).Methods("GET", "POST")
	r.Handle("/login", handlers.LoginHandler(&env)).Methods("GET", "POST")
	r.HandleFunc("/logut", handlers.LogoutHandler).Methods("GET", "POST")

	// routes	// REST API endpoints
	r.Handle("/restapi/get-user-profile", middleware.GatedContentHandler(endpoints.GetUserProfileEndpoint(&env))).Methods("GET", "POST")
	r.Handle("/restapi/save-user-profile", middleware.GatedContentHandler(endpoints.SaveUserProfileEndpoint(&env))).Methods("POST")
	r.Handle("/restapi/save-user-profile-image", middleware.GatedContentHandler(endpoints.SaveUserProfileImageEndpoint(&env))).Methods("POST")
	r.Handle("/restapi/find-gophers", middleware.GatedContentHandler(endpoints.FindGophersEndpoint(&env))).Methods("GET", "POST")
	r.Handle("/restapi/follow-gopher", middleware.GatedContentHandler(endpoints.FollowGopherEndpoint(&env))).Methods("GET", "POST")
	r.Handle("/restapi/unfollow-gopher", middleware.GatedContentHandler(endpoints.UnfollowGopherEndpoint(&env))).Methods("GET", "POST")
	r.Handle("/restapi/get-friends-list", middleware.GatedContentHandler(endpoints.FriendsListEndpoint(&env))).Methods("GET", "POST")
	r.Handle("/restapi/save-post", middleware.GatedContentHandler(endpoints.SavePostEndpoint(&env))).Methods("GET", "POST")
	r.Handle("/restapi/fetch-posts", middleware.GatedContentHandler(endpoints.FetchPostsEndpoint(&env))).Methods("GET", "POST")
	r.Handle("/restapi/get-gopher-profile", middleware.GatedContentHandler(endpoints.GetGopherProfileEndpoint(&env))).Methods("GET", "POST")

	r.Handle("/js/client.js", isokit.GopherjsScriptHandler(WEBAPPROOT))
	r.Handle("/js/client.js.map", isokit.GopherjsScriptMapHandler(WEBAPPROOT))
	r.Handle("/template-bundle", handlers.TemplateBundleHandler(&env))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	loggedRouter := ghandlers.LoggingHandler(os.Stdout, r)
	stdChain := alice.New(middleware.PanicRecoveryHandler)
	http.Handle("/", stdChain.Then(loggedRouter))

	http.ListenAndServe(WEBSERVERPORT, r)
}
