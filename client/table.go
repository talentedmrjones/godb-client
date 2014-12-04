package client

func (table *Table) NewRecord () *Record {
	return &Record{table, make(map[string][]byte), ""}
}
