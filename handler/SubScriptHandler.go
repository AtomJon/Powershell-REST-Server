package handler

import (
	"net/http"

	"github.com/AtomJon/subscriptrestserver/executor"
	"github.com/AtomJon/subscriptrestserver/resource"
)

type SubScriptHandler struct {
	ResourceFinder resource.ResourceFinder
}

func (handler SubScriptHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	reqHandler := RequestHandler{w, *request, handler.ResourceFinder, executor.ExecuteResource}
	reqHandler.Handle()
}