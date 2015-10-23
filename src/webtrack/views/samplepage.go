package views

import (
	"net/http"
	"html/template"

	"webtrack"
)

func SamplePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")

	tmpl, err := template.ParseFiles(webtrack.GetResource("templates", "samplepage.html"), )
	webtrack.CheckErrorPanic(err)

	err = tmpl.Execute(w, nil);
	webtrack.CheckErrorPanic(err)
}
