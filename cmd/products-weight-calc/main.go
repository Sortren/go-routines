package main

import (
	"github.com/Sortren/go-routines/apps/products-weight-calc/products"
	"github.com/Sortren/go-routines/pkg/files"
	"sync"
)

func main() {
	ch := make(chan *products.Product)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	fileProvider := files.NewFileProvider("products.txt")
	weightCalculator := products.NewWeightCalculator()
	productsCreator := products.NewProductsCreator(fileProvider)

	go productsCreator.Produce(wg, ch)
	go weightCalculator.Consume(wg, ch)

	wg.Wait()
}
