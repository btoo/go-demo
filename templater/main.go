package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>dis be da index page</h1>")
}

type NewsAggPage struct {
	Title string
	News string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "this is the title of an instance of a NewsAggPage", News: "this is the news content of an instance of a NewsAggPage"}
	t, err := template.ParseFiles("template.html")
	fmt.Println(err)
	t.Execute(w, p)

	// this line will print the error if there is one (such as when there is an err interpolating the template)
	// fmt.Println(t.Execute(w, p))
}

func main() {
	
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)

}