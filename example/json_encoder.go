package main

import (
	"encoding/json"
	"fmt"
	"io"
	_ "log"
	_ "os"
	"strings"
)

func main() {
	/*dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k, v1 := range v {
			if k != "Title" {
				v[k] = v1
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}*/
	BindJson()
}

func BindJson() {
	var obj interface{}
	/*type person struct {
		a int
		b int
	}*/
	var r io.Reader
	r = strings.NewReader("abcdefghijklmn")

	fmt.Println(r.Len())

	decoder := json.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", obj)
}
