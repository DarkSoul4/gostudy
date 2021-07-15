package main

import (
	"fmt"
	"strings"
)


func main(){
	//字符串用"",单个字符用''
	a1 := "pillow"
	a2 := '滚'
	a3 := 'p'
	fmt.Println(a1)
	fmt.Printf("%c, %c", a2, a3)
	a4 := `
	第一行
	第二行
	第三行`
	fmt.Println(a4)
	//字符串拼接
	a5 := "你说"
	a6 := "是否"
	a7 := a5 + a6
	fmt.Println(a5, a6, a7)
	//分隔
	a8 := strings.Split(a7, "是")
	fmt.Println(a8)
	// 字符串修改
	b1 := "你好世界"
	b2 := []rune(b1)
	fmt.Println(string(b2))
	b2[0] = 'h'
	fmt.Println(string(b2))
	b3 := "hello沙河小王子"
	b4 := []rune(b3)
	for _, c := range b4{
		fmt.Printf("%c", c)
	}
}
