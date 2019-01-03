package main

import(
	"fmt"
)

//利用条件运算符的嵌套来完成此题：学习成绩>=90分的同学用A表示，60-89分之间的用B表示，60分以下的用C表示。
//go不支持条件运算符，改为if else判断

func main () {
	score := 61

	if (score >= 90) {
		fmt.Println("A等")
	} else if ( score < 60 ) {
		fmt.Println("C等")
	} else {
		fmt.Println("B等")
	}

}