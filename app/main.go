package main

import (
	"fmt"
	"net/http"
	"html/template"
	"encoding/xml"
	"io/ioutil"
	"sync"
)

var waitGroup sync.WaitGroup

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>dis be da index page</h1>")
}

type NewsAggPage struct {
	Title string
	News map[string]NewsMap
}

func newsRoutine(c chan News, Location string) {
	defer waitGroup.Done()
	var n News
	// fmt.Printf("\n%s", Location)
	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n) // unmarshal the data into the news address
	resp.Body.Close()

	c <- n
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	
	var s SitemapIndex

	newsMap := make(map[string]NewsMap) // make a map where the keys are strings and the values are NewsMaps

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s) // unmarshal the data into the news address
	queue := make(chan News, 30)

	for _, Location := range s.Locations {
		waitGroup.Add(1)
		go newsRoutine(queue, Location)
	}

	waitGroup.Wait()
	close(queue)

	for elem := range queue {
		for idx, _ := range elem.Titles {
			newsMap[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
		}
	}

	for idx, data := range newsMap {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
	
	p := NewsAggPage{Title: "this is the title of an instance of a NewsAggPage", News: newsMap}
	t, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Execute(w, p))

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}