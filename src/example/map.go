package main

import "fmt"

/*
映射概述

映射（map）是键值对形成的集合
键值的类型都是任意的
可以根据键快速的查询值，而无须遍历
创建map

可以通过声明map[keytype]valuetype的方式创建map
也可以通过make(map[keytype]valuetype,size)的方式来创建，size不指定时默认长度为0
以map[key]和map[key]=value的方式可以访问和修改键值
*/

func demo81() {
	//var mmp = map[string]int{}
	//var mmp map[string]int
	//var mmp = map[string]int{"代码量":10000,"酒量":20}
	//通过内建函数make创建
	//mmp := make(map[string]int)

	//创建一个string为键，值为任意类型的map
	mmp := make(map[string]interface{})

	//丢入键值对到map中
	mmp["name"] = "西门东"
	mmp["sex"] = "male"
	mmp["hobby"] = "female"
	mmp["rmb"] = 0.5

	fmt.Printf("type=%T,value=%v,len=%d\n", mmp, mmp, len(mmp))
}

/*
删除键值对
*/
func demo82() {
	//创建一个string为键，值为任意类型的map
	mmp := make(map[string]interface{})
	//丢入键值对到map中
	mmp["name"] = "西门东"
	mmp["sex"] = "male"
	mmp["hobby"] = "female"
	mmp["rmb"] = 0.5
	//删除hobby对应的键值对
	delete(mmp, "hobby")
	fmt.Println(mmp)
}

/*
修改键值对
*/
func demo83() {
	//创建一个string为键，值为任意类型的map
	mmp := make(map[string]interface{})
	//丢入键值对到map中
	mmp["name"] = "西门东"
	mmp["sex"] = "male"
	mmp["hobby"] = "female"
	mmp["rmb"] = 0.5
	//修改特定的键值对
	mmp["hobby"] = [...]string{"抽烟", "喝酒", "去浪"}
	mmp["rmb"] = -29999.5
	mmp["最爱的书籍"] = "平凡的世界"
	fmt.Println(mmp)
}

/*

map的查询和遍历
* 可以通过value := map[key]的方式得到key对应的value，当key不存在时得到的是nil
* 为了避免访问一个并不存在的键得到空值带来的逻辑错误，可以通过value,ok:=map[key]的方式进行带校验的查询，当ok为true时代表有值，否则代表key并不存在
* range关键字对map的遍历，可以同时得到键值对的key和value，也可以仅仅得到key

*/

func demo84() {
	//创建一个string为键，值为任意类型的map
	mmp := make(map[string]interface{})
	//丢入键值对到map中
	mmp["name"] = "西门东"
	mmp["sex"] = "male"
	mmp["hobby"] = "female"
	mmp["rmb"] = 0.5

	//没有校验的查询，当key不存在时value为<nil>
	value1 := mmp["sex"]
	fmt.Println("sex=", value1)
	value1 = mmp["sex2"]
	//sex2= <nil>
	fmt.Println("sex2=", value1)

	//带有校验的查询
	value2, ok := mmp["最爱的书"]
	if ok == true {
		fmt.Println(value2)
	} else {
		fmt.Println("东哥：洒家没有最爱的书，查你妹")
	}

	//遍历键值
	for key, value := range mmp {
		fmt.Println(key, value)
	}

	//遍历所有键
	fmt.Println("key := range mmp")
	for key := range mmp {
		fmt.Println(key)
	}
}

func main() {
	demo84()
	m := map[int64]string{1:"天仙",2:"人才",3:"哈哈"}
	m[1]="小天真"
	m[2]="小仙女"
	fmt.Print(m)

}
