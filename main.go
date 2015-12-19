package main

import (
	"fmt"
	"github.com/talentedmrjones/godb-client/client"
	"time"
)

func main() {

	connection := client.NewConnection("127.0.0.1", "6000")
	db := connection.Database("tmj")
	table := db.Table("users")

	start := time.Now()
	var i uint32
	i = 0

	for i < 10 {
		record := table.NewRecord()
		record.SetField("age", i)
		record.Create()
		i++
	}

	dur := time.Since(start).Seconds()
	fmt.Printf("Created %d records in %vs (avg %v/s)\n", i, dur, float64(i)/dur)

	// TODO figure out why this keeps process open
	for reply := range connection.Replies {
		fmt.Printf("%#v\n", reply)
	}
	connection.Close()
}
