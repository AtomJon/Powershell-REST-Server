package handler

import "net/http"

type SubScriptHandler struct {
	http.Handler
}

func (SubScriptHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	reqHandler := _RequestHandler{w, *request}
	reqHandler.Handle()
}