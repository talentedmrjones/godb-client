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

	records := make(map[string]client.Record)
	start := time.Now()
	var i uint32
	i = 0
	for  time.Since(start).Seconds() <= 1 {
  	record := table.NewRecord();
		record.SetUint32("age", i);
		record.Create();
		records[record.String("id")] = record
		i++
  }
	dur := time.Since(start).Seconds()
	fmt.Printf("Created %d records in %vs (avg %v/s)\n", i, dur, float64(i)/dur)

	start = time.Now()
	for id, _ := range records {
		query := table.NewQuery()
		query.SetString("id", id)
		_ = query.Find()
	}
	dur = time.Since(start).Seconds()
	j := len(records)
	fmt.Printf("Read %d records in %vs (avg %v/s)\n", j, dur, float64(j)/dur)

	start = time.Now()
	for id, _ := range records {
		query := table.NewQuery()
		query.SetString("id", id)
		_ = query.Delete()
	}
	dur = time.Since(start).Seconds()

	fmt.Printf("Deleted %d records in %vs (avg %v/s)\n", j, dur, float64(j)/dur)



}
