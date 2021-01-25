/*
@Time : 2021/1/25 22:21
@Author : Zhengxingtao
@File : 01_函数.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	fmt.Println("先执行main函数")
	test_1()

	x := test_2(1, 4)
	fmt.Println(x)

	y := test_3(2, 4)
	fmt.Println(y)

	i, j := test_4(5, 4)
	fmt.Println(i, j)

	// 不想接收这个返回值的时候，直接使用 _ 占位
	q, _ := test_4(5, 4)
	fmt.Println(q)
}

// 无参无返回值
func test_1() {
	fmt.Println("Go的第一个函数")
}

// 有参有返回值
func test_2(a int, b int) int {
	// 声明函数时的参数叫形参
	// 调用函数时的参数叫实参
	return a + b
}

// 只return
func test_3(a int, b int) (y int) {
	y = a + b
	return

}

// 多返回值函数
func test_4(a int, b int) (x int, y int) {
	x = a + b
	y = a - b
	return
}
