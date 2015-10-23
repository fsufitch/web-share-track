package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"webtrack"
	"webtrack/views"
)

func main() {
	webtrack.SetResourceDir(os.Args[1])

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/hello", views.SamplePage)
	router.HandleFunc("/js", views.SampleJs)
	router.HandleFunc("/track/{trackId}/track.js", views.RenderTrackingJs)

	fmt.Println(webtrack.GetResource("templates/foo.tmpl"))

	log.Fatal(http.ListenAndServe(":8080", router))

}
