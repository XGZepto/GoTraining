package main

import (
	"fmt"
	"sort"
)

const maxn, mod = 1e5 + 1, 1e9 + 9

var prefix, tree, sorted [maxn]int
var n int

func addTree(pos, val int) {
	for ; pos <= n+2; pos += (pos & (-pos)) {
		tree[pos] = (tree[pos] + val) % mod
	}
}

func sumTree(pos int) (res int) {
	for ; pos > 0; pos -= (pos & (-pos)) {
		res = (res + tree[pos]) % mod
	}
	return
}

func getRank(array []int, target, length int) (res int) {
	l, r := 0, length-1
	for l <= r {
		mid := (l + r) >> 1
		if array[mid] >= target {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return
}

func calculate() (res int) {
	for i := 1; i <= n; i++ {
		res = sumTree(prefix[i])
		addTree(prefix[i], res)
	}
	return
}

func main() {
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&prefix[i])
		prefix[i] += prefix[i-1]
		sorted[i] = prefix[i]
	}
	sort.Ints(sorted[0 : n+1])
	for i := 0; i <= n; i++ {
		prefix[i] = getRank(sorted[0:n+1], prefix[i], n+1) + 1
	}
	addTree(prefix[0], 1)
	fmt.Println(calculate())
}
