package main

import (
	"fmt"
)

func main() {

	// 多重赋值
	a, b := 10, 20
	fmt.Println(a, b)

	// 数据交换
	a, b = b, a
	fmt.Println(a, b)

	// _ 匿名变量
	// 丢弃数据不处理
	// 配合函数返回值使用，才有优势（不完全接收返回值中的数据）
	var tem int
	tem, _ = a, b
	fmt.Println(tem)
}
