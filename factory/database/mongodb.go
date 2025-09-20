package database


type mongoDB struct {
	Database
}

func NewMongoDB() IDatabase {
	return &mongoDB{
		Database: Database{
			database: make(map[string]string),
			name:    "MongoDB",
		},
	}
} 