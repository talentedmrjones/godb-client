package client

type Database struct {
	connection 	*Connection
	name 				string
}

func (db *Database) Table(tableName string) *Table {
	return &Table{db, tableName}
}
