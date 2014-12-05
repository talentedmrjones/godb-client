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

	julia := table.NewRecord()
	julia.SetString("name", "Julia")
	julia.SetUint32("age", 30)
	result := julia.Create()
	fmt.Printf("%d Created %s in %vs\n", result.Status, julia.String("name"), time.Since(start).Seconds())

	start = time.Now()

	query := table.NewQuery()
	query.SetString("id", julia.String("id"))
	result = query.Find()
	fmt.Printf("%s is %d year old in %vs\n", result.Result[0].String("name"), result.Result[0].Uint32("age"), time.Since(start).Seconds())

	start = time.Now()
	jules := table.NewRecord()
	jules.SetString("id", julia.String("id"))
	jules.SetString("name", "Jules!")
	result = jules.Update()
	fmt.Printf("%v in %vs", result, time.Since(start).Seconds())

}
