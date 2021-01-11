package main // 必须有一个main包

import "fmt" // fmt.Println()，导入包后必须要使用

func main() {
	// 变量：程序运行期间，可以改变得量

	// 1、声明格式 var 变量名 类型，变量声明了必须要使用
	// 2、只是声明，没有初始化的变量，默认值为0
	// 3、同一个{}中，声明的变量名是唯一的
	// 4、可以声明多个变量
	var a int
	a = 1

	var b int
	b = 2

	// 变量的赋值
	// 变量的初始化，声明变量的同时赋值
	var c int = 30

	// 自动推导类型
	// 必须初始化，通过初始化的值确定类型
	d := 100

	fmt.Println(a + b - c + d)
	fmt.Printf("%T", d)

}
