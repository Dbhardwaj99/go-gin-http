package main

// import (
// 	"log"
// 	"net/http"

// 	"golang.org/x/oauth2/google/controller"
// 	"golang.org/x/oauth2/google"
// )

// func mainTest() {
// 	// load configs
// 	config.LoadEnv()
// 	config.LoadConfig()

// 	// create a router
// 	mux := http.NewServeMux()

// 	// define routes
// 	mux.HandleFunc("/google_login", controller.GoogleLogin)
// 	mux.HandleFunc("/google_callback", controller.GoogleCallback)
// 	mux.HandleFunc("/fb_login", controller.FbLogin)
// 	mux.HandleFunc("/fb_callback", controller.FbCallback)

// 	// run server
// 	log.Println("started server on :: http://localhost:8080/")
// 	if oops := http.ListenAndServe(":8080", mux); oops != nil {
// 		log.Fatal(oops)
// 	}
// }
