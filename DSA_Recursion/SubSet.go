package dsarecursion

import "fmt"

func SubSet(unproceesednums []int, nums []int) {
	if len(nums) == 0 {
		fmt.Print(unproceesednums)
		return
	}
	include := append([]int{}, unproceesednums...)
	include = append(include, nums[0])
	nums = nums[1:]
	SubSet(include, nums)
	SubSet(unproceesednums, nums)
}
