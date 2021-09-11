package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type FindResourceFunc func (string) ([]byte, error)

type _RequestHandler struct {
	writer http.ResponseWriter
	request http.Request

	findResource FindResourceFunc
}

func (handler _RequestHandler) Handle() {
	resourceName := handler.request.RequestURI;

	log.Println("Request: " + resourceName);
	
	content, err := handler.findResource(resourceName);
	if (err != nil) {
		log.Printf("Error: %v\n", err);

		if (errors.Is(err, _ResourceNotFoundError{})) {
			handler.Reply(404, "Cannot find resource");
		} else {
			handler.Reply(500, "Resource could not be read");
		}
	} else {
		handler.Reply(200, string(content));
	}
}

func (handler _RequestHandler) Reply(code int, s string) {
	handler.writer.WriteHeader(code);
	fmt.Fprint(handler.writer, s);
}