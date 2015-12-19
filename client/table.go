package client

/*
Table ...
*/
type Table struct {
	db   *Database
	name string `default:""`
}

/*
NewRecord ...
*/
func (table *Table) NewRecord() Record {
	return Record{
		table,
		make(map[string]interface{}),
		"",
	}
}

// NewQuery returns Query
// func (table *Table) NewQuery(connection *Connection) Query {
// 	return Query{
// 		Patient{
// 			connection,
// 			table,
// 			make(map[string]interface{}),
// 			"r",
// 		},
// 	}
// }
