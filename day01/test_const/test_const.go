package main

import "fmt"


//常量批量声明
const (
	s1 = 100
	s2 = "qq"
	s3
)
//iota使用, 出现const关键字iota置0，之后每出现一行常量声明iota+1
const (
	a1 = iota	//0
	a2 = 100	
	a3 = iota	//2
)

func main() {
	
	fmt.Print(s2)				   //普通打印
	fmt.Println()
	fmt.Println(s1, s2, s3)		//打印后换行
	fmt.Printf("name: %d\n", s1)	   //使用占位符打印变量
	// 简短变量声明,只能在方法中使用
	fmt.Println(a1,a2,a3)
	
}