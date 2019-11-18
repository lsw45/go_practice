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

	/*person := struct {
		a int
		b int
	}{
		a: 11, b: 11,
	}*/

	var r io.ReadCloser
	r = strings.NewReader("abcdefghijklmn")

	decoder := json.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", obj)
}

func ReadBody(resp *http.Response) {
	resBody := resp.Body
	buf := new(bytes.Buffer)
	buf.ReadFrom(resBody)
	fmt.Println(buf.String())
	fmt.Println(buf.Bytes())
}
