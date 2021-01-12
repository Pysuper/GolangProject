/*
@Time : 2021/1/13 0:54
@Author : Zhengxingtao
@File : 19_类型别名.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	type bigint int64 // 给int64起一个别名叫:bigint

	var a bigint // 等价于 int64 a
	a = 54
	fmt.Printf("%T\n", a)

	type (
		long int64
		char byte
	)

	var b long = 11
	var ch char = 'b'

	fmt.Printf("%T, %T", b, ch)

	fmt.Println(a, b, ch)

}
