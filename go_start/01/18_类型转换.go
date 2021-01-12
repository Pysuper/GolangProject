/*
@Time : 2021/1/13 0:09
@Author : Zhengxingtao
@File : 18_类型转换.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	var flag bool

	flag = true

	fmt.Printf("%t\n", flag)

	// 这种不能转换的类型，就叫不兼容类型
	// bool类型不能转换为int
	//fmt.Printf("%d\n", int(flag))

	// 0就是假，非0就是真
	// 整形也不能转换为bool
	//flag = bool(1)

	var ch byte
	ch = 'a' // 字符类型本质上就是整型

	var t int
	t = int(ch) // 类型转换，把ch的值取出来，转成int，再给t赋值
	fmt.Println(t)
}
