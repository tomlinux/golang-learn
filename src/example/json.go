package main

import (
	"encoding/json"
	"fmt"
)

/*
定义结构体
*/
type Person struct {
	Name    string
	Age     int
	Sex     bool
	Hobbies []string
}

/*
定义数据
*/
var (
	zhangqd, mutz Person //单个map对象
	ps            []Person //多个map对象
	mMap, yMap    map[string]interface{}
	mSlice        []map[string]interface{}
)

/*
初始化数据
 */
func init() {
	//初始化结构体数据
	zhangqd = Person{"张全蛋", 20, true, []string{"质检", "维护流水线", "打饭"}}
	mutz = Person{"穆铁花", 23, false, []string{"搞破坏", "维护流水线", "打人"}}
	ps = make([]Person, 0)
	ps = append(ps, zhangqd, mutz)

	//初始化map数据
	mMap = make(map[string]interface{})
	mMap["name"] = "张全蛋"
	mMap["age"] = 20
	mMap["sex"] = true
	mMap["hobbies"] = []string{"质检", "维护流水线", "打饭"}
	yMap = make(map[string]interface{})
	yMap["name"] = "穆铁柱"
	yMap["age"] = 23
	yMap["sex"] = false
	yMap["hobbies"] = []string{"破坏", "维护流水线", "打人"}

	//初始化切片数据
	mSlice = make([]map[string]interface{}, 0)
	mSlice = append(mSlice, mMap, yMap)
}

/*
1.序列化结构体
*/
func main021() {
	//bytes, err := json.Marshal(zhangqd)
	bytes, err := json.Marshal(ps)
	if err != nil {
		fmt.Println("序列化失败,err=", err)
	} else {
		fmt.Printf("bytes的值:%v \n", bytes)
		jsonStr := string(bytes)
		fmt.Println("序列化结果", jsonStr)
	}
	//序列化结果 [{"Name":"张全蛋","Age":20,"Sex":true,"Hobbies":["质检","维护流水线","打饭"]},{"Name":"穆铁花","Age":23,"Sex":false,"Hobbies":["搞破坏","维护流水线","打人"]}]
}
/*
2.序列化map
*/
func main022(){
	bytes, err := json.Marshal(mMap)
	if err != nil {
		fmt.Println("序列化失败,err=", err)
	} else {
		jsonStr := string(bytes)
		fmt.Println("序列化结果：", jsonStr)
	}
}
// 序列化结果： {"age":20,"hobbies":["质检","维护流水线","打饭"],"name":"张全蛋","sex":true}
/*
3.序列化切片
*/

func main(){
		bytes, err := json.Marshal(mSlice)
		if err != nil {
			fmt.Println("序列化失败,err=", err)
		} else {
			jsonStr := string(bytes)
			fmt.Println("序列化结果：", jsonStr)
		}
		//序列化结果： [{"age":20,"hobbies":["质检","维护流水线","打饭"],"name":"张全蛋","sex":true},{"age":23,"hobbies":["破坏","维护流水线","打人"],"name":"穆
		//
}
