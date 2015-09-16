package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

//Page represents a Page
type Page struct {
	Title string
	Body  template.HTML
}

const (
	tmpDir  = "tmp/"
	dataDir = "data/"
)

var templates = template.Must(template.ParseGlob("tmp/*"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.Handle("/view/css/", http.StripPrefix("/view/css", http.FileServer(http.Dir("tmp/css"))))
	http.Handle("/view/img/", http.StripPrefix("/view/img", http.FileServer(http.Dir("tmp/img"))))
	http.HandleFunc("/", frontPageHandler)
	log.Printf("Starting server to listen on port: 8989...")
	http.ListenAndServe(":8989", nil)
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	log.Printf("Saving Page: %s", filename)
	return ioutil.WriteFile(dataDir+filename, []byte(p.Body), 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	log.Printf("Loading Page: %s", filename)
	body, err := ioutil.ReadFile(dataDir + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: template.HTML(body)}, nil
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func frontPageHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Redirecting to FrontPage:")
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
	return
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	log.Printf("Edit handler: %s %s %s %s", r.RemoteAddr, r.Method, r.URL, title)
	renderTemplate(w, "edit", p)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	log.Printf("View handler: %s %s %s %s", r.RemoteAddr, r.Method, r.URL, title)
	renderTemplate(w, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	var interLink = regexp.MustCompile(`\[(\w+)\]`)
	body := r.FormValue("body")
	body = interLink.ReplaceAllString(body, "<a href=\"/view/$1\">Link to Page $1</a>")
	body = template.HTMLEscapeString(body)
	p := &Page{Title: title, Body: template.HTML(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Save handler: %s %s %s %s", r.RemoteAddr, r.Method, r.URL, title)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	log.Printf("Rendering template: %s %s", tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//String represents a page to string method
func (p *Page) String() string {
	return fmt.Sprintf("Title: %s; Body: %s", p.Title, string(p.Body))
}
