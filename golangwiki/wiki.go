package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Page represents a Page
type Page struct {
	Title string
	Body  []byte
}

func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("This is some test content.")}
	// p1.save()
	// p2, _ := loadPage(p1.Title)
	// fmt.Println(string(p2.Body))
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8989", nil)
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
