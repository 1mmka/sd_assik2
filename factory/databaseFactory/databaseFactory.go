package databaseFactory

import (
	"factory/factory/database"
	"fmt"
)

func GetDatabase(dbType string) (database.IDatabase, error) {
	switch dbType {
		case "sqlite":
			return database.NewSQLite(), nil
		case "mongodb":
			return database.NewMongoDB(), nil
		default:
			return nil, fmt.Errorf("unknown database type: %s", dbType)
	}
}