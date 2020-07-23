package ptrAndint

import (
	"encoding/json"
	"testing"
)

var resp = `{"payload":{"expires_in":1595110967},"status_code":0}`

var pay = `{
    "status_code": 0,
    "payload": {
        "code": 200,
        "message": "success",
        "data": {
            "ticket_id": "311093981649832687",
            "tp_order_id": "202007141540",
            "flow_status": "CANCELING"
        }
    }
}`

type Resp1 struct {
	StatusCode *int        `json:"status_code"` // 反序列化,没有这个字段，会赋nil
	Payload    interface{} `json:"payload,omitempty"`
}

type Resp2 struct {
	StatusCode int         `json:"status_code,omitempty"` // 反序列化,没有这个字段，会赋0
	Payload    interface{} `json:"payload,omitempty"`
}

type Payload struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		TicketID   string `json:"ticket_id"`
		TpOrderID  string `json:"tp_order_id"`
		FlowStatus string `json:"flow_status"`
	} `json:"data"`
}

func TestPtr(t *testing.T) {
	data := &Resp1{}
	err := json.Unmarshal([]byte(resp), data)
	t.Log(err)
	t.Logf("%+v", data)             //&{StatusCode:0xc00000a3f8 Payload:map[expires_in:1.595110967e+09]}
	t.Logf("%+v", *data.StatusCode) // 0
}

func TestInt(t *testing.T) {
	data := &Resp2{}
	err := json.Unmarshal([]byte(resp), data)
	t.Log(err)
	t.Logf("%+v", data) //&{StatusCode:0 Payload:map[expires_in:1.595110967e+09]}
}

func TestPay(t *testing.T) {
	data := &Resp2{}
	payl := &Payload{}
	data.Payload = payl
	err := json.Unmarshal([]byte(pay), data)
	t.Log(err)
	if payl == nil {
		t.Logf("%+v", payl)
	}
}
