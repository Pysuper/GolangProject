/*
@Time : 2021/1/13 1:13
@Author : Zhengxingtao
@File : 21_if的使用.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	s := "屌丝"

	if s == "王思聪" {
		fmt.Println("有钱!")
	} else {
		fmt.Println("铁憨憨")
	}

	// if支持一个初始化语句，初始化语句 和 条件判断语句 用 ; 隔开
	if a := 10; a == 10 {
		fmt.Println("对的")
	}

}
