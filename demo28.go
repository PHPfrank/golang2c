package main

import(
	"fmt"
)

//一个5位数，判断它是不是回文数。即12321是回文数，个位与万位相同，十位与千位相同。

func main () {
	var num int
	fmt.Scanf("%d", &num)
	//fmt.Println(a)
	//万位数
	res1 := num / 10000
	//千位数
	res2 := num % 10000 / 1000
	//百位数
	//res3 := num % 1000 / 100
	//十位数
	res4 := num % 100 / 10
	//个位数
	res5 := num % 10

	if (res1 == res5 && res2 == res4) {
		fmt.Println("我是回文数")
	} else {
		fmt.Println("我不是回文数")
	}
	
}