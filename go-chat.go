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
	"github.com/shunsukw/go-chat/handlers"
	"github.com/shunsukw/go-chat/middleware"
)

const (
	// WEBSERVERPORT ...
	WEBSERVERPORT = ":8443"
)

var WEBAPPROOT = os.Getenv("GOPHERFACE_APP_ROOT")

// TODO: slt / asyncq

func main() {
	db, err := datastore.NewDatastore(datastore.MYSQL, "gochat:gochat@/gochatdb")
	//db, err := datastore.NewDatastore(datastore.MONGODB, "localhost:27017")
	//db, err := datastore.NewDatastore(datastore.REDIS, "localhost:6379")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	env := common.Env{}
	isokit.TemplateFilesPath = WebAppRoot + "/templates"
	isokit.TemplateFileExtension = ".html"
	ts := isokit.NewTemplateSet()
	ts.GatherTemplates()
	env.TemplateSet = ts
	env.DB = db

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	// Signin Login
	r.Handle("/signin", handlers.SignupHandler(&env)).Methods("GET", "POST")
	r.Handle("/login", handlers.LoginHandler(&env)).Methods("GET", "POST")
	r.HandleFunc("/logut", handlers.LogoutHandler).Methods("GET", "POST")

	// routes
	r.Handle("/feed", middleware.GatedContentHandler(handlers.FeedHandler)).Methods("GET")
	r.Handle("/friends", middleware.GatedContentHandler(handlers.FriendsHandler)).Methods("GET,POST")
	r.Handle("/myprofile", middleware.GatedContentHandler(handlers.MyProfileHandler)).Methods("GET")
	r.Handle("/profile/{username}", middleware.GatedContentHandler(handlers.ProfileHandler)).Methods("GET")

	// REST API endpoints

	r.Handle("/js/client.js", isokit.GopherjsScriptHandler(WebAppRoot))
	r.Handle("/js/client.js.map", isokit.GopherjsScriptMapHandler(WebAppRoot))
	r.Handle("/template-bundle", handlers.TemplateBundleHandler(&env))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	loggedRouter := ghandlers.LoggingHandler(os.Stdout, r)
	stdChain := alice.New(middleware.PanicRecoveryHandler)
	http.Handle("/", stdChain.Then(loggedRouter))

	http.ListenAndServe(WEBSERVERPORT, r)
}
