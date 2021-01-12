/*
@Time : 2021/1/13 0:05
@Author : Zhengxingtao
@File : 17_输入的使用.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	var a int // 声明变量
	fmt.Printf("请输入变量a:")

	// 阻塞等待用户输入
	//fmt.Scanf("%d", &a)
	fmt.Scan(&a)

	fmt.Println(a)
}
