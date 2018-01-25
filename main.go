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
	if err := http.ListenAndServe(":8080", nil); err == nil {
		log.Fatal(err)
	}
}

func handleAddBookmark(w http.ResponseWriter, req *http.Request) {
	full_url := req.PostFormValue("fullUrl")
	url, err := url.Parse(full_url)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	full_url = url.String()

	__ := fmt.Sprintf("%X", time.Now().Unix())
	bookmarks[__] = &Bookmark{FullUrl: full_url, TinyName: __}
	w.Write([]byte(__))
}

func handleGetBookmark(w http.ResponseWriter, req *http.Request) {
	httpGetRequestPostedTinyName := req.URL.Path
	if httpGetRequestPostedTinyName == "/" {
		w.Write([]byte("<h1>Welcome to MyTinyUrl</h1>"))
		return
	}
	httpGetRequestPostedTinyName = httpGetRequestPostedTinyName[1:]

	var fullUrl string
	for key := range bookmarks {
		if bookmarks[key].TinyName == httpGetRequestPostedTinyName {
			fullUrl = bookmarks[key].FullUrl
		}
	}

	if fullUrl == "" {
		http.Error(w, "Not Found", 404)
		return
	}
	http.Redirect(w, req, fullUrl, 301)
}
