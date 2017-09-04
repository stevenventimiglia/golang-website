package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	webServer()
}

type defaultContext struct {
	Title       string
	Section     string
	Year        int
	ErrorMsgs   string
	SuccessMsgs string
}

//var themeName = getThemeName()
var staticPages = populateStaticPages()

func serveContent(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	pageAlias := urlParams["pageAlias"]
	t := time.Now()

	if pageAlias == "" {
		pageAlias = "home"
	}

	context := defaultContext{}
	context.Title = strings.Title(pageAlias)
	context.Section = pageAlias
	context.Year = t.Year()
	context.ErrorMsgs = ""
	context.SuccessMsgs = ""

	staticPage := staticPages.Lookup(pageAlias + ".html")
	if staticPage == nil {
		context.Title = strings.Title("Whoops!")
		staticPage = staticPages.Lookup("404.html")
		w.WriteHeader(404)
	}

	staticPage.Execute(w, context)
}
