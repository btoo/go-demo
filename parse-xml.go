package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

// var washPostXML = []byte(`
// 	<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
// 		 <sitemap>
// 				<loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
// 		 </sitemap>
// 		 <sitemap>
// 				<loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
// 		 </sitemap>
// 		 <sitemap>
// 				<loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
// 		 </sitemap>
// 	</sitemapindex>
// `)

// these three structs:
// type SitemapIndex struct {
// 	Locations []Location `xml:"sitemap"`
// }

// type Location struct {
// 	Loc string `xml:"loc"`
// }

// func (l Location) String() string {
// 	return fmt.Sprintf(l.Loc)
// }

// can just be reduced to this struct:
type SitemapIndex struct {
	Locations []string `xml:"sitemap > loc"`
}

type News struct {
	Titles []string `xml:"url > news > title"`
	Keywords []string `xml:"url > news > keywords"`
	Locations []string `xml:"url > loc"`
}

func main() {

	var s SitemapIndex
	var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// fmt.Println(s.Locations) // when printing, go will need a "to string" method, do let's add it to our Location struct
	
	// fmt.Printf("Here %s some %s", "are", "variables") // string formatting with placeholders
	for _, Location := range s.Locations {
		// fmt.Printf("\n%s", Location)
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n) // unmarshal the data into the news address
	}

}