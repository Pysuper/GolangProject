/*
@Time : 2021/1/14 19:58
@Author : Zhengxingtao
@File : 25_数组.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	// 二维数组
	arr := [3][3]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}
	fmt.Println(arr)

	arr0 := arr[0]
	fmt.Println(arr[0])
	fmt.Println(arr0[0], arr0[1], arr0[2]) // 取出第一组元素

	// 三维数组
	arr3 := [2][2][2]int{
		{{1, 2}, {2, 3}},
		{{3, 4}, {4, 5}},
	}
	fmt.Println(arr3)
}
