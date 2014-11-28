package client

type Command struct {
	// TODO support unique ID for async replies
	Action 	string
	Db			string
	Table 	string
	Data		map[string][]byte
}

func NewCommand (command string, record *Record) *Command {

	return &Command{
		command,
		record.table.db.name,
		record.table.name,
		record.data,
	}
}
