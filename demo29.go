package main

import(
	"fmt"
)

//请输入星期几的第一个字母来判断一下是星期几，如果第一个字母一样，则继续判断第二个字母。

func main () {
	var a string
	var b string
	fmt.Println("请输入第一个字母")
	fmt.Scanf("%s", &a)
	if (a == "m") {
		fmt.Println("monday")
	} else if (a == "w") {
		fmt.Println("wednesday")
	} else if (a == "f") {
		fmt.Println("friday")
	} else if (a == "t") {
		fmt.Println("请输入第二个字母")
		fmt.Scan(&b)
		if (b == "u") {
			fmt.Println("tuesday")
		}
        if (b == "h") {
			fmt.Println("thursday")
		}
	} else if (a == "s") {
		fmt.Println("请输入第二个字母")
		fmt.Scan(&b)
		if (b == "a") {
			fmt.Println("saturday")
		}
        if (b == "u") {
			fmt.Println("sunday")
		}
	} else {
		fmt.Println("error")
	}

}