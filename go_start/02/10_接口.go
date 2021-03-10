/*
@Time : 2021/3/10 23:59
@Author : Zhengxingtao
@File : 10_接口.go
@Software: GoLand
*/

package main

import "fmt"

/*
type 接口名 interface{
	方法名(参数列表) 返回值列表
}
*/

// 接口
type Eat interface {
	eat()
}

type Live interface {
	run(run int)

	// 这里也是使用继承来实现的
	Eat
}

// 结构体
type Who struct {
	string
}

type Animate struct {
}

func (p *Who) run(run int) {
	fmt.Println(p.string, "跑步", run)

}

func (p *Who) eat() {
	fmt.Println("人的那些事儿")
}

func (a *Animate) eat() {
	fmt.Println("动物吃饭")
}
func main() {
	// 调用的时候, 先实例化一个对象
	peo := Who{"zhang"}
	peo.run(100)
	peo.eat()

	// 都是eat()方法, 但是输出确实不一样的
	ani := Animate{}
	ani.eat()
}
