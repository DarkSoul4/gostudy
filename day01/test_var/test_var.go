package main

import "fmt"

//声明全局变量，可赋值不使用，局部变量必须声明赋值并使用
var s1 string

//批量声明
var (
	s2 int
	s3 bool = true
)

func foo()(int, string){
	return 10, " ee"
}

func main() {
	s1 = "qq"
	s2 = 446
	//声明变量并赋值，系统自动推导类型
	var s4 = 0.5
	s1 = "ww"
	//go语言中变量声明必须使用，否则报错
	fmt.Print(s2)				   //普通打印
	fmt.Println()
	fmt.Println(s1, s2, s3, s4)		//打印后换行
	fmt.Printf("name: %s\n", s1)	   //使用占位符打印变量
	// 简短变量声明,只能在方法中使用
	s6 := 666
	fmt.Println(s6)
	//匿名变量用_
	x, _ := foo()
	_, y := foo()
	fmt.Println(x,y)
	
}
