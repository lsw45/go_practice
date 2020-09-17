package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func TestRangMap(t *testing.T) {
	//多维map的声明与实现方法
	//方法1 初始化一个空的多维映射
	mainMapA := map[string]map[string]string{}
	subMapA := map[string]string{"A_Key_1": "A_SubValue_1", "A_Key_2": "A_SubValue_2"}
	mainMapA["MapA"] = subMapA
	fmt.Println("MultityMapA")
	for keyA, valA := range mainMapA {
		for subKeyA, subValA := range valA {
			fmt.Printf("mapName=%s	Key=%s	Value=%s\n", keyA, subKeyA, subValA)
		}
	}

	//方法2 使用make声明一个多维映射(等同一般声明)
	//var mainMap map[string]map[string]string
	mainMapB := make(map[string]map[string]string)
	//内部容器必须再次初始化才能使用
	subMapB := make(map[string]string)
	subMapB["B_Key_1"] = "B_SubValue_1"
	subMapB["B_Key_2"] = "B_SubValue_2"
	mainMapB["MapB"] = subMapB
	fmt.Println("\nMultityMapB")

	for keyB, valB := range mainMapB {
		for subKeyB, subValB := range valB {
			fmt.Printf("mapName=%s	Key=%s	Value=%s\n", keyB, subKeyB, subValB)
		}
	}

	/* 方法3 使用interface{}初始化一个一维映射
	 * 关键点：interface{} 可以代表任意类型
	 * 原理知识点:interface{} 就是一个空接口，所有类型都实现了这个接口，所以它可以代表所有类型
	 */
	//mainMapC := make(map[string]interface{})
	mainMapC := map[string]interface{}{}
	subMapC := make(map[string]string)
	subMapC["C_Key_1"] = "C_SubValue_1"
	subMapC["C_Key_2"] = "C_SubValue_2"
	mainMapC["MapC"] = subMapC
	fmt.Println("\nMultityMapC")
	for keyC, valC := range mainMapC {
		//此处必须实例化接口类型，即*.(map[string]string)
		//subMap := valC.(map[string]string)
		for subKeyC, subValC := range valC.(map[string]string) {
			fmt.Printf("mapName=%s	Key=%s	Value=%s\n", keyC, subKeyC, subValC)
		}
	}

	requestBodyMap := make(map[string]interface{})
	userInfoMap := make(map[string]interface{})
	/*OrderExtraInfoVoMap := make(map[string]interface{})
	payInfoMap := make(map[string]interface{})
	orderTypeMap := make(map[string]interface{})*/
	dishesItemMap := make(map[string]interface{})
	var itemSlice []map[string]interface{}

	requestBodyMap["isCheackOut"] = 0
	requestBodyMap["isThirdPay"] = 2
	requestBodyMap["bankCode"] = "weChat"
	requestBodyMap["isSentMsg"] = 0
	requestBodyMap["groupID"] = 1000000000
	var order map[string]interface{}
	requestBodyMap["shopID"], _ = strconv.ParseInt(order["shopId"].(string), 10, 64)
	fmt.Printf("%v", requestBodyMap)
	subMap := make(map[string]string)

	//request.shopID =
	body := `{"shopId":"2018080300077000000060712413","tableNo":"0001","userId":"5c0b2fdb62dd7e66a6b45a43","memberId":"","diningMode":"immediately","merchantCode":"2088912057676397","customerCount":3,"goodsCount":2,"totalOrigAmt":2000,"totalAmt":2000,"baskets":[{"shopId":"2018080300077000000060712413","tableNo":"0001","goodsId":"5bc18d64c209136c2898a9cf","userId":"5c0b2fdb62dd7e66a6b45a43","item":{"count":2,"user":{"userId":"5c0b2fdb62dd7e66a6b45a43","nickName":"c5d4ftmn","avatar":"https://cdn.xunliandata.com/avatar5.png","userFrom":"WEB"},"goods":{"goodsId":"5bc18d64c209136c2898a9cf","no":"000003","name":"3号套餐","nameEng":"3 set meal","nameJp":"3番セット","logo":"https://cdn.xunliandata.com/15394112986432018080300077000000060712413.jpg","price":1000,"originalPrice":1000,"spicy":"0","description":"煲仔饭默认选择黑椒牛柳煲+蛋；碗碗菜默认选择泡菜魔芋肉丝；米线面条类默认选择五香牛肉面。啦啦啦啦啦啦啦啦啦啦啦啦啦啦啦绿绿绿绿绿绿啦啦啦啦啦啦啦啦绿绿绿绿绿绿绿绿绿","descriptionEng":"Boiled rice default choice black pepper beef steak + eggs; bowl dishes default choice of kimchi konjac shredded meat; rice noodles default choice spiced beef noodles.","descriptionJp":"炊飯器は、黒椒の牛柳鍋+卵を選んでいます。茶碗のお椀料理はキムチのこんにゃく肉の糸を選んでいます。","status":"available","createTime":"2018-10-13T14:15:00.801+08:00","updateTime":"2018-11-30T10:09:25.622+08:00","shopId":"2018080300077000000060712413","goodsScopeName":"套餐","isCombo":true,"comboDetails":[{"name":"煲仔饭类","pickCount":1,"multiChecked":false,"comboItems":[{"count":0,"extraPrice":0,"checked":false,"goodsId":"5b7d2bcdc209130d245c8228","no":"B33","name":"酱香牛肉煲+桃橙果汁","logo":"https://cdn.xunliandata.com/15366546334212018080300077000000060712413.jpg","price":1,"spicy":"3","status":"available","createTime":"2018-08-22T17:24:29.562+08:00","updateTime":"2018-09-17T11:01:25.677+08:00","shopId":"2018080300077000000060712413","goodsScopeId":"5b7d2bcdc209130d245c8205","goodsScopeName":"煲仔饭类","isCombo":false,"boxType":""},{"count":0,"extraPrice":0,"checked":false,"goodsId":"5b7d2bcdc209130d245c8226","no":"B31","name":"酱香牛肉煲+蛋","logo":"https://cdn.xunliandata.com/15366544529892018080300077000000060712413.jpg","price":1,"spicy":"3","status":"available","createTime":"2018-08-22T17:24:29.564+08:00","updateTime":"2018-09-13T20:33:47.198+08:00","shopId":"2018080300077000000060712413","goodsScopeId":"5b7d2bcdc209130d245c8205","goodsScopeName":"煲仔饭类","isCombo":false,"boxType":""},{"count":1,"extraPrice":0,"checked":false,"goodsId":"5b7d2bcdc209130d245c8229","no":"B34","name":"黑椒牛柳煲+蛋","logo":"https://cdn.xunliandata.com/15366546424902018080300077000000060712413.jpg","price":1,"spicy":"1","status":"available","createTime":"2018-08-22T17:24:29.561+08:00","updateTime":"2018-09-17T11:01:35.157+08:00","shopId":"2018080300077000000060712413","goodsScopeId":"5b7d2bcdc209130d245c8205","goodsScopeName":"煲仔饭类","isCombo":false,"boxType":""},{"count":0,"extraPrice":0,"checked":false,"goodsId":"5b7d2bcdc209130d245c8208","no":"B01","name":"川香鸡丁煲+蛋","nameEng":"Sichuan chicken with egg + egg","nameJp":"川の香りと卵+卵","logo":"https://cdn.xunliandata.com/15366466134952018080300077000000060712413.jpg","price":1,"spicy":"1","description":"湘菜历来重视原料互相搭配，滋味互相渗透。湘菜调味尤重酸辣。因地理位置的关系，湖南气候温和湿润，故人们多喜食辣椒，用以提神去湿。用酸泡菜作调料，佐以辣椒烹制出来的菜肴，开胃爽口，深受青睐，独具特色","descriptionEng":"Hunan cuisine has always attached importance to the mutual matching of raw materials and mutual penetration of flavor. The flavor of Hunan cuisine is extremely sour and spicy. Because of the geographi","descriptionJp":"湖南料理は従来原料を重視し合って、味は互いに浸透しています。湖南料理は味がよくて辛いです。地理的な位置の関係のため、湖南の気候は温和で湿潤で、だから人々は多く唐辛子を食べて、神を引き立てるために濡れます。酸味のあるキムチで調味料を作って、唐辛子を使って調理した料理は、口を開けてさっぱりとして、とても人気があります。","status":"available","createTime":"2018-08-22T17:24:29.594+08:00","updateTime":"2018-11-30T10:09:25.522+08:00","shopId":"2018080300077000000060712413","goodsScopeId":"5b7d2bcdc209130d245c8205","goodsScopeName":"煲仔饭类","isCombo":false,"boxType":"useless"}]},{"name":"碗碗菜类","pickCount":1,"multiChecked":false,"comboItems":[{"count":0,"extraPrice":0,"checked":false,"goodsId":"5b7d2bcdc209130d245c8236","no":"B47","name":"笋干肉丝","logo":"https://cdn.xunliandata.com/15393241940522018080300077000000060712413.jpg","price":1,"status":"available","createTime":"2018-08-22T17:24:29.548+08:00","updateTime":"2018-10-12T14:03:16.448+08:00","shopId":"2018080300077000000060712413","goodsScopeId":"5b7d2bcdc209130d245c8206","goodsScopeName":"碗碗菜类","isCombo":false,"boxType":""},{"count":0,"extraPrice":0,"checked":false,"goodsId":"5b7d2bcdc209130d245c8237","no":"B48","name":"五彩鸡丁","logo":"https://cdn.xunliandata.com/15366547498682018080300077000000060712413.jpg","price":1,"status":"available","createTime":"2018-08-22T17:24:29.547+08:00","updateTime":"2018-11-27T19:49:10.117+08:00","shopId":"2018080300077000000060712413","goodsScopeId":"5b7d2bcdc209130d245c8206","goodsScopeName":"碗碗菜类","isCombo":false,"boxType":"default"},{"count":1,"extraPrice":0,"checked":false,"goodsId":"5b7d2bcdc209130d245c8238","no":"B49","name":"泡菜魔芋肉丝","logo":"https://cdn.xunliandata.com/15366547588952018080300077000000060712413.jpg","price":1,"status":"available","createTime":"2018-08-22T17:24:29.546+08:00","updateTime":"2018-09-11T16:33:23.142+08:00","shopId":"2018080300077000000060712413","goodsScopeId":"5b7d2bcdc209130d245c8206","goodsScopeName":"碗碗菜类","isCombo":false,"boxType":""}]},{"name":"米线面条类","pickCount":1,"multiChecked":false,"comboItems":[{"count":0,"extraPrice":0,"checked":false,"goodsId":"5b7d2bcdc209130d245c8249","no":"B66","name":"三鲜米线","logo":"https://cdn.xunliandata.com/15393243434932018080300077000000060712413.jpg","price":1,"status":"available","createTime":"2018-08-22T17:24:29.529+08:00","updateTime":"2018-10-12T14:05:45.093+08:00","shopId":"2018080300077000000060712413","goodsScopeId":"5b7d2bcdc209130d245c8207","goodsScopeName":"米线面条类","isCombo":false,"boxType":""},{"count":1,"extraPrice":0,"checked":false,"goodsId":"5b7d2bcdc209130d245c8247","no":"B64","name":"五香牛肉面","logo":"https://cdn.xunliandata.com/15393243638152018080300077000000060712413.jpg","price":1,"status":"available","createTime":"2018-08-22T17:24:29.531+08:00","updateTime":"2018-10-12T14:06:06+08:00","shopId":"2018080300077000000060712413","goodsScopeId":"5b7d2bcdc209130d245c8207","goodsScopeName":"米线面条类","isCombo":false,"boxType":""},{"count":0,"extraPrice":0,"checked":false,"goodsId":"5b7d2bcdc209130d245c8248","no":"B65","name":"杂酱米线","logo":"https://cdn.xunliandata.com/15393243558642018080300077000000060712413.jpg","price":1,"status":"available","createTime":"2018-08-22T17:24:29.53+08:00","updateTime":"2018-10-12T14:05:57.696+08:00","shopId":"2018080300077000000060712413","goodsScopeId":"5b7d2bcdc209130d245c8207","goodsScopeName":"米线面条类","isCombo":false,"boxType":""},{"count":0,"extraPrice":0,"checked":false,"goodsId":"5b7d2bcdc209130d245c8233","no":"B44","name":"板栗烧鸡","logo":"https://cdn.xunliandata.com/15366547201272018080300077000000060712413.jpg","price":1,"status":"available","createTime":"2018-08-22T17:24:29.551+08:00","updateTime":"2018-09-17T11:00:33.572+08:00","shopId":"2018080300077000000060712413","goodsScopeId":"5b7d2bcdc209130d245c8206","goodsScopeName":"碗碗菜类","isCombo":false,"boxType":""}]}],"boxType":""},"price":1000,"originalPrice":0,"mealBoxFee":3,"content":"五香牛肉面X1,泡菜魔芋肉丝X1,黑椒牛柳煲+蛋X1","contentJp":"X1,X1,X1","contentEng":"X1,X1,X1","contentHash":"D5BF40CB0B49F76296582338091C34149F40A45BFD6A747C6EE14401FE30D6C3"}}],"credits":0,"creditsDisCount":0}`
	if err := json.Unmarshal([]byte(body), &order); err != nil {
		fmt.Println("error:", err)
	}
	for k, v := range order {
		fmt.Println(k)
		if k == "baskets" && v != nil {
			for k1, v1 := range v.(map[string]interface{}) {
				if k1 == "item" && v1 != nil {
					for k2, v2 := range v1.(map[string]interface{}) {
						subMap = v2.(map[string]string)

						if k2 == "user" {
							userInfoMap["shopOpenID"] = "avaterksdjljfs9439009" //小程序/公众号对应人的openID，拉取微信支付必传
							userInfoMap["userName"] = subMap["nickName"]
							userInfoMap["shopMpID"] = "sjfsdkldl" //小程序/公众号appID，拉取微信支付必传
						}

						if k2 == "goods" {
							dishesItemMap = make(map[string]interface{})
							dishesItemMap["foodID"] = subMap["no"]
							dishesItemMap["foodName"] = subMap["name"]
							dishesItemMap["isSetFood"] = subMap["isSetFood"]
						}
					}
					if dishesItemMap != nil {
						itemSlice = append(itemSlice, dishesItemMap)
					}
				}
			}

		}
	}
}

func TestFloatStrconv(t *testing.T) {
	var v = 1.6
	fmt.Println("%v", v.Float())
}
