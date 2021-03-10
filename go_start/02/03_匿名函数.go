/*
@Time : 2021/1/27 20:56
@Author : Zhengxingtao
@File : 03_匿名函数.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	// 申明再其他函数的内容

	// 无参数匿名函数
	func() {
		fmt.Println("我是匿名函数！！！")
	}() // 这里立马进行调用

	// 有参数的匿名函数
	func(s string) {
		fmt.Println(s)
	}("我是有参数的匿名函数")

	// 有参数 有返回值的匿名函数
	xx := func(x int) int {
		x += 10
		return x
	}(23)
	fmt.Println(xx)
}
