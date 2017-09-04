package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

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
