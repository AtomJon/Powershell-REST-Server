package handler

import (
	"net/http"

	"github.com/AtomJon/Subscript-REST-Server/executor"
	"github.com/AtomJon/Subscript-REST-Server/resource"
)

type SubScriptHandler struct {
	http.Handler
}

func (SubScriptHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	reqHandler := RequestHandler{w, *request, resource.FindResource, executor.ExecuteResource}
	reqHandler.Handle()
}