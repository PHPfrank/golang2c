// 有1、2、3、4个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？
package main

import(
	"fmt"
)


func main () {
	//计数
	count := 0
	//遍历保证每个三位数各不相同
	//百位数
	for a := 1 ; a <= 4  ; a++ {
		//十位数
		for b := 1 ; b <= 4  ; b++{
			//个位数
			for c := 1 ; c <= 4  ; c++{
				//百，十，个位数各不相同
				if ( a!= b && a != c && b != c){
					//简单输出结果
					fmt.Println(a,b,c)
					//计数加1
					count++
				}
			}
		}
	}
	//输出总数
	fmt.Println(count)
}