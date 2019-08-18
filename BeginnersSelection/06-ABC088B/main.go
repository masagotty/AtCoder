package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, alice, bob, sub int
	fmt.Scan(&n)
	cards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cards[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cards)))

	for i := 0; i < n; i += 2 {
		alice += cards[i]
		if i+1 < n {
			bob += cards[i+1]
		}
	}
	sub = alice - bob
	fmt.Println(sub)

}

/*
N枚のカードがあります.
i枚目のカードには,
aiという数が書かれています.
AliceとBobは,
これらのカードを使ってゲームを行います.
ゲームでは,
AliceとBobが交互に1枚ずつカードを取っていきます.
Aliceが先にカードを取ります.
2人がすべてのカードを取ったときゲームは終了し,
取ったカードの数の合計がその人の得点になります.
2人とも自分の得点を最大化するように最適な戦略を取った時,
AliceはBobより何点多く取るか求めてください.
*/
