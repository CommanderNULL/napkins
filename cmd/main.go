package main

import (
	"github.com/CommanderNULL/napkins/internal/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	//start frontend server
	//fixme later add custom path for different enviroments
	fserver := http.FileServer(http.FS(os.DirFS("./assets/frontend")))

	mux := http.NewServeMux()
	mux.Handle("/", fserver)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/test", handler.HelloWorld)

	go func() {
		log.Println("serving frontend at 14141...")
		err := http.ListenAndServe(":14141", mux)
		if err != nil {
			log.Fatalf("cant serve frontend: %w", err)
		}
	}()

	log.Println("serving backend at 14142...")
	err := http.ListenAndServe(":14142", mux2)
	if err != nil {
		log.Fatalf("cant serve backend: %w", err)
	}

	//encrypt everythong and store key in memory?
	//tls ssl?
	//start backend handlers
}
