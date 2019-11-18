package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type PayInfo struct {
	PaymentSubjectID   int64  `json:"paymentSubjectID,omitempty"`
	PaymentSubjectName string `json:"paymentSubjectName,omitempty"`
	PayorAccount       string `json:"payorAccount,omitempty"`
	PayeeAccount       string `json:"payeeAccount,omitempty"`
	CardTypeID         string `json:"cardTypeID,omitempty"`
	DueAmount          string `json:"dueAmount,omitempty"`
	PaymentStatus      int    `json:"paymentStatus,omitempty"`
	PayWay             int    `json:"payWay,omitempty"`
}

func TestJson1(t *testing.T) {
	pay := PayInfo{
		// PaymentSubjectID:   11,
		// PaymentSubjectName: "xx",
		PayWay: 1,
	}
	fmt.Printf("%+v", pay)

	j, _ := json.Marshal(pay)
	fmt.Println(string(j))
	fmt.Println(len(j))
	fmt.Println(j[0], j[0]+32, string(j[0]), string(j[0]-32))

	var i uint8 = 10
	var s int8 = 10

	fmt.Println(i == uint8(s)) //true

}
