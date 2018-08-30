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

// TODO: middleware / slt / asyncq

func main() {
	db, err := datastore.NewDatastore(datastore.MYSQL, "gochat:gochat@/gochatdb")
	//db, err := datastore.NewDatastore(datastore.MONGODB, "localhost:27017")
	//db, err := datastore.NewDatastore(datastore.REDIS, "localhost:6379")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	env := common.Env{DB: db}

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	// Signin Login
	r.Handle("/signin", handlers.SignupHandler(&env)).Methods("GET", "POST")
	r.Handle("/login", handlers.LoginHandler(&env)).Methods("GET", "POST")
	r.HandleFunc("/logut", handlers.LogoutHandler()).Methods("GET", "POST")

	// routes
	r.Handle("/feed", http.HandlerFunc(handlers.FeedHandler)).Methods("GET")
	r.Handle("/friends", http.HandlerFunc(handlers.FriendsHandler)).Methods("GET,POST")
	r.Handle("/find", http.HandlerFunc(handlers.FindHandler)).Methods("GET")
	r.Handle("/profile", http.HandlerFunc(handlers.MyProfileHandler)).Methods("GET")
	r.Handle("/profile/{username}", http.HandlerFunc(handlers.ProfileHandler)).Methods("GET")
	r.Handle("/postpreview", http.HandlerFunc(handlers.PostPreviewHandler)).Methods("GET", "POST")
	r.Handle("/upload-image", http.HandlerFunc(handlers.UploadImageHandler)).Methods("GET", "POST")
	r.Handle("/upload-video", http.HandlerFunc(handlers.UploadVideoHandler)).Methods("GET", "POST")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	loggedRouter := ghandlers.LoggingHandler(os.Stdout, r)
	stdChain := alice.New(middleware.PanicRecoveryHandler)
	http.Handle("/", stdChain.Then(loggedRouter))

	http.ListenAndServe(WEBSERVERPORT, r)
}
