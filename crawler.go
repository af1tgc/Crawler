package main

import (
//    "fmt"
    "log"
//    "io/ioutil"
    "net/http"
    "html/template"
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
    p := &Page{Title: title, Body: []byte(body)}
    renderTemplate(w, "page/back_run", p)
}



//TODO: Render the Tag lists for selection
//TODO: Save to Files (csv) || Insert to Database

func main() {
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/getUrl", getCrawlUrlHandler)
    http.HandleFunc("/back_run", backCrawlHandler)
    log.Fatal(http.ListenAndServe(":8787", nil))
}
