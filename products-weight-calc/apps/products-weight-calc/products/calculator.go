package products

import (
	"fmt"
	"sync"
)

type WeightCalculator struct{}

func NewWeightCalculator() *WeightCalculator {
	return &WeightCalculator{}
}

func (c *WeightCalculator) Consume(wg *sync.WaitGroup, products <-chan *Product) {
	defer wg.Done()

	count := 0
	sum := 0

	for {
		select {
		case product, ok := <-products:
			if !ok {
				fmt.Printf("total weight sum of products: %d\n", sum)
				return
			}
			sum += product.Weight
			count++

			if count%100 == 0 {
				fmt.Printf("calculated weight of %d products\n", count)
			}
		}
	}
}
