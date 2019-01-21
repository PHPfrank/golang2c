package main

import(
	"fmt"
)

//有5个人坐在一起，问第五个人多少岁？他说比第4个人大2岁。问第4个人岁数，他说比第3个人大2岁。
//问第三个人，又说比第2人大两岁。问第2个人，说比第一个人大两岁。
//最后问第一个人，他说是10岁。请问第五个人多大？

func main () {
	recursion(1,10)
}

func recursion (i , age int) {
	
	if (i == 5) {
		fmt.Println("第",i,"个人",age,"岁")
		return
	} else {
		age = age + 2
		i = i + 1
		recursion(i,age)
	}

}