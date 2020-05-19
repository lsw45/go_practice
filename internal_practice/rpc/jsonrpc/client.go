package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

const (
	URL = "localhost:12321"
)

type Args struct {
	A, B int
}

func main() {

	client, err := jsonrpc.Dial("tcp", URL)
	defer client.Close()

	if err != nil {
		fmt.Println(err)
	}

	args := Args{7, 2}
	var reply int
	err = client.Call("Arith.Multiply", &args, &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)
}
