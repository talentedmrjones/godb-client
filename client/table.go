package client


// TODO support accepting map[string]interface as param
func (table *Table) NewRecord () Record {
	return Record{
		Patient{
			table,
			make(map[string][]byte),
			"",
		},
	}
}

// NewQuery returns Query
func (table *Table) NewQuery () Query {
	return Query{
		Patient{
			table,
			make(map[string][]byte),
			"r",
		},
	}
}
