/*
@Time : 2021/1/25 22:47
@Author : Zhengxingtao
@File : 02_可变参数函数.go
@Software: GoLand
*/

package main

import "fmt"

func test_5(hover ...string) {
	for i, n := range hover {
		fmt.Println(i, n)
	}
}

func main() {
	test_5("看书", "学习")
}
