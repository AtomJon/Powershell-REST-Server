package handler

import (
	"log"

	"github.com/AtomJon/subscriptrestserver/resource"
)

type FindResourceFunc func (string) (resource.Resource, error)
type ExecuteResourceFunc func (resource.Resource, resource.ScriptExecutionRequest) (string, error)

type RequestHandler struct {
	ResourceFinder resource.ResourceFinder
	ExecuteResource ExecuteResourceFunc
}

type HandlerReply struct {
	Code int
	Message string
}

func (handler RequestHandler) Handle(executionRequest resource.ScriptExecutionRequest) HandlerReply {

	resourceName := executionRequest.ScriptName

	log.Println("Request: " + resourceName);
	
	content, err := handler.ResourceFinder.FindResource(resourceName);
	if (err != nil) {
		log.Printf("Error: %v\n", err);

		switch err.(type) {
			case resource.ResourceNotFoundError:
				return Reply(404, err.Error());
			case resource.ResourceNotUniqueError:
				return Reply(406, err.Error());
			default:
				return Reply(500, "Resource could not be read");
		}

	} else {

		parameters := executionRequest.Parameters

		log.Printf("Executing script '%s', with parameters: '%v'", resourceName, parameters)

		reply, err := handler.ExecuteResource(content, executionRequest);
		if (err != nil) {
			return Reply(500, err.Error());
		} else {
			return Reply(200, reply);
		}
	}
}

func Reply(code int, s string) HandlerReply {
	return HandlerReply{
		Code: code,
		Message: s,
	}
}