/*
@Time : 2021/3/10 22:54
@Author : Zhengxingtao
@File : 07_结构体.go
@Software: GoLand
*/

package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func main() {
	//jiegou()
	zhizheng()
}

func jiegou() {
	var peo People

	// 赋值
	// 方式一
	//peo = People{"zhang", 17}
	//fmt.Println(peo)

	// 方式二
	//peo2 := People{Name: "zheng"}
	//fmt.Println(peo2)

	// 方式三
	peo.Name = "zhang"
	peo.Age = 12
	fmt.Println(peo)

	// 当实例化两个结构体之后，如果两个结构体，所有的属性的值都是一样的，那么这两个结构体就是相等的
	// 当时两个结构体指向的内存空间，还是不一样的
}

func zhizheng() {
	// 结构体指针
	// 使用new() 创建一个结构体指针 ==> 指向的是一个内存地址
	peo := new(People)
	fmt.Println(peo)

	peo1 := &People{"zheng", 18}
	peo2 := &People{"zheng", 18}
	fmt.Printf("%p %p \n", peo1, peo2)
	fmt.Println(peo1 == peo2)

	peo3 := People{"zheng", 18}
	fmt.Println(*peo1 == peo3)
}
