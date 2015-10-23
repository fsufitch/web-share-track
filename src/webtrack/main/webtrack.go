package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"webtrack"
	"webtrack/views"
)

func main() {
	webtrack.BuildConfiguration(os.Args[1])
	webtrack.SetResourceDir(os.Args[2])

	webtrack.InitRouter()
	webtrack.AddRoute("samplepage", "GET", "/hello", views.SamplePage)
	webtrack.AddRoute("samplejs", "GET", "/js", views.SampleJs)
	webtrack.AddRoute("trackjs", "GET", "/track/{userKey}/track.js", views.RenderTrackingJs)
	webtrack.AddRoute("track_callback", "POST", "/track/{userKey}/{origTrackId:[0-9a-zA-Z]*}", views.TrackingCallback)

	listen := fmt.Sprintf("%s:%d", webtrack.Configuration.ListenHostname, webtrack.Configuration.ListenPort)
	fmt.Printf("Listening to %s\n", listen)
	log.Fatal(http.ListenAndServe(listen, webtrack.Router))

}
