package main 

import(
	"fmt"
)

//输出国际象棋棋盘
func main () {
	//8*8的棋盘
	for i := 0; i < 8 ; i++ {
		//*与空格代表黑白格
		for j := 0; j < 8 ; j++{
			if ((i+j)%2==0) {
				fmt.Printf("[*]");
			}else {
				fmt.Printf("[ ]");
			}			
		}
		fmt.Printf("\n")		
	}
}