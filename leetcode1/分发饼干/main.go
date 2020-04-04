// 分发饼干 project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1, 2, 3}
	s := []int{3}
	fmt.Println(findContentChildren(g, s))
}

/*
执行用时 :40 ms, 在所有 Go 提交中击败了74.32%的用户
内存消耗 :6 MB, 在所有 Go 提交中击败了43.96%的用户

先排序，然后遍历两个数组，只要s[j]>=g[i]
说明可以分配，res++，最后返回res
给一个孩子的饼干应当尽量小并且又能满足该孩子，这样大饼干才能拿来给满足度比较大的孩子。
因为满足度最小的孩子最容易得到满足，所以先满足满足度最小的孩子。


在以上的解法中，我们只在每次分配时饼干时选择一种看起来是当前最优的分配方法，
但无法保证这种局部最优的分配方法最后能得到全局最优解。我们假设能得到全局最优解，
并使用反证法进行证明，即假设存在一种比我们使用的贪心策略更优的最优策略。
如果不存在这种最优策略，表示贪心策略就是最优策略，得到的解也就是全局最优解。

证明：假设在某次选择中，贪心策略选择给当前满足度最小的孩子分配第 m 个饼干，
第 m 个饼干为可以满足该孩子的最小饼干。假设存在一种最优策略，可以给该孩子分配第 n 个饼干，
并且 m < n。我们可以发现，经过这一轮分配，贪心策略分配后剩下的饼干一定有一个比最优策略来得大。
因此在后续的分配中，贪心策略一定能满足更多的孩子。
也就是说不存在比贪心策略更优的策略，即贪心策略就是最优策略。


*/
func findContentChildren1(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	l1 := len(g)
	l2 := len(s)
	j := 0
	res := 0
	for i := 0; i < l1; i++ {
		for ; j < l2; j++ {
			if s[j] >= g[i] {
				res++
				j++
				break
			}
		}
	}
	return res
}

/*
执行用时 :40 ms, 在所有 Go 提交中击败了74.32%的用户
内存消耗 :6 MB, 在所有 Go 提交中击败了74.73%的用户
*/
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	j, i, l1, l2 := 0, 0, len(g), len(s)
	res := 0
	for i < l1 && j < l2 {
		if s[j] >= g[i] {
			i++
			j++
			res++
		} else {
			j++
		}
	}
	return res
}
