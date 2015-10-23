package views

import (
	"fmt"
	"net/http"
)

func SampleJs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/javascript")
	fmt.Fprintf(w, "alert('hi');")
}
