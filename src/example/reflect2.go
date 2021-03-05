package main

import (
	"fmt"
	"os"
	"reflect"
)

/*
需求摘要

所有的商品都有一些共性，例如都有品名、价格，个性则无千无万；
自行封装出三种商品（以模拟30万种商品）
随意给出一个商品的集合，将每件商品的所有属性值输出到《品名.txt》文件中；
需求分析
该需求的难点在于，给过来的商品是什么类型都有的，每种不同类型的商品具体有些什么属性值我们完全无法预知——所以我们可以通过反射来得到这些属性的名值；
*/
type Computer struct {
	Name   string
	Price  float64
	Cpu    string
	Memory int
	Disk   int
}

type TShirt struct {
	Name  string
	Price float64
	Color string
	Size  int
	Sex   bool
}

type Car struct {
	Name  string
	Price float64
	Cap   int
	Power string
}

func WriteObj2File(obj interface{}, filename string) error {
	//打开文件
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	//获得实例信息
	oType := reflect.TypeOf(obj)
	oValue := reflect.ValueOf(obj)

	//获取obj的所有属性名值
	for i := 0; i < oType.NumField(); i++ {

		//属性名
		fName := oType.Field(i).Name

		//获得属性值的字符串形式
		fValue := oValue.FieldByName(fName).Interface()
		fValueStr := fmt.Sprint(fValue)

		//写入文件
		_, err = file.WriteString(fName + ":" + fValueStr + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	products := make([]interface{}, 0)
	products = append(products, Computer{"未来人类", 10000, "英特尔i7", 16, 1024})
	products = append(products, TShirt{"爆款T恤", 10000, "红色", 40, false})
	products = append(products, Car{"兰博比基尼", 10000, 5, "油电混合"})
	fmt.Println(products)
	//bytes, _ := json.Marshal(products)
	//fmt.Println(string(bytes))
	// 获得所有商品的名称
	//这里我们要反射获得一个叫做Name的属性
	for _, p := range products {
		//先拿到p的值（不在乎类型）
		pValue := reflect.ValueOf(p)

		//再通过p值，拿到p下面的名为Name的属性值
		nameValue := pValue.FieldByName("Name")

		//Name属性的值，反射形式化为正射形式(interface{}类型)
		nameInterface := nameValue.Interface()

		//将空接口断言为string
		name := nameInterface.(string)
		fmt.Println("+++++++++++++++++++=")
		fmt.Println(name)
		fmt.Println(p)

		//WriteObj2File(p, "D:/BJBlockChain180x/demos/W2/day4/file/"+name+".txt")
	}

	for _,p := range products{
		name2 := reflect.ValueOf(p).FieldByName("Name").Interface().(string)
		fmt.Println(name2,p)
		//WriteObj2File(p,"D:/BJBlockChain180x/demos/W2/day4/file/"+name+".txt")
	}


}
