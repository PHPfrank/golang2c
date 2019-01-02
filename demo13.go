package main

import(
	"fmt"
)

//将一个正整数分解质因数。例如：输入90,打印出90=2*3*3*5。

func main () {
	
	num := 90
	//过滤1
	if (num == 1) {
		fmt.Println(1*1)
	}

	for i:= 2 ; i <= num ;i++ {
		//满足条件输出因子
		if (num % i == 0) {
			fmt.Printf("%d",i)
			num/=i;
			if(num != 1) {
				fmt.Printf("*");
			}
			//保持i的值不变
			i = i - 1	
        }
	}

}


