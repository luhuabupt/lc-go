package main

import (
	"fmt"
	"sort"
	"strconv"
)

func compress(chars []byte) int {
	i, j, l, pre := 0, 0, 0, byte(0)
	for j < len(chars) {
		if chars[j] == pre {
			j++
			l++
		} else {
			if l > 1 {
				i++
				if l > 999 {
					chars[i] = byte(l / 1000)
					l %= 1000
					i++
				}
				if l > 99 {
					chars[i] = byte(l / 100)
					l %= 100
					i++
				}
				if l > 9 {
					chars[i] = byte(l / 10)
					l %= 10
					i++
				}
				chars[i] = byte(l)
				i++
			}
			pre = chars[j]
			chars[i] = chars[j]
			l = 1
		}
	}
	return i
}

func findFarmland(land [][]int) [][]int {
	ans := [][]int{}
	for i, arr := range land {
		for j, v := range arr {
			if v == 1 {
				r1, c1, r2, c2 := i, j, i, j
				for r2+1 < len(land) && land[r2+1][c2] == 1 {
					r2++
				}
				for c2+1 < len(land[0]) && land[r2][c2+1] == 1 {
					c2++
				}
				for ti := i; ti <= r2; ti++ {
					for tj := j; tj <= c2; tj++ {
						land[ti][tj] = 0
					}
				}
				ans = append(ans, []int{r1, c1, r2, c2})
			}
		}
	}
	return ans
}

type LockingTree struct {
	lock   []int
	parent []int
	son    [][]int
}

func Constructor_(parent []int) LockingTree {
	son := make([][]int, len(parent))
	for i := 1; i < len(parent); i++ {
		son[parent[i]] = append(son[parent[i]], i)
	}
	lock := make([]int, len(parent))

	return LockingTree{
		lock:   lock,
		parent: parent,
		son:    son,
	}
}

func (this *LockingTree) Lock(num int, user int) bool {
	if this.lock[num] == 0 {
		this.lock[num] = user
		return true
	}
	return false
}

func (this *LockingTree) Unlock(num int, user int) bool {
	if this.lock[num] == user {
		this.lock[num] = 0
		return true
	}
	return false
}

func (this *LockingTree) Upgrade(num int, user int) bool {
	// 祖先
	fmt.Println(this)
	for p := num; p != -1; p = this.parent[p] {
		if this.lock[p] != 0 {
			fmt.Println("parent", p)
			return false
		}
	}
	// 子孙
	flag := false
	son := this.son[num]
	for len(son) > 0 {
		h := son[0]
		son = son[1:]
		fmt.Println("son", h)
		if this.lock[h] != 0 {
			fmt.Println("lock", h)
			flag = true
			this.lock[h] = 0
		}
		if len(this.son[h]) > 0 {
			son = append(son, this.son[h]...)
		}
	}
	if flag {
		this.lock[num] = user
	}
	return flag
}

func numberOfGoodSubsets(nums []int) int {
	no := map[int]int{0: 0, 1: 0, 2: 1, 3: 2, 4: 0, 5: 4, 6: 3, 7: 8, 8: 0, 9: 0, 10: 5,
		11: 16, 12: 0, 13: 32, 14: 9, 15: 6, 16: 0, 17: 64, 18: 0, 19: 128, 20: 0,
		21: 10, 22: 17, 23: 256, 24: 0, 25: 0, 26: 33, 27: 0, 28: 0, 29: 512, 30: 7}
	dp, one, mod, ans := map[int]int{}, 0, int(1e9+7), 0
	for _, v := range nums {
		if v == 1 {
			one++
			continue
		}
		if no[v] == 0 {
			continue
		}
		x := no[v]
		dp[x] += one + 1
		for k, _ := range dp {
			if k&x == 0 {
				dp[x|k] += dp[k]
				dp[x|k] %= mod
			}
		}
	}
	for _, v := range dp {
		ans += v
		ans %= mod
	}
	return ans
}

func numberOfWeakCharacters(properties [][]int) int {
	ans, max := 0, 0
	sort.Slice(properties, func(i, j int) bool {
		return properties[i][0] > properties[j][0] || properties[i][0] == properties[j][0] && properties[i][1] < properties[j][1]
	})
	for _, arr := range properties {
		if max > arr[1] {
			ans++
		}
		if arr[1] > max {
			max = arr[1]
		}
	}
	return ans
}

func firstDayBeenInAllRooms_(nextVisit []int) int {
	n, vt, vc, i, day := len(nextVisit), map[int]int{}, 0, 0, -1
	for vc != n {
		day++
		if day >= int(1e9+7) {
			day -= int(1e9 + 7)
		}
		if vt[i] == 0 { // first
			vc++
		}
		vt[i]++

		// next
		if vt[i]%2 == 1 {
			i = nextVisit[i]
		} else {
			i = (i + 1) % n
		}
	}
	return day
}

func firstDayBeenInAllRooms(nextVisit []int) int {
	dp := map[int]int{0: 0, 1: 2}
	n := len(nextVisit)
	mod := int(1e9) + 7
	for i := 2; i < n; i++ {
		dp[i] += 2*dp[i-1] - dp[nextVisit[i-1]] + 2 + mod
		dp[i] %= mod
	}
	return dp[n-1]
}

func findNumberOfLIS(nums []int) int {
	dp, maxL := make([][]int, len(nums)), 0
	for k, v := range nums {
		max := 0
		for i := 0; i < k; i++ {
			if v > nums[i] && dp[i][0] > max {
				max = dp[i][0]
			}
		}
		dp[k] = []int{max + 1, 0}
		if max == 0 {
			dp[k][1] = 1
		}
		if max+1 > maxL {
			maxL = max + 1
		}
		for i := 0; i < k; i++ {
			if v > nums[i] && dp[i][0] == max {
				dp[k][1] += dp[i][1]
			}
		}
	}
	ans := 0
	for _, arr := range dp {
		if arr[0] == maxL {
			ans += arr[1]
		}
	}
	return ans
}

func longestSubsequenceRepeatedK(s string, k int) string {
	sa := []byte(s)
	sc := make([]int, 26)
	for _, v := range sa {
		sc[v-'a']++
	}

	var dfs func(cur, left []int) []int
	dfs = func(cur, left []int) []int {
		if len(left) == 1 {
			cur = append(cur, left[0])
			if checkS(cur, s, k) {
				return cur
			}
		}
		for i, v := range left {
			nc := []int{}
			nc = cur
			nc = append(nc, v)
			nf := []int{}
			if i == len(left) {
				nf = left[:i]
			} else {
				nf = append(left[:i], left[i+1:]...)
			}

			x := dfs(nc, nf)
			nc, nf = nil, nil
			if len(x) > 0 {
				return x
			}
		}
		return []int{}
	}

	eff := []int{}
	for ki, v := range sc {
		for i := 0; i < v/k; i++ {
			eff = append(eff, ki)
			v -= k
		}
	}
	sort.Slice(eff, func(i, j int) bool {
		return eff[i] > eff[j]
	})
	for i := 1; i < 1<<len(eff); i++ {
		has := []int{}
		for j := 0; j < len(eff); j++ {
			if 1<<j&i > 0 {
				has = append(has, eff[j])
			}
		}
		x := dfs([]int{}, has)
		if len(x) != 0 {
			ans := ""
			for _, si := range x {
				ans += string(si + 'a')
			}
		}
	}
	return ""
}

func checkS(arr []int, s string, k int) bool {
	c, ci := 0, 0
	for _, v := range s {
		if arr[ci] == int(v-'a') {
			ci++
			if ci == len(arr) {
				c++
				if c == k {
					return true
				}
				ci = 0
			}
		}
		//if len(s) - i < (k-c+1) * len(arr) {
		//	return false
		//}
	}

	return false
}

func longestSubsequenceRepeatedK_(s string, k int) (ans string) {
	n := len(s)
	pos := [26]int{}
	for i := range pos {
		pos[i] = n
	}
	nxt := make([][26]int, n)
	cnt := [26]int{}
	for i := n - 1; i >= 0; i-- {
		nxt[i] = pos
		pos[s[i]-'a'] = i
		cnt[s[i]-'a']++
	}

	// 计算所有可能出现在答案中的字符，包括重复的
	// 倒着统计，这样下面计算排列时的第一个合法方案就是答案，从而提前退出
	a := []byte{}
	for i := 25; i >= 0; i-- {
		for c := cnt[i]; c >= k; c -= k {
			a = append(a, 'a'+byte(i))
		}
	}

	for m := len(a); m > 0 && ans == ""; m-- { // 从大到小枚举答案长度 m
		permutations(len(a), m, func(ids []int) bool { // 枚举长度为 m 的所有排列
			t := make([]byte, m)
			for i, id := range ids {
				t[i] = a[id]
			}
			i, j := 0, 0
			if t[0] == s[0] {
				j = 1
			}
			for {
				i = nxt[i][t[j%m]-'a']
				if i == n {
					break
				}
				j++
			}
			if j >= k*m {
				ans = string(t)
				return true // 提前退出
			}
			return false
		})
	}
	return
}

// 模板：生成 n 选 r 的排列
func permutations(n, r int, do func(ids []int) bool) {
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	if do(ids[:r]) {
		return
	}
	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}
	for {
		i := r - 1
		for ; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				tmp := ids[i]
				copy(ids[i:], ids[i+1:])
				ids[n-1] = tmp
				cycles[i] = n - i
			} else {
				j := cycles[i]
				ids[i], ids[n-j] = ids[n-j], ids[i]
				if do(ids[:r]) {
					return
				}
				break
			}
		}
		if i == -1 {
			return
		}
	}
}

type TripleInOne struct {
	arr       []int
	idx       [3]int
	stackSize int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		make([]int, stackSize*3+3),
		[3]int{0, stackSize + 1, 2*stackSize + 2},
		stackSize,
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	if this.idx[stackNum] == this.stackSize*(stackNum+1) {
		return
	}
	this.idx[stackNum]++
	this.arr[this.idx[stackNum]] = value
}

func (this *TripleInOne) Pop(stackNum int) int {
	if this.idx[stackNum] == (this.stackSize+1)*stackNum {
		return -1
	}
	x := this.arr[this.idx[stackNum]]
	this.arr[this.idx[stackNum]] = 0
	this.idx[stackNum]--
	return x
}

func (this *TripleInOne) Peek(stackNum int) int {
	if this.idx[stackNum] == (this.stackSize+1)*stackNum {
		return -1
	}
	return this.arr[this.idx[stackNum]]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.idx[stackNum] == (this.stackSize+1)*stackNum
}

func fraction(cont []int) []int {
	ans := []int{1, cont[len(cont)-1]}
	for i := len(cont) - 2; i >= 0; i-- {
		ans[0] = cont[i]*ans[1] + ans[0]
		ans[0], ans[1] = ans[1], ans[0]
	}
	for i := 2; i <= ans[0] && i <= ans[1]; {
		if ans[0]%i == 0 && ans[1]%i == 0 {
			ans[0] /= i
			ans[1] /= i
		} else {
			i++
		}
	}
	ans[0], ans[1] = ans[1], ans[0]
	return ans
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	dfs(root)
	return root
}

func dfs(p *Node) (last *Node) {
	for p != nil {
		if p.Child != nil {
			next := p.Next
			p.Next, p.Child = p.Child, nil
			p.Next.Prev = p

			p := dfs(p.Next)
			p.Next = next
			if next != nil {
				next.Prev = p
			}
		}
		p, last = p.Next, p
	}
	return
}

func numDecodings(s string) int {
	sa := []byte(s)
	n := len(sa)
	dp := map[int]int{}
	mod := int(1e9) + 7

	dp[-1] = 1
	if sa[0] == '*' {
		dp[0] = 9
	} else {
		dp[0] = 1
	}

	for i := 1; i < n; i++ {
		if sa[i] == '*' && sa[i-1] == '*' {
			dp[i] = dp[i-1]*9 + dp[i-2]*15
		} else if sa[i] == '*' {
			dp[i] = dp[i-1] * 9
			if sa[i-1] == '1' {
				dp[i] += dp[i-2] * 9
			} else if sa[i-1] == '2' {
				dp[i] += dp[i-2] * 6
			}
		} else if sa[i-1] == '*' {
			if sa[i] > '0' {
				dp[i] = dp[i-1]
			}
			if sa[i] <= '6' {
				dp[i] += dp[i-2] * 2
			} else {
				dp[i] += dp[i-2]
			}
		} else {
			if sa[i] > '0' {
				dp[i] += dp[i-1]
			}
			if sa[i-1] == '1' || sa[i-1] == 2 && sa[i] <= '6' {
				dp[i] += dp[i-2]
			}
		}
		dp[i] %= mod
	}
	return dp[n-1]
}

func findMinMoves(machines []int) int {
	n := len(machines)

	sum := 0
	for _, v := range machines {
		sum += v
	}
	if sum%n != 0 {
		return -1
	}

	avg := sum / n
	sum = 0
	ans := 0

	for _, v := range machines {
		v -= avg
		sum += v
		if v > ans {
			ans = v
		}
		if sum > ans {
			ans = sum
		}
		if -sum > ans {
			ans = -sum
		}
	}

	return ans
}

func scoreOfStudents(s string, answers []int) int {
	n := len(s)

	stack := []int{int(s[0] - '0')}
	rs := 0
	for i := 1; i <= n-2; i += 2 {
		if s[i] == '+' {
			stack = append(stack, int(s[i+1]-'0'))
		} else {
			top := len(stack) - 1
			stack[top] *= int(s[i+1] - '0')
		}
	}
	for _, v := range stack {
		rs += v
	}

	dp := make([][]map[int]bool, n)
	for j := 0; j < n; j += 2 {
		dp[j] = make([]map[int]bool, n)
		for i := j; i >= 0; i -= 2 {
			dp[i][j] = map[int]bool{}
			if i == j {
				dp[i][j][int(s[i]-'0')] = true
				continue
			}
			for k := j; k >= i; k -= 2 {
				for a, _ := range dp[i][k] {
					for b, _ := range dp[k+2][j] {
						if s[k+1] == '+' {
							if a+b <= 1000 {
								dp[i][j][a+b] = true
							}
						} else {
							if a*b <= 1000 {
								dp[i][j][a*b] = true
							}
						}
					}
				}
			}
		}
	}

	ans := 0
	for _, v := range answers {
		if v == rs {
			ans += 5
		} else if dp[0][n-1][v] {
			ans += 2
		}
	}

	return ans
}

func toHex(num int) string {
	ans := ""
	if num == 0 {
		return ans
	}
	if num < 0 {
		num = -num
		complement := 1
		for i := 0; i < 32; i++ {
			if num&1<<i == 0 {
				complement += 1 << i
			}
		}
		fmt.Println(num, complement)
		return toHex(complement)
	}

	a := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	for num > 0 {
		ans = a[num%16] + ans
		num /= 16
	}

	return ans
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator < 0 && denominator < 0 {
		numerator, denominator = -numerator, -denominator
	}
	if numerator < 0 {
		return "-" + fractionToDecimal(-numerator, denominator)
	}
	if denominator < 0 {
		return "-" + fractionToDecimal(numerator, -denominator)
	}
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}

	ans := []byte{}
	if numerator/denominator > 0 {
		ans = []byte(strconv.Itoa(numerator / denominator))
	} else {
		ans = []byte{'0'}
	}
	ans = append(ans, '.')

	m := map[int]int{}
	numerator %= denominator
	i, s := len(ans), len(ans)
	for {

		if m[numerator] > 0 {
			s = m[numerator]
			ans = append(ans, ')')
			break
		}
		m[numerator] = i

		numerator *= 10
		for numerator < denominator {
			numerator *= 10
			ans = append(ans, '0')
		}

		ans = append(ans, uint8(numerator/denominator+'0'))
		i = len(ans)
		numerator = numerator % denominator
		if numerator == 0 {
			return string(ans)
		}
	}

	return string(ans[:s]) + "(" + string(ans[s:])
}

func removeKdigits(num string, k int) string {
	stack := []byte{num[0]}
	for i := 1; i < len(num); i++ {
		for ; k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1]; k-- {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]

	i := 0
	for i < len(stack) && stack[i] == '0' {
		i++
	}

	ans := string(stack[i:])
	if len(ans) == 0 {
		ans = "0"
	}

	return ans
}

func removeDuplicateLetters(s string) string {
	cnt := map[int32]int{}
	for _, v := range s {
		cnt[v]++
	}

	m := map[int32]bool{}
	ans := []int32{}
	for _, v := range s {
		cnt[v]--
		for len(ans) > 0 && ans[len(ans)-1] > v && cnt[ans[len(ans)-1]] > 0 {
			m[ans[len(ans)-1]] = false
			ans = ans[:len(ans)-1]
		}
		if m[v] {
			continue
		}
		ans = append(ans, v)
		m[v] = true
	}

	return string(ans)
}

func peakIndexInMountainArray(arr []int) int {
	return sort.Search(len(arr)-1, func(i int) bool { return arr[i] > arr[i+1] })
}

func countAndSay(n int) string {
	a := []string{"", "1"}
	for i := 2; i <= n; i++ {
		x := a[i-1]
		cur := []int32{}
		tmp := 1
		pre := ' '
		for k, v := range x {
			if k == 0 {
				pre = v
				continue
			}
			if v == pre {
				tmp++
			} else {
				cur = append(cur, []int32(strconv.Itoa(tmp))...)
				cur = append(cur, pre)
				pre = v
				tmp = 1
			}
		}
		cur = append(cur, []int32(strconv.Itoa(tmp))...)
		cur = append(cur, pre)
		a = append(a, string(cur))
	}

	return a[n]
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	m := make([]map[int]int, n)
	m[n-1] = map[int]int{nums[n-1]: 1}
	for i := n - 2; i >= 0; i-- {
		m[i] = map[int]int{}
		for k, v := range m[i+1] {
			m[i][k] = v
		}
		m[i][nums[i]]++
	}
	ans := [][]int{}
	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			for x := j + 1; x < n-1; x++ {
				if m[x+1][target-nums[i]-nums[j]-nums[x]] > 0 {
					ans = append(ans, []int{nums[i], nums[j], nums[x], target - nums[i] - nums[j] - nums[x]})
				}
			}
		}
	}

	if len(ans) == 0 {
		return ans
	}

	return dup(ans)
}

func dup(a [][]int) [][]int {
	res := [][]int{}
	m := map[int]map[int]map[int]map[int]bool{}
	for _, v := range a {
		if m[v[0]] == nil {
			m[v[0]] = map[int]map[int]map[int]bool{}
		}
		if m[v[0]][v[1]] == nil {
			m[v[0]][v[1]] = map[int]map[int]bool{}
		}
		if m[v[0]][v[1]][v[2]] == nil {
			m[v[0]][v[1]][v[2]] = map[int]bool{}
		}
		if !m[v[0]][v[1]][v[2]][v[3]] {
			res = append(res, v)
		}
	}
	return res
}

// 记忆搜索
func shoppingOffers(price []int, special [][]int, needs []int) int {
	ns := [][]int{}
	n := len(price)
	for _, arr := range special {
		x := 0
		for i := 0; i < n; i++ {
			if arr[i] > needs[i] {
				x = -1
				break
			}
			x += price[i] * arr[i]
		}
		if arr[n] < x {
			ns = append(ns, arr)
		}
	}

	dp := map[string]int{}

	var dfs func(s []byte)
	dfs = func(s []byte) {
		min := 0
		for i, v := range s {
			min += price[i] * int(v)
		}
		for _, arr := range ns { // 大礼包
			flag := true
			todo := []byte{} // 还差
			for i, v := range s {
				if arr[i] > int(v) {
					flag = false
					break
				}
				todo = append(todo, v-byte(arr[i]))
			}
			if !flag { // 无效
				continue
			}

			if _, ok := dp[string(todo)]; !ok { // 记忆搜索
				dfs(todo)
			}
			if dp[string(todo)]+arr[n] < min {
				min = dp[string(todo)] + arr[n]
			}
		}
	}

	x := []byte{}
	for _, v := range needs {
		x = append(x, byte(v))
	}
	dfs(x)

	return dp[string(x)]
}

// 有向图dfs
func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	max := 0
	hasLoop := false
	vis := make([]int, n)
	dp := make([][26]int, n)

	c := make([]int, n)
	for i, v := range colors {
		c[i] = int(v - 'a')
	}

	notRoot := map[int]bool{}
	sub := make([][]int, n)
	for _, arr := range edges {
		sub[arr[0]] = append(sub[arr[0]], arr[1])
		notRoot[arr[1]] = true
	}
	root := []int{}
	for i := 0; i < n; i++ {
		if !notRoot[i] {
			root = append(root, i)
		}
	}
	if len(root) == 0 {
		return -1
	}

	var dfs func(pre, x int)
	dfs = func(pre, x int) {
		vis[x] = 1

		for _, v := range sub[x] {
			if vis[v] == 1 {
				hasLoop = true
				return
			}
			if vis[v] == 0 {
				dfs(x, v)
			}
			for char := 0; char < 26; char++ {
				if dp[v][char] > dp[x][char] {
					dp[x][char] = dp[v][char]
				}
			}
		}

		dp[x][c[x]]++
		if dp[x][c[x]] > max {
			max = dp[x][c[x]]
		}

		vis[x] = 2
	}

	for _, v := range root {
		dfs(-1, v)
	}
	if hasLoop {
		return -1
	}
	for i := 0; i < n; i++ {
		if vis[i] != 2 {
			return -1
		}
	}

	return max
}

func shiftingLetters(s string, shifts []int) string {
	n := len(s)
	sum := make([]int, n)
	for i, v := range shifts {
		if i == n-1 {
			sum[i] = v
		} else {
			sum[i] = sum[i+1] + v
		}
	}
	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		ans[i] = byte((int(s[i]-'a')+sum[i])%26 + 'a')
	}
	return string(ans)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxSumBST(root *TreeNode) int {
	var dfs func(p *TreeNode) (ban bool, sum, min, max int)
	ans := 0

	dfs = func(p *TreeNode) (ban bool, sum, min, max int) {
		if p.Left == nil && p.Right == nil {
			if p.Val > ans {
				ans = p.Val
			}
			return true, p.Val, p.Val, p.Val
		}

		b1, b2, s1, s2, i1, a1, i2, a2 := true, true, 0, 0, 100000, -100000, 100000, -100000
		if p.Left != nil {
			b1, s1, i1, a1 = dfs(p.Left)
			if a1 >= p.Val {
				b1 = false
			}
		}
		if p.Right != nil {
			b2, s2, i2, a2 = dfs(p.Right)
			if a2 <= p.Val {
				b2 = false
			}
		}

		if b1 && b2 {
			ban, sum = true, s1+s2+p.Val
			if sum > ans {
				ans = sum
			}
		}

		min = min3(i1, i2, p.Val)
		max = max3(a1, a2, p.Val)

		return
	}

	dfs(root)
	return ans
}

func min3(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}
func max3(a, b, c int) int {
	if a > b && a > c {
		return a
	}
	if b > c {
		return b
	}
	return c
}

type MapSum struct {
	next [26]*MapSum
	val  int
}

func Constructor() MapSum {
	return MapSum{
		next: [26]*MapSum{},
		val:  0,
	}
}

func (this *MapSum) Insert(key string, val int) {
	p := this
	for _, v := range key {
		if p.next[v-'a'] == nil {
			p.next[v-'a'] = &MapSum{}
		}
		p = p.next[v-'a']
	}
	p.val = val
}

func (this *MapSum) Sum(prefix string) int {
	ans := 0
	for _, v := range prefix {
		ans += this.next[v-'a'].val
		this = this.next[v-'a']
	}
	return ans
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

func findClosestElements(arr []int, k int, x int) []int {
	idx := sort.SearchInts(arr, x)
	n := len(arr)

	m := map[int]bool{}
	for l, r := idx-1, idx; k > 0; {
		if r == n || x-arr[l] <= arr[r]-x {
			m[arr[l]] = true
			l--
		} else {
			m[arr[r]] = true
			r++
		}
		k--
	}

	ans := []int{}
	for _, v := range arr {
		if m[v] {
			ans = append(ans, v)
		}
	}

	return ans
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getDirections(root *TreeNode, startValue int, destValue int) string {
	res := map[int][][]int{}
	res[startValue] = [][]int{}
	res[destValue] = [][]int{}

	var dfs func(find int, p *TreeNode) (path, road []int)
	dfs = func(find int, p *TreeNode) (path, road []int) {
		if p == nil {
			return
		}
		if p.Val == find {
			return []int{p.Val}, []int{}
		}
		pa, ro := dfs(find, p.Left)
		if len(pa) > 0 {
			return append(pa, p.Val), append(ro, 0)
		}
		par, ror := dfs(find, p.Left)
		if len(par) > 0 {
			return append(par, p.Val), append(ror, 1)
		}
		return
	}

	sp, sr := dfs(startValue, root)
	ep, er := dfs(destValue, root)

	re(sp)
	re(sr)
	re(ep)
	re(er)

	idx := -1
	for i := 0; i < len(ep) && i < len(sp); i++ {
		if ep[i] != sp[i] {
			break
		}
		idx++
	}

	ans := []byte{}
	for i := len(sr) - 1; i >= idx; i-- {
		ans = append(ans, 'U')
	}
	for i := idx; i < len(er); i++ {
		if er[i] == 0 {
			ans = append(ans, 'L')
		} else {
			ans = append(ans, 'R')
		}
	}
	return string(ans)
}

func re(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func superPow(a int, b []int) int {
	if a == 1 {
		return 1
	}
	M := int(1e9 + 7)

	var pow func(a, n int) int
	pow = func(a, n int) int {
		fmt.Println(a, n)
		ans := 1
		for n > 0 {
			if n%2 == 1 {
				ans *= a
			}
			n /= 2
			a *= a
			a %= M
			ans %= M
		}
		fmt.Println(ans)
		return ans
	}

	p, p0 := a%M, pow(a%M, 9)%M
	ans := 0
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] > 0 {
			ans += pow(p, b[i])
			ans %= M
		}
		p = p * p0 % M
	}

	return ans
}
