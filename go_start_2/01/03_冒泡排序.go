/*
@Time : 2021/1/14 22:27
@Author : Zhengxingtao
@File : 03_冒泡排序.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	// 冒泡排序（升序 > ; 降序 < ;）
	arr := [5]int{2, 5, 3, 8, 4}

	/*
		// 把8放到最后
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		fmt.Println(arr)
		// [2 3 5 4 8]

		// 把5放到最后
		for i := 0; i < len(arr)-1-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		fmt.Println(arr)
		// [2 3 4 5 8]

		// 把4放到最后
		for i := 0; i < len(arr)-1-2; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		fmt.Println(arr)
		// [2 3 4 5 8]

		// 把3放到最后
		for i := 0; i < len(arr)-1-2; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		fmt.Println(arr)
		// [2 3 4 5 8]
	*/

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
