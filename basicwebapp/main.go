package main

import (
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

func main() {
<<<<<<< HEAD
	//Port is provided by heroku which we need to bind to.
=======
>>>>>>> 36f013f142130430590313c0ab4a1b375578533c
	port := os.Getenv("PORT")
	if port == "" {
		port = "8484"
	}

	http.HandleFunc("/markdown", generateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("assets")))
	http.ListenAndServe(":"+port, nil)
}

func generateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
