package views

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"

	"webtrack"
)

type TrackJsData struct {
	Key string
	Callback string
	OrigTrackId string
	CurrTrackId string
}

func RenderTrackingJs(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(webtrack.GetResource("templates", "track.js"))
	webtrack.CheckErrorPanic(err)

	w.Header().Set("Content-type", "application/javascript")

	user_key := mux.Vars(r)["userKey"]
	cb_url, err := webtrack.Router.Get("track_callback").URL(
		"userKey", user_key,
		"origTrackId", "",
	)

	webtrack.CheckErrorPanic(err)

	data := TrackJsData{
		Key: user_key,
		Callback: cb_url.String(),
	}

	var maxjs bytes.Buffer 
	tmpl.Execute(&maxjs, data)

	minified := webtrack.Minify(maxjs.String(), "js")
	w.Write([]byte(minified))
}

func TrackingCallback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	data := TrackJsData{
		Key: mux.Vars(r)["userKey"],
		Callback: webtrack.Configuration.CallbackHost,
		OrigTrackId: mux.Vars(r)["origTrackId"],
		CurrTrackId: makeTrackingId(),
	}

	fmt.Printf("Share %s -> %s\n", data.OrigTrackId, data.CurrTrackId)

	output, err := json.Marshal(data)
	webtrack.CheckErrorPanic(err)
	
	w.Write(output)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func makeTrackingId() string {
	b := make([]rune, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
