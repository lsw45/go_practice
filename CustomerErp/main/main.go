package main

import (
	"go_code/softwaredev/customer/code/view"
)

func main() {
	var vcv *view.CustomerView
	vcv = view.CreateCustomerView()
	vcv.CustomerInforManger()
}
