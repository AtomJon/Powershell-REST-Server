package resource

type ResourceFinder interface {
	FindResource(string) (Resource, error)
}