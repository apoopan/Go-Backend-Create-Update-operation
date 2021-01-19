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
	dbName := "appversion"
	fmt.Println("connection info:", host, dbName)
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Println("session err:", err)
		os.Exit(2)
	}
	db = session.DB(dbName)
	return db
}
