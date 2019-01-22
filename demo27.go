package main

import(
	"fmt"
)

//给一个不多于5位的正整数，要求：一、求它是几位数，二、逆序打印出各位数字。

func main () {
	var num int
	fmt.Scanf("%d", &num)
	res := num / 10000
	if (res != 0) {
		fmt.Println("是一个五位数")
	}else if (num % 10000 / 1000 != 0) {
		fmt.Println("是一个四位数")
	}else if (num % 1000 / 100 != 0) {
		fmt.Println("是一个三位数")
	}else if (num % 100 / 10 != 0) {
		fmt.Println("是一个两位数")
	}else if (num % 10 != 0) {
		fmt.Println("是一个个位数")
	}else {
		fmt.Println("超纲拉！")
	}
}