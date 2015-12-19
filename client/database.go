package client

/*
Database
*/
type Database struct {
	connection *Connection // TODO change to chan *Connection to support pooling
	name       string
}

/*
Table creates the specified Table from this database
*/
func (db *Database) Table(tableName string) *Table {
	return &Table{db, tableName}
}
