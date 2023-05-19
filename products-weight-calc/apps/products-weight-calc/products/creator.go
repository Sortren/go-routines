package products

import (
	"bufio"
	"fmt"
	"github.com/Sortren/go-routines/products-weight-calc/pkg/files"
	"log"
	"strconv"
	"strings"
	"sync"
)

type Creator struct {
	fileProvider *files.FileProvider
}

func NewProductsCreator(fileProvider *files.FileProvider) *Creator {
	return &Creator{fileProvider: fileProvider}
}

func (p *Creator) parseProductPayload(payload string) (id string, weight int) {
	data := strings.Split(payload, " ")

	id = data[0]
	weight, err := strconv.Atoi(data[1])

	if err != nil {
		log.Fatal(err)
	}
	return id, weight
}
func (p *Creator) Produce(wg *sync.WaitGroup, products chan<- *Product) {
	defer wg.Done()
	defer close(products)

	file, closeFunc := p.fileProvider.File()
	defer closeFunc(file)

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		id, weight := p.parseProductPayload(scanner.Text())
		product := NewProduct(id, weight)

		products <- product
		count++

		if count%200 == 0 {
			fmt.Printf("created %d product objects\n", count)
		}
	}

}
