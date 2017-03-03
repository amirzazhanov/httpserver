// https://golang.org/doc/articles/wiki/
// http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
// https://thenewstack.io/make-a-restful-json-api-go/
//
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Page structure simple
type Page struct {
	Title string
	Body  []byte
}

// HandlerImages - handler for /images/ url
func HandlerImages(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Requested Image Is %s!", r.URL.Path[1:])
}

// HandlerHTML - handler for /rss/ url
func HandlerHTML(w http.ResponseWriter, r *http.Request) {
	pTmp := &Page{Title: "TestPage", Body: []byte("This is a sample simple Page.")}
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, pTmp)
	log.Printf("%s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/images/", HandlerImages)
	http.HandleFunc("/html/", HandlerHTML)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
