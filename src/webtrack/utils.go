package webtrack

import (
	"log"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
)

func CheckErrorFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckErrorPanic(err error) {
	if err != nil {
		log.Print(err)
		panic(err)
	}
}

func CheckErrorPrint(err error) {
	if err != nil {
		log.Print(err)
	}
}

var minifier *minify.Minify
func Minify(input string, mime string) (result string) {
	if minifier == nil {
		log.Print("initializing minifier")
		minifier = minify.New()
		minifier.AddFunc("js", js.Minify)
	}
	result, err := minify.String(minifier, mime, input)
	CheckErrorPanic(err)
	return
}
