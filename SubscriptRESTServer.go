package subscriptrestserver

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/AtomJon/Subscript-REST-Server/handler"
)

func StartServer() {
	dateString := time.Now().Local().Format("2006-01-02")

	os.Mkdir("log", os.ModeDir);
	logPath := path.Join("log", dateString + ".txt")

	file, err := os.Create(logPath);
	if (err != nil) {
		log.Panicln(err);
	}

	writer := io.MultiWriter(os.Stdout, file)
	log.Default().SetOutput(writer)

	log.Println("Starting listener");

	_handler := handler.SubScriptHandler{};
	
	err = http.ListenAndServe(":8000", _handler);
	if (err == nil) {
		log.Fatal(err);
	}

	log.Println("Exiting");
}