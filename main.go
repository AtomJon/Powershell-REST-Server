package main

import (
	"log"
	"net/http"

	"github.com/AtomJon/Powershell-REST-Server/handler"
)

func main() {
	log.Println("Starting listener");

	_handler := handler.SubScriptHandler{};
	
	err := http.ListenAndServe(":8000", _handler);
	if (err == nil) {
		log.Fatal(err);
	}

	log.Println("Exiting");
}