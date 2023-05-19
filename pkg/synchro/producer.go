package synchro

import "sync"

type Producer[T any] interface {
	Produce(wg *sync.WaitGroup, resource chan<- *T)
}
