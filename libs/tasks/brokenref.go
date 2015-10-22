package tasks

import (
	"log"

	"github.com/nut-abctech/zon-db-cleanup/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"github.com/nut-abctech/zon-db-cleanup/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
)

type refPost struct {
	PostID bson.ObjectId `bson:"post_id"`
	ID     bson.ObjectId `bson:"_id"`
}

func CleanRateBrokenRef(rate *mgo.Collection, post *mgo.Collection) {
	var records []refPost
	rate.Find(bson.M{}).All(&records)
	var cleanC int
	for _, item := range records {
		c, _ := post.FindId(item.PostID).Count()
		if c == 0 {
			err := rate.RemoveId(item.ID)
			if err != nil {
				log.Println(err)
			}
			cleanC++
		}
	}
	log.Printf("Wipeout %d rates with broken link to post", cleanC)
}

func CleanReplyBrokenRef(reply *mgo.Collection, post *mgo.Collection) {
	var records []refPost
	reply.Find(bson.M{}).All(&records)
	var cleanC int
	for _, item := range records {
		c, _ := post.FindId(item.PostID).Count()
		if c == 0 {
			err := reply.RemoveId(item.ID)
			if err != nil {
				log.Println(err)
			}
			cleanC++
		}
	}
	log.Printf("Wipeout %d replies with broken link to post", cleanC)
}
