package main

import(
	"fmt"
)

//打印出所有的"水仙花数"，所谓"水仙花数"是指一个三位数，其各位数字立方和等于该数 本身。
//例如：153是一个"水仙花数"，因为153=1的三次方＋5的三次方＋3的三次方。

func main () {

	for i:= 100 ; i < 1000 ; i++ {
		//百位数
		hundreds := i / 100 % 10
		//十位数
		tens := i / 10 % 10
		//个位数
		single := i%10

		if (hundreds * hundreds* hundreds + tens * tens * tens + single * single * single == i) {
			fmt.Println("水仙花数是：",i)
		}
	}
}

