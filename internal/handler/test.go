package handler

import (
	"log"
	"net/http"
)

func HelloWorld(http.ResponseWriter, *http.Request) {
	log.Println("kek!")
}
