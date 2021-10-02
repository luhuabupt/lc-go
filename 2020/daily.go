package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//fmt.Println(countArrangement(15))
	fmt.Println(checkValidString("(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"))
}

// 863 二叉树中所有距离为 K 的结点
// https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree/
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var path []*TreeNode
	var ans []int
	m := map[int]bool{}

	// 递归出path
	var dfs func(p *TreeNode, tmp []*TreeNode)
	dfs = func(p *TreeNode, tmp []*TreeNode) {
		if p == nil {
			return
		}

		var nt []*TreeNode
		nt = append(nt, p)
		nt = append(nt, tmp...)
		if p == target {
			path = nt
			return
		}

		dfs(p.Left, nt)
		dfs(p.Right, nt)
	}

	// 计算每个节点的自己点是否k=0
	var getK func(p *TreeNode, k int)
	getK = func(p *TreeNode, k int) {
		if k < 0 || p == nil || m[p.Val] {
			return
		}
		if k == 0 {
			ans = append(ans, p.Val)
			return
		}
		getK(p.Left, k-1)
		getK(p.Right, k-1)
	}

	dfs(root, []*TreeNode{})
	for i, v := range path {
		getK(v, k-i)
		m[v.Val] = true // 标记已经遍历过的父节点
	}

	return ans
}

// 1104. 二叉树寻路
// https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree/
func pathInZigZagTree(label int) []int {
	level, idx, ans := 0, 0, []int{label}
	for {
		if 1<<(level+1) > label {
			break
		}
		level++
	}
	if level%2 == 0 { // 偶 左
		idx = label - (1 << level)
	} else {
		idx = (1 << (level + 1)) - 1 - label
	}
	for i := level - 1; i >= 0; i-- {
		idx = idx / 2
		if i%2 == 0 { // 偶 左
			ans = append(ans, (1<<i)+idx)
		} else {
			ans = append(ans, (1<<(i+1))-1-idx)
		}
	}

	for i, n := 0, len(ans)-1; i <= n/2; i++ {
		ans[i], ans[n-i] = ans[n-i], ans[i]
	}

	return ans
}

// 743 网络延迟时间
// https://leetcode-cn.com/problems/network-delay-time/
// Dijkstra 迪杰斯特拉算法 有向图最短路径
// not ac
func networkDelayTime(times [][]int, n int, k int) int {
	m, s, u, dp := map[int]map[int]int{}, map[int]bool{k: true}, map[int]bool{}, map[int]int{}
	for i := 1; i <= n; i++ {
		dp[i], u[i] = -1, true
	}
	dp[k] = 0
	delete(u, k)

	for _, arr := range times {
		if m[arr[0]] == nil {
			m[arr[0]] = map[int]int{}
		}
		m[arr[0]][arr[1]] = arr[2]
	}

	for len(u) > 0 {
		for si, _ := range s {
			for i, v := range m[si] { // 更新U
				if dp[i] < 0 || dp[i] > dp[si]+v {
					dp[i] = dp[si] + v
				}
			}
		}
		minU, minUi := int(1e6), -1
		for ui, _ := range u {
			if dp[ui] > 0 && dp[ui] < minU {
				minU, minUi = dp[ui], ui
			}
		}
		if minUi == -1 {
			return -1
		}
		s[minUi] = true
		delete(u, minUi)
	}

	ans := 0
	for _, v := range dp {
		if v == -1 {
			return -1
		}
		if v > ans {
			ans = v
		}
	}
	return ans
}

// 611. 有效三角形的个数
// https://leetcode-cn.com/problems/valid-triangle-number/
// 双指针
func triangleNumber(nums []int) int {
	ans := 0
	sort.Ints(nums)
	nums = append(nums, 10000)
	for k := 0; k < len(nums)-3; k++ {
		for i, j := k+1, k+2; i < len(nums)-2; i++ {
			for {
				if nums[j] >= nums[i]+nums[k] {
					if j > i+1 {
						ans += j - i - 1
					}
					break
				}
				j++
			}
			if i == j {
				j++
			}
		}
	}
	return ans
}

// https://leetcode-cn.com/problems/find-eventual-safe-states/
// 802. 找到最终的安全状态
// 三色标记 | 拓扑排序
func eventualSafeNodes_(graph [][]int) []int {
	color := map[int]int{} // 0-未访问, 1-访问过, 2-安全
	var safe func(x int) bool
	safe = func(x int) bool {
		if len(graph[x]) == 0 || color[x] == 2 {
			return true
		}
		if color[x] == 1 {
			return false
		}
		color[x] = 1
		for _, v := range graph[x] {
			if !safe(v) {
				return false
			}
		}
		color[x] = 2
		return true
	}

	ans := []int{}
	for i := 0; i < len(graph); i++ {
		if safe(i) {
			ans = append(ans, i)
		}
	}
	return ans
}
func eventualSafeNodes(graph [][]int) []int {
	g := make([][]int, len(graph))
	inDeg := make([]int, len(graph))
	list := []int{}
	ans := []int{}

	for k, arr := range graph {
		inDeg[k] = len(arr) // 入度
		if len(arr) == 0 {
			list = append(list, k)
		}
		for _, v := range arr { // 反向图
			g[v] = append(g[v], k)
		}
	}
	for len(list) > 0 {
		for _, v := range g[list[0]] { // 拆边
			inDeg[v]--
			if inDeg[v] == 0 {
				list = append(list, v)
			}
		}
		list = list[1:]
	}
	for i, v := range inDeg {
		if v == 0 {
			ans = append(ans, i)
		}
	}

	return ans
}

func numberOfArithmeticSlices_(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	diff, ans, t := nums[1]-nums[0], 0, 0
	nums = append(nums, 3000)
	for i := 2; i < len(nums); i++ {
		if diff == nums[i]-nums[i-1] {
			t++
		} else {
			t, diff = 0, nums[i]-nums[i-1]
		}
		ans += t
	}
	return ans
}

// https://leetcode-cn.com/problems/shortest-path-visiting-all-nodes/
// 847. 访问所有节点的最短路径
// 图 广度优先
func shortestPathLength(graph [][]int) int {
	n := len(graph)
	type tuple struct{ u, mask, dist int }
	q := []tuple{}
	seen := []map[int]bool{}

	for i := 0; i < n; i++ {
		q = append(q, tuple{i, 1 << i, 0})
		seen = append(seen, map[int]bool{1 << i: true})
	}

	for {
		//t, q := q[0], q[1:]
		t := q[0]
		q = q[1:]
		if t.mask == (1<<n)-1 {
			return t.dist
		}
		for _, v := range graph[t.u] {
			mask := (1 << v) | t.mask
			if !seen[v][mask] {
				seen[v][mask] = true
				q = append(q, tuple{v, mask, t.dist + 1})
			}
		}
	}
}

// https://leetcode-cn.com/problems/minimum-total-space-wasted-with-k-resizing-operations/
func minSpaceWastedKResizing(nums []int, k int) int {
	n := len(nums)
	dp, g, w, sum := make([][]int, n), make([][]int, n), make([][]int, n), 0
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		w[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		g[i][i], sum = nums[i], nums[i]
		for x := i - 1; x >= 0; x-- {
			sum += nums[x]
			g[x][i] = maxTwo(g[x+1][i], nums[x])
			w[x][i] = g[x][i]*(i-x+1) - sum
		}
	}

	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		dp[i][0] = w[0][i]
		for j := 1; j <= k; j++ {
			dp[i][j] = w[1][i]
			for x := 1; x <= i-1; x++ {
				dp[i][j] = minTwo(dp[x][j-1]+w[x+1][i], dp[i][j])
			}
			dp[i][j] = minTwo(dp[i][j], dp[i][j-1])
		}
	}

	return dp[n-1][k]
}

func maxTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minTwo(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// https://leetcode-cn.com/problems/finding-the-users-active-minutes/
func findingUsersActiveMinutes(logs [][]int, k int) []int {
	m, u, ans := map[int]map[int]bool{}, map[int]int{}, make([]int, k)
	for _, arr := range logs {
		if m[arr[0]] == nil {
			m[arr[0]] = map[int]bool{}
		}
		if !m[arr[0]][arr[1]] {
			m[arr[0]][arr[1]] = true
			u[arr[0]]++
		}
	}

	for _, num := range u {
		ans[num-1]++
	}

	return ans
}

// https://leetcode-cn.com/problems/arithmetic-slices-ii-subsequence/
func numberOfArithmeticSlices(nums []int) int {
	n, ans := len(nums), 0
	dp := make([]map[int]int, n)
	for i, v := range nums {
		dp[i] = map[int]int{}
		for j, x := range nums[:i] {
			ans += dp[j][v-x]
			dp[i][v-x] += dp[j][v-x] + 1
		}
	}
	return ans
}

// https://leetcode-cn.com/problems/maximum-score-from-removing-substrings/
func maximumGain(s string, x int, y int) int {
	a, b, ca, cb, ans := 0, 0, 'a', 'b', 0
	if x < y {
		ca, cb, x, y = 'b', 'a', y, x
	}
	s += "x"
	for _, v := range s {
		if v == ca || v == cb {
			if v == cb {
				if a > 0 {
					ans += x
					a--
				} else {
					b++
				}
			}
			if v == ca {
				a++
			}
		} else {
			ans += minTwo(a, b) * y
			a, b = 0, 0
		}
	}
	return ans
}

// https://leetcode-cn.com/problems/longest-palindromic-subsequence/solution/
func longestPalindromeSubseq(s string) int {
	n, arr := len(s), []byte(s)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1-i)
	}
	for i := 0; i < n; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			if arr[j] == arr[j+i-1] {
				dp[i][j] = dp[i-2][j+1] + 2
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j+1])
			}
		}
	}
	return dp[n][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countArrangement(n int) int {
	f := map[int]int{0: 1}
	for mask := 1; mask < 1<<n; mask++ {
		k := oc(mask)
		for i := 0; i < n; i++ {
			if (mask & (1 << i)) > 0 { // 第i位是1
				pre := mask ^ 1<<i                // 第i位设为0
				if (i+1)%k == 0 || k%(i+1) == 0 { // 第k个数整除
					f[mask] += f[pre]
				}
			}
		}
	}
	return f[(1<<n)-1]
}

func oc(x int) (ans int) {
	for x > 0 {
		x -= x & -x
		ans++
	}
	return
}

func rearrangeArray(nums []int) []int {
	sort.Ints(nums)
	ans := []int{}
	for i := 0; i < len(nums)/2; i++ {
		ans = append(ans, nums[i], nums[len(nums)-1-i])
	}
	if len(nums)%2 == 1 {
		ans = append(ans, nums[len(nums)/2])
	}
	return ans
}

func wiggleSort(nums []int) {
	sort.Ints(nums)
	ans := []int{}
	for i := 0; i < len(nums)/2; i++ {
		ans = append(ans, nums[(len(nums)-1)/2-i], nums[len(nums)-1-i])
	}
	if len(nums)%2 == 1 {
		ans = append(ans, nums[0])
	}
	for i, v := range ans {
		nums[i] = v
	}
}

// https://leetcode-cn.com/problems/student-attendance-record-ii/
func checkRecord(n int) int {
	mod := int(1e9) + 7
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, 3) // 0-P 1-L 2-A
		for j := 0; j <= 2; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	dp[0][0][0] = 1
	dp[1][0][0] = 1
	dp[1][1][0] = 1
	dp[1][2][1] = 1

	for i := 2; i <= n; i++ {
		// P
		dp[i][0][0] = (dp[i-1][0][0] + dp[i-1][1][0]) % mod
		dp[i][0][1] = (dp[i-1][0][1] + dp[i-1][1][1] + dp[i-1][2][1]) % mod

		// L
		dp[i][1][0] = (dp[i-1][0][0] + dp[i-2][0][0]) % mod
		dp[i][1][1] = (dp[i-1][2][1] + dp[i-1][0][1] + dp[i-2][0][1] + dp[i-2][2][1]) % mod

		// A
		dp[i][2][1] = (dp[i-1][0][0] + dp[i-1][1][0]) % mod
	}

	ans := 0
	for _, arr := range dp[n] {
		for _, v := range arr {
			ans += v
		}
	}
	return ans
}

// hello
func reverseVowels(s string) string {
	arr, m := []byte(s), map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	i, j := 0, len(s)-1
	for {
		for !m[arr[i]] {
			if i == len(s)-1 {
				return string(arr)
			}
			i++
		}
		for !m[arr[j]] {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return string(arr)
}

func reversePrefix(word string, ch byte) string {
	w := []byte(word)
	for i := 0; i < len(w); i++ {
		if w[i] == ch {
			for j := 0; j <= i/2; j++ {
				w[j], w[i-j] = w[i-j], w[j]
			}
			break
		}
	}
	return string(w)
}

func interchangeableRectangles(rectangles [][]int) int64 {
	m := map[float64]int64{}
	for _, v := range rectangles {
		m[(float64(v[0])/float64(v[1]))]++
	}
	ans := int64(0)
	for _, v := range m {
		ans += v * (v - 1) / 2
	}
	return ans
}

func checkValidString(s string) bool {
	l, star, cl := 0, 0, 0
	for _, v := range s {
		if v == '*' {
			star++
			cl++ // 可能的右
		} else if v == '(' {
			l++
		} else if v == ')' {
			if l > 0 {
				l--
			} else if star > 0 {
				star--
			} else {
				return false
			}
		}
		if cl > l { // 在L之前出现的删除
			cl = l
		}
	}
	return l <= cl
}

func getSeq(s []byte) {
	for i := 1; i < 1<<len(s)-1; i++ {
		for j := 1; j < len(s); j++ {

		}
	}
}
func getTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func topKFrequent(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	h := &hp{}
	for w, c := range cnt {
		heap.Push(h, pair{w, c})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(pair).w
	}
	return ans
}

func isPrefixString(s string, words []string) bool {
	idx := 0
	for _, str := range words {
		if idx+len(str) > len(s) {
			return false
		}
		for i := 0; i < len(str); i++ {
			if str[i] != s[idx] {
				return false
			}
			idx++
		}
		if idx == len(s) {
			return true
		}
	}
	return idx == len(s)
}

func get(v int, z *[]int) int {
	ans := 0
	for v != 1 {
		for i, x := range *z {
			if v%x == 0 {
				ans |= 1 << i
				v /= x
			}
		}
	}
	return ans
}

func countSubstrings(s string) int {
	dp, ans := map[int]map[int]bool{}, 0
	for i := 1; i <= len(s); i++ {
		dp[i] = map[int]bool{}
	}
	for k, _ := range s {
		dp[1][k] = true
		ans++
	}
	for i := 0; i <= len(s)-2; i++ {
		if s[i] == s[i+1] {
			dp[2][i] = true
			ans++
		}
	}
	for l := 3; l <= len(s); l++ {
		for i := 0; i <= len(s)-l; i++ {
			if dp[l-2][i+1] && s[i] == s[i+l-1] {
				dp[l][i] = true
				ans++
			}
		}
	}
	return ans
}

func maxProduct(s string) (max int) {
	sa, n, max := []byte(s), len(s), 0
	for i := 1; i < 1<<n-1; i++ {
		cur, supp := []byte{}, []byte{}
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				cur = append(cur, sa[j])
			} else {
				supp = append(supp, sa[j]) // 补集
			}
		}
		if check(cur) {
			x := len(cur) * maxSeq(string(supp))
			max = twoMax(x, max)
		}
	}
	return
}
func check(s []byte) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
func maxSeq(s string) int {
	max := 1
	for i := 1; i < 1<<len(s)-1; i++ {
		tmp := []byte{}
		for j := 1; j < len(s); j++ {
			if i&(1<<j) > 0 {
				tmp = append(tmp, s[j])
			}
		}
		if check(tmp) {
			max = twoMax(max, len(tmp))
		}
	}
	return max
}
func maxSeq_(s string) int {
	n, arr := len(s), []byte(s)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1-i)
	}
	for i := 0; i < n; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			if arr[j] == arr[j+i-1] {
				dp[i][j] = dp[i-2][j+1] + 2
			} else {
				dp[i][j] = twoMax(dp[i-1][j], dp[i-1][j+1])
			}
		}
	}
	return dp[n][0]
}
func twoMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func smallestMissingValueSubtree(parents []int, nums []int) []int {
	n := len(parents)
	son := make([][]int, n)
	ans := make([]int, n)

	for i := 1; i < n; i++ {
		son[parents[i]] = append(son[parents[i]], i)
	}

	var f func(i int) map[int]bool
	f = func(i int) map[int]bool {
		inSet := map[int]bool{}
		for _, v := range son[i] {
			sm := f(v)
			if len(sm) > len(inSet) {
				sm, inSet = inSet, sm
			}
		}
	}
	f(0)

	return ans
}
func tm(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findLongestWord(s string, dictionary []string) string {
	n := len(s)
	m := make([][26]int, n+1)
	for x := 'a'; x <= 'z'; x++ {
		m[n][x-'a'] = -1
	}
	for i := n - 1; i >= 0; i-- {
		m[i] = m[i+1]
		m[i][s[i]-'a'] = i
	}

	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) > len(dictionary[j]) || len(dictionary[i]) == len(dictionary[j]) && dictionary[i] < dictionary[j]
	})

outer:
	for i := 0; i < len(dictionary); i++ {
		j := 0
		for _, v := range dictionary[i] {
			if m[j][v-'a'] == -1 {
				continue outer
			}
			j = m[j][v-'a'] + 1
		}
		return dictionary[i]
	}
	return ""
}
