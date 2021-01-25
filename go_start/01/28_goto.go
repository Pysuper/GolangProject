/*
@Time : 2021/1/14 23:40
@Author : Zhengxingtao
@File : 04_goto.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	// TODO：在for和if中的使用效果不同
	i := 6
	if i >= 6 {
		fmt.Println(i)

		// 执行goto之后，不管后面是什么，都不在执行了，等于就翻篇儿了
		goto Loop
	}

	fmt.Println("执行结束！")

Loop:
	if i < 2 {
		goto Loop1
	}
	fmt.Println("执行了Loop语句") // Loop里面的语句还是执行完的！

Loop1:
	fmt.Println("i报错?")
}
