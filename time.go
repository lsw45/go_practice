package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Item 购物车中的每个菜品
type Item struct {
	Count         int    `json:"count"`           // 份数
	Size          string `json:"size,omitempty"`  // 规格
	Tasty         string `json:"tasty,omitempty"` // 口味
	Tip           bool   `json:"tip"`             // 单品备注
	Goods         Goods  `json:"goods"`           // 商品信息
	Price         int64  `json:"price"`           // 价格
	OriginalPrice int64  `json:"originalPrice"`   // 原价
	MealBoxFee    int64  `json:"mealBoxFee"`      // 餐盒费
	Content       string `json:"content"`         // 商品内容
	ContentJp     string `json:"contentJp"`
	ContentEng    string `json:"contentEng"`
	ContentHash   string `json:"contentHash"` // 内容hash
}

type Goods struct {
	Id             string        `json:"goodsId,omitempty" bson:"_id" description:"商品id"`
	No             string        `json:"no,omitempty" bson:"no,omitempty" required:"true" description:"商品编号，即对接系统id"`
	Name           string        `json:"name,omitempty" bson:"name,omitempty" required:"true" description:"商品名称"`
	NameEng        string        `json:"nameEng,omitempty" bson:"nameEng,omitempty" description:"商品英文名称"`
	NameJp         string        `json:"nameJp,omitempty" bson:"nameJp,omitempty" description:"商品日文名称"`
	Logo           string        `json:"logo,omitempty" bson:"logo,omitempty" description:"商品图片"`
	Score          int8          `json:"score,omitempty" bson:"score,omitempty" description:"评分"`
	SuggestReason  string        `json:"suggestReason,omitempty" bson:"suggestReason,omitempty" description:"推荐建议"`
	Unit           string        `json:"unit,omitempty" bson:"unit,omitempty" description:"商品单位，例如个、盒等"`
	Price          int64         `json:"price" bson:"price,omitempty" required:"true" description:"价格"`
	OriginalPrice  int64         `json:"originalPrice,omitempty" bson:"originalPrice" description:"原价"`
	Spicy          string        `json:"spicy,omitempty" bson:"spicy,omitempty" description:"辣度"`
	Description    string        `json:"description,omitempty" bson:"description,omitempty" description:"说明"`
	DescriptionEng string        `json:"descriptionEng,omitempty" bson:"descriptionEng,omitempty" description:"英文说明"`
	DescriptionJp  string        `json:"descriptionJp,omitempty" bson:"descriptionJp,omitempty" description:"日文说明"`
	Status         string        `json:"status,omitempty" bson:"status,omitempty" required:"true" description:"状态，在售 available，停售soldStop 等"`
	Size           string        `json:"size,omitempty" bson:"size,omitempty"  description:"规格，大份/小份"`
	Quota          int           `json:"quota,omitempty" bson:"quota,omitempty" description:"销量限制"`
	CreateTime     *time.Time    `json:"createTime,omitempty" bson:"createTime,omitempty"`
	UpdateTime     *time.Time    `json:"updateTime,omitempty" bson:"updateTime,omitempty"`
	ShopId         string        `json:"shopId,omitempty" bson:"shopId,omitempty" required:"true" description:"门店id"`
	GoodsScopeId   string        `json:"goodsScopeId,omitempty" bson:"goodsScopeId,omitempty" description:"商品分类id"`
	GoodsScopeName string        `json:"goodsScopeName,omitempty" bson:"goodsScopeName,omitempty" required:"true" description:"商品分类名称"`
	IsCombo        bool          `json:"isCombo" bson:"isCombo" required:"true" description:"是否套餐，默认不是"`
	ComboDetails   []*ComboGroup `json:"comboDetails" bson:"comboDetails,omitempty" description:"套餐明细列表"`
	Flag           string        `json:"-" bson:"flag,omitempty"` // 批量添加时标记

	Available bool `json:"-" bson:"available"` // 是否可用, 在售和今日售完是true, 停售是false， 排序用字段

	Sort      int  `json:"-" bson:"sort"`      // 商品排序
	ScopeSort *int `json:"-" bson:"scopeSort"` // 分类的排序
}

type ComboGroup struct {
	Name         string       `json:"name" bson:"name,omitempty" required:"true" description:"商品名称"`
	NameEng      string       `json:"nameEng,omitempty" bson:"nameEng,omitempty" description:"商品英文名称"`
	NameJp       string       `json:"nameJp,omitempty" bson:"nameJp,omitempty" description:"商品日文名称"`
	PickCount    int          `json:"pickCount" bson:"pickCount" description:"选择数量"`
	MultiChecked bool         `json:"multiChecked" bson:"multiChecked" description:"可以重复选，默认不可重复选"`
	ComboItems   []*ComboItem `json:"comboItems" bson:"comboItems,omitempty" description:"套餐商品列表"`
}

// Item 购物车中的每个菜品
type ComboItem struct {
	Count      int   `json:"count" bson:"count" description:"菜品数量"`
	ExtraPrice int64 `json:"extraPrice" bson:"extraPrice" description:"加价"`
	Checked    bool  `json:"checked" bson:"checked" description:"默认选择，false为默认选择"`
	Goods
}

func TestXX(t *testing.T) {
	p := fmt.Println

	item := Item{}
	fmt.Printf("%+v", item)
	fmt.Printf("1111111111%v2222\n", item.Tip)
	good, _ := json.Marshal(item)
	p(string(good))
	os.Exit(2)

	// 这里有一个根据RFC3339来格式化日期的例子
	t := time.Now()
	fmt.Printf("%T", t)
	p(t.Format("2006-01-02T15:04:05Z07:00"))

	// Format 函数使用一种基于示例的模式匹配方式，
	// 它使用已经格式化的时间模式来决定所给定参数
	// 的输出格式

	p("3:04PM:")
	p(t.Format("3:04PM"))
	p("Mon Jan _2 15:04:05 2006:")
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p("2006-01-02T15:04:05.999999-07:00:")
	p(t.Format("20060102150405"))
	// 对于纯数字表示的时间来讲，你也可以使用标准
	// 的格式化字符串的方式来格式化时间
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// 时间解析也是采用一样的基于示例的方式
	withNanos := "2006-01-02T15:04:05.999999999-07:00"
	t1, e := time.Parse(
		withNanos,
		"2012-11-01T22:08:41.117442+00:00")
	p(t1)
	kitchen := "3:04PM"
	t2, e := time.Parse(kitchen, "8:41PM")
	p(t2)

	// Parse将返回一个错误，如果所输入的时间格式不对的话
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)

	// 你可以使用一些预定义的格式来格式化或解析时间
	p(t.Format(time.Kitchen))
}
