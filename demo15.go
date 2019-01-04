package main

import(
	"fmt"
)

//输入两个正整数m和n，求其最大公约数和最小公倍数。

func main () {
	m := 26 //默认m比n大
	n := 12
	//最大公约数
	a := divisor(m,n)
	//最小公倍数（最大公约数*最小公倍数 = m * n）
	b := m * n / a
	fmt.Println("最大公约数：",a)
	fmt.Println("最小公倍数：",b)
}

func divisor (m,n int ) (result int) {
	a := 0  //最大公约数，辗转相除法
	for {
		a = m % n
		if (a == 0) {
			return n
		}
		m = n
		n = a 
	} 
}