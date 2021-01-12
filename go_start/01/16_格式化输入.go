/*
@Time : 2021/1/12 23:58
@Author : Zhengxingtao
@File : 16_格式化输入.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	a := 10
	b := "abc"
	c := 'a'
	d := 3.14

	// %T 变量所属类型
	fmt.Printf("%T,%T,%T,%T\n", a, b, c, d)

	// %d 整型
	// %s 字符串
	// %c 字符个数
	// %f 浮点型
	fmt.Printf("%d,%s,%c,%f\n", a, b, c, d)

	// 自动匹配格式，字符类型==>整形
	fmt.Printf("%v,%v,%v,%v\n", a, b, c, d)
}
