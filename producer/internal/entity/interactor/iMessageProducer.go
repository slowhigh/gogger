package interactor

type MessageProducer interface {
	Produce(msg any) error
}
