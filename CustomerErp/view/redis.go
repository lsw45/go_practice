package view

import (
	"fmt"
	_ "fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	_ "log"
)

func SaveDBData(data []byte) {

	//1. 链接到redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	CheckError(err)
	defer conn.Close()

	//2、通过go 向redis写入数据 string类型数据
	fmt.Println(string(data))
	_, err = conn.Do("set", "customer", data)
	CheckError(err)

}

func ReadDBData() (data []byte) {
	//1. 链接到redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	CheckError(err)
	defer conn.Close()
	//3. 通过go 向redis读取数据string类型数据
	data, err = redis.Bytes(conn.Do("get", "customer"))
	CheckError(err)
	fmt.Println(string(data))

	return
}

//错误检查
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
