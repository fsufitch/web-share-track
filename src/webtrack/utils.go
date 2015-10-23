package webtrack

import (
	"log"
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
