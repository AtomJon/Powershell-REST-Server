package resource

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func _FindResource(resourceName string) (Resource, error) {
	resourcePath := filepath.Join("./routes/", resourceName)

	if len(filepath.Ext(resourcePath)) < 1 {
		resourcePath += ".*"
	}

	matches, err := filepath.Glob(resourcePath)
	if err != nil {
		return _ReturnErrorWithString("Invalid pattern '%s':\n%e\n", resourcePath, err)
	}

	if len(matches) < 1 {
		return Resource{}, ResourceNotFoundError{fmt.Sprintf("Could not find resource %s\n", resourcePath)}
	}

	if len(matches) > 1 {
		reply := fmt.Sprintf("Resource is not unique, specify between resources: %s", strings.Join(matches, ", "))
		return Resource{}, ResourceNotUniqueError{reply}
	}

	content, err := os.ReadFile(matches[0])
	if err != nil {
		return _ReturnErrorWithString("Cannot read resource '%s':\n%v\n", resourcePath, err)
	}

	return Resource{content, filepath.Ext(resourcePath)}, nil
}

func _ReturnErrorWithString(format string, a ...interface{}) (Resource, error) {
	return Resource{}, fmt.Errorf(format, a...)
}