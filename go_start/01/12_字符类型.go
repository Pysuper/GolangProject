package main

import "fmt"

func main() {
	//	声明字符类型
	var ch byte
	ch = 97

	fmt.Println(ch)

	// 格式化输出：%c, 按照字符串形式打印，%d，以整形方式打印
	fmt.Printf("%c, %d\n", ch, ch)

	// 字符是''
	ch = 'a'
	fmt.Printf("%c, %d\n", ch, ch)

	// 大写转小写，小写转大写，大小写相差，小写大
	fmt.Printf("%d, %d\n", 'A', 'a')
	fmt.Printf("%c, %c\n", 'A', 'a')
	fmt.Printf("大写转小写：%c\n", 'A'+32)
	fmt.Printf("小写转大写：%c\n", 'a'-32)

	// 以'\'反斜杠开头的字符是转义字符，'\n'代表换行
	fmt.Printf("Hello World!\n")
	fmt.Printf("Hello Zheng!\n")
}
