/*
@Time : 2021/1/11 23:09
@Author : Zhengxingtao
@File : 13_字符串类型.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	// 声明变量
	var str_1 string
	str_1 = "时间啊"

	fmt.Println(str_1)

	// 自动推导类型
	str_2 := "怎么啦"
	fmt.Printf("%T\n", str_2)

	// 内建函数， len()可以直接测字符串的长度，有多少个字符
	fmt.Println(len(str_2))
}
