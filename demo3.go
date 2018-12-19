package main

import(
	"fmt"
	"math"
	"strconv"
	"strings"
)

//一个整数，它加上100后是一个完全平方数，再加上168又是一个完全平方数，请问该数是多少？


func main () {

	for i := 0; i < 10000; i++ {		
		a := math.Sqrt(float64(i + 100))
		a_result := strconv.FormatFloat(a, 'f', -1, 64)
		b := math.Sqrt(float64(i + 100 + 168))
		b_result := strconv.FormatFloat(b, 'f', -1, 64)	
		if(!strings.Contains(a_result, ".") && !strings.Contains(b_result, ".")){
			fmt.Println("这个数字是",i)
			return
		}
	 }

}
