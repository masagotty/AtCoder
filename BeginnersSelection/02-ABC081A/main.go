package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a int
	fmt.Scanf("%d", &a)
	b := strconv.Itoa(a)
	c := strings.Count(b, "1")
	fmt.Printf("%d\n", c)
}
