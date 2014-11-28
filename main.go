package main

import (
	//"encoding/binary"
	"github.com/talentedmrjones/godb-client/client"
)

type User struct {

}

func main () {

	var users [4]map[string]string

	users[0] = map[string]string{"id":"123", "name":"Tait"}
	users[1] = map[string]string{"id":"124", "name":"Eden"}
	users[2] = map[string]string{"id":"125", "name":"Mike"}
	users[3] = map[string]string{"id":"126", "name":"Dick"}

	connection := client.NewConnection("127.0.0.1", "6000")
	db := connection.Database("tmj")
	table := db.Table("users")
	for _, user := range users {

		record := table.NewRecord()
		for field, value := range user {
			record.SetFieldString(field, value)
		}
		record.Create()
	}
}
