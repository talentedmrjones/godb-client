package client



func (db *Database) Table(tableName string) *Table {
	return &Table{db, tableName}
}
