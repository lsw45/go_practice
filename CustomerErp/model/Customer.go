package model

import (
	"fmt"
)

//Customer结构体
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//新建结构体的方法，相当于Java中构造函数

func CreateCustomer(id int, name string, gender string, age int, phone string, email string) Customer {

	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//客户信息显示，相当于Java中tosting 方法
func (this *Customer) CustomerInforShow() string {
	str := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",
		this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return str

}

/*各个字段的get set方法*/

func (this *Customer) GetID() int {
	return this.Id
}

func (this *Customer) SetID(id int) {
	this.Id = id
}

func (this *Customer) GetName() string {
	return this.Name
}

func (this *Customer) SetName(name string) {
	this.Name = name
}

func (this *Customer) GetAge() int {
	return this.Age
}

func (this *Customer) SetAge(age int) {
	this.Age = age
}

func (this *Customer) GetGender() string {
	return this.Gender
}

func (this *Customer) SetGender(gender string) {
	this.Gender = gender
}

func (this *Customer) GetPhone() string {
	return this.Phone
}

func (this *Customer) SetPhone(phone string) {
	this.Phone = phone
}

func (this *Customer) GetEmail() string {
	return this.Email
}

func (this *Customer) SetEmail(email string) {
	this.Email = email
}
