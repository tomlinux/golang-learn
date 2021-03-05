package main

import "fmt"

/*
多态概述

* 多态是指，一个父类接口可以拥有多种子类实现形态
* 这些子类形态之间有相同点，那就是它们都实现了父类接口中的方法
* 不同点则是大家对父类接口方法的实现方式各不相同，演绎何为千姿百态
* 例如，定义一个父类接口战士，凡战士者势必能攻能守，所以定义两个抽象方法：进攻与防守，但不做具体实现
* 骑兵、步兵、法师都能杀敌护国，他们是战士接口的三个子类实现
* 所以他们的共性是：都有进攻方法与防守方法
* 而他们的不同点是：具体的进攻方式各不相同，具体的防守方式也各不相同
* 所以当我拥有一个战士的集合时，我可以无视其具体类型，管你骑兵步兵还是法师弓箭手，都给劳资进攻，这就是调度它们的共性
* 当敌人的大军来势汹汹时，最好的问候就是先给他们来一轮万箭齐发，这时我要从战士的集合里找出弓箭手这个具体实现，使用弓箭手的进攻方法（而非步兵的进攻方法因为无法满足需求），这就是调度它们的个性
* 多态、共性与个性，使程序变得既丰富又简单

*/
/*

1.定义父类接口
*/
//战士接口
type Soldier interface {
	Attack() (bloodLoss int)
	Defend()
}

//2.定义子类实现一 骑兵
//骑兵：战士的实现形态一
type Rider struct{}

//骑兵的进攻与防守
func (r *Rider) Attack() (bloodLoss int) {
	fmt.Println("发动战争践踏，撞死你，踩死你，扎死你")
	return 50
}
func (r *Rider) Defend() {
	fmt.Println("不能跑，我得快点走")
}

//3.定义子类实现二
//法师：战士的实现形态二
type Master struct{}
//法师的进攻与防守
func (r *Master) Attack() (bloodLoss int) {
	fmt.Println("天崩地裂，电闪雷鸣，群星坠落，万籁俱寂，遍布下去了...")
	return 10000
}
func (r *Master) Defend() {
	fmt.Println("回城")
}


//不同场景下，时而共性，时而个性


func main() {

	//定义一个战士的集合，加入不同的子类实现
	soldiers := make([]Soldier, 0)
	soldiers = append(soldiers, new(Rider))
	soldiers = append(soldiers, new(Master))

	//强调共性时
	fmt.Println("全体进攻")
	for _, f := range soldiers {
		//所有子类实现都调用共同的父类方法
		f.Attack()
	}

	//强调个性时
	fmt.Println("敌寇势大，法师进攻，骑士防守")
	for _, s := range soldiers {

		/*      //类型断言方式1
		        switch s.(type) {
		        case *Rider:
		            s.Defend()
		        case *Master:
		            s.Attack()
		        default:
		            fmt.Println("来了个什么鬼...")
		        }*/

		//类型断言方式2
		//判断是当前s实例是不是骑士
		// 如果是:ok为true，riderPtr有值
		//如果不是：ok为false，riderPtr为nil
		if riderPtr, ok := s.(*Rider); ok {
			riderPtr.Defend()
		}
		//判断是当前s实例是不是法师
		// 如果是:ok为true，masterPtr有值
		//如果不是：ok为false，masterPtr为nil
		if masterPtr, ok := s.(*Master); ok {
			masterPtr.Attack()
		}

	}

}
