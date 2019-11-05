package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	_ "time"
)

type int33 = int
type inth int

func main() {
	var li int33 = 23
	var rou inth = 22
	alias(li)  // int 23
	alias(rou) // 报错： cannot use rou (type inth) as type int in argument to alias
}

func alias(pa int) {
	fmt.Printf("%T\n", pa)
	fmt.Printf("%v\n", pa)

}

func GobCopy(src, dest interface{}) error {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	err := enc.Encode(src)
	if err != nil {
		return err
	}
	err = dec.Decode(dest)
	if err != nil {
		return err
	}
	return nil
}
