/*
@Time : 2021/1/15 23:28
@Author : Zhengxingtao
@File : 09_sort排序.go
@Software: GoLand
*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	// int类型的切片排序
	num := []int{23, 54, 27, 56, 34}

	// 升序排序
	sort.Ints(num)
	fmt.Println(num)

	// 类型转换
	var a int8 = 1
	var b int64 = 2
	// a =b 不同类型，无法直接赋值
	b = int64(a)
	fmt.Println(b)

	// 降序排序
	// 类型转换之后再排序
	// sort.IntSlice(num) ==> 进行类型转换
	// sort.Reverse() ==> 指明要对哪个对象进行排序
	// sort.Sort() ==> 降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Println(num)

	sort.Sort(sort.IntSlice(num)) // 直接使用sort.Sort()进行升序排序 等效于上面的sort.Ints(num)
	fmt.Println(num)

	fmt.Println("===============================================")

	// float类型的切片排序
	f := []float64{1.2, 3.4, 4.2, 4.1}

	// 升序排序
	sort.Float64s(f)
	fmt.Println(f)

	sort.Sort(sort.Float64Slice(f))
	fmt.Println(f)

	// 降序排序
	sort.Sort(sort.Reverse(sort.Float64Slice(f)))
	fmt.Println(f)

	fmt.Println("===============================================")

	// string类型的切片排序
	// 首字符对应码表中的数字
	str_sli := []string{"张三", "李四", "王二", "麻子"}

	// 升序排序
	sort.Strings(str_sli)
	fmt.Println(str_sli)

	sort.Sort(sort.StringSlice(str_sli))
	fmt.Println(str_sli)

	// 降序排序
	sort.Sort(sort.Reverse(sort.StringSlice(str_sli)))
	fmt.Println(str_sli)

	// 必须是一个升序排序的切片
	// 如果存在：返回指定值，在切片中的索引
	// 如果不存在：返回指定值，插入到什么位置才能保证依旧是升序排序的切片
	str_ := sort.SearchStrings(str_sli, "王二")
	fmt.Println(str_)

}
