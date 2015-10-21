package main

import (
	"log"

	"github.com/nut-abctech/zon-db-cleanup/Godeps/_workspace/src/gopkg.in/mgo.v2" //TODO configuration
	tasks "github.com/nut-abctech/zon-db-cleanup/libs/tasks"
)

const (
	server       = "mongodb://localhost"
	databaseName = "radius"
)

var (
	session  *mgo.Session
	database *mgo.Database
)

func init() {
	var err error
	session, err = mgo.Dial(server)
	if err != nil {
		log.Panicf("Error dail server : %s", err)
	}
	log.Printf("Connect to db: %s/%s", server, databaseName)
	database = session.DB(databaseName)
}

func main() {
	log.Printf("Run clean up database")
	defer session.Close()
	tasks.CleanReplyUnlinkPost(database.C("replies"), database.C("posts"))
	tasks.CleanRateUnlinkPost(database.C("rates"), database.C("posts"))
	tasks.DeleteUncompletedPost(database.C("posts"))
}
