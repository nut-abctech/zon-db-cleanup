package tasks

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type replyPost struct {
	PostID  bson.ObjectId `bson:"post_id"`
	ReplyID bson.ObjectId `bson:"_id"`
}

func CleanReplyUnlinkPost(reply *mgo.Collection, post *mgo.Collection) {
	var records []replyPost
	reply.Find(bson.M{}).All(&records)
	var cleanC int
	for _, item := range records {
		c, _ := post.FindId(item.PostID).Count()
		if c == 0 {
			err := reply.RemoveId(item.ReplyID)
			if err != nil {
				log.Println(err)
			}
			cleanC++
		}
	}
	log.Printf("Clean %d reply with broken link to post", cleanC)
}
