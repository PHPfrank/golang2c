package main

import (
	"fmt"
)

//求1+2!+3!+...+20!的和。

func main () {

	total := 0
	//二十次
	for i := 1 ;i <= 20 ;i++ {
		num := i
		if (i > 1) {
			for j := i ; j > 1; j-- {
				//fmt.Println(j)
				num = num * (j - 1)
			}
		}		
		total = total + num
	} 
	fmt.Println(total)
}