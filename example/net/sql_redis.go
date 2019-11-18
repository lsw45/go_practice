package main

import (
	"github.com/jmoiron/sqlx" //是我们需要用到的主要类库

	_ "github.com/go-sql-driver/mysql" //是作为MySQL的驱动程序存在的，我们只需要执行包的init方法即可

	"fmt"
)

type Person struct {
	Board       int32  `db:"board"`
	Title       string `db:"title"`
	Discussion  int    `db:"discussion"`
	Color       string `db:"color"`
	Content     string `db:"content"`
	Author      string `db:"author"`
	Editor      string `db:"editor"`
	AddedDate   string `db:"addedDate"`
	EditedDate  string `db:"editedDate"`
	Readonly    int    `db:"readonly"`
	Views       int    `db:"views"`
	Stick       int    `db:"stick"`
	StickTime   string `db:"stickTime"`
	StickBold   int    `db:"stickBold"`
	Replies     int    `db:"replies"`
	RepliedBy   string `db:"repliedBy"`
	RepliedDate string `db:"repliedDate"`
	ReplyID     int    `db:"replyID"`
	Status      string `db:"status"`
}

func main() {
	db, _ := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/chanzhi")
	fmt.Printf("%T,%v", db, db)

	defer db.Close()
	var person []Person
	err := db.Select(&person, "select board,title,discussion,color,content,author,editor,addedDate,editedDate from eps_thread where id=?;", 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T,%v", person, person)
}
