package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// WebServer - Server Function w/ Routers
func webServer() {
	route := mux.NewRouter()

	route.HandleFunc("/", serveContent)
	route.HandleFunc("/{pageAlias}", serveContent) // Dynamic URL

	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/js/", serveResource)
	http.HandleFunc("/img/", serveResource)

	http.Handle("/", route)

	portNumber := "8088"
	fmt.Println("------------------------------------------------------------")
	log.Println("Server started at http://localhost:" + portNumber)
	fmt.Println("------------------------------------------------------------")
	http.ListenAndServe(":"+portNumber, nil)

}
