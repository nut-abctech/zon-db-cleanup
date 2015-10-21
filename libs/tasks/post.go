package tasks

import (
	"log"

	"github.com/nut-abctech/zon-db-cleanup/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"github.com/nut-abctech/zon-db-cleanup/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
)

func DeleteUncompletedPost(post *mgo.Collection) {
	info, err := post.RemoveAll(bson.M{"text": bson.M{"$exists": false}})
	if err != nil {
		log.Println(err)
	}
	log.Printf("Wipeout %d uncompleted posts data", info.Removed)
}
