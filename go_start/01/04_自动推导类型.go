package main // 必须有一个main包

import "fmt" // fmt.Println()，导入包后必须要使用

func main() {
	// 赋值，赋值前，必须先声明变量
	var a int
	a = 10
	fmt.Println(a)

	// :== 自动推导类型：先声明b的，再给b赋值为20，go自动推导b的类型
	b := 20
	fmt.Println(b)

	// b:=30 前面已经有变量b，不能再新建同名变量
}
