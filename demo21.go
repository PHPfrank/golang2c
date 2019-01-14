package main

import(
	"fmt"
)

//打印出如下图案（菱形）。
//    *    1-1
//   ***   2-3
//  *****  3-5
// ******* 4-7
//  *****  5-5
//   ***   6-3
//    *    7-1
func main () {
	max := 7
	//递增
	for i := 1 ; i <= 4 ; i++  {
		num := 2*i - 1
		//输出空格
		for k:= 0 ; k < (max-num)/2 ; k++ {
			fmt.Printf(" ")
		}
		//输出图形
		for j := 0 ; j < num ; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n");
	}
	//递减
	for i := 3 ; i >= 1 ; i--  {
		num := 2*i - 1
		//输出空格
		for k:= 0 ; k < (max-num)/2 ; k++ {
			fmt.Printf(" ")
		}
		//输出图形
		for j := 0 ; j < num ; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n");
	}



}