package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/AtomJon/subscriptrestserver/resource"
)

type FindResourceFunc func (string) (resource.Resource, error)
type ExecuteResourceFunc func (resource.Resource, resource.ScriptExecutionRequest) (string, error)

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

		var executionRequest resource.ScriptExecutionRequest

		if handler.Request.Method == http.MethodPost {

			body, err := io.ReadAll(handler.Request.Body)
			if (err != nil) {
				log.Printf("Error while reading request body: %v", err)
			}

			err = json.Unmarshal(body, &executionRequest)

			if (err != nil) {
				log.Printf("Error while parsing request body: %v", err)
			}
		}

		log.Printf("Executing script '%s', with parameters: '%v'", resourceName, executionRequest.Parameters)

		reply, err := handler.ExecuteResource(content, executionRequest);
		if (err != nil) {
			log.Printf("Error while executing resource: %v", err)
			handler.Reply(500, err.Error());
		} else {
			handler.Reply(200, reply);
		}
	}
}

func (handler RequestHandler) Reply(code int, s string) {
	handler.Writer.WriteHeader(code);
	fmt.Fprint(handler.Writer, s);
}