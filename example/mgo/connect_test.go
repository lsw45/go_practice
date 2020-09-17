package mgo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

var db *mgo.Database
var session *mgo.Session

var url = "mongodb://127.0.0.1:27017/local"
var personCol = "person"
var logCol = "log"

var day time.Time
var endTime time.Time
var beginTime time.Time

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

func TestInsert(t *testing.T) {
	col := db.C(personCol)
	// 插入
	if err := col.Insert(&Person{Name: "carey", Age: 26}, &Person{Name: "wangJin", Age: 29}); err != nil {
		panic(err)
	}
}

func TestQuery(t *testing.T) {
	col := db.C(personCol)

	// 查单条
	result := Person{}
	if err := col.Find(bson.M{"name": "carey"}).One(&result); err != nil {
		// panic(err)
		fmt.Println("err:", err)
	}
	fmt.Printf("%+v\n", result)
	// 按id查
	id := "5d9098a77cb1ed9a6c1c355c"
	objId := bson.ObjectIdHex(id)
	col.Find(bson.M{"_id": objId}).One(&result)
	fmt.Printf("%+v\n", result)
	// 或者FindId()
	col.FindId(objId).One(&result)
	fmt.Printf("%+v\n", result)

	// 查多条
	perA := []Person{}
	if err := col.Find(bson.M{"name": "carey"}).Limit(5).Skip(0).All(&perA); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", perA)

	// !=
	col.Find(bson.M{"name": bson.M{"$ne": "carey"}}).All(&perA)
	fmt.Printf("!=:%+v\n", perA)

	// >
	col.Find(bson.M{"_id": bson.M{"$gt": 11111}}).All(&perA)
	fmt.Printf(">:%+v\n", perA)

	// 多条件and
	col.Find(bson.M{"name": "carey", "age": 26}).All(&perA)
	fmt.Printf("多条件and:%+v\n", perA)

	// or
	col.Find(bson.M{"$or": []bson.M{bson.M{"name": "carey"}, bson.M{"age": 28}}})
	fmt.Printf("or:%+v\n", perA)

	// "$regex"表示字符串匹配，"$options": "$i"表示不区分大小写
	col.Find(bson.M{"name": bson.M{"$regex": bson.RegEx{Pattern: "/a/", Options: "im"}}}).All(&perA)
	fmt.Printf("like:%+v\n", perA)

	// 混合查询
	day := time.Now().Format("2006-01-02")
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s 00:00:00", day), time.Local)
	beginTime := endTime.Add(-24 * time.Hour)

	beginId := bson.NewObjectIdWithTime(beginTime)
	endId := bson.NewObjectIdWithTime(endTime)
	query := bson.M{
		"name": "carey",
		"age":  27,
		"_id":  bson.M{"$gte": beginId, "$lt": endId},
	}
	col.Find(query).All(&perA)
	fmt.Printf("混合查询:%+v\n", perA)

}

func TestUpdate(t *testing.T) {
	col := db.C(logCol)
	id := "5d919f9898d7e81fd4ce73ea"
	objId := bson.ObjectIdHex(id)
	err := col.Update(bson.M{"_id": objId}, bson.M{"$set": bson.M{"trace": "ca001", "time": 12}})
	if err != nil {
		panic(err)
	}
	// 多条件
	day := time.Now().Format("2006-01-02")
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s 00:00:00", day), time.Local)
	beginTime := endTime.Add(-24 * time.Hour)

	beginId := bson.NewObjectIdWithTime(beginTime)
	endId := bson.NewObjectIdWithTime(endTime)
	query := bson.M{
		"trace": "carey",
		"time":  27,
		"_id":   bson.M{"$gte": beginId, "$lt": endId},
	}
	col.Update(query, bson.M{"$set": bson.M{"trace": "CAREY"}})

	// inc
	err = col.Update(bson.M{"trace": "ca001", "time": 12}, bson.M{"$inc": bson.M{"time": -1}})
	if err != nil {
		panic(err)
	}
	// 批量更新con.UpdateAll(selector, update)

	// 更新或插入数据con.Upsert(selector, update)

	// push
	col.Update(bson.M{"_id": objId}, bson.M{"$push": bson.M{"总裁室": "C21", "zhou": "chao"}})

	// pull
	col.Update(bson.M{"_id": objId}, bson.M{"$pull": bson.M{"zhou": "chao"}})

	log := []Log{}
	col.Find(bson.M{"_id": objId}).All(&log)
	fmt.Printf("query:%+v", log)

}
