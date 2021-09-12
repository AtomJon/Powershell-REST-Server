package resource

type ResourceNotFoundError struct {
	ErrorString string
}

type ResourceNotUniqueError struct {
	ErrorString string
}

func (err ResourceNotFoundError) Error() string {
	return err.ErrorString
}

func (err ResourceNotUniqueError) Error() string {
	return err.ErrorString
}

type Resource struct {
	Content           []byte
	ResourceExtension string
}