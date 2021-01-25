/*
@Time : 2021/1/20 23:17
@Author : Zhengxingtao
@File : 11_双向链表.go
@Software: GoLand
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 实例化list
	my_list := list.New()
	fmt.Println(my_list)

	// 添加
	my_list.PushFront("a") // 在当前的双向链表 最前面 添加一个元素 ['a']
	my_list.PushBack("b")  // 在当前的双向链表 最后面 添加一个元素 ['a', 'b']
	// Back() 最后一个元素; Front() 第一个元素
	my_list.InsertAfter("c", my_list.Back())  // 在 最后一个元素 之后添加 c
	my_list.InsertAfter("d", my_list.Front()) // 在 第一个元素 之前添加 d

	// 查看--遍历
	for e := my_list.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	// 查看
	fmt.Println(" ")

	fmt.Println(my_list.Front().Value)
	fmt.Println(my_list.Back().Value)

	// 取出单个元素
	//n := 3
	//first := my_list.Front()
	//for i := 0; i < n; i++ {
	//	fmt.Println(first.Next().Value)
	//}

	// 移动元素
	my_list.MoveBefore(my_list.Front(), my_list.Back())
	my_list.MoveAfter(my_list.Front(), my_list.Back())
	my_list.MoveToFront(my_list.Back())
	my_list.MoveToBack(my_list.Front())
	for e := my_list.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println("")

	// 删除元素
	my_list.Remove(my_list.Back())
	for e := my_list.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
}
