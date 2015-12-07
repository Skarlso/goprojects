package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
)

func main() {
	n := negroni.New(negroni.NewRecovery(), negroni.HandlerFunc(myMiddleWare), negroni.NewLogger(), negroni.NewStatic(http.Dir("assets")))

	n.Run(":8484")
}

func myMiddleWare(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Logging on the way in...")

	if r.URL.Query().Get("password") == "secret123" {
		next(rw, r)
	} else {
		http.Error(rw, "Not Authorized", 401)
	}

	log.Println("Logging on the way back...")
}
