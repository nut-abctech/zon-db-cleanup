package tasks

import (
	"log"

	"github.com/nut-abctech/zon-db-cleanup/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"github.com/nut-abctech/zon-db-cleanup/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
)

type ratePost struct {
	PostID bson.ObjectId `bson:"post_id"`
	RateID bson.ObjectId `bson:"_id"`
}

func CleanRateUnlinkPost(rate *mgo.Collection, post *mgo.Collection) {
	var records []ratePost
	rate.Find(bson.M{}).All(&records)
	var cleanC int
	for _, item := range records {
		c, _ := post.FindId(item.PostID).Count()
		if c == 0 {
			err := rate.RemoveId(item.RateID)
			if err != nil {
				log.Println(err)
			}
			cleanC++
		}
	}
	log.Printf("Clean %d rate with broken link to post", cleanC)
}
