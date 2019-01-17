package main

import(
	"fmt"
)

//利用递归方法求5!。

func main () {
    recursion(5,1)
}

func recursion (i , sum int) {
	//递归出口
    if (i > 0) {
		sum = sum * i
		j := i - 1
		recursion(j,sum)
    } else {
		fmt.Println(sum)
		return
	}
}