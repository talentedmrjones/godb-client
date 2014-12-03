package client

type Table struct {
	db 		*Database
	name	string
}

func (table *Table) NewRecord () *Record {
	return &Record{table, make(map[string][]byte), ""}
}
