package main

import(
	"fmt"
	"strings"
)

//删除一个字符串中的指定字母，如：字符串 "aca"，删除其中的 a 字母。

func main () {
	var str string
	fmt.Println("请输入字符串")
	fmt.Scanf("%s", &str)
	fmt.Println(strings.Replace(str, "a","",-1))
}