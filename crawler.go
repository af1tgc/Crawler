package main

import (
//    "fmt"
    "log"
    "net/http"
    "html/template"
    "github.com/PuerkitoBio/goquery"
)

type Page struct {
    Title string
    Body  []byte
}

// Render Handler
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
    t, err := template.ParseFiles("static/" + tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Routers
func rootHandler(w http.ResponseWriter, r *http.Request) {
    title := "index"
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    renderTemplate(w, "root", p)
}


//TODO: Get Url to Crawling
func getCrawlUrlHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[len("/getUrl"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    renderTemplate(w, "page/getUrl", p)
}

func backCrawlHandler(w http.ResponseWriter, r *http.Request){
    title := "Go Ahead"
    body := r.FormValue("body")
    test := Crawl(body)
    p := &Page{Title: title, Body: []byte(test)}
    
    renderTemplate(w, "page/back_run", p)
}


//TODO: web-crawl
func Crawl(url string) string {
    test := ""
//    res, err := http.Get(url)
    res, err := http.Get("http://metalsucks.net")
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()
    if res.StatusCode != 200 {
        log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
    }

    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        log.Fatal(err)
    }
/*
    doc.Find(".col-md-9 content").Each(func(i int, s *goquery.Selection){
        test = s.Find("h3").Text()
    })

    return test
*/

  doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
    band := s.Find("a").Text()
    title := s.Find("i").Text()
    test = "Review " + band + title
  })
    return test
}

//TODO: Render the Tag lists for selection
//TODO: Save to Files (csv) || Insert to Database

func main() {
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/getUrl", getCrawlUrlHandler)
    http.HandleFunc("/back_run", backCrawlHandler)
    log.Fatal(http.ListenAndServe(":8787", nil))
}
