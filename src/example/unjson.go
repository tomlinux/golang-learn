package main

import (
	"encoding/json"
	"fmt"
)

/*
 定义human结构体
 */
type Human struct {
	Name    string
	Age     int
	Sex     bool
	Hobbies []string
}

/*
定义数据类型
 */

var jsonStr,jsonStr2,jsonStr3 string
var zqd Human
var theMap map[string]interface{}
var theSlice []map[string]interface{}


/*
初始化数据
 */
func init() {
	jsonStr = "{\"age\":20,\"hobbies\":[\"质检\",\"维护流水线\",\"打饭\"],\"name\":\"张全蛋\",\"sex\":true}"
	jsonStr2 = "[{\"age\":20,\"hobbies\":[\"质检\",\"维护流水线\",\"打饭\"],\"name\":\"张全蛋\",\"sex\":true},{\"age\":23,\"hobbies\":[\"破坏\",\"维护流水线\",\"打人\"],\"name\":\"穆铁柱\",\"sex\":false}]"
	jsonStr3 = "{\"age\":20,\"hobbies\":[\"质检\",\"维护流水线\",\"打饭\"],\"name\":\"张全蛋\",\"sex\":true}"
}
// json 压缩转义小工具 ： https://www.json.cn/json/jsonzip.html
//func main()  {
//	fmt.Println(jsonStr3)
//}

//1.反序列化为结构体
func main() {
	err := json.Unmarshal([]byte(jsonStr), &zqd)
	if err!=nil{
		fmt.Println("反序列化失败，err=",err)
	}else {
		fmt.Printf("反序列化成功：%#v\n",zqd)
	}
	//反序列化成功：main.Human{Name:"张全蛋", Age:20, Sex:true, Hobbies:[]string{"质检", "维护流水线", "打饭"}}
}


//2.反序列化为map
func main22() {
	err := json.Unmarshal([]byte(jsonStr), &theMap)
	if err!=nil{
		fmt.Println("反序列化失败，err=",err)
	}else {
		fmt.Printf("反序列化成功：%#v\n", theMap)
	}
	//反序列化成功：map[string]interface {}{"age":20, "hobbies":[]interface {}{"质检", "维护流水线", "打饭"}, "name":"张全蛋", "sex":true}
}

//3.反序列化为切片
func main23() {
	err := json.Unmarshal([]byte(jsonStr2), &theSlice)
	if err!=nil{
		fmt.Println("反序列化失败，err=",err)
	}else {
		fmt.Printf("反序列化成功：%v\n", theSlice)
	}
	//反序列化成功：[map[name:张全蛋 sex:true age:20 hobbies:[质检 维护流水线 打饭]] map[age:23 hobbies:[破坏 维护流水线 打人] name:穆铁柱 sex:false]]
}
