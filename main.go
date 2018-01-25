package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Bookmark struct {
	TinyName string
	FullUrl  string
}

var (
	bookmarks = map[string]*Bookmark{}
)

func main() {
	http.HandleFunc("/_add_bookmark", handleAddBookmark)
	http.HandleFunc("/", handleGetBookmark)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleAddBookmark(w http.ResponseWriter, req *http.Request) {
	fullUrl := req.PostFormValue("fullUrl")
	url, err := url.Parse(fullUrl)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fullUrl = url.String()
	tinyName := fmt.Sprintf("%X", time.Now().Unix())
	bookmarks[tinyName] = &Bookmark{FullUrl: fullUrl, TinyName: tinyName}
	w.Write([]byte(tinyName))
}

func handleGetBookmark(w http.ResponseWriter, req *http.Request) {
	tinyName := req.URL.Path[1:]
	if tinyName == "" {
		w.Write([]byte("<h1>Welcome to MyTinyUrl</h1>"))
		return
	}
	bookmark := bookmarks[tinyName]
	if bookmark == nil {
		http.Error(w, "Not Found", 404)
		return
	}
	http.Redirect(w, req, bookmark.FullUrl, 301)
}
