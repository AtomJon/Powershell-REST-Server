package resource

type ResourceFinder interface {
	CreateOrModifyResource(string, []byte) error
	FindResource(string) (Resource, error)
	DeleteResource(string) error
}