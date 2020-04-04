// 餐厅过滤器 project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	restaurants := [][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}
	veganFriendly := 1
	maxPrice := 50
	maxDistance := 10
	fmt.Println(filterRestaurants(restaurants, veganFriendly, maxPrice, maxDistance))
}

/*
执行用时 :72 ms, 在所有 Go 提交中击败了50.00%的用户
内存消耗 :6.8 MB, 在所有 Go 提交中击败了100.00%的用户

排序，按照评分降序序号升序，然后判断条件，把符合条件的添加到res里

*/
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	sort.Slice(restaurants, func(i, j int) bool {
		if restaurants[i][1] == restaurants[j][1] {
			return restaurants[i][0] > restaurants[j][0]
		}
		return restaurants[i][1] > restaurants[j][1]
	})
	res := []int{}
	for i := 0; i < len(restaurants); i++ {
		if veganFriendly == 1 {
			if restaurants[i][2] == 1 && restaurants[i][3] <= maxPrice && restaurants[i][4] <= maxDistance {
				res = append(res, restaurants[i][0])
			}
		} else {
			if restaurants[i][3] <= maxPrice && restaurants[i][4] <= maxDistance {
				res = append(res, restaurants[i][0])
			}
		}
	}
	return res
}
