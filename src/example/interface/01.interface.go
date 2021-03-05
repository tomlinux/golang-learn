package main

import "fmt"

/*

@接口与实现

接口定义了一组共性
这些共性体现为接口的抽象方法
抽象方法就是只有方法定义，没有方法实现的方法
接口里有且只有抽象方法
接口可以有多种不同的具体子类实现
接口的作用是为子类实现提供统一的API

 */

/*
1.定义接口
*/
//定义【USB设备】接口
//接口：只包含抽象方法
type USBDevice interface {

	/*定义了三个【抽象方法】：定义而不做实现*/

	//连接电脑，返回值代表连接是否成功
	Connect() bool
	//读取数据，读到的数据
	ReadData() []byte
	//断开连接，返回值代表连接是否成功
	Disconnct() bool
}

/*
2. 定义U盘设备
 */
//USB设备的优盘实现接口方法
type Upan struct {
	//定义容量
	storage int
}

/*优盘实现USBDevice接口中定义的全部抽象方法=实现了USBDevice接口*/
func (up *Upan)Connect() bool  {
	fmt.Println("优盘连接成功")
	return true
}
func (up *Upan)ReadData() []byte  {
	fmt.Println("优盘连接成功")
	return []byte("东哥撸无码艺术.avi")
}
func (up *Upan)Disconnct() bool  {
	fmt.Println("优盘已断开连接")
	return true
}

//用子类实现去赋值父类对象
func main() {
	//定义接口对象
	var usbDevice USBDevice

	//用具体实现类去为接口赋值
	up := Upan{8 * 1024 * 1024 * 1024}
	usbDevice = &up

	//操作usb设备
	usbDevice.Connect()
	fmt.Println(usbDevice.ReadData())
	usbDevice.Disconnct()
}


