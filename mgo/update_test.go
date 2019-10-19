package mgo

import (
	"fmt"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

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

func TestUpdateAll(t *testing.T) {

}

func TestUpdateId(t *testing.T) {

}

func TestUpsert(t *testing.T) {

}

func TestUpsertId(t *testing.T) {

}

func TestUpdateTime(t *testing.T) {
	fromDate := time.Date(2014, time.November, 4, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2014, time.November, 5, 0, 0, 0, 0, time.UTC)

	var sales_his []Log
	err := db.C(logCol).Find(
		bson.M{
			"sale_date": bson.M{
				"$gt": fromDate,
				"$lt": toDate,
			},
		}).All(&sales_his)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", sales_his)
}
