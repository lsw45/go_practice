package mgo

import (
	"fmt"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

var db *mgo.Database
var session *mgo.Session

// mongodb://angrycard:angrycard@114.80.87.245:27019/angrycard", 100*time.Second
var url = "mongodb://127.0.0.1:27017/local"
var personCol = "person"
var logCol = "log"

type Log struct {
	ID    bson.ObjectId `bson:"_id"`
	Age   int           `bson:"age"`
	Time  string        `bson:"time"`
	Trace string        `bson:"trace"`
}

type Person struct {
	Name string
	Age  int
}

/*
func TestMain(m *testing.M) {
	fmt.Println("begin……………………………………………………………………………………………………………………")
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// 强调实时性——一般用于多实例
	session.SetMode(mgo.Strong, true)
	// session.SetSafe(&mgo.Safe())

	db = session.DB("local")

	m.Run()
	fmt.Println("end…………………………………………………………………………………………………………………………")
}
*/
func TestUpdate(t *testing.T) {
	col := db.C(logCol)
	// 插入
	if err := col.Insert(&Person{Name: "carey", Age: 26}, &Person{Name: "wangJin", Age: 29}); err != nil {
		panic(err)
	}
}

func TestUpdateAll(t *testing.T) {

}

func TestUpdateId(t *testing.T) {

}

func TestUpsert(t *testing.T) {

}

func TestUpsertId(t *testing.T) {

}

func TestUpdateTime() {
	fromDate := time.Date(2014, time.November, 4, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2014, time.November, 5, 0, 0, 0, 0, time.UTC)

	var sales_his []Sale
	err = c.Find(
		bson.M{
			"sale_date": bson.M{
				"$gt": fromDate,
				"$lt": toDate,
			},
		}).All(&sales_his)
}
