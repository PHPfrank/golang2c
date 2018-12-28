package main

import(
	"fmt"
)

//判断101到200之间的素数

func main () {
	
	i,j := 0,0

	for i = 101 ; i <= 200 ; i++ {
		for j = 2; j < i; j++ {
        	// 如果j能被i整出在跳出循环
            if (i % j == 0) {
				break;
			}             
		}
		// 判断循环是否提前跳出，如果j<i说明在2~j之间,i有可整出的数
		if (j >= i) {
			fmt.Println(i)
		}    
	}

}