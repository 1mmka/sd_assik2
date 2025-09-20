package database


type sqlite struct {
	Database
}

func NewSQLite() IDatabase {
	return &sqlite{
		Database: Database{
			database: make(map[string]string),
			name:    "SqLite",
		},
	}
} 