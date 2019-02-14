package main

import(
	"fmt"
)

//判断一个数字是否为质数。

func main () {
	var n int
	fmt.Println("请输入数字")
	fmt.Scanf("%d", &n)
	IsPrime(n)
}

func IsPrime(n int) {
	if n <= 1 {
		fmt.Println("不是质数")
		return
	}

	//从2遍历到n-1，看看是否有因子
	for i := 2; i < n; i++ {
		if n%i == 0 {
			fmt.Println("不是质数")
			return
		}
	}
	fmt.Println("是质数")
}