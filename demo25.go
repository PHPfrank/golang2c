package main

import(
	"fmt"
	//"reflect"
)

//利用递归函数调用方式，将所输入的5个字符，以相反顺序打印出来。

func main () {
	str := "abcde"
	recursion(len(str),str)
}

func recursion (i int,str string) {
	
		//递归出口
    if (i != 0) {
				j := i - 1
				s := str[j]
				fmt.Println(string(s))
				recursion(j,str)	
    	} else {
				return
			}
			
}