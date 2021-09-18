package subscriptrestserver

import (
	"net/http"

	"github.com/AtomJon/subscriptrestserver/executor"
	"github.com/AtomJon/subscriptrestserver/handler"
	"github.com/AtomJon/subscriptrestserver/resource"
)

type SubScriptHandler struct {
	ResourceFinder resource.ResourceFinder
}

func (subScriptHandler SubScriptHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	reqHandler := handler.RequestHandler{
		Writer: w,
		Request: *request,
		ResourceFinder: subScriptHandler.ResourceFinder,
		ExecuteResource: executor.ExecuteResource,
	}

	reqHandler.Handle()
}