package abstract_factory

import (
	"fmt"
)

type ICommand interface {
	Execute() string
}

type Command struct {
	DbType string
}

func (comm *Command) Execute() string {
	return fmt.Sprintf("executed command on %s database", comm.DbType)
}