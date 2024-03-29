package main

import (
	"fmt"
	"sort"
)

type tree struct {
	arr []int
}

func (t tree) add(i, v int) {
	for i < len(t.arr) {
		t.arr[i] = t.arr[i] + v
		i += lowbit(i)
	}
}

func (t tree) sum(i int) (ans int) {
	for i > 0 {
		ans += t.arr[i]
		i -= lowbit(i)
	}
	return
}

func (t tree) query(l, r int) (ans int) {
	return t.sum(r) - t.sum(l-1)
}

func lowbit(v int) int {
	return v & -v
}

// 离线离散映射
func offline(a []int) []int {
	//a := []int{-1e9, -100, -1, 0, 1, 33, 1e8, 1e9}
	n := len(a)

	b := [][]int{}
	for i, v := range a {
		b = append(b, []int{v, i})
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i][0] < b[j][0] || b[i][0] == b[j][0] && b[i][1] < b[j][1]
	})
	rea := make([]int, n)
	for i, idx := 0, 0; i < n; i++ {
		if i == 0 || b[i][0] > b[i-1][0] {
			idx++
		}
		rea[b[i][1]] = idx
	}

	return rea
}

func main() {
	fmt.Println(offline([]int{-1e9, -100, -1, -1, 0, 1, 33, 1e8, 1e9}))
	x := []int{1, 2, 3, 4, 5, 1}
	t := tree{make([]int, len(x)+1)}

	for i, v := range x {
		t.add(i+1, i+v+1000)
	}
	fmt.Println(t.arr)
	fmt.Println(t.query(1, 2))
	fmt.Println(t.query(2, 3))
	fmt.Println(t.query(1, 5))
}
