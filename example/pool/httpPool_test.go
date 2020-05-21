package pool_test

import (
	"../pool"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
)

func TestPutHttpPool(t *testing.T) {

}

func TestGetHttpPool(t *testing.T) {
	pool := pool.DefaultHttpPool
	con1 := pool.Get()

	con3, _ := net.Dial("http", ":8090")
	pool.Put(con3)
	if con1 == con3 {
		fmt.Println(true)
		return
	}

	fmt.Println(false)
	_ = con1.Close()

}

func DoRequest(req *http.Request) (string, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
