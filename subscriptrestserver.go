package subscriptrestserver

import (
	"log"
	"net/http"

	"github.com/AtomJon/subscriptrestserver/handler"
	"github.com/AtomJon/subscriptrestserver/resource"
)

func StartServer(addr string, scriptsFolder string) {
	log.Println("Starting rest server");

	_handler := handler.SubScriptHandler{
		resource.DirectoryResourceFinder{scriptsFolder},
	};
	
	err := http.ListenAndServe(addr, _handler);
	if (err == nil) {
		log.Fatal(err);
	}

	log.Println("Exiting");
}
