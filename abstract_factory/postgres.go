package abstract_factory


type PostgreSQLFactory struct{}

func (p *PostgreSQLFactory) MakeConnection() IConnection {
	return &PostgresConnection{Connection{DbType: "PostgreSQL"}}
}

func (p *PostgreSQLFactory) MakeCommand() ICommand {
	return &PostgresCommand{Command{DbType: "PostgreSQL"}}
}