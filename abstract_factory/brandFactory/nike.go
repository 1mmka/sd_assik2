package brandFactory

import (
	"factory/abstract_factory/abstractProduct"
	"factory/abstract_factory/products"
)

type Nike struct {}

func (a *Nike) MakeShirt() abstractProduct.IShirt {
	return &products.AdidasShirt{
		Shirt: abstractProduct.Shirt{
			Logo: "Adidas",
			Size: 0,
		},
	}
}

func (a *Nike) MakeShoe() abstractProduct.IShoe {
	return &products.AdidasShoe{
		Shoe: abstractProduct.Shoe{
			Logo: "Adidas",
			Size: 0,
		},
	}
}