package main
 
import (
	"fmt"
)

//输入三个整数x,y,z，请把这三个数由小到大输出。

func main() {
	sort(3,5,1)
}

func sort (a,b,c int) {
	
	if (a > b && a > c){
		//第一个数最大
		if ( b > c){
			fmt.Println(c,b,a)
		}else{
			fmt.Println(b,c,a)
		}	
	}else if (b > a && b > c) {
		//第二个数最大
		if ( a > c){
			fmt.Println(c,a,b)
		}else{
			fmt.Println(a,c,b)
		}
	}else {
		//第三个数最大
		if ( a > b){
			fmt.Println(b,a,c)
		}else{
			fmt.Println(a,b,c)
		}
	}
}
 


 