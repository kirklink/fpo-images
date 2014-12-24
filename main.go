package myapp

import (
	//	"fmt"
	"html/template"
	"net/http"

	"appengine"
	"appengine/user"
)

type PageContent struct {
	Title   string
	SomeNum int
}

func init() {
	//	http.HandleFunc("/", urlCapture)
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	PageStuff := PageContent{Title: "The Title", SomeNum: 4}
	//	fmt.Fprint(w, pageStuff.someNum)
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, PageStuff)
}

//func urlCapture(w http.ResponseWriter, r *http.Request) {
//	//	c := appengine.NewContext(r)
//	item := r.URL.Path[1:]
//	fmt.Fprint(w, item)
//}
