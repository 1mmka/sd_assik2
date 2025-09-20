package database


type IDatabase interface {
	GetName() string
	SetName(name string)
	GetData(key string) string
	SetData(key, value string)
}

type Database struct {
	name string
	database map[string]string
}

func (d *Database) GetName() string {
	return d.name
}

func (d *Database) SetName(name string) {
	d.name = name
}

func (d *Database) GetData(key string) string {
	return d.database[key]
}

func (d *Database) SetData(key, value string) {
	d.database[key] = value
}


