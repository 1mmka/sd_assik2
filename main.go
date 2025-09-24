package main

import (
	"fmt"
	// f "factory/factory/databaseFactory"
	af "factory/abstract_factory"
)

func main () {

	// Factory Method Pattern

	// mongodb, _ := f.GetDatabase("mongodb")
	// sqlite, _ := f.GetDatabase("sqlite")

	// fmt.Println(mongodb.GetName())
	// fmt.Println(sqlite.GetName())

	// mongodb.SetData("password", "bgdad6t1.adHDAJDa")
	// fmt.Println(mongodb.GetData("password"))

	
	// Abstract Factory Pattern
	mysqlFactory, _ := af.GetDatabaseFactory("MySQL")

	mysqlConnection := mysqlFactory.MakeConnection()
	mysqlCommand := mysqlFactory.MakeCommand()

	fmt.Println(mysqlConnection.Connect())
	fmt.Println(mysqlCommand.Execute())

	fmt.Println(mysqlConnection.Close())
}