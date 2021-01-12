/*
@Time : 2021/1/11 23:14
@Author : Zhengxingtao
@File : 14_字符和字符串的区别.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	var ch byte
	var str string

	// 字符
	// 1、单引号
	// 2、往往都只有一个字符，转义字符除外
	ch = 'a'

	// 字符串
	// 1、双引号
	// 2、有一个或者多个字符组成
	// 3、字符串都隐藏了一个结束符，'\0'
	str = "什么" // + '\0'

	// 字符串切片
	fmt.Println(str[0], str[1])

	fmt.Println(ch, str)
}
