package handler

import (
	"net/http"
)

type SubScriptHandler struct {
	http.Handler
}

func (SubScriptHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	reqHandler := RequestHandler{w, *request, func(s string) (Resource, error) {return Resource{}, nil}, _ExecuteContent}
	reqHandler.Handle()
}



func _ExecuteContent(resource Resource) (error, string) {

	return nil, "Success";
}