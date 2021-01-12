/*
@Time : 2021/1/13 1:00
@Author : Zhengxingtao
@File : 20_运算符.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	fmt.Println(4 > 3)
	fmt.Println(4 == 3)

	// 取反
	fmt.Println(!(4 > 3))
	fmt.Println(!(4 == 3))

	// 与，并且，都为真
	fmt.Println((4 > 3) && (4 == 3))

	// 或，或者，一个为真
	fmt.Println((4 > 3) || (4 == 3))

	// go的bool和int不兼容
	a := 10
	fmt.Println((0 < a) && (a < 20))
}
