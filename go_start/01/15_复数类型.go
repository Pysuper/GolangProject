/*
@Time : 2021/1/12 20:36
@Author : Zhengxingtao
@File : 15_复数类型.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	var t complex128 //申明

	t = 2.1 + 3.14i // 赋值
	fmt.Println(t)

	// 自动推导类型
	t2 := 3.3 + 4.4i
	fmt.Printf("%T\n", t2)

	// 通过内建函数，获取实部和虚部
	fmt.Println(real(t2), imag(t2))
}
