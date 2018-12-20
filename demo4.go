package main

import(
	"fmt"
	"strings"
	"strconv"
)
//输入某年某月某日，判断这一天是这一年的第几天？

func main () {
	date := "2016-10-09"
	result := getDay(date)
	fmt.Println("该日期是这一年的第",result,"天")
}

func getDay (date string) (result int) {
	new_date := strings.Split(date, "-")
	//年
	year , _ := strconv.Atoi(new_date[0])
	//月
	month , _ := strconv.Atoi(new_date[1])
	//日
	day , _ := strconv.Atoi(new_date[2])
	//总天数
	sum := 0
	//闰年标识
	res := 0

	switch month {
		case 1:
			sum = 0
		case 2:
			sum = 31
		case 3:
			sum = 59
		case 4:
			sum = 90
		case 5:
			sum = 120
		case 6:
			sum = 151
		case 7:
			sum = 181
		case 8:
			sum = 212
		case 9:
			sum = 243
		case 10:
			sum = 273
		case 11:
			sum = 304
		case 12:
			sum = 334
		default:
			sum = 0	
	}
	
	sum = sum + day

	//闰年判断

	if ((year%4 == 0 && year%100 != 0) || year%400 == 0) {
		res = 1
	}

	if (res == 1 && month > 2 ) {
		sum++
	}

	return sum
}