/*
@Time : 2021/1/27 21:43
@Author : Zhengxingtao
@File : 04_函数变量.go
@Software: GoLand
*/

package main

import "fmt"

/*
函数作为参数

func main() {
	my_do(func(name string) {
		// 这里变成的匿名函数的申明
		fmt.Println(name)
	})

}

func my_do(arg func(name string))  {
	fmt.Println("我是my_do")

	// 这里实际调用
	arg("张三")
}
*/

// 函数作为返回值
func main() {
	result := a()
	num := result()
	fmt.Println(num)
}

func a() func() int {
	// 这里的func() int 是func a()的返回值
	return func() int {
		return 11
	}
}
