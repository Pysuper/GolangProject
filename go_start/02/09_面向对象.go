/*
@Time : 2021/3/10 23:28
@Author : Zhengxingtao
@File : 09_面向对象.go
@Software: GoLand
*/

package main

import "fmt"

/*
封装
封装主要体现在两个方面封装数据、封装业务
Go语言中通过首字母大小控制访问权限属性首字母小写对外提供访问方法是封装数据最常见的实现方式
可以通过方法封装业务
	提出方法是封装
	控制结构体属性访问对外提供访问方法也是封装
在面向对象中封装的好处:
	安全性结构体属性访问受到限制必须按照特定访问渠道
	可复用性封装的方法实现可复用性
	可读写多段增加代码可读性

继承
一、
	按图传统面对想就是把一类事物出共间为让子美可以美的可内容
	继有多种实现方式
		通过关键字继合实方式
		合式松合继方式
	使用过java或C的应该知道量少用是使用组合代承，可以使用高内聚内耦合
	java支付在一次采的到候也说过如果给一次机会重新做java，他最希望修改的地方就是继承
	Go语言中的承是通过组合实现

二、匿名属性
	在Go语言中支持匿名属性(结构体中属性名字)，但是是每个最多只能存在匿名属性
	编译器认为类型就是属性名，我们在使用时就把类型当做属性名进行使用

多态
*/

// 匿名属性
type Other struct {
	string
	int
}

type Student struct {
	classroom string
	//peo       Other
	Other
}

// 通过这种方式完成 继承
type Teacher struct {
	classroom string
	//peo       Other
	Other
}

func main() {
	// 实例化对象
	stu := Student{"302", Other{"zhang", 12}}
	fmt.Println(stu)
	fmt.Println(stu.classroom)
	fmt.Println(stu.string)

	tea := Student{"302", Other{"zhang", 12}}
	fmt.Println(tea)
	fmt.Println(tea.classroom)
	fmt.Println(tea.string)
}
