package main

import(
	"fmt"
)

//一球从100米高度自由落下，每次落地后反跳回原高度的一半；再落下，求它在第10次落地时，共经过多少米？第10次反弹多高？

func main () {

	var height float64 = 100 

	for i := 0 ; i < 10 ; i++ {
		height = height / 2
	} 

	fmt.Println(height)

}