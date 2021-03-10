/*
@Time : 2021/1/31 12:59
@Author : Zhengxingtao
@File : 05_闭包.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	f := closure()

	// 这里返回的 f 是一个函数的引用
	fmt.Println(f())

	// 可以通过闭包的方式，访问其他函数的局部变量
	// 重新调用：会使函数局部变量 发生变化
}

func closure() func() int {
	i := 1
	return func() int {
		i += 1
		return i
	}
}
