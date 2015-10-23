package views

import (
	"html/template"
	"net/http"

	"webtrack"
)

func SamplePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")

	tmpl, err := template.ParseFiles(webtrack.GetResource("templates", "samplepage.html"))
	webtrack.CheckErrorPanic(err)

	script_url, err := webtrack.Router.Get("trackjs").URL("userKey", "XXXXXXXXXX")
	webtrack.CheckErrorPanic(err)

	err = tmpl.Execute(w, script_url)
	webtrack.CheckErrorPanic(err)
}
