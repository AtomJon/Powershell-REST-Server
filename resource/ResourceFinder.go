package resource

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindResource(resourceName string) (Resource, error) {
	resourcePath := filepath.Join("./routes/", resourceName)

	if len(filepath.Ext(resourcePath)) < 1 {
		resourcePath += ".*"
	}

	matches, err := filepath.Glob(resourcePath)
	if err != nil {
		return _ReturnErrorWithString("invalid pattern '%s':\n%e\n", resourcePath, err)
	}

	if len(matches) < 1 {
		return Resource{}, ResourceNotFoundError{fmt.Sprintf("could not find resource %s\n", resourcePath)}
	}

	if len(matches) > 1 {
		reply := fmt.Sprintf("resource is not unique, specify between resources: %s", strings.Join(matches, ", "))
		return Resource{}, ResourceNotUniqueError{reply}
	}

	resourcePath = matches[0];

	content, err := os.ReadFile(resourcePath)
	if err != nil {
		return _ReturnErrorWithString("cannot read resource '%s':\n%v\n", resourcePath, err)
	}

	return Resource{string(content), filepath.Ext(resourcePath)}, nil
}

func _ReturnErrorWithString(format string, a ...interface{}) (Resource, error) {
	return Resource{}, fmt.Errorf(format, a...)
}