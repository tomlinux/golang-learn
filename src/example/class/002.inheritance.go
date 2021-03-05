package main

import "fmt"

/*
继承性概述
* 继承的意义在于低成本地扩展和改造原有代码
* 首先，子类一行代码即可拥有父类的全部成员（属性和方法）
* 继承的目的，在于扩展和改造父类
* 扩展，是指子类发展出自己独有的新属性和新方法，以适应自身需求
* 改造，是指子类覆写和覆盖父类成员（属性和方法），以适应自身需求

*/
/*
1.定义父类Doggy
*/
type Doggy struct {
	name string
	sex  bool
}

//定义父类方法
func (d *Doggy) bite() {
	fmt.Printf("%s要咬你了啊\n", d.name)
}

/*
2.定义子类PoliceDog
*/
type PoliceDog struct {
	// 持有一个Doggy对象，并继承Doggy的全部属性和方法--->父类Doggy
	Doggy

	//警犬所独有的属性
	skill string
}

//覆写父类方法 子类可以覆盖父类的方法
func (pd *PoliceDog) bite() {
	fmt.Printf("%s:还没出嘴你就熏晕过去了，我他妈刚喝了三瓶茅台\n", pd.name)
}

//发展出新的方法
func (pd *PoliceDog) doPoliceJob() {
	fmt.Printf("%s正在执行警务工作,普通的狗狗是做不到滴", pd.name)
}

/*
3.继承了Doggy的PoliceDog拥有父类的全部属性和方法
*/
func demo31() {
	pdPtr := new(PoliceDog)
	pdPtr.name = "战狼"
	pdPtr.bite()
}
/*
4.继承的目的是为了有新的扩展
 */
func demo32() {
	pdPtr := new(PoliceDog)
	//这里是从父类继承过来的属性和方法
	pdPtr.name = "战狼"
	pdPtr.bite()
	//访问独有的属性和方法
	pdPtr.skill = "徒手接RPG"
	fmt.Println(pdPtr.skill)
	pdPtr.doPoliceJob()
}

/*
5.直接在声明子类时完成属性的重定义


*/

func main() {
	pd := PoliceDog{Doggy{"战狼II", true}, "豪饮"}
	pd.bite()
}


