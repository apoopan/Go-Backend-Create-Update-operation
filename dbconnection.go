//dbconnection.go
package main

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
)

// for mongodb
var db *mgo.Database

//GetMongoDB is ...
func GetMongoDB() *mgo.Database {
	host := "MONGO_HOST"
	dbName := "admin"
	fmt.Println("connection info:", host, dbName)
	//session, err := mgo.Dial("mongodb://localhost:27012,localhost:27013,localhost:27011")
	session, err := mgo.Dial("mongodb://test-user:test-pass@localhost:27012,localhost:27013,localhost:27011")
	if err != nil {
		fmt.Println("session err:", err)
		os.Exit(2)
	}
	db = session.DB(dbName)
	return db
}
