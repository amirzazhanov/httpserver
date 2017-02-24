// https://golang.org/doc/articles/wiki/
// http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
// https://thenewstack.io/make-a-restful-json-api-go/
//
package main

import (
	"fmt"
	"net/http"
)

// HandlerImages - handler for /images/ url
func HandlerImages(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Requested Image Is %s!", r.URL.Path[1:])
}

// HandlerRSS - handler for /rss/ url
func HandlerRSS(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Requested RSS is %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/images/", HandlerImages)
	http.HandleFunc("/rss/", HandlerRSS)
	http.ListenAndServe(":8080", nil)
}
