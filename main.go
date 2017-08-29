package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	serveWeb()
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

func serveWeb() {
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

func serveContent(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	pageAlias := urlParams["pageAlias"]
	t := time.Now()

	if pageAlias == "" {
		pageAlias = "index"
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

//func getThemeName() string {
//	return "default"
//}

func populateStaticPages() *template.Template {
	result := template.New("templates")
	templatePaths := new([]string)

	basePath := "content"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()
	templatePathsRaw, _ := templateFolder.Readdir(-1)

	for _, pathInfo := range templatePathsRaw {
		log.Println(pathInfo.Name())
		*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
	}

	basePath = "includes"
	templateFolder, _ = os.Open(basePath)
	defer templateFolder.Close()
	templatePathsRaw, _ = templateFolder.Readdir(-1)

	for _, pathInfo := range templatePathsRaw {
		log.Println(pathInfo.Name())
		*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
	}

	result.ParseFiles(*templatePaths...)
	return result
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	//path := "assets" + themeName + req.URL.Path
	path := "assets" + req.URL.Path
	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css; charset=utf-8"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript; charset=utf-8"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png; charset=utf-8"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg; charset=utf-8"
	} else if strings.HasSuffix(path, ".svg") {
		contentType = "image/svg+xml; charset=utf-8"
	} else {
		contentType = "text/plain; charset=utf-8"
	}

	log.Println(path)
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
