package pattern

import (
	"fmt"
)

type Factorer interface {
	createProduct(string) Producter
}

type Producter interface {
	work()
}

type Factory struct {
}

func (f *Factory) createProduct(pType string) Producter {
	var product Producter
	if pType == "epic" {
		product = &epicProduct{}
	} else {
		product = &basicProduct{}
	}
	return product
}

type epicProduct struct {
}

func (ep *epicProduct) work() {
	fmt.Println("epic product working")
}

type basicProduct struct{}

func (bp *basicProduct) work() {
	fmt.Println("basic product working")
}
