package main

import(
	"fmt"
)

//有一分数序列：2/1，3/2，5/3，8/5，13/8，21/13...求出这个数列的前20项之和。

func main () {

	var sum,a,b,t float64 = 0,2,1,0
	
    for i := 1 ; i<=20 ; i++ {
        sum=sum+a/b
        t=a
        a=a+b
        b=t
    }
    fmt.Println(sum);

}