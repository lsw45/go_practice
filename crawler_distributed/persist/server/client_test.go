package main

import (
	"../../persist"
	"../../rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	go serverRpc(host, "test1")
	time.Sleep(1 * time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := persist.Item{
		Type:  "string",
		Id:    "test1",
		Value: "xmsksksk",
	}
}
