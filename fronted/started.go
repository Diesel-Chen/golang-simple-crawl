package main

import (
	"golang-simple-crawl/fronted/controller"
	"log"
	"net/http"
)

func main() {
	//localhost:9999/search?q=未婚&from=10
	http.Handle("/", http.FileServer(http.Dir("fronted/view")))
	http.Handle("/search", controller.CreateSerachResult())
	log.Fatal(http.ListenAndServe(":9999", nil))
}
