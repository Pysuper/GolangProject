/*
@Time : 2021/1/15 23:52
@Author : Zhengxingtao
@File : 10_map类型.go
@Software: GoLand
*/

package main

import "fmt"

func main() {

	// 申明一个map类型的变量
	var m map[string]string
	fmt.Println(m == nil) // 空指针
	fmt.Printf("%p\n", m) // 0x0 --> 没有内存地址

	// 实例化一个map的对象
	// 实例化之后，map类型的变量就有内存了，并且不再是nil了
	ma := make(map[string]string)
	fmt.Println(ma == nil)
	fmt.Printf("%p\n", ma)

	// 直接使用赋值的方式创建map
	// key:value， 二者的类型必须和定义的map参数一样
	mp := map[string]string{
		"id":   "1",
		"name": "zhang",
	}
	fmt.Println(mp)

	// 增
	mp["phone"] = "1990"
	fmt.Println("增: ", mp)

	// 删
	delete(mp, "id")
	fmt.Println("删: ", mp)
	// 删除的时候，如果不存在，不报错
	delete(mp, "di")
	fmt.Println("删不存在: ", mp)

	// 改
	mp["phone"] = "1999"
	fmt.Println("改: ", mp)

	// 查
	fmt.Println("查: ", mp["name"])
	// 如果key不存在，返回value的默认值 ==> 空字符串
	fmt.Println("查不存在：", mp["pp"])
	// key--value对应的值，okk表示key是否存在
	key, okk := mp["id"]
	fmt.Println(key, okk)

	// 循环遍历map所有内容
	for key, value := range mp {
		fmt.Println(key, value)
	}
}
