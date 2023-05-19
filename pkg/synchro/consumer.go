package synchro

import "sync"

type Consumer[T any] interface {
	Consume(wg *sync.WaitGroup, resource <-chan *T)
}
