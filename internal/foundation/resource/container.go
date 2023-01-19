package resource

type resourceContainer struct {
	storage map[string][]byte
}

func newResourceContainer() *resourceContainer {
	return &resourceContainer{storage: make(map[string][]byte)}
}

func (resourceContainer *resourceContainer) Has(filename string) bool {
	if _, ok := resourceContainer.storage[filename]; ok {
		return true
	}
	return false
}

func (resourceContainer *resourceContainer) Get(filename string) ([]byte, bool) {
	if f, ok := resourceContainer.storage[filename]; ok {
		return f, ok
	}
	return nil, false
}

func (resourceContainer *resourceContainer) Add(filename string, content []byte) {
	resourceContainer.storage[filename] = content
}

var resources = newResourceContainer()

func Get(filename string) ([]byte, bool) {
	return resources.Get(filename)
}

func Add(filename string, content []byte) {
	resources.Add(filename, content)
}

func Has(filename string) bool {
	return resources.Has(filename)
}
