package handler

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Handler struct {
	http.Handler
}

type RequestHandler struct {
	writer http.ResponseWriter
	request http.Request
}

func (Handler) _ServeHTTP(w http.ResponseWriter, request *http.Request ) {

	reqHandler := RequestHandler{w, *request};
	reqHandler.Handle();
}

func (handler RequestHandler) Handle() {
	resourceName := handler.request.RequestURI;

	log.Println("Request: " + resourceName);
	
	resourcePath := filepath.Join("./routes/", resourceName);

	file, err := os.OpenFile(resourcePath, os.O_RDONLY, 0);
	if (err != nil) {
		if (errors.Is(err, os.ErrNotExist)) {
			log.Printf("Could not find resource %s, replying 404\n", resourcePath);
			
			handler.Reply(404, "Resource not found :(");
		} else {
			log.Printf("Error while opening resource '%s':\n%e\n", resourcePath, err);
		}

		return;
	}

	stat, err := file.Stat();
	if (err != nil) {
		log.Printf("Error while getting status of resource '%s':\n%v\n", resourcePath, err);
		return;
	}

	if (stat.IsDir()) {
		log.Printf("Resource '%s', is a dir, replying 404\n", resourcePath);
		return;
	}

	content, err := io.ReadAll(file);
	if (err != nil) {
		log.Printf("Error while reading resource '%s':\n%v\n", resourcePath, err);
		return;
	}

	handler.writer.WriteHeader(200);
	fmt.Fprint(handler.writer, content);
}

func (handler RequestHandler) Reply(code int, s string) {
	handler.writer.WriteHeader(code);
	fmt.Fprint(handler.writer, s);
}