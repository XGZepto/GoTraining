package main

import "fmt"

const maxn = 10001
const maxm = 200001

type Edge struct {
	to, next int
}

type Node struct {
	x, y int
}

var q [maxn * 4]Node
var l [maxm]Edge
var head [maxn]int
var light, tag [maxn]bool

var qh, qt int
var n, m, tot int
var x1, y1, x2, y2 int

func addEdge(from, to int) {
	tot++
	l[tot] = Edge{to, head[from]}
	head[from] = tot
}

func getIndex(x, y int) int {
	return (x-1)*n + y
}

func getNode(Index int) (res Node) {
	res.x = (Index-1)/n + 1
	res.y = Index % n
	if res.y == 0 {
		res.y += n
	}
	return
}

func pushback(x Node) {
	qt++
	q[qt] = x
	tag[getIndex(x.x, x.y)] = true
}

func checkVal(a, b Node) bool {
	if a.x*a.y*b.x*b.y == 0 {
		return false
	}
	if a.x > n || a.y > n || b.x > n || b.y > n {
		return false
	}
	return true
}

func breadthSearch() (res int) {
	dx := [4]int{0, 0, 1, -1}
	dy := [4]int{1, -1, 0, 0}

	qh, qt, res = 1, 1, 1
	q[qh] = Node{1, 1}
	light[1], tag[1] = true, true
	for ; qh <= qt; qh++ {
		top := q[qh]

		for i := 0; i < 4; i++ {
			nxt := Node{top.x + dx[i], top.y + dy[i]}
			nxtID := getIndex(nxt.x, nxt.y)
			if checkVal(top, nxt) && light[nxtID] && !tag[nxtID] {
				pushback(nxt)
			}
		}

		for i := head[getIndex(top.x, top.y)]; i > 0; i = l[i].next {
			nxt := getNode(l[i].to)
			if !light[l[i].to] {
				res++
			}
			light[l[i].to] = true

			if tag[l[i].to] {
				continue
			}

			flag := false
			for i := 0; i < 4 && !flag; i++ {
				adj := Node{nxt.x + dx[i], nxt.y + dy[i]}
				if checkVal(nxt, adj) && tag[getIndex(adj.x, adj.y)] {
					flag = true
				}
			}
			if flag {
				pushback(nxt)
			}
		}
	}
	return
}

func main() {
	fmt.Scan(&n, &m)
	for i := 1; i <= m; i++ {
		fmt.Scan(&x1, &y1, &x2, &y2)
		addEdge(getIndex(x1, y1), getIndex(x2, y2))
	}
	fmt.Println(breadthSearch())
}
