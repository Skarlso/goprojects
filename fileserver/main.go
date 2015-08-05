package main

import "net/http"

func main() {
	http.ListenAndServe(":8484", http.FileServer(http.Dir(".")))
}
