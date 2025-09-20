package brandFactory

import (
	"factory/abstract_factory/abstractProduct"
	"factory/abstract_factory/products"
)

type Adidas struct {}

func (a *Adidas) MakeShirt() abstractProduct.IShirt {
	return &products.AdidasShirt{
		Shirt: abstractProduct.Shirt{
			Logo: "Adidas",
			Size: 0,
		},
	}
}

func (a *Adidas) MakeShoe() abstractProduct.IShoe {
	return &products.AdidasShoe{
		Shoe: abstractProduct.Shoe{
			Logo: "Adidas",
			Size: 0,
		},
	}
		
}