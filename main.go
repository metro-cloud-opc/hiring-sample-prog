package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Hashtable map[string]*Bookmark

type Bookmark struct {
	TinyName string
	FullUrl  string
}

var (
	bookmarks = Hashtable{}
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
	reqTinyName := req.URL.Path
	if reqTinyName == "/" {
		w.Write([]byte("<h1>Welcome to MyTinyUrl</h1>"))
		return
	} else {
		reqTinyName = reqTinyName[1:]
	}

	var bookmark *Bookmark
	for key := range bookmarks {
		if bookmarks[key].TinyName == reqTinyName {
			bookmark = bookmarks[key]
		}
	}

	if bookmark == nil {
		http.Error(w, "Not Found", 404)
		return
	}
	http.Redirect(w, req, bookmark.FullUrl, 301)
}
