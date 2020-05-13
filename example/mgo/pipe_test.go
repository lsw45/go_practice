package mgo

import (
	"fmt"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
	_ "time"
)

// var db *mgo.Database
// var session *mgo.Session

// var url = "mongodb://127.0.0.1:27017/local"
// var personCol = "person"
// var logCol = "log"

// type Log struct {
// 	ID    bson.ObjectId `bson:"_id"`
// 	Age   int           `bson:"age"`
// 	Time  string        `bson:"time"`
// 	Trace string        `bson:"trace"`
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

type LogPipe []struct {
	// ID       bson.ObjectId `bson:"_id"`
	Total    int      `bson:"total"`
	Currency string   `bson:"_id"`
	Year     int      `bson:"year"`
	Tags     []string `bson:"tags"`
}

func TestPipe(t *testing.T) {
	var result LogPipe

	// day, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	// var query = bson.M{"createTime": bson.M{"$gt": day}}
	// query["state"] = 1

	query := bson.M{"year": 2014}
	// query["_id"] = 2

	_ = db.C(logCol).Pipe([]bson.M{
		{"$match": query},
		{"$project": bson.M{"_id": 1, "total": 1, "tags": 1}},
		{"group": bson.M{"$total": bson.M{"$sum": "$total"}}},
	}).All(&result)

	fmt.Printf("%+v\n", result)
}
