package handler

import (
	"fmt"
	"log"
	"net/http"
)

type FindResourceFunc func (string) (Resource, error)
type ExecuteContentFunc func (Resource) (error, string)

type RequestHandler struct {
	writer http.ResponseWriter
	request http.Request

	findResource FindResourceFunc
	executeConent ExecuteContentFunc
}

func (handler RequestHandler) Handle() {
	resourceName := handler.request.RequestURI;

	log.Println("Request: " + resourceName);
	
	content, err := handler.findResource(resourceName);
	if (err != nil) {
		log.Printf("Error: %v\n", err);

		switch err.(type) {

		case ResourceNotFoundError:
			handler.Reply(404, err.Error());

		case ResourceNotUniqueError:
			handler.Reply(406, err.Error());

		default:
			handler.Reply(500, "Resource could not be read");

		}

	} else {
		handler.executeConent(content);
		// handler.Reply(200, string(content));
	}
}

func (handler RequestHandler) Reply(code int, s string) {
	handler.writer.WriteHeader(code);
	fmt.Fprint(handler.writer, s);
}