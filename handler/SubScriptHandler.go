package handler

import (
	"net/http"

	"github.com/AtomJon/subscriptrestserver/executor"
	"github.com/AtomJon/subscriptrestserver/resource"
)

type SubScriptHandler struct {
	http.Handler
}

func (SubScriptHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	reqHandler := RequestHandler{w, *request, resource.FindResource, executor.ExecuteResource}
	reqHandler.Handle()
}