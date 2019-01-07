package main

import(
	"fmt"
)

//计算s=a+aa+aaa+aaaa+aa...a的值，其中a是一个数字。例如2+22+222+2222+22222(此时共有5个数相加)

func main () {

	a := 2
	num := 5
	b := a 
	total := 0
	for i := 1 ; i <= num ; i++ {
		//次数大于1
		if (i > 1) {
			//要加的数(上个数*10加上各位的原数)
			b = b * 10 + a
			total = total + b
		}else {
			total = total + a
		}		
	}

	fmt.Println("加和结果是",total)

}