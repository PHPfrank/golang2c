package main

import(
	"fmt"
)

//一个数如果恰好等于它的因子之和，这个数就称为"完数"。例如6=1＋2＋3.编程找出1000以内的所有完数。

func main () {

	sum := 0

	for i := 1 ; i <= 1000 ; i++ {
		//重置因数和
		sum = 0

		for j := 1 ; j <= i / 2 ; j++ {
			//所有的因数
			if (i%j == 0) {
				sum = sum + j 
			}

		}	
		//因数和等于这个数输出结果
		if (sum == i) {
			fmt.Println(i)
		}

	}


}