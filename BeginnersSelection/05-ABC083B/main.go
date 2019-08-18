package main

import (
	"fmt"
)

func sumOfDigits(n int) int {
	var sumOf int
	for n > 0 {
		sumOf += n % 10
		n /= 10
	}
	return sumOf
}

func main() {
	var n, a, b, sum int
	fmt.Scan(&n, &a, &b)
	for i := 1; i < n+1; i++ {
		if v := sumOfDigits(i); a <= v && v <= b {
			sum += i
		}
	}
	fmt.Println(sum)
}

/*
1以上N以下の整数のうち、
10進法での各桁の和がA以上B以下であるものの総和を求めてください。
*/
