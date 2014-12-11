package protocol

type Container struct {
	Body []byte
	sig []byte
}

func NewContainer() (r Container) {
	r = Container{}
	return r
}
