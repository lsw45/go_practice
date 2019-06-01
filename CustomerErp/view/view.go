package view

import (
	"encoding/json"
	"fmt"
	"go_code/softwaredev/customer/code/control"
	"go_code/softwaredev/customer/code/model"
	"io/ioutil"
	"os"
	"strings"
)

/*主界面实现：
1、主界面
-----------------客户信息管理软件-----------------

                            1 添 加 客 户
                            2 修 改 客 户
                            3 删 除 客 户
                            4 客 户 列 表
                            5 退           出

                            请选择(1-5)：_
需求分析：
每个客户的信息被保存在Customer对象中。
以一个Customer类型的数组来记录当前所有的客户
每次“添加客户”（菜单1）后，客户（Customer）对象被添加到数组中。
每次“修改客户”（菜单2）后，修改后的客户（Customer）对象替换数组中原对象。
每次“删除客户”（菜单3）后，客户（Customer）对象被从数组中清除。
执行“客户列表 ”（菜单4）时，将列出数组中所有客户的信息
*/
/*
分析：
1、menu() 显示界面
2、for 循环实现重复执行的逻辑
3、switch 实现功能切换
*/
//CustomerView结构体，调用service的方法
type CustomerView struct {
	cs   *control.CustomerService
	loop bool   //功能选择
	key  string //是否退出系统标志
}

func CreateCustomerView() *CustomerView {

	return &CustomerView{
		cs:   control.CreateCustomerService(),
		loop: true,
	}
}

func (this *CustomerView) CustomerInforManger() {

	for {
		menuShow()
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.addUser()
		case "2":
			// updateUser(this)
			this.updateUser1()
		case "3":
			this.deleteUser()
		case "4":
			this.userList()
		case "5":
			this.quaryUser()
		case "7":
			this.saveExist()
		case "8":
			this.restoreData()
		case "6":
			this.exit()
		default:
			fmt.Println("输入有误！！！")
		}
		if !this.loop {
			fmt.Println("你退出了客户关系管理系统...")
			break
		}
	}
}

func (this *CustomerView) restoreData() {
	customerData := restoreFileData()

	// fmt.Println(str)
	// this.details = str
	// fmt.Println(len(customerData))
	for _, value := range customerData {
		this.cs.AddUser(&value)
	}
}

func (this *CustomerView) saveExist() {
	// 这里加入 提示是否退出，并要求是y/n
	fmt.Println("你确定要退出吗? y/n")
	var choice string
	for {
		fmt.Scanln(&choice)
		//判断choice 是不是  y/n [Y/n]
		choice = strings.ToLower(choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你确定要退出吗? y/n")
	}

	if choice == "y" {
		this.loop = false
		customerData := this.cs.UserLists()
		data, err := json.Marshal(&customerData)
		if err != nil {
			fmt.Println("序列化失败 err=", err)
			return
		}
		// fmt.Println("序列化的字符串 = ",string(data))
		// saveFileData(data)
		SaveDBData(data)
	}
}

//退出
func (this *CustomerView) exit() {
	// 这里加入 提示是否退出，并要求是y/n
	fmt.Println("你确定要退出吗? y/n")
	var choice string
	for {
		fmt.Scanln(&choice)
		//判断choice 是不是  y/n [Y/n]
		choice = strings.ToLower(choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你确定要退出吗? y/n")
	}

	//这里再去判断choice
	if choice == "y" {
		this.loop = false
	}
}
func (this *CustomerView) deleteUser() {
	index := -1
	var isDelete string
	fmt.Println("---------------------删除客户---------------------")
	fmt.Println("请选择待删除客户编号(-1退出)：")
	fmt.Scanln(&index)
	if index == -1 {
		return
	}
	fmt.Println("确认是否删除(Y/N)：")
	for {
		fmt.Scanln(&isDelete)
		if strings.ToLower(isDelete) != "y" || strings.ToLower(isDelete) != "n" {
			break
		}
		fmt.Println("确认是否删除(Y/N)：")
	}

	if strings.ToLower(isDelete) == "y" {
		if this.cs.DeleteUser(index) {
			fmt.Println("---------------------删除完成---------------------")
		} else {
			fmt.Println("---------------------删除失败---------------------")
		}
	} else {
		fmt.Println("---------------------删除失败---------------------")
	}

}

func (this *CustomerView) quaryUser() {

	fmt.Println("---------------------查询客户---------------------")
	key := -1
	fmt.Println("请选择查询客户方式(-1 退出, 0(ID)/1(name))：")
	for {
		fmt.Scanln(&key)
		if key == 0 || key == 1 || key == -1 {
			break
		}
		fmt.Println("请选择查询客户方式(0(ID)/1(name))：")
	}

	if key == 0 {
		index := -1
		fmt.Println("请选择待查询客户编号：")
		for {
			fmt.Scanln(&index)
			if index != -1 {
				break
			}
			fmt.Println("请选择待查询客户编号：")
		}
		if index != -1 {
			flag, customerData := this.cs.QuaryUserByID(index)
			if flag {
				fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
				fmt.Println(customerData.CustomerInforShow())
				fmt.Println("---------------------查询客户完成---------------------")
			} else {
				fmt.Println("---------------------查询客户失败---------------------")
			}
		}
	}
	if key == 1 {
		name := ""
		isBlurry := "n"
		fmt.Println("确认模糊查询(Y/N)：")
		for {
			fmt.Scanln(&isBlurry)
			if strings.ToLower(isBlurry) != "y" || strings.ToLower(isBlurry) != "n" {
				break
			}
			fmt.Println("确认确认模糊查询(Y/N)：")
		}
		flag := false
		if strings.ToLower(isBlurry) == "y" {
			flag = true
		}
		fmt.Println("请选择待查询客户姓名：")
		for {
			fmt.Scanln(&name)
			if name != "" {
				break
			}

			fmt.Println("请选择待查询客户姓名：")
		}
		isSucced, customerData := this.cs.QuaryUserByName(flag, name)
		if isSucced {
			fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
			for _, value := range customerData {
				fmt.Println(value.CustomerInforShow())
			}
			fmt.Println("---------------------查询客户完成---------------------")
		} else {
			fmt.Println("---------------------查询客户失败---------------------")
		}

	}

}

func (this *CustomerView) updateUser() {
	index := -1
	fmt.Println("---------------------修改客户---------------------")
	fmt.Println("请选择待修改客户编号(-1退出)：")
	fmt.Scanln(&index)
	if index == -1 {
		return
	}
	var name, gender, phone, email string
	var age int
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入性别（男/女）")
	fmt.Scanln(&gender)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入电话")
	fmt.Scanln(&phone)
	fmt.Println("请输入邮箱")
	fmt.Scanln(&email)
	Customers := model.CreateCustomer(index, name, gender, age, phone, email)
	if this.cs.UpdateUser(&Customers) {
		fmt.Println("---------------------修改完成---------------------")
	} else {
		fmt.Println("---------------------修改失败---------------------")
	}

}

func (this *CustomerView) updateUser1() {
	index := -1
	fmt.Println("---------------------修改客户---------------------")
	fmt.Println("请选择待修改客户编号(-1退出)：")
	fmt.Scanln(&index)
	if index == -1 {
		return
	}
	var name, gender, phone, email string
	var age int
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入性别（男/女）")
	fmt.Scanln(&gender)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入电话")
	fmt.Scanln(&phone)
	fmt.Println("请输入邮箱")
	fmt.Scanln(&email)

	_, customerData := this.cs.QuaryUserByID(index)
	if name == "" {
		name = customerData.GetName()
	}

	if gender == "" {
		gender = customerData.GetGender()
	}

	if age == 0 {
		age = customerData.GetAge()
	}

	if email == "" {
		email = customerData.GetEmail()
	}

	if phone == "" {
		phone = customerData.GetPhone()
	}

	Customers := model.CreateCustomer(customerData.GetID(), name, gender, age, phone, email)
	if this.cs.UpdateUser1(&Customers) {
		fmt.Println("---------------------修改完成---------------------")
	} else {
		fmt.Println("---------------------修改失败---------------------")
	}

}

func (this *CustomerView) addUser() {
	fmt.Println("---------------------添加客户---------------------")
	var name, gender, phone, email string
	var age int

	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入性别（男/女）")
	fmt.Scanln(&gender)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入电话")
	fmt.Scanln(&phone)
	fmt.Println("请输入邮箱")
	fmt.Scanln(&email)

	Customers := model.CreateCustomer(0, name, gender, age, phone, email)
	if this.cs.AddUser(&Customers) {
		fmt.Println("---------------------添加完成---------------------")
	} else {
		fmt.Println("---------------------添加失败---------------------")
	}
}

func (this *CustomerView) userList() {

	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	customerData := this.cs.UserLists()
	for _, value := range customerData {

		fmt.Println(value.CustomerInforShow())
	}
	fmt.Println("-------------------------客户列表完成-------------------------")
	fmt.Print("\n\n")

}

func menuShow() {
	fmt.Println("------------客户信息管理软件------------")
	fmt.Println("              1、添加用户")
	fmt.Println("              2、修改用户")
	fmt.Println("              3、删除用户")
	fmt.Println("              4、客户列表")
	fmt.Println("              5、用户查询")
	fmt.Println("              6、退出系统")
	fmt.Println("              7、保存客户信息退出")
	fmt.Println("              8、恢复客户信息")
	fmt.Println("              请选择(1-8)")
	fmt.Println("--------------------------------------")
}

func saveFileData(details []byte) {
	filePath := "customer.txt"
	//1、打开文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 066) // For read access.
	//2、关闭文件
	defer file.Close()
	CheckError(err)
	// 3、写入数据
	err = ioutil.WriteFile(filePath, details, 066)
	CheckError(err)
}

func restoreFileData() (customers []model.Customer) {
	filePath := "customer.txt"
	//1、打开文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 066) // For read access.
	//2、关闭文件
	defer file.Close()
	CheckError(err)

	//4、读取文件
	res, err := ioutil.ReadFile(filePath)
	CheckError(err)
	res = ReadDBData()
	err = json.Unmarshal(res, &customers)
	if err != nil {
		fmt.Println("unmarshal err=", err)
	}
	// fmt.Printf("反序列化struct: monster=%v\n", customers)
	return
}
