package interactor

type Producer interface {
	Produce(msg any) error
}
