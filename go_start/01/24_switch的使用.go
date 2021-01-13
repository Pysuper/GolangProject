/*
@Time : 2021/1/13 21:13
@Author : Zhengxingtao
@File : 24_switch的使用.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	num := 0

	// switch 后面写的是变量本身
	switch num {
	case 1: // 判断是否为true，如果为true，则执行当前case中的语句
		fmt.Printf("%d\n", num)
		break // go保留了break关键字，跳出switch语句，默认已经包含了==> 不写
	case 2:
		fmt.Printf("%d\n", num)
		break
	case 3:
		fmt.Printf("%d\n", num)
		break
	case 4:
		fmt.Printf("%d\n", num)
		break
	default: // 最后都不符合上面的情况，则执行default里面的内容
		fmt.Println("上面的情况都为False")
	}
}
