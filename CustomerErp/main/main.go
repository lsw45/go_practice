package main

import (
	"customer/code/view"
)

func main() {
	var vcv *view.CustomerView
	vcv = view.CreateCustomerView()
	vcv.CustomerInforManger()
}
