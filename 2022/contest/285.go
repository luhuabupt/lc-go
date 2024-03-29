package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	a := []byte{'"'}
	for i := 0; i < 100000; i++ {
		if i == 10000 {
			a = append(a, 'b')
		} else if i > 10000 {
			a = append(a, byte('a'+i%26))
		} else {
			a = append(a, 'a')
		}
	}
	a = append(a, '"')

	b := []byte{'"'}
	for i := 0; i < 100000; i++ {
		if i%2 == 0 {
			b = append(b, 'a')
		} else {
			b = append(b, 'b')
		}
	}
	b = append(b, '"')

	c := "[10000"
	for i := 0; i < 100000-1; i++ {
		c += ",10000"
	}
	c += "]"

	fmt.Println(string(a))
	fmt.Println(string(b))
	fmt.Println(c)

	////fmt.Println(longestRepeating("babacc", "bcb", []int{1, 3, 3}))
	////fmt.Println(longestRepeating("abyzz", "aa", []int{2,1}))
	////fmt.Println(longestRepeating("mm", "bfviuwsr", []int{0, 0, 1, 0, 0, 1, 1, 0}))
	////fmt.Println(longestRepeating("geuqjmt", "bgemoegklm	", []int{3,4,2,6,5,6,5,4,3,2}))
	////fmt.Println(longestRepeating("zlcbw", "ygwcnshib	", []int{3,1,3,3,4,0,3,0,0}))
	//fmt.Println(longestRepeating("iiiiiccmmmmmmmmggghhhhhhhbbzzzzzzztttaaaagggcccccccccccccckkkkkkkfuuuuuuuuiiiiiiiqqqqqwwwwwwwddddrvvvvttttttttkkkkkkfffddmmmmmxxxxxxxxmmmmmeeeeeefffssssssissxddddddpppppppzzzccnzzzxxxxxxxrrrrmmpppvvvvvvuuuuuuussssssssyyyynnnnnnnnffffoovvvvvvvvqqqqqqqkkkdddddddkppppppt", "gbbbbbbbbbbbbbbbbqccccccccccddddddddddddddddddddooooooooooooooggggggggggggggffffffffffffffffffffgggggggggggggggghhhiiiiiwwwwwwbbbbbbbbbbbbbbbbbbbbbbbbzzzzzzzzzzzzzzzzxxxxxxxxxxxxxxxxxxxxxxcccccccccccfffffffffffffffffffffffjjjjjjjjjjjjjjjjjjjjjjjgggggggggggggggggggghhhhhhhhhiiiiiiiiiiiiiiiiiyyyyyyypppppppppppppppppppttttttttttttttttttttbbbbbbbbbbfffffffffffffffffvvvvvvvvvvvkkkkkkkkkkkkkkkkkkkkkoooooooooooooooeeeeeeeeeeeeeeeeeeeeeeeecccccccccccccccllllllllllllllllllllcccccccccccccccvvvvvvvvvvvvvvvvvvvvvvdddddiiiiiiiiiiiiiiiiiiiimmmmmmmmmmmmmmmmmoooollllllllllllleeeeeeeeeeeerrrrrrrrggggggggggggggggggggggkkkkkkkkkkkkkkkkkkkkkkkkkkoooovvvvvvvvvvvvuuuugggggggggggggggggggglllllllllllliiiiiiiiiiiiiiiiiiiiiissssssssssssswwwwwwwwwwwwwwwwwnnnnnnnnnwwwwwwwwwwgggggggggggggggggggqqqqqqllllllllllllllllllllllyyyyyyyyyyyyyyyyyyyyyyyyiiiiiiiiiiiiiiiiiiiiiikkkkkkkkkkkkkkkkkkkkkkkkffaaaaaaaaaaaaaaaaaaas", []int{130,69,76,73,71,67,81,70,74,68,75,72,78,82,80,79,77,225,222,217,215,218,219,224,221,220,216,223,117,115,129,127,121,134,124,122,123,131,128,118,116,126,130,133,120,125,119,132,195,196,184,190,193,186,189,188,187,183,191,192,185,194,93,96,95,99,89,88,97,87,90,92,98,94,86,91,95,97,101,88,106,105,90,102,96,104,94,87,99,100,91,103,92,89,98,93,89,92,91,81,84,79,85,93,87,82,90,78,83,86,80,88,81,83,82,74,75,73,72,71,206,208,210,205,207,209,135,118,119,127,129,128,122,136,117,123,134,126,120,133,125,115,132,116,131,137,114,124,121,130,142,143,140,147,136,138,133,135,148,146,145,144,141,139,134,137,252,254,259,244,262,248,261,249,250,247,243,264,253,251,263,257,260,245,246,255,258,256,154,158,152,155,151,159,156,153,157,150,160,61,55,59,62,69,77,71,57,58,67,73,76,56,74,60,70,72,65,63,66,68,64,75,240,235,229,245,224,238,237,236,227,234,231,241,239,230,244,226,242,246,233,232,225,243,228,78,74,79,81,75,82,73,77,76,84,66,68,65,80,70,72,69,71,67,83,6,10,9,7,5,3,4,8,11,28,42,29,32,30,31,39,43,38,40,33,35,41,34,44,36,37,116,119,117,118,115,120,121,69,67,76,85,81,78,82,79,77,73,84,74,72,70,71,68,83,75,80,89,80,87,79,81,94,86,88,78,91,84,83,85,95,92,93,96,82,90,77,193,198,190,199,197,195,194,192,191,196,181,178,182,186,179,183,174,187,176,188,184,180,175,185,190,177,189,25,29,31,28,26,27,33,30,34,24,32,211,215,216,202,218,209,207,213,201,206,205,217,221,212,214,219,220,210,203,208,204,139,132,127,131,138,129,137,134,133,126,128,135,136,140,130,102,95,94,106,111,103,100,93,108,113,92,99,109,105,97,98,101,112,114,110,104,107,96,91,42,51,44,45,39,41,46,47,48,49,40,50,52,53,43,243,245,247,242,248,234,231,232,230,239,233,240,244,249,238,237,241,235,246,236,75,72,77,70,69,67,66,73,71,78,76,80,68,74,79,42,29,45,30,40,34,32,26,37,44,39,38,33,31,41,24,43,28,25,36,35,27,5,4,3,2,1,236,240,233,229,238,230,224,227,239,228,231,234,237,222,225,235,223,226,232,221,176,173,181,175,177,182,183,189,179,186,187,184,178,174,180,185,188,197,195,198,196,69,80,68,73,71,76,75,70,78,74,72,79,77,204,201,199,200,207,203,209,206,202,205,210,208,195,194,196,192,190,191,189,193,170,188,176,187,175,182,169,177,174,179,171,180,181,183,172,173,178,189,185,168,184,186,134,132,142,154,153,131,138,148,147,140,133,155,141,137,143,151,139,156,152,136,144,146,145,150,149,135,154,156,155,157,107,105,104,102,108,110,113,106,103,111,112,109,32,30,29,31,45,35,42,46,47,43,48,33,36,52,51,38,49,39,34,44,41,37,50,40,87,91,90,94,96,85,95,93,89,86,88,92,253,243,237,248,242,256,258,247,254,239,255,246,251,252,241,244,240,249,257,238,245,250,252,253,246,251,242,244,247,249,248,250,245,243,241,113,114,108,123,117,115,111,112,118,119,121,122,109,110,120,116,107,150,148,149,146,152,154,147,151,153,95,88,87,93,92,91,94,90,96,89,221,219,217,228,232,234,225,218,227,233,223,224,230,231,229,222,216,226,220,123,122,124,127,125,126,95,106,92,105,90,89,88,86,99,100,103,94,102,98,96,101,85,87,97,91,104,93,250,232,230,236,238,246,247,249,243,240,228,234,239,235,227,231,242,248,237,233,244,245,229,241,38,42,44,51,55,47,52,50,39,54,41,46,53,40,45,57,58,43,49,59,48,56,256,237,238,236,239,258,252,248,253,246,249,250,245,257,247,251,254,235,243,241,240,244,255,242,123,122,172,183,174,179,186,169,168,175,182,176,171,178,173,184,170,181,180,177,185,195}))
}

type hl []int

func (h hl) Len() int            { return len(h) }
func (h hl) Less(i, j int) bool  { return h[i] > h[j] } // 大顶堆
func (h hl) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hl) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *hl) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func longestRepeating(s string, qc string, qi []int) []int {
	sa := []byte(s)
	st := [][2]int{}
	n := len(s)
	del := make([]int, n+1) // 标记已删除的区间

	h := &hl{}
	for pre, i := 0, 1; i <= n; i++ {
		if i == n {
			st = append(st, [2]int{pre, n - 1})
			heap.Push(h, n-pre)
			break
		}
		if s[i] == s[i-1] {
			continue
		}
		st = append(st, [2]int{pre, i - 1})
		heap.Push(h, i-pre)

		pre = i
	}

	getMax := func() int {
		for {
			x := heap.Pop(h).(int)
			if del[x] == 0 {
				heap.Push(h, x)
				return x
			}
			del[x]--
		}
	}

	ans := make([]int, len(qi))

	for i, v := range qi {
		pos := sort.Search(len(st), func(i int) bool {
			return st[i][0] > v
		})
		pos--

		cur := st[pos]
		if sa[v] == qc[i] {
		} else if cur[0] == cur[1] { // 只有一个点的区间
			if pos > 0 && pos < len(st)-1 && sa[v-1] == qc[i] && sa[v+1] == qc[i] {
				// 左合并 and 右合并
				del[1]++
				del[st[pos-1][1]-st[pos-1][0]+1]++
				del[st[pos+1][1]-st[pos+1][0]+1]++

				or := st[pos+1][1]
				x := append([][2]int{}, st[pos+2:]...)
				st = append(st[:pos], x...)
				st[pos-1][1] = or

				heap.Push(h, st[pos-1][1]-st[pos-1][0]+1)
			} else if pos > 0 && sa[v-1] == qc[i] {
				// 左合并
				del[st[pos-1][1]-st[pos-1][0]+1]++

				st[pos-1][1] = v
				x := st[pos+1:]
				st = append(st[:pos], x...)

				heap.Push(h, st[pos-1][1]-st[pos-1][0]+1)
			} else if pos < len(st)-1 && sa[v+1] == qc[i] {
				// 右合并
				del[st[pos+1][1]-st[pos+1][0]+1]++

				st[pos+1][0] = v
				x := st[pos+1:]
				st = append(st[:pos], x...)
				heap.Push(h, st[pos][1]-st[pos][0]+1)
			}
		} else if cur[0] == v { // 在左端点上
			if pos > 0 && sa[v-1] == qc[i] {
				// 左合并
				del[st[pos-1][1]-st[pos-1][0]+1]++
				del[st[pos][1]-st[pos][0]+1]++

				st[pos-1][1] = v
				st[pos][0] = v + 1

				heap.Push(h, st[pos-1][1]-st[pos-1][0]+1)
				heap.Push(h, st[pos][1]-st[pos][0]+1)
			} else {
				// 左拆
				del[st[pos][1]-st[pos][0]+1]++

				tmp := append([][2]int{}, st[pos:]...)
				st = append(append(st[:pos], [2]int{v, v}), tmp...)
				st[pos+1][0] = v + 1

				heap.Push(h, 1)
				heap.Push(h, st[pos+1][1]-st[pos+1][0]+1)
			}
		} else if cur[1] == v { // 在右端点上
			if pos < len(st)-1 && sa[v+1] == qc[i] {
				// 右合并
				del[st[pos+1][1]-st[pos+1][0]+1]++
				del[st[pos][1]-st[pos][0]+1]++

				st[pos+1][0] = v
				st[pos][1] = v - 1

				heap.Push(h, st[pos][1]-st[pos][0]+1)
				heap.Push(h, st[pos+1][1]-st[pos+1][0]+1)
			} else {
				// 右拆
				del[st[pos][1]-st[pos][0]+1]++

				tmp := append([][2]int{}, st[pos+1:]...)
				st = append(append(st[:pos+1], [2]int{v, v}), tmp...)
				st[pos][1] = v - 1

				heap.Push(h, 1)
				heap.Push(h, st[pos][1]-st[pos][0]+1)
			}
		} else { // 在中间，区间就1拆3
			// 左拆 and 右拆
			del[st[pos][1]-st[pos][0]+1]++

			op := st[pos]
			tmp := append([][2]int{}, st[pos+1:]...)
			st = append(append(st[:pos+1], [][2]int{{v, v}, {v + 1, op[1]}}...), tmp...)
			st[pos][1] = v - 1

			heap.Push(h, st[pos][1]-st[pos][0]+1)
			heap.Push(h, 1)
			heap.Push(h, st[pos+2][1]-st[pos+2][0]+1)
		}

		ans[i] = getMax()
		sa[v] = qc[i]
	}

	return ans
}
