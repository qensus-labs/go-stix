package stix

// Builder creates STIX bundles.
type Builder struct {
	bundle *Bundle

	objects map[string]STIXObject
}

// NewBuilder creates a new builder.
func NewBuilder() *Builder {

	return &Builder{

		bundle: NewBundle(),

		objects: make(
			map[string]STIXObject,
		),
	}
}

// Add adds an object.
//
// Duplicate IDs are ignored.
func (b *Builder) Add(
	object STIXObject,
) STIXObject {

	if object == nil {
		return nil
	}

	id := object.GetID()

	if existing, ok := b.objects[id]; ok {

		return existing
	}

	b.objects[id] = object

	b.bundle.Add(object)

	return object
}

// Get retrieves an object by ID.
func (b *Builder) Get(
	id string,
) (STIXObject, bool) {

	object, ok := b.objects[id]

	return object, ok
}

// Count returns number of unique objects.
func (b *Builder) Count() int {

	return len(b.objects)
}

// Objects returns all objects currently registered in the builder.
func (b *Builder) Objects() []STIXObject {

	objects := make(
		[]STIXObject,
		0,
		len(b.objects),
	)

	for _, object := range b.objects {
		objects = append(
			objects,
			object,
		)
	}

	return objects
}

// Bundle returns the generated bundle.
func (b *Builder) Bundle() *Bundle {

	return b.bundle
}

// JSON serializes the bundle.
func (b *Builder) JSON() ([]byte, error) {

	return b.bundle.JSON()
}
