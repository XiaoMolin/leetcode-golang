// 翻转对 project main.go
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2, 3, 1}
	//nums := []int{2, 4, 3, 5, 1}
	//nums := []int{-5, -5}
	fmt.Println(reversePairs(nums))
}

func reversePairs(nums []int) int {
	if nums == nil || len(nums) < 2 {
		// 数组为空，或长度小于2，肯定不存在逆序对
		return 0 // 返回0
	}
	tmp := make([]int, len(nums))                  // 定义辅助数组tmp
	return sortAndCount(nums, 0, len(nums)-1, tmp) // 一边做归并排序，一边统计逆序对数量
}

// 辅助函数 归并两个有序子数组
// 输入数组arr，low到mid对应第一个有序子数组，mid+1到high对应第二个有序子数组，辅助数组tmp
func merge(arr []int, low, mid, high int, tmp []int) {
	// 定义游标i，j，k分别指向两个有序子数组和辅助数组tmp
	i, j, k := low, mid+1, 0    // 初始化为各自的开始下标
	for i <= mid && j <= high { // 当满足条件时不断执行以下操作
		if arr[i] <= arr[j] { // 如果arr[i]小于arr[j]
			tmp[k] = arr[i] // 把i指向的数字放到tmp中
			i, k = i+1, k+1 // i和k都加一
		} else { // 否则说明arr[i]>arr[j]产生逆序对
			tmp[k] = arr[j] // 同时也需要把较小的数字加入到tmp[k]中
			k, j = k+1, j+1 // 且j和k都需要加1
		}
	}
	// 退出循环后，为了避免其中一个子数组还有未进行对比的数字，所以对两个子数组进行检查
	for i <= mid { // 把下标i指向的数字以及后面的数字依次加入到tmp中
		tmp[k] = arr[i]
		k, i = k+1, i+1
	}
	for j <= high { // 把子数组剩下的数字加入到tmp中
		tmp[k] = arr[j]
		k, j = k+1, j+1
	}
	// 这时tmp是合并后的有序数组,再拷贝回arr数组
	copy(arr[low:high+1], tmp)
}

// 辅助函数 统计重要逆序对
func count(arr []int, low, mid, high int) int {
	i, j := low, mid+1
	cnt := 0
	for i <= mid && j <= high {
		if arr[i] <= 2*arr[j] {
			i++
		} else {
			cnt += (mid - i + 1)
			j++
		}
	}
	return cnt
}

// 辅助函数 同时进行归并排序和统计逆序对数量
func sortAndCount(arr []int, low, high int, tmp []int) int {
	cnt := 0        // 定义cnt统计逆序对数量
	if low < high { // 当low小于high时执行以下操作
		mid := low + (high-low)/2                  // 计算low和high中间下标mid
		cnt += sortAndCount(arr, low, mid, tmp)    // 递归调用自己处理low到mid之间的子数组，同时把返回的逆序对数量加到cnt上
		cnt += sortAndCount(arr, mid+1, high, tmp) // 同样的处理mid+1到high子数组,同时把返回的逆序对数量加到cnt上
		cnt += count(arr, low, mid, high)          // 把返回的逆序对数量加到cnt上
		merge(arr, low, mid, high, tmp)            // 最后调用辅助函数去合并两个子数组
	}
	return cnt // 最后返回cnt即可
}

/*
暴力求解，超出时间限制
*/
func reversePairs1(nums []int) int {
	res := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		a := nums[i]
		for j := i + 1; j < l; j++ {
			b := 2 * nums[j]
			//fmt.Println(a, b)
			if a > b {
				res++
			}
		}
	}
	return res
}
