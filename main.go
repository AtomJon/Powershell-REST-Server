package main

import (
	"log"
	"net/http"

	"github.com/AtomJon/Powershell-REST-Server/handler"
)

func main() {
	log.Println("Starting listener");
	
	handler := handler.Handler{};
	
	err := http.ListenAndServe(":8000", handler);
	if (err == nil) {
		log.Fatal(err);
	}

	log.Println("Exiting");
}