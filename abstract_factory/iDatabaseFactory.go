package abstract_factory

import (
	"fmt"
)

type IDatabaseFactory interface {
	MakeConnection() IConnection
	MakeCommand() ICommand
}

func GetDatabaseFactory(DbType string) (IDatabaseFactory, error) {
	switch DbType {
	case "MySQL":
		return &MySQLFactory{}, nil
	case "PostgreSQL":
		return &PostgreSQLFactory{}, nil
	default:
		return nil, fmt.Errorf("unknown database type: %s", DbType)
	}
}