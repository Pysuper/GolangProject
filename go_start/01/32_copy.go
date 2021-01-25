/*
@Time : 2021/1/15 23:15
@Author : Zhengxingtao
@File : 08_copy.go
@Software: GoLand
*/

package main

import "fmt"

func main() {

	// copy不会影响目标切片的长度
	s1 := []int{1, 2}
	s2 := []int{3, 4, 5, 6}

	copy(s2, s1) // 短 --> 长
	fmt.Println(s2)

	copy(s1, s2) // 长 --> 短
	fmt.Println(s1)

	copy(s1, s2[1:3]) // 这里的s2[1:3]也是一个切片，同样可以这样操作的
	fmt.Println(s1)

	// 使用copy完成删除元素
	// copy可以保证原切片内容不变
	j := 2 // 要删除元素的索引
	num := []int{1, 2, 3, 4, 5, 6}
	newSlice := make([]int, j)
	copy(newSlice, num[0:j])
	newSlice = append(newSlice, num[j+1:]...)
	fmt.Println(newSlice)
}
