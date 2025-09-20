package abstractFactory

import (
	"fmt"
	"factory/abstract_factory/abstractProduct"
	"factory/abstract_factory/brandFactory"
)

type ISportsFactory interface {
	MakeShoe() abstractProduct.IShoe
	MakeShirt() abstractProduct.IShirt
}

func GetSportsFactory(brandType string) (ISportsFactory, error) {
	switch brandType {
	case "adidas":
		return &brandFactory.Adidas{}, nil
	case "nike":
		return &brandFactory.Nike{}, nil
	default:
		return nil, fmt.Errorf("invalid brand type: %s", brandType)
	}
}