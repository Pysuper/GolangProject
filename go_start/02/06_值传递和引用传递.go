/*
@Time : 2021/1/31 13:10
@Author : Zhengxingtao
@File : 06_值传递和引用传递.go
@Software: GoLand
*/

package main

import "fmt"

/*
go 语言中的 5个 引用类型变量 ==> 浅拷贝：形参和实参都指向同一个地址
slice
map
channel
interface
func()

其他都是值类型变量 ==> 深拷贝：形参改变，实参不变，形参会开辟一个新的地址，与实参指向的地址不同
*/

func main() {
	a := 119
	b := "new string"
	c := []int{1, 2}
	d := 120
	demo(a, b, c, &d)
	fmt.Println(a, b, c, d)
}

func demo(i int, s string, arr []int, content *int) {
	i = 100
	s = "string... "
	arr[0] = 100
	arr[1] = 200
	*content = 111 // 这里使用的引用，在函数内部修改之后，会改变本身
}
