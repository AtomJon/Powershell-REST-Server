package resource

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mattn/go-zglob"
)

type DirectoryResourceFinder struct {
	Folder string
}

func (finder DirectoryResourceFinder) CreateOrModifyResource(resourceName string, contents []byte) (error) {
	path, err := finder.FindResourcePath(resourceName)

	if err != nil {
		path = finder.Folder + "/" + resourceName
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(contents)
	return err
}

func (finder DirectoryResourceFinder) FindResource(resourceName string) (Resource, error) {
	resourcePath, err := finder.FindResourcePath(resourceName)

	if err != nil {
		return Resource{}, err
	}

	content, err := os.ReadFile(resourcePath)
	if err != nil {
		return _ReturnErrorWithString("cannot read resource '%s':\n%v\n", resourcePath, err)
	}

	return Resource{string(content), filepath.Ext(resourcePath)}, nil
}

func (finder DirectoryResourceFinder) DeleteResource(resourceName string) (error) {
	resourcePath, err := finder.FindResourcePath(resourceName)

	if err != nil {
		return err
	}

	err = os.Remove(resourcePath)
	return err
}

func (finder DirectoryResourceFinder) FindResourcePath(resourceName string) (string, error) {
	if (len(resourceName) < 1) {
		return "", fmt.Errorf("no name for the resource was provided")
	}

	if (resourceName[0] == '/') {
		resourceName = resourceName[1:]
	}

	resourcePath := finder.Folder + "/**/" + resourceName

	if len(filepath.Ext(resourcePath)) < 1 {
		resourcePath += ".*"
	}

	matches, err := zglob.Glob(resourcePath)
	if err != nil {
		return "", fmt.Errorf("invalid pattern '%s':\n%e", resourcePath, err)
	}

	if len(matches) < 1 {
		return "", ResourceNotFoundError{fmt.Sprintf("could not find resource %s\n", resourcePath)}
	}

	if len(matches) > 1 {
		reply := fmt.Sprintf("resource is not unique, specify between resources: %s", strings.Join(matches, ", "))
		return "", ResourceNotUniqueError{reply}
	}

	resourcePath = matches[0]

	return resourcePath, nil
}

func _ReturnErrorWithString(format string, a ...interface{}) (Resource, error) {
	return Resource{}, fmt.Errorf(format, a...)
}