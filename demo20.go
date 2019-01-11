package main

import(
	"fmt"
)

//两个乒乓球队进行比赛，各出三人。甲队为a,b,c三人，乙队为x,y,z三人。
//已抽签决定比赛名单。有人向队员打听比赛的名单。a说他不和x比，c说他不和x,z比，请编程序找出三队赛手的名单。
//c-y a-z b-x

func main () {

	for i := 'x'; i <= 'z'; i++ {
		for j := 'x'; j <= 'z'; j++ {
			for k := 'x'; k <= 'z'; k++ {
				if (i == 'x'||j == i||k == 'x'||k == 'z'||j == k||i == k) {
					continue;
				}
				fmt.Printf("a对战%c,b对战%c,c对战%c",i,j,k);
			}
		}
	}

}
