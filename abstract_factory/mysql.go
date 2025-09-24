package abstract_factory

type MySQLFactory struct{}

func (mysql *MySQLFactory) MakeConnection() IConnection {
	return &MySQLConnection{Connection{DbType: "MySQL"}}
}

func (mysql *MySQLFactory) MakeCommand() ICommand {
	return &MySQLCommand{Command{DbType: "MySQL"}}
}