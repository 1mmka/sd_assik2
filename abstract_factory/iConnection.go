package abstract_factory

import (
	"fmt"
)

type IConnection interface {
	Connect() string
	Close() string
}

type Connection struct {
	DbType string
}

func (conn *Connection) Connect() string {
	return fmt.Sprintf("connected to %s database", conn.DbType)
}

func (conn *Connection) Close() string {
	return fmt.Sprintf("closed connection to %s database", conn.DbType)
}