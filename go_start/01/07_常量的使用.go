package main

import "fmt"

func main() {
	// 变量：程序运行期间，可以改变的量	变量声明--> var
	// 常量：程序运行期间，不可以改变的量	常量声明--> const
	const a int = 10
	// a = 20	// err 常量不允许赋值

	const b = 20 // 没有使用:=
	fmt.Println(a)
	fmt.Println(b)
}
