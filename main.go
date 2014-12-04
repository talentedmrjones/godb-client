package main

import (
	"fmt"
	"time"
	"github.com/talentedmrjones/godb-client/client"
)

type User struct {

}

func main () {

	connection := client.NewConnection("127.0.0.1", "6000")
	db := connection.Database("tmj")
	table := db.Table("users")

	start := time.Now()

	user := table.NewRecord()
	user.SetString("name", "Julia")
	user.SetUint32("age", 30)
	result := user.Create()
	fmt.Printf("Created %s in %s\n", user.String("id"), time.Since(start))

	start = time.Now()

	query := table.NewQuery()
	query.SetString("id", user.String("id"))
	result = query.Find()
	julia := result.Result[0]


	fmt.Printf("%s is %d year old in %s\n", julia.String("name"), julia.Uint32("age"), time.Since(start))



}
