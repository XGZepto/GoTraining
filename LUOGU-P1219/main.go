package main

import (
	"fmt"
)

var n, tot int
var a [15]int
var v0, v1, v2 [30]bool

func work(v int) {
	if v > n {
		if tot++; tot <= 3 {
			for i := 1; i <= n; i++ {
				fmt.Print(a[i], " ")
			}
			fmt.Print("\n")
		}
		return
	}
	for i := 1; i <= n; i++ {
		if v0[i] || v1[v+i] || v2[v-i+n] {
			continue
		}
		v0[i], v1[v+i], v2[v-i+n] = true, true, true
		a[v] = i
		work(v + 1)
		v0[i], v1[v+i], v2[v-i+n] = false, false, false
	}
}

func main() {
	fmt.Scanln(&n)
	work(1)
	fmt.Println(tot)
}
