package handler

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type _ResourceNotFoundError struct {
	ResourceName string
}

func (err _ResourceNotFoundError) Error() string {
	return err.ResourceName;
}

type SubScriptHandler struct {
	http.Handler
}

func (SubScriptHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	reqHandler := _RequestHandler{w, *request, _FindResource}
	reqHandler.Handle()
}

func _FindResource(resourceName string) ([]byte, error) {
	resourcePath := filepath.Join("./routes/", resourceName);

	file, err := os.OpenFile(resourcePath, os.O_RDONLY, 0);
	if (err != nil) {
		if (errors.Is(err, os.ErrNotExist)) {
			return nil, _ResourceNotFoundError{fmt.Sprintf("Could not find resource %s\n", resourcePath)};
		} else {
			return _ReturnErrorWithString("Cannot open resource '%s':\n%e\n", resourcePath, err);
		}
	}

	stat, err := file.Stat();
	if (err != nil) {
		return _ReturnErrorWithString("Cannot get status of resource '%s':\n%v\n", resourcePath, err);
	}

	if (stat.IsDir()) {
		return _ReturnErrorWithString("Resource '%s' is dir", resourcePath);
	}

	content, err := io.ReadAll(file);
	if (err != nil) {
		return _ReturnErrorWithString("Cannot read resource '%s':\n%v\n", resourcePath, err);
	}

	return content, nil
}

func _ReturnErrorWithString(format string, a ...interface{}) ([]byte, error) {
	log.Printf(format, a...);
	return nil, fmt.Errorf(format, a...);
}