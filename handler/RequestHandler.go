package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AtomJon/subscriptrestserver/resource"
)

type FindResourceFunc func (string) (resource.Resource, error)
type ExecuteResourceFunc func (resource.Resource) (string, error)

type RequestHandler struct {
	Writer http.ResponseWriter
	Request http.Request

	ResourceFinder resource.ResourceFinder
	ExecuteResource ExecuteResourceFunc
}

func (handler RequestHandler) Handle() {
	resourceName := handler.Request.RequestURI;

	log.Println("Request: " + resourceName);
	
	content, err := handler.ResourceFinder.FindResource(resourceName);
	if (err != nil) {
		log.Printf("Error: %v\n", err);

		switch err.(type) {

		case resource.ResourceNotFoundError:
			handler.Reply(404, err.Error());

		case resource.ResourceNotUniqueError:
			handler.Reply(406, err.Error());

		default:
			handler.Reply(500, "Resource could not be read");

		}

	} else {
		reply, err := handler.ExecuteResource(content);
		if (err != nil) {
			msg := fmt.Sprintf("Error while executing %v", err)
			
			log.Print(msg)
			handler.Reply(500, msg);
		} else {
			handler.Reply(200, reply);
		}
	}
}

func (handler RequestHandler) Reply(code int, s string) {
	handler.Writer.WriteHeader(code);
	fmt.Fprint(handler.Writer, s);
}