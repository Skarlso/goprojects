package main

import (
	"net/http"

	"github.com/russross/blackfriday"
)

func main() {
	http.HandleFunc("/markdown", generateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("assets")))
	http.ListenAndServe(":8484", nil)
}

func generateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
