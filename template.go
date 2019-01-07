package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// NewsAggPage template for html page
type NewsAggPage struct {
	Title string
	News  string
}

func main() {
	http.HandleFunc("/", newsAgg)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi There!")
}

func newsAgg(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "Some News"}
	t, _ := template.ParseFiles("basicTemplating.html")

	t.Execute(w, p)
}
