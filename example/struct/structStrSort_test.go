package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"log"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

type Cat struct {
	Catty string `json:"catty,omitempty"`
	Lili  string `json:"lili,omitempty"`
}

type Fish struct {
	Mouth  []string `json:"Mouth,omitempty"`
	Tail   string   `json:"Tail,omitempty"`
	Age    int      `json:"Age,omitempty"`
	Length float32  `json:"length,omitempty"`
	Cat    []Cat    `json:"cat"`
}

type TestCommon struct {
	InCommon string
	Fish     Fish
}

type TestStruct struct {
	TestCommon
	InStruct int
}

var s = TestStruct{
	TestCommon: TestCommon{
		InCommon: "test",
		Fish: Fish{
			Mouth: []string{"mouth", "mouth1"},
			Tail:  "tail",
			Age:   23,
			// Length: 0,
			Cat: []Cat{Cat{Catty: "喵~", Lili: "lili~"}, Cat{Catty: "", Lili: "heihei~"}},
		},
	},
	InStruct: 1,
}

// []Fish{Fish{},Fish{}}

func TestGetFieldName(t *testing.T) {
	field := GetFieldName(nil, s)

	// fmt.Printf("%T", field)
	fmt.Printf("%v\n", field)
	field_map := map[string]string{}

	for _, field_name := range field {
		split := strings.Split(strings.ToLower(field_name), ":")
		fmt.Printf("%T:%v\n", split[1], split[1])

		if field_map[split[0]] == "" && split[1] != "" && split[1] != "0" {
			field_map[split[0]] = split[1]
		}
		// field_buffer.WriteString(strings.ToLower(field_name))
	}
	fmt.Printf("%v", field_map)

}

//遍历获取结构体中字段的名称
func GetFieldName(result []string, structName interface{}) []string {
	refValue := reflect.ValueOf(structName) // value
	refType := reflect.TypeOf(structName)   // type
	// t := reflect.TypeOf(structName)
	// fmt.Println("orange refValue:", refValue)
	// fmt.Println("orange refType:", refType)

	if refType.Kind() == reflect.Ptr {
		refType = refType.Elem()
		refValue = refValue.Elem()
	}

	fieldNum := refType.NumField()
	for i := 0; i < fieldNum; i++ {
		// fmt.Println("------------------------------------------------------------", i)
		// fmt.Printf("%+v\n", refType.Field(i))
		// fmt.Printf("%+v\n", refType.Field(i).Index)
		// fmt.Printf("%+v\n", refValue.Field(i))

		// fmt.Println("field type:", refType.Field(i).Type)
		// fmt.Println("field basic type:", refValue.Field(i).Kind())
		if refValue.Field(i).Kind() == reflect.Struct { //判断是结构体并且可以导出
			if refValue.Field(i).CanInterface() {
				result = GetFieldName(result, refValue.Field(i).Interface())
			}
		} else if refValue.Field(i).Kind() == reflect.Slice || refValue.Field(i).Kind() == reflect.Array { //数组
			var value reflect.Value
			for x := 0; x < refValue.Field(i).Len(); x++ {
				value = refValue.Field(i).Index(x)

				// 判断是否是结构体，如果是就递归
				if value.Kind() == reflect.Struct {
					result = GetFieldName(result, value.Interface())
				} else {
					val := formatAtom(value)
					result = append(result, refType.Field(i).Name+":"+val)
				}
			}
		} else {
			val := formatAtom(refValue.Field(i))
			result = append(result, refType.Field(i).Name+":"+val)
			// result = append(result, refValue.Field(i).String())
		}
	}
	return result
}

// 引用go语言圣经：https://wizardforcel.gitbooks.io/gopl-zh/content/ch12/ch12-02.html
/*
Interface() interface {}	将值以 interface{} 类型返回，可以通过类型断言转换为指定类型
Int() int64	将值以 int 类型返回，所有有符号整型均可以此方式返回
Uint() uint64	将值以 uint 类型返回，所有无符号整型均可以此方式返回
Float() float64	将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回
Bool() bool	将值以 bool 类型返回
Bytes() []bytes	将值以字节数组 []bytes 类型返回
String() string	将值以字符串类型返回
*/
// 为简洁起见省略了float和complex
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64) // 浮点转字符串，精度有待优化
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

//获取结构体中Tag的值，如果没有tag则返回字段值

func GetTagName(structName interface{}) []string {
	refType := reflect.TypeOf(structName) // type

	t := reflect.TypeOf(structName)
	if refType.Kind() == reflect.Ptr {
		t = refType.Elem()
	}
	if refType.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := refType.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		tagName := refType.Field(i).Name
		tags := strings.Split(string(t.Field(i).Tag), "\"")
		if len(tags) > 1 {
			tagName = tags[1]
		}
		result = append(result, tagName)
	}
	return result
}

func TestGetFieldByte(t *testing.T) {
	var f interface{}
	b := []byte(`[{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}]`)
	json.Unmarshal(b, &f)

	fmt.Printf("%T\n", f) //[]interface {}
	// data, ok := f.([]map[string]interface{}) // 这里不能使用f.([]map[string]interface{})，这样是无法判断成功的
	data, ok := f.([]interface{})
	if ok {
		fmt.Printf("%+v\n", data)
		return
	}
}

func TestQuery(t *testing.T) {
	buf, err := Query(s)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(buf.String())
}

func Query(s interface{}, excludes ...string) (buf bytes.Buffer, err error) {
	if s == nil {
		return
	}

	v, err := query.Values(s)
	fmt.Printf("%+v\n", v)
	if err != nil {
		return buf, err
	}
	v.Del("sign")

	return QueryValues(v), nil
}

// QueryValues implements encoding of values into URL query parameters without escape
func QueryValues(v url.Values) (buf bytes.Buffer) {
	if v == nil {
		return
	}

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		vs := v[k]
		prefix := k
		for _, v := range vs {

			buf.WriteString(prefix)
			buf.WriteString(v)
		}
	}
	return buf
}
