package control

import (
	_ "fmt"
	"go_code/softwaredev/customer/code/model"
	_ "strings"
)

//CustomerService结构体用于处理Model数据
type CustomerService struct {
	//新建Customer切片存储客户信息
	customers []model.Customer
	num       int
}

func (this *CustomerService) GetID() []model.Customer {
	return this.customers
}

//新建 CustomerService方法，便于view使用
func CreateCustomerService() *CustomerService {
	var cs CustomerService
	//初始化一个数据
	cs.num++
	initCustomer := model.CreateCustomer(cs.num, "张三", "男", 20, "110", "zs@sohu.com")
	cs.customers = append(cs.customers, initCustomer)
	return &cs
}

func (this *CustomerService) UserLists() []model.Customer {
	return this.customers
}

func (this *CustomerService) QuaryUserByID(id int) (flag bool, cs model.Customer) {
	id = this.findIDByIndex(id)

	if id == -1 {
		flag = false
	} else {
		flag = true
		cs = this.customers[id]
	}
	return
}

func (this *CustomerService) QuaryUserByName(isbulrry bool, name string) (flag bool, cs []model.Customer) {
	r := []rune(name)
	if isbulrry { //模糊查询  [u4e00-u9fa5]
		for _, value := range this.customers {
			//搜索中文
			res := value.GetName()
			for _, val := range res {
				// fmt.Println("val=",val)
				// fmt.Println("name=",[]rune(name))
				for i := 0; i < len(r); i++ {
					// fmt.Println("r[i]=",r[i])
					if val == r[i] {
						flag = true
						break
					} else {
						flag = false
					}
				}
				if flag {
					break
				}
			}
			if flag {
				cs = append(cs, value)
			}
			//只能搜索英文
			// if strings.Contains(value.GetName(),name) {
			//  cs = append(cs,value)
			//  flag = true
			// }else{
			//  flag = false
			// }
		}
	} else { //标准查询
		for _, value := range this.customers {

			if value.GetName() == name {
				cs = append(cs, value)
				flag = true
				break
			} else {
				flag = false
			}
		}
	}

	return
}

func (this *CustomerService) DeleteUser(index int) bool {

	index = this.findIDByIndex(index)

	if index == -1 {
		return false
	} else {

		this.customers = append(this.customers[:index], this.customers[index+1:]...)
		return true
	}
}

func (this *CustomerService) UpdateUser1(newCustomer *model.Customer) bool {
	index := newCustomer.GetID()
	index = this.findIDByID(index)

	this.customers[index] = *newCustomer
	return true
}
func (this *CustomerService) UpdateUser(newCustomer *model.Customer) bool {
	index := newCustomer.GetID()
	// fmt.Println("index:",index)
	index = this.findIDByID(index)
	// fmt.Println("index:",index)
	if index == -1 {
		return false
	} else {
		name := newCustomer.GetName()
		if name != "" {
			this.customers[index].SetName(name)
		}
		gender := newCustomer.GetGender()
		if gender != "" {
			this.customers[index].SetGender(gender)
		}
		age := newCustomer.GetAge()
		if age != 0 {

			this.customers[index].SetAge(age)
		}
		email := newCustomer.GetEmail()
		if email != "" {
			this.customers[index].SetEmail(email)
		}
		phone := newCustomer.GetPhone()
		if phone != "" {
			this.customers[index].SetPhone(phone)
		}
		return true
	}

}

func (this *CustomerService) AddUser(newCustomer *model.Customer) bool {
	this.num++
	newCustomer.SetID(this.num)
	this.customers = append(this.customers, *newCustomer)
	return true
}

func (this *CustomerService) findIDByIndex(index int) (id int) {
	for key, value := range this.customers {
		// fmt.Println("value:",value.GetID())
		if value.GetID() == index {
			id = key
			break
		} else {
			id = -1
		}
	}
	return
}
func (this *CustomerService) findIDByID(index int) (id int) {
	res := -1
	for key, value := range this.customers {
		// fmt.Println("value:",value.GetID())
		res = value.GetID()
		if res == index {
			id = key
			break
		} else {
			id = -1
		}
	}
	return
}
