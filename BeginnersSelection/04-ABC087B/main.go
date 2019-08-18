package main

import (
	"fmt"
)

func main() {
	var a, b, c, x, count int
	fmt.Scan(&a, &b, &c, &x)

	for i := 0; i < a+1; i++ {
		for j := 0; j < b+1; j++ {
			for k := 0; k < c+1; k++ {
				// fmt.Printf("i: %d, j: %d, k: %d\n", i, j, k)
				if i*500+j*100+k*50 == x {
					// fmt.Println("Success!")
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

/*
あなたは、500円玉をA枚、100円玉をB枚、50円玉をC枚持っています。
これらの硬貨の中から何枚かを選び、
合計金額をちょうどX円にする方法は何通りありますか。

500a
100b
 50c
*/
