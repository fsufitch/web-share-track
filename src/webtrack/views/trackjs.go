package views

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"

	"webtrack"
)

func RenderTrackingJs(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(webtrack.GetResource("templates", "track.js"))
	webtrack.CheckErrorPanic(err)

	trackId := mux.Vars(r)["trackId"]
	w.Header().Set("Content-type", "application/javascript")
	tmpl.Execute(w, trackId)
}
