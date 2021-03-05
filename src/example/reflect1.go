package main

import (
	"fmt"
	"reflect"
)

/*
什么是反射？

常规的代码逻辑是：先写好剧本，再照着剧本演；
• 即：将要使用一个什么类型的对象（“演员”），然后访问和设置其什么属性，调用其什么方法，都是预先想好的；
• 这一切的前提是：你必须预先知道对象的类型！
• 而反射则提供了完全不同的方式；
• 对于运行时内存中的任何一个对象，你不需要预先知道其类型是什么，也能访问其全部属性，调用其全部方法；
• 即反射支持【运行时的动态类型检测】；
• 这极大地提高了代码的灵活性，使得我们可以在代码时去动态操作一些不可以预知的未来；
• 反射的主要作用在于编写通用的框架；
• 框架的要求就是：不管你来的是什么类型，我都支持你——这就需要用到【运行时的动态类型检测】，即反射；

反射功能一览

当我们在运行时动态捕获一个实例，我们可以：

获得检测其类型和值
获得其所有属性的类型和名值
获得其所有方法的类型和名值
访问其任意属性
调用其任意方法
# 反射应用场景举例：导出商品列表到Excel

需求是：不管用户在界面上看到什么商品列表，当它捅一下导出按钮，就将该商品的所有属性和值写出为文件；
本例的难点是：我们无法预知用户会选择导出什么类型的商品数据、它有哪些属性，也就无法相应地去创建Excel数据表的列；
因为商品的种类太多，如果用“正射”去做，那么有多少商品类型我们就要写多少个switch或if分支，然后在每一个分支里根据当前分支的具体商品类型去构造相应的数据列，这显然是狠蹩脚、狠难维护和扩展的；
而通过反射来做就易如反掌了，管你用户要导出的是什么商品实例，我可以动态地解析其类型、动态获知其所有属性和方法，然后根据再根据其具体属性名称去创建相应的表格列，并将属性值填入其中；

eg:接下来我们通过在运行时捕获一个Human2的实例（我事先并不知道类型），来说明反射的功能；

*/

//1.定义结构体
//这里我们为Human设置了2个属性+3个方法
type Human2 struct {
	Name string "姓名"
	Age  int    "年龄"
}

func (h *Human2) GetName() string {
	fmt.Println("GetName called", h.Name)
	return h.Name
}

func (h *Human2) SetName(name string) {
	fmt.Println("SetName called:", name)
	h.Name = name
}

func (h *Human2) Eat(food string, grams int) (power int) {
	fmt.Println("Eat called:", food, grams)
	return 5 * grams
}

//1.获取任意对象的类型和值
func getObjTypeValue(obj interface{}) {
	oType := reflect.TypeOf(obj) //main.Human2
	oKind := oType.Kind()        //struct
	fmt.Println("+++++++++++++")

	fmt.Println(oType, oKind)

	oValue := reflect.ValueOf(obj)
	fmt.Println(oValue)
	fmt.Println("++++++++++++")
}

//2.获得任意对象的所有属性
func getObjFields(obj interface{}) {
	oType := reflect.TypeOf(obj)
	oValue := reflect.ValueOf(obj)

	fieldsCount := oType.NumField()
	for i := 0; i < fieldsCount; i++ {

		//从对象值中获取第i个属性的值，进而值的“正射”形式
		fValue := oValue.Field(i).Interface()
		structField := oType.Field(i)
		fmt.Println(structField.Name, structField.Type, structField.Tag, fValue)
	}
}

//3.获取任意对象的所有方法
func getObjMethods(obj interface{}) {
	oType := reflect.TypeOf(obj)
	fmt.Println(oType.NumMethod())
	for i := 0; i < oType.NumMethod(); i++ {
		method := oType.Method(i)
		fmt.Println(method.Name, method.Type)
	}
}

//4.修改对象任意属性的值
func modifyFieldValue(objPtr interface{}) {
	//得到【指针的Value】
	oPtrValue := reflect.ValueOf(objPtr)
	//得到【指针所指向的值（结构体）的Value】
	oValue := oPtrValue.Elem()
	fmt.Println(oValue)

	//遍历所有属性的值
	for i := 0; i < oValue.NumField(); i++ {

		//根据序号拿到【属性的Value】
		fValue := oValue.Field(i)
		//拿到属性值的原生类型
		fKind := fValue.Kind()
		//fmt.Println(fKind)

		//根据不同的原生类型设置为不同的值
		switch fKind {
		case reflect.String:
			fValue.SetString("张全蛋")
		case reflect.Int:
			fValue.SetInt(99)
		case reflect.Bool:
			fValue.SetBool(false)
		default:
			fmt.Println("设置为其他值...")
		}
	}

}

//5.动态调用对象的任意方法
func callMethods(objPtr interface{}) {

	//要通过对象的oType拿取方法的参数列表(oType.In(i))
	oType := reflect.TypeOf(objPtr)
	//要通过对象的oValue拿取方法的具体实现(methodValue)
	oValue := reflect.ValueOf(objPtr)

	//根据方法的数量进行遍历
	for i := 0; i < oType.NumMethod(); i++ {

		//预定义要传值的参数切片[]Value
		args := make([]reflect.Value, 0)

		//从对象的oType拿到当前的方法的methodType
		methodType := oType.Method(i).Type

		//根据参数个数进行遍历
		//为每个参数乱怼一个同种类型反射值，丢入参数列表
		//内层循环走完时，当前方法的参数列表就形成了
		for j := 0; j < methodType.NumIn(); j++ {

			//从方法的methodType获取当前参数artType
			artType := methodType.In(j)
			//再获取参数类型artType的原生类型
			argKind := artType.Kind()

			//根据不同的参数原生类型乱怼相同类型的值
			switch argKind {

			case reflect.String:
				//获取字符串"一坨翔"的反射值Value，丢入参数列表
				args = append(args, reflect.ValueOf("一坨翔"))

			case reflect.Int:
				args = append(args, reflect.ValueOf(100))
			case reflect.Bool:
				args = append(args, reflect.ValueOf(false))
			}
		}

		//从对象值oValue中获取当前方法值methodValue
		methodValue := oValue.Method(i)

		//使用参数列表+方法值，反射调用当前方法
		methodValue.Call(args)
	}
}

//然后假设我们在程序运行时，捕获到一个Human实例，在事先根本不知情的情况下，让我们把这个实例的老底扒个底朝天吧：
func main() {

	//创建对象（在实际代码运行中可以是任意未知类型）
	h := Human2{"bill", 60}

	//获取任意对象的类型和值
	getObjTypeValue(h)

	//获得任意对象的所有属性
	getObjFields(h)

	//获取任意对象的所有方法
	getObjMethods(&h)

	//修改对象任意属性的值
	modifyFieldValue(&h)

	//动态调用对象的任意方法
	callMethods(&h)
	fmt.Println("after:h=", h)
}
