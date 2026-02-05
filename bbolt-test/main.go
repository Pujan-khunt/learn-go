package main

import (
	"fmt"
	"log"
	"time"

	"go.etcd.io/bbolt"
)

func main() {
	db, err := bbolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatalf("Error opening database file: %v\n", err)
	}
	defer db.Close()
	fmt.Println(db) // The first process running go run . will instantly print db and then wait for 10 seconds. The second process(with same go run . command) which ran right after the first process will have to wait until the 10 second timer for the first process expires, only then it will have access to that locked file and only then will it able to fmt.Println(db)
	<-time.NewTimer(10 * time.Second).C

}
