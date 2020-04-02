package michel

import (
	"io"
)

type Thing interface {
	DoAnotherThing()
}

type ThingImpl struct {
}

func (t ThingImpl) Read(thing []byte) (int, error) {
	return 1, nil
}

var michel string

func init() {
	michel = "Michel"

	thing := ThingImpl{}
	ThingConsumer(thing)
}

func ThingConsumer(thing io.Reader) {
	thing.Read([]byte("String"))
}
