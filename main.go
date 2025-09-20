package main

import (
	"fmt"
	"factory/factory/databaseFactory"
	"factory/abstract_factory/abstractFactory"
	"factory/abstract_factory/abstractProduct"
)

func main () {

	// Factory Method Pattern

	mongodb, _ := databaseFactory.GetDatabase("mongodb")
	sqlite, _ := databaseFactory.GetDatabase("sqlite")

	fmt.Println(mongodb.GetName())
	fmt.Println(sqlite.GetName())

	mongodb.SetData("password", "bgdad6t1.adHDAJDa")
	fmt.Println(mongodb.GetData("password"))

	
	// Abstract Factory Pattern
	adidasFactory, _ := abstractFactory.GetSportsFactory("adidas")
	nikeFactory, _ := abstractFactory.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	nikeShoe.SetSize(42)
	nikeShirt.SetSize(40)

	adidasShoe.SetSize(44)
	adidasShirt.SetSize(42)

	printShoeDetails(nikeShoe)
    printShirtDetails(nikeShirt)

    printShoeDetails(adidasShoe)
    printShirtDetails(adidasShirt)
}

func printShirtDetails(shirt abstractProduct.IShirt) {
	fmt.Printf("Logo: %s", shirt.GetLogo())
	fmt.Println()
    fmt.Printf("Size: %d", shirt.GetSize())
	fmt.Println()
}

func printShoeDetails(shoe abstractProduct.IShoe) {
	fmt.Printf("Logo: %s", shoe.GetLogo())
	fmt.Println()
    fmt.Printf("Size: %d", shoe.GetSize())
	fmt.Println()
}