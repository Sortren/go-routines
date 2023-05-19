package products

type Product struct {
	Id     string
	Weight int
}

func NewProduct(id string, weight int) *Product {
	return &Product{Id: id, Weight: weight}
}
