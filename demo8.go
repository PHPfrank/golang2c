package main

import(
	"fmt"
)

//输出9*9口诀。
func main () {

	//for循环打出第一层
	for i := 1; i <= 9 ; i++ {
		for j := 1; j <= i ; j++{
			fmt.Printf("%d",j)
			fmt.Printf("%s","*")
			fmt.Printf("%d",i)
			fmt.Printf("%s","=")
			fmt.Printf("%d  ",j*i)
		}
		fmt.Printf("\n")		
	}

}
